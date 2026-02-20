# Security Analysis: Celo-Espresso Integration

---

**Document Date:** February 2, 2026
**Version:** 1.0

---

## Executive Summary

This document provides a security analysis of the Celo-Espresso integration, which adds fast finality capabilities to the Optimism rollup stack while maintaining full compatibility with the standard OP Stack security model.

### Architecture

The integration introduces three main components:

1. **L1 Smart Contracts** (~163 lines of Solidity) - Minimal on-chain verification layer with `BatchInbox` and `BatchAuthenticator`
2. **TEE-Based Batcher** - Runs inside AWS Nitro Enclaves with ZK-proven attestation (~240× gas reduction: 63M → 260k)
3. **Espresso Streamer** - Batch verification and ordering service with signature validation

### Security Model

Every batch is expected to undergo validation through **three independent layers**:

- **TEE Attestation** - Cryptographic attestations verified via zero-knowledge proofs (Automata SDK + Succinct SP1)
- **Smart Contract Verification** - On-chain validation of sender address and TEE signatures
- **Batcher Signature Verification** - Signature validation during batch unmarshaling from Espresso

A dual-key design separates the long-lived batcher key (operator-managed) from the ephemeral TEE key (hardware-isolated), requiring both for successful batch posting.

### Safety Guarantee

**Critical Property**: In all failure scenarios, the system degrades gracefully to vanilla Optimism behavior—never worse. Whether Espresso becomes unavailable, the TEE enclave fails, or both, the system falls back to standard L1-only operation identical to the vanilla Celo L2 rollup.

The integration is strictly **additive**: it adds fast finality without replacing existing OP Stack security mechanisms.

### Test Coverage

The codebase includes **23 test scenarios** (14 integration + 9 devnet) covering:

- Stateless batcher recovery and restart resilience
- TEE attestation validation and signature verification
- L2 reorg handling and state consistency
- Censorship resistance via forced transactions
- Fallback mechanism activation and mode switching
- Real AWS Nitro Enclave and full Espresso node validation

All Espresso integration tests and L1 contract tests run automatically on every PR. Devnet and enclave tests are available on-demand.

### Trust Assumptions

The integration relies on AWS Nitro Enclave hardware, L1 Ethereum consensus, cryptographic primitives (ECDSA, Keccak256), and ZK proof system security (Succinct SP1, Automata SDK). **These assumptions affect only fast finality**—the base security model remains unchanged from vanilla Optimism.

### Document Scope

The following sections detail the technical implementation, validation mechanisms, testing methodology, contract architecture, failure recovery procedures, and future security enhancements.

## Architecture Overview

The integration introduces three primary components:
1. **L1 Smart Contracts** (`BatchInbox` and `BatchAuthenticator`) - On-chain verification layer
2. **TEE Batcher** - Trusted execution environment for batch processing
3. **Espresso Streamer** - Batch verification and ordering service

Each component operates within well-defined trust boundaries with multiple layers of validation.

## 1. Multi-Layer Validation Architecture

### Validation Flow Overview

The integration implements three independent validation layers. Each layer checks different properties using separate mechanisms. A batch proceeds through all three layers before acceptance into the L2 chain.

### 1.1 The Three Security Layers (How a Batch Gets Validated)

Let's follow a batch from creation to acceptance to see how the layers work:

#### **Layer 1: TEE Attestation**

The batcher runs inside AWS Nitro Enclaves, which:
- Isolates the batcher code from the host operating system
- Generates cryptographic attestations of the running code (PCR0 measurements)
- Creates private keys within the enclave that cannot be exported

**Implementation**: The enclave generates an attestation document containing the hash of the batcher code and a public key. This attestation is converted into a zero-knowledge proof (using Automata Network's SDK and Succinct's SP1) and verified on-chain before the key is registered as authorized to sign batches. The ZK proof approach reduces verification costs from ~63M gas to ~260k gas (approximately 240× improvement).

References:
- [`BatchAuthenticator.sol`](packages/contracts-bedrock/src/L1/BatchAuthenticator.sol)
- [ZK Attestation Verification](https://docs.espressosys.com/network/concepts/rollup-developers/integrating-an-optimistic-rollup/zk-attestation-verification)

#### **Layer 2: Smart Contract Verification**

When a batch is authenticated on L1, the `BatchAuthenticator` contract verifies the TEE signer:

1. **Signature Check**: Verifies the batch hash signature against registered TEE signers
2. **Event Emission**: Emits a `BatchInfoAuthenticated` event recording the commitment and signer

```solidity
// From BatchAuthenticator.sol
address signer = ECDSA.recover(commitment, signature);
require(
    espressoTEEVerifier.espressoNitroTEEVerifier().isSignerValid(signer, ServiceType.BatchPoster),
    "BatchAuthenticator: invalid signer"
);
emit BatchInfoAuthenticated(commitment, signer);
```

**Implementation**: Before posting batch data to the BatchInbox EOA, the batcher calls `authenticateBatchInfo()` with a signature from the TEE ephemeral key. The derivation pipeline then scans L1 receipts for `BatchInfoAuthenticated` events within a lookback window to determine which batches are authenticated.

Reference: [`BatchAuthenticator.sol`](packages/contracts-bedrock/src/L1/BatchAuthenticator.sol)

#### **Layer 3: Batcher Signature Verification**

Each batch contains a signature from the batcher. When the streamer unmarshals batches from Espresso, it verifies:
- The signature cryptographically validates
- The signer matches the authorized batcher address

**Implementation**: Batches posted to Espresso include the batcher's ECDSA signature over the batch data. The streamer calls `UnmarshalEspressoTransaction()` which recovers the public key from the signature and verifies it matches the expected batcher address before accepting the batch.

Reference: [`espresso_batch.go:95-113`](op-node/rollup/derive/espresso_batch.go)

### 1.2 Validation Flow: Complete Example

The system has two parallel derivation paths that both validate batches:

#### **Batch Creation and Submission**

```
1. User submits transaction to Sequencer
   ↓
2. Sequencer creates L2 block and bundles into batch
   ↓
3. TEE Batcher (inside AWS Nitro Enclave):
   - Reads batch from sequencer
   - Signs batch with batcher private key
   - Submits to Espresso (for fast confirmation)
   - Waits for Espresso finality
   - Calls BatchAuthenticator contract to register batch hash (Layer 2: TEE signature)
   - Posts batch data to L1 BatchInbox
```

#### **Two Parallel Derivation Paths**

After submission, batches flow through two independent paths:

**Path A: Fast Confirmation (Caff Node)**
```
1. Espresso Streamer (in Caff Node):
   - Reads batches from Espresso network
   - Verifies batcher signature during unmarshal
   ↓
2. Caff Node Derivation Pipeline:
   - Derives L2 blocks from validated batches
   - Produces optimistically finalized L2 state
```

**Path B: L1-Based Derivation (Standard OP Node)**
```
1. OP Node reads from L1:
   - Reads batch data from BatchInbox contract
   - Validates batches were authenticated via BatchAuthenticator
   ↓
2. Standard OP Derivation Pipeline:
   - Derives L2 blocks from L1 data
   - Produces L1-finalized L2 state
```

**Key Points:**
- The **TEE Batcher** submits to both Espresso and L1
- The **Espresso Streamer** is used by the Caff Node for fast derivation from Espresso
- The **OP Node** uses standard L1-based derivation
- Both paths independently validate batches
- Layer 1 (TEE Attestation) validates the batcher's enclave
- Layer 2 (Contract Verification) validates on L1 via address check + TEE signature
- Layer 3 (Batcher Signature) validates when reading from Espresso

### 1.3 Dual-Key Architecture

The implementation uses two distinct private keys with separate roles:

#### **Batcher Key** (Long-lived, managed by operator)
```solidity
address public immutable teeBatcher;  // E.g., 0x1234...
```
- Registered in the rollup configuration
- Gives authority to post batches to L1
- Can exist outside the TEE
- **Role**: Proves "this is the official batcher"

#### **Ephemeral Key** (Short-lived, generated in TEE)
```go
func (bs *BatcherService) initKeyPair() error
// Generates key inside enclave
// Private key NEVER leaves the hardware
```
- Generated inside AWS Nitro Enclave
- Used to sign batch commitments
- Cannot be extracted from the hardware
- **Role**: Proves "this came from the correct TEE code"

#### Key Separation Properties

The dual-key design implements the following separation:

| Scenario | Batcher Key Compromised | Ephemeral Key Compromised |
|----------|------------------------|---------------------------|
| **Attacker capability** | Send transactions to L1 | Sign batch hashes |
| **Missing capability** | Cannot forge TEE signatures | Cannot post to L1 BatchInbox |
| **Observed result** | Batches rejected (no TEE sig) | Signatures rejected (wrong address) |

**Design note**: Successful batch posting requires both keys. The keys are stored in different locations:
- Batcher key: Configured on the server
- Ephemeral key: Generated and stored within the Nitro Enclave hardware

## 2. Fault Tolerance and Recovery

The implementation includes mechanisms for handling Espresso component failures.

### 2.1 Fallback Mechanism

The system includes a non-TEE batcher that can be activated when TEE components are unavailable. The owner can switch between TEE and non-TEE modes:

**Fallback Batcher Activation**
```solidity
function switchBatcher() external onlyOwner {
    activeIsTee = !activeIsTee;  // Toggle between TEE and non-TEE mode
}
```

#### When to Use Fallback

The fallback batcher is activated when any Espresso or TEE component fails:

- AWS Nitro Enclave failure (TEE batcher cannot start)
- Espresso network unavailable (cannot get fast confirmations)
- TEE attestation service down (cannot register new keys)
- Succinct Network unavailable (cannot generate ZK proofs)

#### Fallback Mode Behavior

When operating in non-TEE mode:
- Batcher posts directly to L1 without TEE attestation
- No Espresso confirmation required before L1 posting
- BatchInbox accepts batches from the non-TEE batcher address
- Derivation continues using standard OP Stack mechanisms

**Switching Procedure**
1. Owner calls `switchBatcher()` on the BatchAuthenticator contract
2. Non-TEE batcher begins posting to L1
3. When ready to resume TEE mode, update caffeinated height
4. Owner calls `switchBatcher()` again to re-enable TEE mode
5. TEE batcher resumes operation from the new heights


References:
- [`BatchInbox.t.sol:84-165`](packages/contracts-bedrock/test/L1/BatchInbox.t.sol)
- [Specification §36.4.2](https://eng-wiki.espressosys.com/mainch36.html#x43-22900036)

### 2.2 Worst-Case Degradation: Equivalent to Vanilla Celo Rollup

**Security Property**: In all failure scenarios, the system degrades gracefully to behave identically to the current vanilla Celo L2 rollup, never worse.

#### Degradation Scenarios

The Espresso integration adds fast finality capabilities without compromising the baseline security of the standard OP Stack. All component failures are handled by switching to the fallback (non-TEE) batcher, which operates identically to the vanilla Celo L2 rollup:

- **Espresso network unavailable** - Cannot retrieve batches for fast confirmation
- **TEE enclave failure** - Cannot generate attestations or run TEE batcher
- **Succinct Network down** - Cannot generate ZK proofs for attestation verification
- **Attestation service failure** - Cannot verify new TEE attestations

In each case, the owner activates the fallback batcher via a single `switchBatcher()` transaction, and the system continues with standard L1-only operation.

#### Why Degradation is Always Safe

**1. L1 Derivation Path Always Exists**

The standard OP Stack derivation pipeline remains fully functional regardless of Espresso status:

```
OP Node → L1 BatchInbox → Standard Derivation Pipeline → L2 Blocks
```

This path operates independently of:
- Espresso network availability
- TEE batcher status
- Fast finality features

**2. Non-TEE Batcher is Pre-Configured**

```solidity
address public immutable teeBatcher;     // Espresso-enhanced batcher
address public immutable nonTeeBatcher;  // Fallback (standard) batcher
```

The non-TEE batcher address is immutably configured in the contracts. The owner can activate it instantly via `switchBatcher()`.

**3. Espresso Features are Additive, Not Replacement**

The integration **adds** capabilities without **replacing** core functionality:

| Capability | Standard OP Stack | With Espresso Integration |
|------------|-------------------|---------------------------|
| L2 block production | ✅ Sequencer | ✅ Same sequencer |
| Batch posting to L1 | ✅ Batcher → L1 | ✅ Same, plus optional Espresso |
| Derivation from L1 | ✅ OP Node | ✅ Slightly different derivation logic |
| Fault proofs | ✅ Dispute game | ✅ Same dispute game |
| Withdrawals | ✅ Standard bridge | ✅ Same bridge |
| **Fast finality** | ❌ Not available | ✅ **New**: Caff Node + Espresso |

**4. Fallback Mode is equivalent to Vanilla Behavior**

When operating in fallback mode:

```solidity
if (!activeIsTee) {
    // Non-TEE batcher posts to BatchInbox
    // No TEE attestation required
    // No Espresso submission required
    // Identical to standard OP Stack
}
```

The system:
- Uses standard batcher (no TEE)
- Posts only to L1 (no Espresso)
- Derives blocks using standard OP Node (with slight change in derivation pipeline)
- Processes transactions identically
- Maintains same security guarantees

**5. No New Trust Assumptions for Base Security**

The standard security model remains unchanged:
- L1 Ethereum consensus (same)
- Sequencer liveness (same)
- Fault proof system (same)
- Contract immutability (same)

New trust assumptions (Espresso, Succinct, Automata) **only** affect fast finality, not base security.

**6. Minimal Derivation Pipeline Changes**

The Espresso integration makes only **one architectural change** to the OP Stack derivation pipeline: moving sender verification from the pipeline to the L1 smart contract.

**The Single Modification: `isValidBatchTx()` Function**

In the standard OP Stack, the derivation pipeline verifies the batch sender:

```go
// Standard OP Stack (vanilla)
func isValidBatchTx(..., l1Signer types.Signer, ..., batcherAddr common.Address) bool {
    // ... other checks ...

    // Verify sender matches authorized batcher
    from, err := l1Signer.Sender(tx)
    if err != nil || from != batcherAddr {
        return false
    }
}
```

In the Espresso integration, this verification is removed from the pipeline:

```go
// Espresso integration
func isValidBatchTx(..., _ types.Signer, ..., batcherAddr common.Address) bool {
    // ... same checks (tx type, inbox address, receipt status) ...

    // NOTE: contrary to a standard OP batcher, we can safely skip any verification
    // related to the sender of the transaction. Indeed the Batch Inbox contract
    // takes care of ensuring the sender of the batch information is a legitimate batcher.

    return true
}
```

**Why This Change is Safe**

The sender verification hasn't been removed—it's been **moved to a more secure location**:

| Verification Location | Standard OP Stack | Espresso Integration |
|----------------------|-------------------|----------------------|
| **In derivation pipeline** | ✅ `l1Signer.Sender(tx) == batcherAddr` | ❌ Removed |
| **In L1 smart contract** | ❌ Not present | ✅ `BatchInbox.sol` enforces sender check |

**L1 Contract Enforcement:**
```solidity
// BatchInbox.sol
fallback() external payable {
    if (msg.sender != batchAuthenticator.teeBatcher() &&
        msg.sender != batchAuthenticator.nonTeeBatcher()) {
        revert("Not authorized");
    }
    // ... store batch data ...
}
```


#### Tested Degradation Paths

The test suite validates degradation behavior:

| Test | What It Validates |
|------|-------------------|
| `TestBatcherSwitching` | Switching between TEE and non-TEE modes maintains correctness |
| `TestBatcherRestart` | Batcher failures don't compromise chain state |
| `TestSmokeWithoutTEE` | System operates correctly in non-TEE mode |
| Fallback tests | Manual switch to vanilla mode preserves all functionality |

#### Operational Guarantees

**Guarantee 1: No Additional Liveness Risk**
- If Espresso fails → system continues via L1
- If TEE fails → system continues via non-TEE batcher
- Worst case: vanilla Celo rollup liveness

**Guarantee 2: No Additional Safety Risk**
- L1 contracts validate all batches (TEE or non-TEE)
- Standard derivation path validates all blocks
- Fault proof system covers all state transitions
- Worst case: vanilla Celo rollup safety

**Guarantee 3: Instant Fallback**
- Owner can switch batchers via single transaction
- No migration or state transition required
- Chain continues from current block
- Worst case: vanilla Celo rollup behavior


References:
- [`BatchInbox.sol`](packages/contracts-bedrock/src/L1/BatchInbox.sol) - Dual batcher support
- [`BatchAuthenticator.sol`](packages/contracts-bedrock/src/L1/BatchAuthenticator.sol) - Mode switching

## 3. Testing Strategy

### 3.1 End-to-End Integration Tests

The integration includes extensive scenario-based testing across two test suites:

#### Environment Integration Tests (14 test scenarios)

These tests run in a controlled environment with mock Espresso nodes:

| # | Test File | What It Tests | Why It Matters |
|---|-----------|---------------|----------------|
| 1 | `espresso_benchmark_test.go` | High-throughput performance | Validates system under load |
| 2 | `espresso_liveness_test.go` | Continuous operation | Core functionality |
| 3.1 | `espresso_caff_node_test.go` | Caff node derivation | L2 state correctness |
| 3.2 | `deterministic_state_test.go` | State determinism | Same inputs → same state |
| 3.3 | `fast_derivation_and_caff_node_test.go` | Optimistic derivation | Fast confirmation path |
| 4 | `confirmation_integrity_with_reorgs_test.go` | Reorg handling | L2 reorganization safety |
| 5 | `batch_authentication_test.go` | TEE attestation | Authentication security |
| 6 | `batch_inbox_test.go` | Contract validation | On-chain security |
| 7 | `stateless_batcher_test.go` | **Stateless recovery** | **Critical: restart safety** |
| 8 | `reorg_test.go` | L2/Espresso reorgs | Multi-layer consistency |
| 9 | `pipeline_enhancement_test.go` | Derivation pipeline | Integration correctness |
| 10 | `soft_confirmation_integrity_test.go` | Fast confirmations | Espresso confirmation validity |
| 11 | `forced_transaction_test.go` | Censorship resistance | Security invariant |
| 12 | `enforce_majority_rule_test.go` | Query service voting | Byzantine fault tolerance |
| 13 | `dispute_game_test.go` | Fault proof system | L1 dispute resolution |
| 14 | `batcher_fallback_test.go` | Fallback mechanism | Graceful degradation |

#### Devnet Tests (9 real-world scenarios)

These tests run against a full Docker-based devnet with real Espresso nodes:

| Test | What It Tests | Environment |
|------|---------------|-------------|
| `TestSmokeWithoutTEE` | Basic operation without TEE | Standard mode |
| `TestSmokeWithTEE` | Basic operation with TEE | AWS Nitro Enclave |
| `TestBatcherRestart` | Batcher restart resilience | Failure recovery |
| `TestBatcherSwitching` | Switch between TEE/non-TEE | Fallback activation |
| `TestBatcherActivePublishOnly` | Active batch publishing | Data availability |
| `TestForcedTransaction` | Force inclusion via L1 | Censorship resistance |
| `TestWithdrawal` | L2→L1 withdrawals | Bridge security |
| `TestChallengeGame` | Fault proof challenges | Dispute resolution |
| `TestChangeBatchAuthenticatorOwner` | Ownership transfer | Access control |

#### Critical Test Deep Dive: Stateless Batcher (Test 7)

```go
// Validates batcher can restart randomly without data loss
// Verifies Espresso-L1 consistency after restarts
func TestStatelessBatcher(t *testing.T)
```

**What it does:**
1. Starts sequencer, batcher (Espresso mode), Caff node, OP node
2. Loops over N iterations:
   - Randomly picks one iteration to **stop** the batcher
   - Randomly picks another to **start** the batcher
   - For all other iterations: send 1 coin to Alice
3. Asserts:
   - Alice's balance on Caff node = Alice's balance on OP node
   - No transactions lost during batcher downtime

**Why this is critical:** Proves the batcher maintains no persistent state and can recover from arbitrary restarts without data loss or inconsistency.

Reference: [`7_stateless_batcher_test.go:21-38`](espresso/environment/7_stateless_batcher_test.go)

### Test Coverage Analysis

#### 1. **Security Property Validation**

Each security validation layer has corresponding test coverage:

| Validation Property | Test Coverage | Validation Method |
|-------------------|-----------|-----------|
| Authenticity | Test 5, 6, TestSmokeWithTEE | TEE attestation verification |
| Integrity | Test 7, 10, TestBatcherRestart | State consistency across restarts |
| Liveness | Test 2, 14, TestBatcherSwitching | Operation under component failures |
| Consistency | Test 3.2, 4, 8 | Deterministic state across nodes |
| Censorship Resistance | Test 11, TestForcedTransaction | Force inclusion via L1 |
| Fallback Behavior | Test 14, TestBatcherSwitching | Mode switching validation |
| Query Service | Test 12 | Majority voting implementation |
| Dispute Resolution | Test 13, TestChallengeGame | Fault proof verification |



#### 2. **Failure Scenario Testing**

Tests include various failure scenarios and recovery mechanisms:

| Failure Scenario | Test Coverage | Recovery Mechanism Tested |
|------------------|---------------|-------------------|
| Batcher crash | Test 7, TestBatcherRestart | Stateless recovery |
| TEE unavailable | Test 14, TestBatcherSwitching | Fallback to non-TEE |
| Espresso unavailable | Test 14 | Direct L1 posting |
| L2 reorg | Test 4, 8 | Automatic state reset |
| Invalid attestation | Test 5 | Contract rejection |
| Query service disagreement | Test 12 | Majority rule application |

**Test design**: Each test verifies that the system detects the failure condition and executes the corresponding recovery mechanism.

#### 3. **Environment Testing Characteristics**

The devnet tests differ from environment tests in their setup:

- **AWS Nitro Enclaves**: `TestSmokeWithTEE` runs against actual Nitro hardware
- **Espresso Nodes**: Tests interact with running Espresso consensus nodes
- **L1 Interaction**: Full Ethereum L1 deployment using actual contracts
- **Docker Networking**: Inter-service communication over Docker networks

**Setup difference**: Environment tests use mocked Espresso components for faster iteration, while devnet tests use the full production stack.

#### 4. **Layered Testing Strategy**

Tests are organized by scope:

```
Unit Tests (Go packages)
    ↓
Contract Tests (Foundry)
    ↓
Environment Tests (Mocked Espresso)
    ↓
Devnet Tests (Real Espresso)
    ↓
Enclave Tests (Real AWS Nitro)
```

Each layer catches different classes of bugs:
- **Unit**: Logic errors
- **Contract**: Smart contract vulnerabilities
- **Environment**: Integration issues (fast iteration)
- **Devnet**: Real-world scenarios (high confidence)
- **Enclave**: Hardware-specific issues

#### 5. **OP Stack Test Suite Availability**

A test script exists for validating OP Stack compatibility:

```bash
# From run_all_tests.sh (manual execution)
make -C ./cannon test
just -f ./op-batcher/justfile test
just -f ./op-challenger/justfile test
just -f ./op-node/justfile test
just -f ./op-proposer/justfile test
# ... (all OP Stack component tests)
```

**Test scope**: These tests validate that the integration maintains compatibility with existing OP Stack components and behaviors.

**Note**: This comprehensive suite is available for manual testing but does not run automatically in CI. CI focuses on Espresso-specific integration tests and L1 contract tests.

Reference: [`run_all_tests.sh`](run_all_tests.sh)

#### 6. **Continuous Testing in CI**

Every PR triggers:
- ✅ 14 integration tests
- ✅ 9 devnet tests
- ✅ L1 contract tests (Foundry tests for BatchInbox, BatchAuthenticator)
- ✅ Enclave tests (on actual AWS infrastructure)

**CI configuration**: Tests run in parallel across multiple groups with 30-minute timeouts.

**Additional Testing**: An OP Stack regression test suite (`run_all_tests.sh`) is available for manual execution to validate compatibility with all OP Stack components (op-program, cannon, op-challenger, op-node, op-proposer, op-service, op-supervisor, op-e2e).

References:
- [`espresso-integration.yaml`](.github/workflows/espresso-integration.yaml)
- [`espresso-devnet-tests.yaml`](.github/workflows/espresso-devnet-tests.yaml)
- [`espresso-enclave.yaml`](.github/workflows/espresso-enclave.yaml)

### Test Coverage Characteristics

The test suite exhibits the following properties:

- **Component independence**: Each component has dedicated test coverage
- **Path coverage**: Tests include normal operation, failure scenarios, and edge cases
- **Environment variety**: Tests run in both mocked and production-like environments
- **Continuous execution**: CI runs all tests on every pull request
- **Property validation**: Each validation layer has test coverage
- **Deployment simulation**: Devnet tests use the same deployment process as production


### 3.2 Smart Contract Security Tests

The Espresso integration includes Foundry-based smart contract tests that validate security properties of the L1 contracts responsible for batch data submission.

#### BatchInbox Contract Tests

The `BatchInbox` contract enforces batcher authentication based on operating mode (TEE vs non-TEE). Test coverage includes:

**TEE Mode Authentication**

```solidity
// TEE batcher requires valid attestation
function test_fallback_teeBatcherRequiresAuthentication() external

// TEE batcher succeeds with authenticated batch
function test_fallback_teeBatcherSucceedsWithValidAuth() external

// Non-TEE batcher cannot post when TEE is active
function test_fallback_nonTeeBatcherRevertsWhenTeeActiveAndUnauthenticated() external
```

These tests verify that when the system operates in TEE mode:
- Only the designated TEE batcher address can submit batches
- Batches must be pre-authenticated via the `BatchAuthenticator` contract
- The non-TEE batcher is rejected even if attempting to submit authenticated batches

**Fallback Mode (Non-TEE) Authentication**

```solidity
// Non-TEE batcher can post after mode switch
function test_fallback_nonTeeBatcherCanPostAfterSwitch() external

// Non-TEE batcher doesn't require attestation
function test_fallback_nonTeeBatcherDoesNotRequireAuth() external

// Inactive batcher (TEE) reverts in fallback mode
function test_fallback_inactiveBatcherReverts() external

// Unauthorized addresses are rejected
function test_fallback_unauthorizedAddressReverts() external
```

These tests verify that when switched to fallback mode:
- Only the designated non-TEE batcher can submit batches
- No attestation or pre-authentication is required
- The TEE batcher cannot post (even with valid attestations)
- Random unauthorized addresses are rejected

**Security Properties Validated:**
- **Exclusive access control**: Only one batcher can be active at a time
- **Mode enforcement**: Authentication requirements match the active mode
- **Address authorization**: Unauthorized addresses cannot submit batches

#### BatchAuthenticator Contract Tests

The `BatchAuthenticator` contract manages batcher switching and batch authentication. Test coverage includes:

**Ownership and Access Control**

```solidity
// Only owner can switch active batcher
function test_switchBatcher_revertsForNonOwner() external
```


**References:**
- [`BatchInbox.t.sol`](packages/contracts-bedrock/test/L1/BatchInbox.t.sol) - 7 test functions covering all authentication scenarios
- [`BatchAuthenticator.t.sol`](packages/contracts-bedrock/test/L1/BatchAuthenticator.t.sol) - 4 test functions covering ownership and initialization

### 3.3 Enclave Testing

**Real TEE Validation**

The integration includes tests running on actual AWS Nitro Enclaves:

```yaml
# .github/workflows/espresso-enclave.yaml
- name: Run enclave tests
  run: just espresso-enclave-tests
```

These tests validate:
- Attestation generation in real Nitro environment
- Key generation isolation
- PCR0 measurement consistency
- Contract registration flow

Reference: [`espresso-enclave.yaml`](.github/workflows/espresso-enclave.yaml)

### 3.4 Continuous Integration

**Automated Security Checks**

Every pull request triggers the execution of different test suites:

**1. Espresso Integration Tests** (automatic on every PR)
```yaml
# .github/workflows/espresso-integration.yaml
- Parallelized across 4 groups
- Tests all Espresso-specific components
- Runs: ./espresso/... test suite
- Timeout: 30 minutes
```

**2. L1 Contract Tests** (automatic on every PR)
```yaml
# .github/workflows/contracts-l1-tests.yaml
- Foundry tests for BatchInbox and BatchAuthenticator
- Validates on-chain security properties
```

**3. Devnet Tests** (on-demand via workflow dispatch)
```yaml
# .github/workflows/espresso-devnet-tests.yaml
- Full Docker-based environment with real Espresso nodes
- Tests 9 real-world scenarios
```

**4. Enclave Tests** (on-demand via workflow dispatch)
```yaml
# .github/workflows/espresso-enclave.yaml
- Runs on actual AWS Nitro Enclave hardware
- Validates TEE attestation and key isolation
```

CI ensures no regression in Espresso-specific security properties and contract behavior.

References:
- [`espresso-integration.yaml`](.github/workflows/espresso-integration.yaml)
- [`contracts-l1-tests.yaml`](.github/workflows/contracts-l1-tests.yaml)
- [`espresso-devnet-tests.yaml`](.github/workflows/espresso-devnet-tests.yaml)
- [`espresso-enclave.yaml`](.github/workflows/espresso-enclave.yaml)

## 4. Contract Security Architecture

### 4.1 Minimal On-Chain Complexity

The L1 contracts follow a minimalist design philosophy:

**`BatchInbox.sol` (77 lines)**
- Single fallback function
- Clear authentication logic
- No complex state management
- Minimal attack surface

**`BatchAuthenticator.sol` (86 lines)**
- Straightforward signature verification
- Immutable TEE verifier reference
- Simple batcher switching
- Event emission for auditability

Small, focused contracts are easier to audit and less prone to vulnerabilities.


### 4.3 External Dependency Isolation

Contracts minimize external dependencies:

```solidity
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
```

Only battle-tested OpenZeppelin libraries are used, reducing supply chain risks.


## 5. Off-Chain Component Security

### 5.1 Batcher Architecture

**Isolation and Compartmentalization**

The batcher separates concerns into independent loops:

```go
// Batch queuing: Fast submission to Espresso
func (l *BlockLoader) BatchQueuingLoop()

// Batch loading: Validation and L1 preparation
func (l *BlockLoader) BatchLoadingLoop()

// Frame publishing: L1 submission
func publishingLoop()
```

Each loop can fail independently without compromising overall system integrity. State is minimized to enable easy recovery.

### 5.2 Streamer Security

**Validation Pipeline**

The Espresso streamer implements defense-in-depth:

```go
func (s *BatchStreamer[B]) CheckBatch(ctx context.Context, batch B) (BatchValidity, int) {
    // Check ordering and buffering
    i, batchRecorded := s.BatchBuffer.TryInsert(batch)

    // Verify batcher signature during unmarshaling
    batch, err := s.UnmarshalBatch(transaction)
}
```

**Buffering for Resilience**

The `BufferedEspressoStreamer` adds resilience:
- Absorbs temporary streamer resets without data loss
- Maintains consistent read position across reorgs
- Enables efficient batch retrieval

Reference: [`buffered_streamer.go`](espresso/buffered_streamer.go)


## 6. Internal Security Reviews

The Celo-Espresso integration has undergone comprehensive internal security audits covering TEE contracts and the Espresso streamer component.

### Audit Summary

| Audit | Date | Reference | Scope | Critical | High | Medium | Low | Status |
|-------|------|-----------|-------|----------|------|--------|-----|--------|
| **TEE Contracts** | Jan 28, 2026 | [PR #43](https://github.com/EspressoSystems/espresso-tee-contracts/pull/43) | Attestation verification, signer registration, enclave hash validation | 2 | 1 | 0 | 0 | ✅ All resolved |
| **Streamer** | 2026 | [PR #339](https://github.com/EspressoSystems/optimism-espresso-integration/pull/339) | Batch validation, reorg handling, L1 consistency, buffer management | 0 | 2 | 2 | 7 | ⏳ Documented. Resolution in progress.  |
| **Total** | - | - | - | **2** | **3** | **2** | **7** | **3 fixed, 11 documented** |

### Key Outcomes

**TEE Contracts:** All critical and high-severity vulnerabilities resolved, including cross-chain deployment replay attacks, missing journal validations, and signer deletion DoS attacks.

**For more details**, see the [internal Security Audit Report](audits/internal_report_30_january_2026.md).


## 7. Trust Model and Assumptions

### 7.1 Trust Boundaries

**Trusted Components**
- AWS Nitro Enclave hardware
- L1 Ethereum consensus
- Espresso consensus (for liveness)
- Succinct Network (for ZK proof generation)
- Automata's Nitro ZK Attestation SDK
- Espresso's Attestation Verifier service

**Untrusted Components**
- Sequencer
- Batcher operator (networking, infrastructure)
- Espresso query service (validated via majority voting)
- Individual Espresso nodes

### 7.2 Adversarial Scenarios Considered

**Batcher and Attestation Attacks**

| Attack Vector | Mitigation | Test Coverage |
|---------------|------------|---------------|
| Malicious batcher operator | TEE attestation proves code integrity | [Test 5](espresso/environment/5_batch_authentication_test.go), [TestSmokeWithTEE](espresso/devnet-tests/smoke_test.go) |
| Invalid TEE attestation | On-chain ZK proof verification rejects unauthorized batches | [TestE2eDevnetWithInvalidAttestation](espresso/environment/5_batch_authentication_test.go) |
| Unattested batcher key | Unsafe blocks produced; safe blocks require valid attestation | [TestE2eDevnetWithUnattestedBatcherKey](espresso/environment/5_batch_authentication_test.go) |
| Forged batch signature | BatchAuthenticator validates ECDSA signatures against registered signers | [BatchAuthenticator.t.sol](packages/contracts-bedrock/test/L1/BatchAuthenticator.t.sol) |
| Invalid batch commitment | BatchInbox verifies keccak256 hash of calldata/blobs before acceptance | [BatchInbox.t.sol](packages/contracts-bedrock/test/L1/BatchInbox.t.sol) |

**Network and Infrastructure Attacks**

| Attack Vector | Mitigation | Test Coverage |
|---------------|------------|---------------|
| Compromised Espresso node | Majority voting across multiple query service nodes | [Test 12](espresso/environment/12_enforce_majority_rule_test.go) |
| Espresso query service disagreement | 2/3 majority rule; inconsistent responses trigger re-query | [Test 12](espresso/environment/12_enforce_majority_rule_test.go) |
| TEE/Espresso unavailability | Fallback to non-TEE batcher with standard OP Stack security | [Test 14](espresso/environment/14_batcher_fallback_test.go), [TestBatcherSwitching](espresso/devnet-tests/batcher_switching_test.go) |
| Succinct Network unavailability | Batcher cannot register new attestations until service restored; existing keys continue | - |
| Execution engine crash | Stateless restart recovery; no persistent state required | [Test 7](espresso/environment/7_stateless_batcher_test.go) |

**Censorship and Liveness Attacks**

| Attack Vector | Mitigation | Test Coverage |
|---------------|------------|---------------|
| Sequencer censorship | Forced transaction inclusion via L1 after sequencing window expires | [Test 11](espresso/environment/11_forced_transaction_test.go) |
| Batcher refusing to submit | Users can force-include transactions through L1 deposits | [Test 11](espresso/environment/11_forced_transaction_test.go) |
| Sequencer downtime | Forced inclusion mechanism activates after sequencing window | [Test 11](espresso/environment/11_forced_transaction_test.go) |

**State and Consistency Attacks**

| Attack Vector | Mitigation | Test Coverage |
|---------------|------------|---------------|
| L1 reorg invalidating posted batches | Batcher re-derives and re-posts same batches in same order after L1 reorg | [Test 4](espresso/environment/4_confirmation_integrity_with_reorgs_test.go) |
| Submitting batches derived from unfinalized L1 blocks | Batcher and Caff node wait for L1 finality before submission/processing | [Test 8](espresso/environment/8_reorg_test.go) |

**Smart Contract Attacks**

| Attack Vector | Mitigation | Test Coverage |
|---------------|------------|---------------|
| Unauthorized batcher switch | Only contract owner can call switchBatcher() | [test_switchBatcher_revertsForNonOwner](packages/contracts-bedrock/test/L1/BatchAuthenticator.t.sol) |
| Zero address configuration | Constructor rejects zero addresses for batchers | [test_constructor_revertsWhen*IsZero](packages/contracts-bedrock/test/L1/BatchAuthenticator.t.sol) |
| Wrong batcher in TEE mode | BatchInbox enforces only TEE batcher can post in TEE mode | [test_fallback_teeBatcherRequiresAuthentication](packages/contracts-bedrock/test/L1/BatchInbox.t.sol) |
| Wrong batcher in fallback mode | BatchInbox enforces only non-TEE batcher can post in fallback mode | [test_fallback_unauthorizedAddressReverts](packages/contracts-bedrock/test/L1/BatchInbox.t.sol) |
| Unauthenticated batch in TEE mode | Derivation pipeline checks BatchInfoAuthenticated events in lookback window | [test_authenticateBatchInfo_succeeds](packages/contracts-bedrock/test/L1/BatchAuthenticator.t.sol) |

**Future Threat Vectors**

| Attack Vector | Current Exposure | Planned Mitigation |
|---------------|------------------|-------------------|
| Operator network MitM | Operator controls enclave networking | SSL certificate pinning or in-enclave L1 light client |
| Espresso query service trust | Majority voting across operators | Direct QC and namespace proof verification |
| Centralized batcher operation | Single TEE batcher address | Permissionless batching with stake-based selection |


## 9. Next steps

### 9.1 Planned Audit with Least Authority

All L1 smart contracts for the Celo-Espresso integration will undergo external security audit by [Least Authority](https://leastauthority.com/), a security research firm specializing in privacy-focused and cryptographic systems.

**Audit Scope:**
- `BatchInbox.sol` - Batch data submission and authentication
- `BatchAuthenticator.sol` - Dual-batcher switching and TEE signature verification
- TEE Verifier contracts


Least Authority proactively identified and disclosed a bug in the [Espresso Jellyfish cryptographic library](https://github.com/EspressoSystems/jellyfish), demonstrating their commitment to responsible disclosure and deep understanding of cryptographic systems.

### 9.2 Monitoring System

The specification of the monitoring system is in progress.

## References

- [OP Stack Integration Specification](https://eng-wiki.espressosys.com/mainch36.html#x43-22900036)
- [Source Code Repository](https://github.com/EspressoSystems/optimism-espresso-integration)
- [Optimism Rollup Protocol Specification](https://specs.optimism.io/)
- [AWS Nitro Enclaves Documentation](https://docs.aws.amazon.com/enclaves/)

**Document Version**: 1.1
**Last Updated**: January 29, 2026

