package core

// PeerType represents the type of a peer
type PeerType string

// EligibleList represents the list of peers who participate in consensus inside a shard
const EligibleList PeerType = "eligible"

// WaitingList represents the list of peers who don't participate in consensus but will join the next epoch
const WaitingList PeerType = "waiting"

// LeavingList represents the list of peers who were taken out of eligible and waiting because of rating
const LeavingList PeerType = "leaving"

// ObserverList represents the list of peers who don't participate in consensus but will join the next epoch
const ObserverList PeerType = "observer"

// NewList -
const NewList PeerType = "new"

// UnVersionedAppString represents the default app version that indicate that the binary wasn't build by setting
// the appVersion flag
const UnVersionedAppString = "undefined"

// NodeType represents the node's role in the network
type NodeType string

// NodeTypeObserver signals that a node is running as observer node
const NodeTypeObserver NodeType = "observer"

// NodeTypeValidator signals that a node is running as validator node
const NodeTypeValidator NodeType = "validator"

// pkPrefixSize specifies the max numbers of chars to be displayed from one publc key
const pkPrefixSize = 12

// FileModeUserReadWrite represents the permission for a file which allows the user for reading and writing
const FileModeUserReadWrite = 0600

// MaxTxNonceDeltaAllowed specifies the maximum difference between an account's nonce and a received transaction's nonce
// in order to mark the transaction as valid.
const MaxTxNonceDeltaAllowed = 30000

// MaxBulkTransactionSize specifies the maximum size of one bulk with txs which can be send over the network
//TODO convert this const into a var and read it from config when this code moves to another binary
const MaxBulkTransactionSize = 2 << 17 //128KB bulks

// ConsensusTopic is the topic used in consensus algorithm
const ConsensusTopic = "consensus"

// HeartbeatTopic is the topic used for heartbeat signaling
const HeartbeatTopic = "heartbeat"

// PathShardPlaceholder represents the placeholder for the shard ID in paths
const PathShardPlaceholder = "[S]"

// PathEpochPlaceholder represents the placeholder for the epoch number in paths
const PathEpochPlaceholder = "[E]"

// PathIdentifierPlaceholder represents the placeholder for the identifier in paths
const PathIdentifierPlaceholder = "[I]"

// MetricCurrentRound is the metric for monitoring the current round of a node
const MetricCurrentRound = "erd_current_round"

// MetricNonce is the metric for monitoring the nonce of a node
const MetricNonce = "erd_nonce"

// MetricProbableHighestNonce is the metric for monitoring the max speculative nonce received by the node by listening on the network
const MetricProbableHighestNonce = "erd_probable_highest_nonce"

// MetricNumConnectedPeers is the metric for monitoring the number of connected peers
const MetricNumConnectedPeers = "erd_num_connected_peers"

// MetricNumConnectedPeersClassification is the metric for monitoring the number of connected peers split on the connection type
const MetricNumConnectedPeersClassification = "erd_num_connected_peers_classification"

// MetricSynchronizedRound is the metric for monitoring the synchronized round of a node
const MetricSynchronizedRound = "erd_synchronized_round"

// MetricIsSyncing is the metric for monitoring if a node is syncing
const MetricIsSyncing = "erd_is_syncing"

// MetricPublicKeyBlockSign is the metric for monitoring public key of a node used in block signing
const MetricPublicKeyBlockSign = "erd_public_key_block_sign"

// MetricShardId is the metric for monitoring shard id of a node
const MetricShardId = "erd_shard_id"

// MetricTxPoolLoad is the metric for monitoring number of transactions from pool of a node
const MetricTxPoolLoad = "erd_tx_pool_load"

// MetricCountLeader is the metric for monitoring number of rounds when a node was leader
const MetricCountLeader = "erd_count_leader"

// MetricCountConsensus is the metric for monitoring number of rounds when a node was in consensus group
const MetricCountConsensus = "erd_count_consensus"

// MetricCountAcceptedBlocks is the metric for monitoring number of blocks that was accepted proposed by a node
const MetricCountAcceptedBlocks = "erd_count_accepted_blocks"

// MetricNodeType is the metric for monitoring the type of the node
const MetricNodeType = "erd_node_type"

// MetricLiveValidatorNodes is the metric for monitoring live validators on the network
const MetricLiveValidatorNodes = "erd_live_validator_nodes"

// MetricConnectedNodes is the metric for monitoring total connected nodes on the network
const MetricConnectedNodes = "erd_connected_nodes"

// MetricCpuLoadPercent is the metric for monitoring CPU load [%]
const MetricCpuLoadPercent = "erd_cpu_load_percent"

// MetricMemLoadPercent is the metric for monitoring memory load [%]
const MetricMemLoadPercent = "erd_mem_load_percent"

// MetricMemTotal is the metric for monitoring total memory bytes
const MetricMemTotal = "erd_mem_total"

// MetricMemUsedGolang is the metric that stores the total memory used by golang in bytes
const MetricMemUsedGolang = "erd_mem_used_golang"

// MetricMemUsedSystem is the metric that stores the total memory used by the system in bytes
const MetricMemUsedSystem = "erd_mem_used_sys"

// MetricNetworkRecvPercent is the metric for monitoring network receive load [%]
const MetricNetworkRecvPercent = "erd_network_recv_percent"

// MetricNetworkRecvBps is the metric for monitoring network received bytes per second
const MetricNetworkRecvBps = "erd_network_recv_bps"

// MetricNetworkRecvBpsPeak is the metric for monitoring network received peak bytes per second
const MetricNetworkRecvBpsPeak = "erd_network_recv_bps_peak"

// MetricNetworkSentPercent is the metric for monitoring network sent load [%]
const MetricNetworkSentPercent = "erd_network_sent_percent"

// MetricNetworkSentBps is the metric for monitoring network sent bytes per second
const MetricNetworkSentBps = "erd_network_sent_bps"

// MetricNetworkSentBpsPeak is the metric for monitoring network sent peak bytes per second
const MetricNetworkSentBpsPeak = "erd_network_sent_bps_peak"

// MetricRoundTime is the metric for round time in seconds
const MetricRoundTime = "erd_round_time"

// MetricEpochNumber is the metric for the number of epoch
const MetricEpochNumber = "erd_epoch_number"

// MetricAppVersion is the metric for the current app version
const MetricAppVersion = "erd_app_version"

// MetricNumTxInBlock is the metric for the number of transactions in the proposed block
const MetricNumTxInBlock = "erd_num_tx_block"

// MetricConsensusState is the metric for consensus state of node proposer,participant or not consensus group
const MetricConsensusState = "erd_consensus_state"

// MetricNumMiniBlocks is the metric for number of miniblocks in a block
const MetricNumMiniBlocks = "erd_num_mini_blocks"

// MetricConsensusRoundState is the metric for consensus round state for a block
const MetricConsensusRoundState = "erd_consensus_round_state"

// MetricCrossCheckBlockHeight is the metric that store cross block height
const MetricCrossCheckBlockHeight = "erd_metric_cross_check_block_height"

// MetricNumProcessedTxs is the metric that stores the number of transactions processed
const MetricNumProcessedTxs = "erd_num_transactions_processed"

// MetricCurrentBlockHash is the metric that stores the current block hash
const MetricCurrentBlockHash = "erd_current_block_hash"

// MetricCurrentRoundTimestamp is the metric that stores current round timestamp
const MetricCurrentRoundTimestamp = "erd_current_round_timestamp"

// MetricHeaderSize is the metric that stores the current block size
const MetricHeaderSize = "erd_current_block_size"

// MetricMiniBlocksSize is the metric that stores the current block size
const MetricMiniBlocksSize = "erd_mini_blocks_size"

// MetricNumShardHeadersFromPool is the metric that stores number of shard header from pool
const MetricNumShardHeadersFromPool = "erd_num_shard_headers_from_pool"

// MetricNumShardHeadersProcessed is the metric that stores number of shard header processed
const MetricNumShardHeadersProcessed = "erd_num_shard_headers_processed"

// MetricNumTimesInForkChoice is the metric that counts how many time a node was in fork choice
const MetricNumTimesInForkChoice = "erd_fork_choice_count"

//MetricHighestFinalBlockInShard is the metric that stores the highest nonce block notarized by metachain for current shard
const MetricHighestFinalBlockInShard = "erd_highest_notarized_block_by_metachain_for_current_shard"

//MetricLatestTagSoftwareVersion is the metric that stores the latest tag software version
const MetricLatestTagSoftwareVersion = "erd_latest_tag_software_version"

//MetricCountConsensusAcceptedBlocks is the metric for monitoring number of blocks accepted when the node was in consensus group
const MetricCountConsensusAcceptedBlocks = "erd_count_consensus_accepted_blocks"

//MetricRewardsValue is the metric that stores rewards value
const MetricRewardsValue = "erd_rewards_value"

//MetricNodeDisplayName is the metric that stores the name of the node
const MetricNodeDisplayName = "erd_node_display_name"

//MetricConsensusGroupSize is the metric for consensus group size
const MetricConsensusGroupSize = "erd_metric_consensus_group_size"

//MetricNumValidators is the metric for the number of validators
const MetricNumValidators = "erd_metric_num_validators"

// MetricPeerType is the metric which tells the peer's type (in eligible list, in waiting list, or observer)
const MetricPeerType = "erd_peer_type"

//MetricLeaderPercentage is the metric for leader rewards percentage
const MetricLeaderPercentage = "erd_metric_leader_percentage"

//MetricDenominationCoefficient is the metric for denomination coefficient that is used in views
const MetricDenominationCoefficient = "erd_metric_denomination_coefficient"

// MetricRoundAtEpochStart is the metric for storing the first round of the current epoch
const MetricRoundAtEpochStart = "erd_round_at_epoch_start"

// MetricRoundsPerEpoch is the metric that tells the number of rounds in an epoch
const MetricRoundsPerEpoch = "erd_rounds_per_epoch"

// MetricRoundsPassedInCurrentEpoch is the metric that tells the number of rounds passed in current epoch
const MetricRoundsPassedInCurrentEpoch = "erd_rounds_passed_in_current_epoch"

//MetricReceivedProposedBlock is the metric that specify the moment in the round when the received block has reached the
//current node. The value is provided in percent (0 meaning it has been received just after the round started and
//100 meaning that the block has been received in the last moment of the round)
const MetricReceivedProposedBlock = "erd_consensus_received_proposed_block"

//MetricCreatedProposedBlock is the metric that specify the percent of the block subround used for header and body
//creation (0 meaning that the block was created in no-time and 100 meaning that the block creation used all the
//subround spare duration)
const MetricCreatedProposedBlock = "erd_consensus_created_proposed_block"

//MetricProcessedProposedBlock is the metric that specify the percent of the block subround used for header and body
//processing (0 meaning that the block was processed in no-time and 100 meaning that the block processing used all the
//subround spare duration)
const MetricProcessedProposedBlock = "erd_consensus_processed_proposed_block"

// MetricMinGasPrice is the metric that specifies min gas price
const MetricMinGasPrice = "erd_min_gas_price"

// MetricChainId is the metric that specifies current chain id
const MetricChainId = "erd_chain_id"

// MetachainShardId will be used to identify a shard ID as metachain
const MetachainShardId = uint32(0xFFFFFFFF)

// AllShardId will be used to identify that a message is for all shards
const AllShardId = uint32(0xFFFFFFF0)

// MegabyteSize represents the size in bytes of a megabyte
const MegabyteSize = 1024 * 1024

// BaseOperationCost represents the field name for base operation costs
const BaseOperationCost = "BaseOperationCost"

// BuiltInCost represents the field name for built in operation costs
const BuiltInCost = "BuiltInCost"

// MetaChainSystemSCsCost represents the field name for metachain system smart contract operation costs
const MetaChainSystemSCsCost = "MetaChainSystemSCsCost"

const (
	// StorerOrder defines the order of storers to be notified of a start of epoch event
	StorerOrder = iota
	// NodesCoordinatorOrder defines the order in which NodesCoordinator is notified of a start of epoch event
	NodesCoordinatorOrder
	// ConsensusOrder defines the order in which Consensus is notified of a start of epoch event
	ConsensusOrder
)

// NodeState specifies what type of state a node could have
type NodeState int

const (
	// NsSynchronized defines ID of a state of synchronized
	NsSynchronized NodeState = iota
	// NsNotSynchronized defines ID of a state of not synchronized
	NsNotSynchronized
	// NsNotCalculated defines ID of a state which is not calculated
	NsNotCalculated
)

// MetricP2PPeerInfo is the metric for the node's p2p info
const MetricP2PPeerInfo = "erd_p2p_peer_info"

// MetricP2PIntraShardValidators is the metric that outputs the intra-shard connected validators
const MetricP2PIntraShardValidators = "erd_p2p_intra_shard_validators"

// MetricP2PCrossShardValidators is the metric that outputs the cross-shard connected validators
const MetricP2PCrossShardValidators = "erd_p2p_cross_shard_validators"

// MetricP2PIntraShardObservers is the metric that outputs the intra-shard connected observers
const MetricP2PIntraShardObservers = "erd_p2p_intra_shard_observers"

// MetricP2PCrossShardObservers is the metric that outputs the cross-shard connected observers
const MetricP2PCrossShardObservers = "erd_p2p_cross_shard_observers"

// MetricP2PUnknownPeers is the metric that outputs the unknown-shard connected peers
const MetricP2PUnknownPeers = "erd_p2p_unknown_shard_peers"

// MetricP2PNumConnectedPeersClassification is the metric for monitoring the number of connected peers split on the connection type
const MetricP2PNumConnectedPeersClassification = "erd_p2p_num_connected_peers_classification"

// HighestRoundFromBootStorage is the key for the highest round that is saved in storage
const HighestRoundFromBootStorage = "highestRoundFromBootStorage"

// TriggerRegistryKeyPrefix is the key prefix to save epoch start registry to storage
const TriggerRegistryKeyPrefix = "epochStartTrigger_"

// TriggerRegistryInitialKeyPrefix is the key prefix to save initial data to storage
const TriggerRegistryInitialKeyPrefix = "initial_value_epoch_"

// NodesCoordinatorRegistryKeyPrefix is the key prefix to save epoch start registry to storage
const NodesCoordinatorRegistryKeyPrefix = "indexHashed_"
