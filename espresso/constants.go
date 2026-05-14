// Package espresso contains constants and helpers shared between the op-node
// derivation pipeline, the batcher, and Caff node integrations.
//
// This file (constants.go) is the only file in the espresso package without
// a build tag. The other files (cli.go, ethclient.go, interface.go) require
// dependencies that are not buildable on mips64 (the op-program fault-proof
// target) and are tagged with //go:build !mips64. Keep this file free of
// imports that are not mips64-clean so that mips64-reachable code (in
// particular op-node/rollup/derive and op-node/rollup) can continue to
// reference these constants.
package espresso

// DefaultBatchAuthLookbackWindow is the default number of L1 blocks before
// the batch submission to scan for a BatchInfoAuthenticated event. The
// authentication transaction must land in this window (or in the same block
// as the batch submission) for the batch to be considered valid.
//
// At ~12s per L1 block, 100 blocks ≈ 20 minutes. This gives the batcher
// time to land the batch data transaction on L1 after the authentication
// transaction, even under L1 congestion or batcher restarts. The window is
// intentionally generous: a tighter window risks rejecting valid batches
// during congestion spikes.
//
// Not exposed as a CLI flag; configured per-chain via rollup.json
// (Config.BatchAuthLookbackWindow) and consumed via
// rollup.Config.BatchAuthLookbackWindowOrDefault().
const DefaultBatchAuthLookbackWindow uint64 = 100
