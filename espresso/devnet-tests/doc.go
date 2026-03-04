// Package devnet_tests runs integration tests against a Docker-based Espresso devnet.
//
// Run with OP_E2E_SKIP_ALLOC_GEN=1 so op-e2e config init skips alloc generation
// (devnet uses Docker and does not need L1/L2 allocs). Otherwise init panics
// during DeploySuperchain. Example:
//
//	OP_E2E_SKIP_ALLOC_GEN=1 go test -timeout 60m -p 1 -run 'TestSmokeWithoutTEE' -v ./espresso/devnet-tests/...
//
// Or use: just devnet-tests (from repo root), which sets the env automatically.
package devnet_tests
