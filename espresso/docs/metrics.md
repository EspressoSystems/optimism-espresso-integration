# Metrics

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
  `"Added L2 block to channel manager"`
- Espresso batch submissions
  `"Submitted transaction to Espresso"`
- L1 batch submissions
  `"Transaction confirmed"`
- Espresso transaction queue size
  `"Espresso transaction submitter queue status"`
- AltDA submissions
  `"Sent txdata to altda layer and received commitment"`
- Espresso batches fetched
  `"Inserting accepted batch"`

### Recoverable Errors

Events that we need to monitor and raise alerts if they're encountered often:

- State reset (even once is suspicious)
  `"Clearing state"`
- Espresso transaction creation failed
  `"Failed to derive batch from block"`
- L1 submission failed
  `"Transaction failed to send"`
- AltDA submission failed
  `"DA request failed"`
- L2 reorg detected
  `"Found L2 reorg"`

### Critical Errors

- L1 finalized height not increasing
- L2 unsafe height not increasing
- L2 safe height not increasing

### Potential Issue Indicators

Non-errors that can indicate preconditions for a problem to occur:

- Gas price too high
  `effectiveGasPrice` field of `"Transaction confirmed"` log
- Espresso transaction backlog is growing
  can be derived from Espresso transaction queue metrics above

## Caff Validator Node

### Key Metrics

- Espresso batches fetched
  `"Inserting accepted batch"`
- New L1 safe blocks
  `"New L1 safe block"`
- New L2 unsafe blocks
  `"Inserted new L2 unsafe block"`
- New L2 safe blocks
  `"safe head updated"`

### Recoverable Errors

- Pipeline errors
  `"Derivation process error"`
- Malformed batch
  `"Dropping batch"`, `"Failed to parse frames"`

### Critical Errors

Events that need to raise urgent alerts as they indicate full chain stall:

- L1 finalized height not increasing
- L2 unsafe height not increasing
- L2 safe height not increasing

## Non-caff Validator Node

### Key Metrics

- New L1 safe blocks
  `"New L1 safe block"`
- New L2 unsafe blocks
  `"Inserted new L2 unsafe block"`
- New L2 safe blocks
  Either `"safe head updated"` or `"Hit finalized L2 head, returning immediately"` with increasing
  L2 safe number. The former is the normal case, and the latter happens after a reset.

### Recoverable Errors

- Pipeline errors
  `"Derivation process error"`
- Malformed batch
  `"Dropping batch"`, `"Failed to parse frames"`

### Critical Errors

Events that need to raise urgent alerts as they indicate full chain stall:

- L1 finalized height not increasing
- L2 unsafe height not increasing
- L2 safe height not increasing

## Sequencer

All events of Decaff Validator Node, and:

### Key Metrics

- Blocks produced
  `"Sequencer sealed block"`

### Recoverable Errors

- Engine failure
  `"Engine failed temporarily, backing off sequencer"`
- Engine reset
  `"Engine reset confirmed, sequencer may continue"`
