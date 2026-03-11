# Project Instructions

## Project
- Description: This is [EspressoSystems/optimism-espresso-integration](https://github.com/EspressoSystems/optimism-espresso-integration), a fork of [celo-org/optimism](https://github.com/celo-org/optimism) (which forks [ethereum-optimism/optimism](https://github.com/ethereum-optimism/optimism)), adding Espresso sequencer integration and TEE-based batch authentication to the OP Stack.
- Architecture: The batcher posts L2 blocks to Espresso (HotShot) for fast soft confirmations, then reads confirmed batches back before submitting to L1. A "Caff node" (caffeinated node) derives L2 state directly from Espresso for faster finality. The batcher runs inside a TEE (AWS Nitro enclave) and attests on-chain via `BatchAuthenticator` that it faithfully posts data from Espresso. A non-TEE fallback batcher handles Espresso outages.

## Language & Framework
- Language: Go 1.23, Solidity 0.8.x
- Framework: Foundry (Solidity), Go standard library + go-ethereum
- Key dependencies: `github.com/EspressoSystems/espresso-network/sdks/go`, OpenZeppelin Contracts v4/v5, Foundry forge/cast, golangci-lint, semgrep

## Build Commands
- Build (contracts): `cd packages/contracts-bedrock && FOUNDRY_PROFILE=lite forge build --deny=never`
- Build (devnet): `just build-devnet`
- Test (all): `just tests` (standard OP stack), `just espresso-tests` (full integration, ~30 min), `just devnet-tests` (Docker Compose, slow)
- Test (fast): `just smoke-tests`, `just fast-tests`
- Test (contracts): `just run-l1-espresso-contracts-tests`
- Test (single integration): `just compile-contracts && go test -timeout 35m -p 1 -count 1 -v -run '^TestName$' ./espresso/environment`
- Test (single devnet): `just build-devnet && U_ID=$(id -u) GID=$(id -g) go test -timeout 30m -p 1 -count 1 -v -run '^TestName$' ./espresso/devnet-tests/...`
- Lint (Go): `just golint`
- Lint (Solidity): `just semgrep`
- Lint (Shell): `just shellcheck`
- Generate Go bindings: `just gen-bindings` (run after modifying `BatchAuthenticator` or `BatchInbox`)
- Clean up test containers: `just remove-containers`

## Code Style
- Import style: Absolute imports for Go; Foundry remappings for Solidity
- Naming convention: `camelCase`/`PascalCase` for Go; `camelCase`/`PascalCase` for Solidity
- File organization: Espresso-specific Go code in `espresso.go` or `espresso_*.go` files within each package; Espresso contracts in `packages/contracts-bedrock/src/L1/`; all Espresso config flags in `espresso/cli.go` with `espresso.` prefix
- Branch naming: `xy/branch-name` (author initials + description); default branch is `celo-integration-rebase-XX.YY`
- Commits: add yourself as co-author (`Co-Authored-By:`)

## Diff Discipline
This is a fork-of-a-fork with regular biweekly rebases via cherry-pick. **Keep the diff addition-heavy:**
- Put new code in Espresso-specific files and directories (`espresso/`, `**/espresso.go`, `**/espresso_*.go`)
- Minimize modifications to core Optimism files — keep changes small and isolated when required
- Avoid large refactors of upstream code; prefer wrapping or extending
- See `docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md` for the full sync procedure

## Where Things Live

| Location | What |
|----------|------|
| `espresso/` | Core Go package: streamer, batch buffer, CLI config, interfaces |
| `op-batcher/batcher/espresso.go` | Batcher write path: submission to Espresso, batch loading loop, L1 posting |
| `op-node/rollup/derive/espresso_batch.go` | `EspressoBatch` type and block↔batch↔Espresso-tx conversions |
| `op-node/rollup/derive/attributes_queue.go` | Caff node derivation (`CaffNextBatch()`) — main hook into derivation pipeline |
| `op-service/crypto/espresso.go` | `ChainSigner` interface for signing transactions and arbitrary data |
| `op-batcher/bindings/` | Generated Go bindings for Espresso L1 contracts |
| `op-deployer/pkg/deployer/*/espresso.go` | Deployment pipeline for Espresso contracts |
| `op-alt-da/cmd/daserver/espresso.go` | Alternative DA store fetching from Espresso |
| `packages/contracts-bedrock/src/L1/BatchAuthenticator.sol` | Batch signature verification, TEE attestation, batcher switching |
| `packages/contracts-bedrock/src/L1/BatchInbox.sol` | Delegates validation to `BatchAuthenticator` |
| `packages/contracts-bedrock/lib/espresso-tee-contracts/` | Git submodule: TEE verifier interfaces and implementations |

## Error Handling
- Errors should: be returned and propagated explicitly
- Never: silently swallow errors, especially in the batcher or streamer (silent failures cause chain stalls or incorrect data)

## Testing
- Test framework: Foundry (Solidity), Go `testing` package
- Test location: `packages/contracts-bedrock/test/L1/` (contracts), `espresso/environment/` (integration, requires Docker), `espresso/devnet-tests/` (devnet, requires Docker + `ghcr.io` auth)
- Requires: Nix (`nix develop .`), Docker for integration/devnet tests

## Pre-commit Hooks
- No auto-install; run `just golint`, `just semgrep`, `just shellcheck` manually before committing

## Security Considerations
- **Derivation pipeline changes** affect canonical chain state. Keep Espresso logic in isolated paths (e.g., `CaffNextBatch()`) rather than altering core derivation flow. Changes here require the most scrutiny.
- **Solidity contracts** — `BatchAuthenticator` is upgradeable (transparent proxy); treat every change with full rigor: reentrancy, access control, upgrade safety
- **Error handling** — silent failures in batcher or streamer cause chain stalls or incorrect data
- See `espresso/SECURITY_ANALYSIS.md` for the full 3-layer security model

## Key Documentation
- [`README_ESPRESSO.md`](README_ESPRESSO.md) — Dev environment, devnet guide, enclave setup
- [`docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md`](docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md) — Biweekly sync procedure
- [`docs/README_ESPRESSO_DEPLOY_CONFIG.md`](docs/README_ESPRESSO_DEPLOY_CONFIG.md) — Deployment configuration
- [`espresso/SECURITY_ANALYSIS.md`](espresso/SECURITY_ANALYSIS.md) — Security analysis
- [`espresso/docs/metrics.md`](espresso/docs/metrics.md) — Monitoring metrics

## Constraints
- Always: keep this file up to date as the project evolves — it's the only context that persists across sessions
