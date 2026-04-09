package logmodule

// Non-dashboard log constants for events referenced in metrics.md that are useful for
// debugging but are not directly monitored by the DataDog dashboard.

const (
	// Batcher

	// AddedL2BlockToChannelManager is emitted each time a new L2 block is enqueued for batching.
	AddedL2BlockToChannelManager = "Added L2 block to channel manager"

	// ClearingState is emitted on a batcher state reset. Even a single occurrence is suspicious.
	ClearingState = "Clearing state"

	// FailedToDeriveBatchFromBlock is emitted when the batcher cannot construct an Espresso transaction from an L2 block.
	FailedToDeriveBatchFromBlock = "Failed to derive batch from block"

	// TransactionFailedToSend is emitted when an L1 submission attempt fails.
	TransactionFailedToSend = "Transaction failed to send"

	// DARequestFailed is emitted when an AltDA submission fails.
	DARequestFailed = "DA request failed"

	// FoundL2Reorg is emitted when the batcher detects an L2 reorg.
	FoundL2Reorg = "Found L2 reorg"

	// Node (Caff, Non-caff, Sequencer)

	// NewL1SafeBlock is emitted each time a new L1 safe block is observed.
	NewL1SafeBlock = "New L1 safe block"

	// InsertedNewL2UnsafeBlock is emitted each time a new L2 unsafe block is inserted.
	InsertedNewL2UnsafeBlock = "Inserted new L2 unsafe block"

	// HitFinalizedL2Head is emitted during a sync reset when the node reaches the finalized L2 head.
	// An increasing L2 safe number here serves as an alternative indicator for "new L2 safe blocks"
	// after a pipeline reset (non-Caff validator node only).
	HitFinalizedL2Head = "Hit finalized L2 head, returning immediately"

	// DerivationProcessError is emitted on a recoverable derivation pipeline error.
	DerivationProcessError = "Derivation process error"

	// DroppingBatch is emitted when a malformed or invalid batch is discarded.
	DroppingBatch = "Dropping batch"

	// FailedToParseFrames is emitted when frame parsing fails for a batch.
	FailedToParseFrames = "Failed to parse frames"

	// Sequencer

	// EngineFailedTemporarily is emitted when the execution engine fails and the sequencer backs off.
	EngineFailedTemporarily = "Engine failed temporarily, backing off sequencer"

	// EngineResetConfirmed is emitted after a successful engine reset, allowing the sequencer to resume.
	EngineResetConfirmed = "Engine reset confirmed, sequencer may continue"
)
