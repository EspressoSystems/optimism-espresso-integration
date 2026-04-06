// Package logmodule defines log message string constants for the Espresso integration.
//
// dashboard_keys.go contains constants for events monitored by the DataDog dashboard.
// These must be kept in sync with any dashboard queries, alerts, or deployment investigation
// runbooks that reference them by name. If you change any of these strings, update the
// DataDog dashboard queries and alerts at the same time.
package logmodule

const (
	// SequencerSealedBlock is emitted by the sequencer each time it seals a new L2 block.
	// Monitored as the primary "blocks produced" metric.
	SequencerSealedBlock = "Sequencer sealed block"

	// CrossSafeHeadUpdated is emitted by the op-node status tracker each time the cross-safe L2
	// head advances. Monitored as "new L2 safe blocks" for both the Caff and non-Caff validator nodes.
	CrossSafeHeadUpdated = "Cross safe head updated"

	// TransactionConfirmedOnEspresso is emitted by the batcher after it verifies that a transaction
	// was included in HotShot consensus.
	TransactionConfirmedOnEspresso = "Transaction confirmed on Espresso"

	// TransactionSuccessfullyPublished is emitted by the tx manager after a transaction is accepted
	// by the L1 RPC. Monitored as "L1 batch submissions".
	TransactionSuccessfullyPublished = "Transaction successfully published"

	// SubmittedTransactionToEspresso is emitted by the batcher each time it sends a transaction to
	// the Espresso sequencer. Monitored as "Espresso batch submissions".
	SubmittedTransactionToEspresso = "Submitted transaction to Espresso"

	// ChannelClosed is emitted by the batcher channel manager when a channel is closed and ready
	// for frame submission.
	ChannelClosed = "Channel closed"

	// ReceivedBlockFromEspresso is emitted by the batcher each time it reads a confirmed L2 block
	// back from the Espresso query service.
	ReceivedBlockFromEspresso = "Received block from Espresso"
)
