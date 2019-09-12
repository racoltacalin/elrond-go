package block

import (
	"sync"

	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/data/state/addressConverters"
	"github.com/ElrondNetwork/elrond-go/data/transaction"
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/sharding"
)

// TxPoolsCleaner represents a pools cleaner that check if a transaction should be in pool
type TxPoolsCleaner struct {
	accounts         state.AccountsAdapter
	shardCoordinator sharding.Coordinator
	dataPool         dataRetriever.PoolsHolder
	addrConverter    *addressConverters.PlainAddressConverter
	numRemovedTxs    uint64
	mutNumRemovedTxs sync.RWMutex
}

// NewTxsPoolsCleaner will return a new transaction pools cleaner
func NewTxsPoolsCleaner(
	accounts state.AccountsAdapter,
	shardCoordinator sharding.Coordinator,
	dataPool dataRetriever.PoolsHolder,
) (*TxPoolsCleaner, error) {
	if accounts == nil || accounts.IsInterfaceNil() {
		return nil, process.ErrNilAccountsAdapter
	}
	if shardCoordinator == nil || shardCoordinator.IsInterfaceNil() {
		return nil, process.ErrNilShardCoordinator
	}
	if dataPool == nil {
		return nil, process.ErrNilDataPoolHolder
	}
	transactionPool := dataPool.Transactions()
	if transactionPool == nil {
		return nil, process.ErrNilTransactionPool
	}
	addrConverter, err := addressConverters.NewPlainAddressConverter(32, "0x")
	if err != nil {
		return nil, err
	}

	return &TxPoolsCleaner{
		accounts:         accounts,
		shardCoordinator: shardCoordinator,
		dataPool:         dataPool,
		addrConverter:    addrConverter,
		numRemovedTxs:    0,
	}, nil
}

// Clean will check if in pools exits transactions with nonce low that transaction sender account nonce
// and if tx have low nonce will be removed from pools
func (tpc *TxPoolsCleaner) Clean(haveTime func() bool) error {
	if haveTime == nil {
		return process.ErrNilHaveTimeHandler
	}

	shardId := tpc.shardCoordinator.SelfId()
	transactions := tpc.dataPool.Transactions()
	numOfShards := tpc.shardCoordinator.NumberOfShards()

	for destShardId := uint32(0); destShardId < numOfShards; destShardId++ {
		cacherId := process.ShardCacherIdentifier(shardId, destShardId)
		txsPool := transactions.ShardDataStore(cacherId)

		for _, key := range txsPool.Keys() {
			if !haveTime() {
				return nil
			}

			obj, ok := txsPool.Peek(key)
			if !ok {
				continue
			}

			tx, ok := obj.(*transaction.Transaction)
			if !ok {
				continue
			}

			sndAddr := tx.GetSndAddress()
			addr, err := tpc.addrConverter.CreateAddressFromPublicKeyBytes(sndAddr)
			if err != nil {
				txsPool.Remove(key)
				tpc.incrementNumRemovedTxs()
				continue
			}

			accountHandler, err := tpc.accounts.GetExistingAccount(addr)
			if err != nil {
				txsPool.Remove(key)
				tpc.incrementNumRemovedTxs()
				continue
			}

			accountNonce := accountHandler.GetNonce()
			txNonce := tx.Nonce
			lowerNonceInTx := txNonce < accountNonce
			if lowerNonceInTx {
				txsPool.Remove(key)
				tpc.incrementNumRemovedTxs()
			}
		}
	}

	return nil
}

func (tpc *TxPoolsCleaner) incrementNumRemovedTxs() {
	tpc.mutNumRemovedTxs.Lock()
	tpc.numRemovedTxs++
	tpc.mutNumRemovedTxs.Unlock()
}

// NumRemovedTxs will return the number of removed txs from pools
func (tpc *TxPoolsCleaner) NumRemovedTxs() uint64 {
	tpc.mutNumRemovedTxs.Lock()
	defer tpc.mutNumRemovedTxs.Unlock()

	return tpc.numRemovedTxs
}
