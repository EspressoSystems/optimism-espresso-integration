# Metrics

> **Log constants**: Log message strings referenced in this document are defined as Go constants in
> [`espresso/logmodule/dashboard_keys.go`](../logmodule/dashboard_keys.go) (DataDog dashboard logs) and
> [`espresso/logmodule/log_keys.go`](../logmodule/log_keys.go) (debugging logs). If you change a dashboard
> log string you must update the corresponding DataDog queries and alerts at the same time.

This document outlines the monitoring framework for our system components, organized into the following categories:

- **Key Metrics**: Metrics that belong on the dashboard for operational visibility
- **Recoverable Errors**: Events that we need to monitor and raise alerts if they're encountered often, but do not necessarily lead to liveness or safety violations
- **Critical Errors**: Events that need to raise urgent alerts as they indicate full chain stall or stall of the particular service
- **Potential Issue Indicators**: Non-errors that can indicate preconditions for a problem to occur

Each indicator points to a log event to monitor.

## Batcher

### Key Metrics

Metrics that belong on the dashboard:

- Blocks enqueued for batching to L1/AltDA:
  `logmodule.AddedL2BlockToChannelManager`
- Espresso batch submissions
  `logmodule.SubmittedTransactionToEspresso`
- Espresso transaction confirmations
  `logmodule.TransactionConfirmedOnEspresso`
- Blocks received from Espresso
  `logmodule.ReceivedBlockFromEspresso`
- Channel sealed for L1 submission
  `logmodule.ChannelClosed`
- L1 batch submissions
  `logmodule.TransactionSuccessfullyPublished`

### Recoverable Errors

Events that we need to monitor and raise alerts if they're encountered often:

- State reset (even once is suspicious)
  `logmodule.ClearingState`
- Espresso transaction creation failed
  `logmodule.FailedToDeriveBatchFromBlock`
- L1 submission failed
  `logmodule.TransactionFailedToSend`
- AltDA submission failed
  `logmodule.DARequestFailed`
- L2 reorg detected
  `logmodule.FoundL2Reorg`

### Critical Errors

- L1 finalized height not increasing
- L2 unsafe height not increasing
- L2 safe height not increasing

### Potential Issue Indicators

Non-errors that can indicate preconditions for a problem to occur:

- Gas price too high
  `effectiveGasPrice` field of `logmodule.TransactionSuccessfullyPublished` log
- Espresso transaction backlog is growing
  can be derived from Espresso transaction queue metrics above

## Caff Validator Node

### Key Metrics

- New L1 safe blocks
  `logmodule.NewL1SafeBlock`
- New L2 unsafe blocks
  `logmodule.InsertedNewL2UnsafeBlock`
- New L2 safe blocks
  `logmodule.CrossSafeHeadUpdated`

### Recoverable Errors

- Pipeline errors
  `logmodule.DerivationProcessError`
- Malformed batch
  `logmodule.DroppingBatch`, `logmodule.FailedToParseFrames`

### Critical Errors

Events that need to raise urgent alerts as they indicate full chain stall:

- L1 finalized height not increasing
- L2 unsafe height not increasing
- L2 safe height not increasing

## Non-caff Validator Node

### Key Metrics

- New L1 safe blocks
  `logmodule.NewL1SafeBlock`
- New L2 unsafe blocks
  `logmodule.InsertedNewL2UnsafeBlock`
- New L2 safe blocks
  Either `logmodule.CrossSafeHeadUpdated` or `logmodule.HitFinalizedL2Head` with increasing L2 safe
  number. The former is the normal case, and the latter happens after a reset.

### Recoverable Errors

- Pipeline errors
  `logmodule.DerivationProcessError`
- Malformed batch
  `logmodule.DroppingBatch`, `logmodule.FailedToParseFrames`

### Critical Errors

Events that need to raise urgent alerts as they indicate full chain stall:

- L1 finalized height not increasing
- L2 unsafe height not increasing
- L2 safe height not increasing

## Sequencer

All events of Decaff Validator Node, and:

### Key Metrics

- Blocks produced
  `logmodule.SequencerSealedBlock`

### Recoverable Errors

- Engine failure
  `logmodule.EngineFailedTemporarily`
- Engine reset
  `logmodule.EngineResetConfirmed`
