package dataprocessor

import (
	"github.com/ElrondNetwork/elastic-indexer-go/workItems"
	storer2ElasticData "github.com/ElrondNetwork/elrond-go/cmd/storer2elastic/data"
	"github.com/ElrondNetwork/elrond-go/cmd/storer2elastic/databasereader"
	"github.com/ElrondNetwork/elrond-go/core/statistics"
	"github.com/ElrondNetwork/elrond-go/data"
	"github.com/ElrondNetwork/elrond-go/data/block"
	"github.com/ElrondNetwork/elrond-go/sharding"
	"github.com/ElrondNetwork/elrond-go/storage"
)

// DatabaseReaderHandler defines the actions that a database reader has to do
type DatabaseReaderHandler interface {
	GetDatabaseInfo() ([]*databasereader.DatabaseInfo, error)
	GetStaticDatabaseInfo() ([]*databasereader.DatabaseInfo, error)
	GetHeaders(dbInfo *databasereader.DatabaseInfo) ([]data.HeaderHandler, error)
	LoadPersister(dbInfo *databasereader.DatabaseInfo, unit string) (storage.Persister, error)
	LoadStaticPersister(dbInfo *databasereader.DatabaseInfo, unit string) (storage.Persister, error)
	IsInterfaceNil() bool
}

// NodesCoordinator defines the actions that a nodes' coordinator has to do
type NodesCoordinator interface {
	sharding.NodesCoordinator
	EpochStartPrepare(metaHdr data.HeaderHandler, body data.BodyHandler)
}

// HeaderMarshalizerHandler defines the actions that a header marshalizer has to do
type HeaderMarshalizerHandler interface {
	UnmarshalShardHeader(headerBytes []byte) (*block.Header, error)
	UnmarshalMetaBlock(headerBytes []byte) (*block.MetaBlock, error)
	IsInterfaceNil() bool
}

// DataReplayerHandler defines the actions that a data replayer has to do
type DataReplayerHandler interface {
	Range(handler func(persistedData storer2ElasticData.RoundPersistedData) bool) error
	IsInterfaceNil() bool
}

// TPSBenchmarkUpdaterHandler defines the actions that a TPS benchmark updater has to do
type TPSBenchmarkUpdaterHandler interface {
	IndexTPSForMetaBlock(metaBlock *block.MetaBlock)
	IsInterfaceNil() bool
}

// RatingProcessorHandler defines the actions that a rating processor has to do
type RatingProcessorHandler interface {
	IndexRatingsForEpochStartMetaBlock(metaBlock *block.MetaBlock) error
	IsInterfaceNil() bool
}

// StorageDataIndexer defines the actions that a storage data indexer has to do
type StorageDataIndexer interface {
	SaveRoundsInfo(roundsInfos []workItems.RoundInfo)
	SaveBlock(body data.BodyHandler, header data.HeaderHandler, txPool map[string]data.TransactionHandler,
		signersIndexes []uint64, notarizedHeadersHashes []string, headerHash []byte)
	SaveValidatorsRating(indexID string, infoRating []workItems.ValidatorRatingInfo)
	SaveValidatorsPubKeys(validatorsPubKeys map[uint32][][]byte, epoch uint32)
	UpdateTPS(tpsBenchmark statistics.TPSBenchmark)
	IsInterfaceNil() bool
}
