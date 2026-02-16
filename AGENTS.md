# AGENTS.md

## Repository Context

This is **NOT** upstream Optimism. This is [EspressoSystems/optimism-espresso-integration](https://github.com/EspressoSystems/optimism-espresso-integration), a fork-of-a-fork:

- **Upstream:** [ethereum-optimism/optimism](https://github.com/ethereum-optimism/optimism)
- **Celo fork:** [celo-org/optimism](https://github.com/celo-org/optimism) (syncs periodically with upstream)
- **This repo:** Espresso's fork of Celo's fork, adding Espresso sequencer integration

## Branch Naming

Default branch: `celo-integration-rebase-XX.YY` (currently `celo-integration-rebase-14.2`).

- `XX` matches Celo's `celo-rebase-XX` branch number
- `YY` is our rebase index on top of Celo's branch (incremented each biweekly sync)

Feature branches: `xy/branch-name`, where `xy` are the author's initials (derive from `git config user.name`).

When creating commits, add yourself as co-author.

See `docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md` for the full sync procedure across this repo, kona, celo-kona, and op-succinct forks.

## Diff Discipline

Because this is a fork-of-a-fork with regular rebases via cherry-pick, **keep the diff addition-heavy**:

- Put new code in **Espresso-specific files and directories** whenever possible (e.g., `espresso/`, `**/espresso.go`, `**/espresso_*.go`)
- Minimize modifications to core Optimism files — when you must touch them, keep changes small and isolated
- Avoid large refactors of upstream code; prefer wrapping or extending
- This makes cherry-picks during biweekly syncs dramatically easier

## Architecture Overview

The integration adds Espresso as a fast confirmation layer for the OP stack. The core idea:

1. The **batcher** posts L2 blocks to Espresso (HotShot consensus) for fast soft confirmations
2. Espresso becomes the **source of truth** for what gets posted to L1 — the batcher reads confirmed batches back from Espresso before submitting them on-chain
3. A new node type, the **Caff node** (caffeinated node), derives L2 state by reading directly from Espresso instead of L1, enabling faster finality
4. The Espresso-enabled batcher runs inside a **TEE** (AWS Nitro enclave) so it can attest on-chain (via `BatchAuthenticator`) that it is faithfully posting data it read from Espresso. A **non-TEE fallback batcher** exists for resilience: if Espresso goes down, the chain switches to the fallback batcher which operates as a standard OP batcher — Caff nodes stop advancing, but the chain survives the outage

### Where Things Live

| Location | What |
|----------|------|
| `espresso/` | Core Go package: streamer (reads from HotShot), batch buffer, CLI config, interfaces |
| `op-batcher/batcher/espresso.go` | Batcher write path: submission to Espresso, batch loading loop, L1 posting with authentication |
| `op-node/rollup/derive/espresso_batch.go` | `EspressoBatch` type and conversion functions (block <-> batch <-> Espresso transaction) |
| `op-node/rollup/derive/attributes_queue.go` | Caff node derivation path (`CaffNextBatch()`) — the main hook into the derivation pipeline |
| `op-service/crypto/espresso.go` | `ChainSigner` interface for signing both transactions and arbitrary data |
| `op-batcher/bindings/` | Generated Go bindings for Espresso L1 contracts |
| `op-deployer/pkg/deployer/*/espresso.go` | Deployment pipeline for Espresso contracts |
| `op-alt-da/cmd/daserver/espresso.go` | Alternative DA store fetching from Espresso |
| `packages/contracts-bedrock/src/L1/BatchAuthenticator.sol` | L1 contract: batch signature verification, TEE attestation, batcher switching |
| `packages/contracts-bedrock/src/L1/BatchInbox.sol` | L1 contract: delegates validation to BatchAuthenticator |
| `packages/contracts-bedrock/lib/espresso-tee-contracts/` | Git submodule: TEE verifier contract interfaces and implementations |

Small modifications exist in core OP files (`op-node/rollup/types.go`, `op-batcher/batcher/driver.go`, `op-batcher/batcher/service.go`, `op-node/service.go`, flag registration files) to wire Espresso in. These are intentionally kept minimal per the diff discipline above.

## Solidity Contracts

Foundry-based, in `packages/contracts-bedrock/`. For fast iteration when testing contract changes, use `just build-dev` in that directory. After modifying contracts that have Go bindings (e.g., `BatchAuthenticator`, `BatchInbox`), regenerate bindings with `just gen-bindings` from the repo root.

## Running Tests

Requires Nix (`nix develop .`). Integration and devnet tests also require Docker (authenticated to `ghcr.io`).

```bash
just smoke-tests              # Fast smoke tests
just espresso-tests           # Full integration tests (~30 min)
just devnet-tests             # Docker Compose-based devnet tests (slow, as it runs build-devnet first)
just tests                    # Standard OP stack tests (no Espresso)
just fast-tests               # Fast subset of OP stack tests
just golint                   # Go linter
just remove-containers        # Clean up stuck test containers
```

### Running Individual Tests

Integration tests (`espresso/environment/`) and devnet tests (`espresso/devnet-tests/`) are the two test suites that require Docker. Both are slow. Prefer not to run devnet tests unless directly working on or debugging one, or when the user requests it.

**Single integration test:**

```bash
just compile-contracts
go test -timeout 35m -p 1 -count 1 -v -run '^TestName$' ./espresso/environment
```

**Single devnet test** (requires building Docker images first):

```bash
just build-devnet
U_ID=$(id -u) GID=$(id -g) go test -timeout 30m -p 1 -count 1 -v -run '^TestName$' ./espresso/devnet-tests/...
```

## Security Considerations

This is blockchain infrastructure code securing real assets. Treat every change accordingly:

- **Derivation pipeline changes** affect what the chain considers canonical state. An incorrect batch accepted by the Caff node or the standard pipeline can cause chain splits or invalid state transitions. Changes to the derivation pipeline require the most scrutiny of anything in this repo, and we should strive to keep our modifications to it at the absolute minimum — prefer adding Espresso logic in isolated paths (like `CaffNextBatch()`) rather than altering the core derivation flow
- **Solidity contracts** are immutable once deployed. `BatchAuthenticator` is upgradeable (transparent proxy), but treat changes with the same rigor — consider reentrancy, access control, and upgrade safety
- **Error handling matters** — silent failures in the batcher or streamer can cause the chain to stall or post incorrect data. Prefer explicit errors over swallowed ones

See `espresso/SECURITY_ANALYSIS.md` for the full security model.

## Key Documentation

- `README_ESPRESSO.md` — Dev environment, devnet guide, enclave setup, OP Succinct dependencies
- `docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md` — Biweekly sync procedure for all forked repos
- `docs/README_ESPRESSO_DEPLOY_CONFIG.md` — Deployment configuration parameters
- `espresso/SECURITY_ANALYSIS.md` — Security analysis (3-layer model, degradation behavior)
- `espresso/docs/metrics.md` — Monitoring metrics for batcher, caff node, verifier, sequencer

## Configuration Flags

All Espresso flags are prefixed with `espresso.` and defined in `espresso/cli.go`. Shared between op-node and op-batcher.

| Flag | Purpose |
|------|---------|
| `espresso.enabled` | Master switch |
| `espresso.urls` | Espresso query service URLs |
| `espresso.light-client-addr` | Light Client contract address on L1 |
| `espresso.namespace` | Espresso namespace (defaults to L2 chain ID) |
| `espresso.origin-height-espresso` | First HotShot block to read from |
| `espresso.origin-height-l2` | L2 height to switch to Espresso derivation |
| `espresso.poll-interval` | HotShot polling interval (default 250ms) |
| `espresso.espresso-attestation-service` | Attestation verifier service URL |
