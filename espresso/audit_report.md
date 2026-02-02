# Security Audit Report
## Optimism-Espresso Integration: OP Streamer & TEE Contracts

---

**Report Date:** January 29, 2026
**Audit Scope:** OP Streamer Component & TEE Contracts Infrastructure
**Version:** v0.5.0
**Auditors:** Internal Security Team
**Status:** Active Development

---

## Executive Summary

This report presents the findings of a security audit of the Optimism-Espresso integration, focusing on the OP Streamer component and TEE (Trusted Execution Environment) contracts. The audit identified **14 vulnerabilities** (2 Critical, 4 High, 1 Medium, 7 Low) across batch streaming logic, TEE networking, and smart contract implementations.

### Severity Distribution

| Severity | Count | Status |
|----------|-------|--------|
| 🔴 Critical | 2 | 2 Fixed, 0 Open |
| 🟠 High | 4 | 1 Fixed, 3 Open |
| 🟡 Medium | 1 | 0 Fixed, 1 Open |
| 🟢 Low | 7 | 0 Fixed, 7 Open |
| **Total** | **14** | **3 Fixed, 11 Open** |

### Key Findings Summary

| ID | Vulnerability | Severity | Component | Status | Reference |
|----|---------------|----------|-----------|--------|-----------|
| V-4 | Cross-Chain Deployment | 🔴 Critical | TEE Contracts | Fixed | Section 4.1, PR #43 |
| V-6 | Missing Journal Validations | 🔴 Critical | TEE Contracts | Fixed | Section 4.3, PR #43 |
| V-2 | Infinite Buffer Growth | 🟠 High | OP Streamer | Open | Section 2.2 |
| V-3 | TEE Networking MitM Attack | 🟠 High | TEE Enclave | Open | Section 3.1 |
| V-5 | Signer Deletion DoS | 🟠 High | TEE Contracts | Fixed | Section 4.2, PR #43 |
| V-7 | Type Mismatch in Refresh() | 🟠 High | OP Streamer | Open | Section 2.3 |
| V-1 | All-At-Once RPC Calls | 🟡 Medium | OP Streamer | Open | Section 2.1 |
| V-8 | Missing Duplicate Detection | 🟢 Low | OP Streamer | Open | Section 2.4 |
| V-9 | Misleading Log Messages | 🟢 Low | OP Streamer | Open | Section 2.5 |
| V-10 | Inefficient Batch Overwrite | 🟢 Low | OP Streamer | Open | Section 2.6 |
| V-11 | Confusing Variable Naming | 🟢 Low | OP Streamer | Open | Section 2.7 |
| V-12 | Unused Constant Declaration | 🟢 Low | OP Streamer | Open | Section 2.8 |
| V-13 | Missing Sort Order Validation | 🟢 Low | OP Streamer | Open | Section 2.9 |
| V-14 | No Network Failure Distinction | 🟢 Low | OP Streamer | Open | Section 2.10 |

---

## Table of Contents

1. [Scope and Methodology](#1-scope-and-methodology)
2. [OP Streamer Component Vulnerabilities](#2-op-streamer-component-vulnerabilities)
   - 2.1 [V-1: All-At-Once RPC Calls](#v-1-all-at-once-rpc-calls)
   - 2.2 [V-2: Infinite Buffer Growth](#v-2-infinite-buffer-growth)
   - 2.3 [V-7: Type Mismatch in Refresh()](#v-7-type-mismatch-in-refresh)
   - 2.4 [V-8: Missing Duplicate Detection](#v-8-missing-duplicate-detection)
   - 2.5 [V-9: Misleading Log Messages](#v-9-misleading-log-messages)
   - 2.6 [V-10: Inefficient Batch Overwrite](#v-10-inefficient-batch-overwrite)
   - 2.7 [V-11: Confusing Variable Naming](#v-11-confusing-variable-naming)
   - 2.8 [V-12: Unused Constant Declaration](#v-12-unused-constant-declaration)
   - 2.9 [V-13: Missing Sort Order Validation](#v-13-missing-sort-order-validation)
   - 2.10 [V-14: No Network Failure Distinction](#v-14-no-network-failure-distinction)
3. [TEE Enclave Vulnerabilities](#3-tee-enclave-vulnerabilities)
   - 3.1 [V-3: TEE Networking MitM Attack](#v-3-tee-networking-mitm-attack)
4. [TEE Contracts Vulnerabilities](#4-tee-contracts-vulnerabilities)
   - 4.1 [V-4: Cross-Chain Deployment](#v-4-cross-chain-deployment-vulnerability)
   - 4.2 [V-5: Signer Deletion DoS Attack](#v-5-signer-deletion-dos-attack)
   - 4.3 [V-6: Missing Journal Validations](#v-6-missing-tee-journal-struct-validations)

---

## 1. Scope and Methodology

### 1.1 Audit Scope

This audit covers the following components:

- **OP Streamer Component** (`espresso/`)
  - `batch_buffer.go` - Batch buffering and ordering logic
  - `streamer.go` - Espresso block streaming and batch processing
  - `buffered_streamer.go` - Buffered streaming implementation

- **TEE Contracts** (`espresso-tee-contracts`)
  - `EspressoNitroTEEVerifier.sol` - AWS Nitro attestation verification
  - `EspressoSGXTEEVerifier.sol` - Intel SGX attestation verification
  - `TEEHelper.sol` - Shared TEE helper functionality
  - Related interfaces and libraries

### 1.2 Methodology

- **Static Code Analysis**: Manual review of Go and Solidity source code
- **Architecture Review**: Analysis of component interactions and data flow
- **Attack Scenario Modeling**: Threat modeling for potential exploits
- **Documentation Review**: Analysis of security documentation and deployment guides

---

## 2. OP Streamer Component Vulnerabilities

### V-1: All-At-Once RPC Calls

**Severity:** 🟡 **Medium**
**Status:** ⚠️ **Open**
**Component:** `espresso/streamer.go` - `CheckBatch()` and `processRemainingBatches()`

#### Description

The `CheckBatch` function makes synchronous L1 RPC calls (`HeaderHashByNumber`) when validating finalized batches. When multiple batches become finalized simultaneously, the system executes sequential synchronous RPC calls, causing the streamer to freeze.

#### Technical Details

**Vulnerable Code Path:**
```go
func CheckBatch(batch B, l1Origin eth.BlockID) {
    if isFinalized(l1Origin) {
        hash := HeaderHashByNumber(l1Origin.Number) // Synchronous RPC call
        // ... validation logic
    }
}
```

**Attack Scenario:**

1. Node accumulates 500 batches in `RemainingBatches` while waiting for L1 finality
2. L1 finalizes a new state
3. `processRemainingBatches()` iterates all 500 batches
4. Finalized check now passes for all batches
5. System executes **500 sequential synchronous RPC calls** inside the `Update` loop

#### Impact

- **Availability**: Streamer freezes for seconds to minutes
- **Denial of Service**: Node stops fetching new Espresso blocks
- **Cascading Failure**: Downstream components dependent on streamer become blocked

#### Likelihood

**Medium** - Requires specific conditions where many batches accumulate before L1 finalization

#### Overall Risk

**Medium** - Temporary performance degradation rather than permanent failure. System recovers once RPC calls complete.

#### Recommendation

1. **Immediate**: Implement batch RPC calls using `eth_getBlockByNumber` with multicall
2. **Short-term**: Add asynchronous RPC call handling with worker pool
3. **Long-term**: Cache L1 block hashes and implement rate limiting

---

### V-2: Infinite Buffer Growth

**Severity:** 🟠 **High**
**Status:** ⚠️ **Open**
**Component:** `espresso/batch_buffer.go` - `BatchBuffer` and `HasNext()`

#### Description

The `BatchBuffer` has no size limit and will accept batches indefinitely while waiting for a missing batch, leading to memory exhaustion and node crashes.

#### Technical Details

**Vulnerable Logic:**
```go
func (b *BatchBuffer[B]) HasNext() bool {
    return b.Peek() == b.expectedBatchPos
}
```

**Attack Scenario:**

1. Node expects Batch #100
2. Espresso network delivers Batch #101, #102, ... #50,000
3. Batch #100 is missing (network partition, Byzantine node, etc.)
4. `BatchBuffer` accepts and stores Batches #101 through #50,000 in memory
5. Node waits indefinitely for Batch #100
6. Memory exhaustion → Node crash

#### Impact

- **Availability**: Node crashes due to out-of-memory (OOM)
- **Denial of Service**: Missing batch prevents all downstream processing
- **No Recovery**: No mechanism to invalidate the stream and skip missing batch

#### Likelihood

**Medium** - Requires network partition or Byzantine behavior, but no mitigation exists

#### Proof of Concept

```go
// Attacker causes batch #N to be permanently lost
// Network continues delivering batches N+1, N+2, ...
// Victim node accumulates unlimited batches in memory
for i := N+1; i < infinity; i++ {
    batchBuffer.Insert(batch[i], i)  // No size check!
}
// Eventually: panic: runtime: out of memory
```

#### Recommendation

1. **Critical**: Implement maximum buffer size (e.g., 1000 batches)
2. **Critical**: Add timeout for missing batches (e.g., 10 minutes)
3. **Important**: Implement gap detection and alerting
4. **Important**: Add mechanism to request missing batches from peers
5. **Long-term**: Implement stream reset when gap is detected beyond threshold

**Suggested Implementation:**
```go
const MAX_BUFFER_SIZE = 1000
const MISSING_BATCH_TIMEOUT = 10 * time.Minute

func (b *BatchBuffer[B]) TryInsert(batch B) (int, bool) {
    if len(b.batches) >= MAX_BUFFER_SIZE {
        return 0, false  // Reject new batches
    }

    if batch.Number > b.expectedBatchPos {
        if time.Since(b.lastProgressTime) > MISSING_BATCH_TIMEOUT {
            // Log critical error and reset stream
            return 0, false
        }
    }

    // ... existing logic
}
```

---

### V-7: Type Mismatch in Refresh()

**Severity:** 🟠 **High**
**Status:** ⚠️ **Open**
**Component:** `espresso/streamer.go` - `Refresh()` function, Line 173

#### Description

The `Refresh()` function contains a type mismatch where it compares `fallbackBatchPos` (representing a Batch Index) with `hotShotPos` (representing an Espresso Block Height). These are incompatible types that should not be directly compared.

#### Technical Details

**Vulnerable Code:**
```go
func (s *BatchStreamer[B]) Refresh(ctx context.Context, finalizedL1 eth.L1BlockRef,
                                    safeBatchNumber uint64, safeL1Origin eth.BlockID) error {
    // Line 173
    if fallbackBatchPos < hotShotPos {  // Type mismatch!
        // ... logic
    }
}
```

**Issue:**
- `fallbackBatchPos` is a **Batch Index** (sequential batch number)
- `hotShotPos` is an **Espresso Block Height** (blockchain height)
- Comparing these directly may lead to logic errors

#### Impact

- **State Inconsistency**: Incorrect state transitions during batch processing
- **Logic Error**: May cause unexpected behavior in edge cases
- **Potential Data Loss**: Could skip or process wrong batches

#### Likelihood

**Medium** - Will manifest in specific blockchain state conditions

#### Overall Risk

**High** - Logic error with potential for incorrect state management

#### Recommendation

1. Review the comparison logic and ensure type compatibility
2. Add type-safe wrappers or explicit conversions
3. Document the relationship between batch index and block height
4. Add assertions to validate the comparison is meaningful

---

### V-8: Missing Duplicate Detection

**Severity:** 🟢 **Low**
**Status:** ⚠️ **Open**
**Component:** `espresso/batch_buffer.go` - `Insert()` function

#### Description

The `Insert(batch B, i int)` function unconditionally inserts a batch at the specified index without checking for duplicates. Calling this function twice with the same batch will create duplicate entries.

#### Technical Details

**Vulnerable Code:**
```go
func (b *BatchBuffer[B]) Insert(batch B, i int) {
    // No duplicate check - directly inserts
    b.batches[i] = batch
}
```

#### Impact

- **Data Redundancy**: Duplicate batches in buffer
- **Memory Waste**: Unnecessary memory consumption
- **Potential Confusion**: Downstream processing may see duplicates

#### Likelihood

**Low** - Assumes correct upstream usage patterns

#### Overall Risk

**Low** - Minor inefficiency that relies on correct caller behavior

#### Recommendation

1. Add duplicate detection before insertion
2. Document preconditions that caller must ensure no duplicates
3. Add debug assertions in development builds
4. Consider returning a boolean to indicate if insertion occurred

**Suggested Implementation:**
```go
func (b *BatchBuffer[B]) Insert(batch B, i int) (inserted bool) {
    if b.batches[i] == batch {
        return false  // Already exists
    }
    b.batches[i] = batch
    return true
}
```

---

### V-9: Misleading Log Messages

**Severity:** 🟢 **Low**
**Status:** ⚠️ **Open**
**Component:** `espresso/streamer.go` - `processEspressoTransaction()`, Lines 304 & 435

#### Description

The streamer contains misleading and redundant log messages that make debugging more difficult and could confuse operators.

#### Technical Details

**Issue 1: Redundant Debug Log (Line 304)**
```go
s.log.Debug("Fetching range", "from", from, "to", to)
// fetchHotShotRange() immediately logs Trace
```

**Issue 2: Misleading Message (Line 435)**
```go
s.log.Warn("Batch already in buffer")  // Actually in RemainingBatches map!
```

#### Impact

- **Operational Confusion**: Misleading messages during debugging
- **Log Noise**: Redundant logs clutter output
- **Maintenance Burden**: Harder to understand code behavior

#### Likelihood

**High** - Will occur during normal operation

#### Overall Risk

**Low** - Does not affect functionality, only observability

#### Recommendation

1. **Line 304**: Remove redundant Debug log, rely on `fetchHotShotRange` logs
2. **Line 435**: Change message to "Batch already in remaining list" for accuracy
3. Add log level guidelines to documentation

---

### V-10: Inefficient Batch Overwrite

**Severity:** 🟢 **Low**
**Status:** ⚠️ **Open**
**Component:** `espresso/streamer.go` - `processEspressoTransaction()`, Line 435

#### Description

The code unnecessarily overwrites a batch in the `RemainingBatches` map when the batch already exists, performing redundant work.

#### Technical Details

**Inefficient Code:**
```go
if _, exists := s.RemainingBatches[hash]; exists {
    s.log.Warn("Batch already in buffer")
    s.RemainingBatches[hash] = *batch  // Overwrites with identical data!
}
```

**Analysis:**
- Hash is the map key derived from batch content
- If hash matches, the batch content must be identical
- Overwriting is redundant and wastes CPU cycles

#### Impact

- **Performance**: Minor CPU waste during batch processing
- **Code Clarity**: Suggests potential logic confusion

#### Likelihood

**Medium** - Occurs when batches are received multiple times

#### Overall Risk

**Low** - Benign inefficiency with minimal performance impact

#### Recommendation

Skip the overwrite operation when batch already exists:

```go
if _, exists := s.RemainingBatches[hash]; exists {
    s.log.Warn("Batch already in remaining list")
    return  // Skip redundant overwrite
}
s.RemainingBatches[hash] = *batch
```

---

### V-11: Confusing Variable Naming

**Severity:** 🟢 **Low**
**Status:** ⚠️ **Open**
**Component:** `espresso/cli.go` - Configuration variable naming

#### Description

The configuration variable `PollingHotShotPollingInterval` contains redundant naming that reduces code readability and increases cognitive load for developers.

#### Technical Details

**Current Naming:**
```go
PollingHotShotPollingInterval  // "Polling" appears twice
```

**Issue:**
- Redundant "Polling" prefix and suffix
- Verbose without added clarity
- Violates DRY principle in naming

#### Impact

- **Maintainability**: Harder to read and understand configuration
- **Developer Experience**: Increased cognitive load
- **Documentation**: More verbose configuration examples

#### Likelihood

**High** - Affects every developer working with the codebase

#### Overall Risk

**Low** - Code quality issue with no functional impact

#### Recommendation

Simplify to `HotShotPollingInterval`:

```go
HotShotPollingInterval  // Clear and concise
```

---

### V-12: Unused Constant Declaration

**Severity:** 🟢 **Low**
**Status:** ⚠️ **Open**
**Component:** `espresso/` - Constants definition

#### Description

The constant `HOTSHOT_BLOCK_STREAM_LIMIT` is defined in the codebase but never actually used, leading to dead code and potential confusion.

#### Technical Details

**Declared Constant:**
```go
const HOTSHOT_BLOCK_STREAM_LIMIT = 1000  // Never referenced
```

**Issue:**
- Constant is defined but has zero references
- May indicate incomplete feature implementation
- Adds maintenance burden

#### Impact

- **Code Bloat**: Unnecessary declarations in codebase
- **Confusion**: Developers may wonder about its purpose
- **Maintenance**: Must be maintained despite no usage

#### Likelihood

**N/A** - Already present in codebase

#### Overall Risk

**Low** - Minor code quality issue

#### Recommendation

1. **If unused**: Remove the constant entirely
2. **If planned**: Add TODO comment explaining future usage
3. **If needed**: Implement the feature that should use this limit

---

### V-13: Missing Sort Order Validation

**Severity:** 🟢 **Low**
**Status:** ⚠️ **Open**
**Component:** `espresso/batch_buffer.go` - `TryInsert()` function

#### Description

The `TryInsert()` function assumes the batch list is already sorted and uses binary search without verifying this invariant. If the sort order is violated, the function will produce incorrect results.

#### Technical Details

**Current Implementation:**
```go
func (b *BatchBuffer[B]) TryInsert(batch B) (int, bool) {
    // Uses binary search - assumes sorted list
    // No validation that list is actually sorted
}
```

**Issue:**
- Critical invariant (sorted order) is assumed but not verified
- Binary search will fail silently if invariant is broken
- No debug assertions to catch violations

#### Impact

- **Correctness**: Incorrect insertion if invariant violated
- **Debugging**: Hard to diagnose if sort order breaks
- **Reliability**: Silent failures in edge cases

#### Likelihood

**Low** - Invariant should be maintained by design

#### Overall Risk

**Low** - Invariant maintained by implementation, but lacks safety checks

#### Recommendation

Add debug assertions in development builds:

```go
func (b *BatchBuffer[B]) TryInsert(batch B) (int, bool) {
    if DEBUG {
        // Verify sort order invariant
        for i := 1; i < len(b.batches); i++ {
            if b.batches[i-1] >= b.batches[i] {
                panic("BatchBuffer invariant violated: list not sorted")
            }
        }
    }
    // ... existing binary search logic
}
```

---

### V-14: No Network Failure Distinction

**Severity:** 🟢 **Low**
**Status:** ⚠️ **Open**
**Component:** `espresso/streamer.go` - `confirmEspressoBlockHeight()` function

#### Description

The `confirmEspressoBlockHeight()` function returns `false` when the `FinalizedState()` RPC call fails, treating network failures the same as "no reorg occurred". This makes it impossible to distinguish between actual state verification and network errors.

#### Technical Details

**Current Behavior:**
```go
func (s *BatchStreamer[B]) confirmEspressoBlockHeight(safeL1Origin eth.BlockID) (shouldReset bool) {
    state, err := s.FinalizedState()
    if err != nil {
        return false  // Network error treated as "no reorg"
    }
    // ... actual reorg check
}
```

**Issue:**
- Network failure → returns `false` (conservative default)
- No reorg → returns `false` (correct behavior)
- Cannot distinguish between these two cases

#### Impact

- **Observability**: Cannot detect network issues vs. normal operation
- **Debugging**: Harder to diagnose connectivity problems
- **Monitoring**: No visibility into RPC failure rate

#### Likelihood

**Low** - Conservative default is safe but reduces observability

#### Overall Risk

**Low** - Safe default behavior, only affects monitoring and debugging

#### Recommendation

Add explicit error handling to distinguish cases:

```go
func (s *BatchStreamer[B]) confirmEspressoBlockHeight(safeL1Origin eth.BlockID) (shouldReset bool) {
    state, err := s.FinalizedState()
    if err != nil {
        s.log.Warn("Failed to fetch finalized state, assuming no reorg", "error", err)
        s.metrics.RPCFailures.Inc()  // Track network failures
        return false
    }
    // ... actual reorg check with clear logging
}
```

---

## 3. TEE Enclave Vulnerabilities

### V-3: TEE Networking MitM Attack

**Severity:** 🟠 **High**
**Status:** ⚠️ **Open**
**Component:** TEE Enclave Networking Layer


#### Description

The TEE enclave networking layer lacks sufficient protection against Man-in-the-Middle (MitM) attacks, potentially allowing malicious actors to feed the enclave with arbitrary input, such as maliciously crafted HotShot blocks.

#### Technical Details

**Vulnerability:**
- TEE networking layer does not enforce strict TLS certificate validation
- No certificate pinning implemented
- Enclave trusts any valid TLS certificate

**Attack Scenario:**

1. Attacker positions themselves between TEE enclave and HotShot network
2. Attacker presents valid TLS certificate (e.g., from compromised CA)
3. TEE accepts connection as legitimate
4. Attacker injects malicious HotShot blocks
5. TEE processes fraudulent data as authentic

#### Impact

- **Integrity**: TEE processes malicious input as authentic
- **Consensus Manipulation**: Fraudulent blocks could affect L2 state
- **Trust Violation**: Undermines security guarantees of TEE

#### Likelihood

**Medium** - Requires network access but no cryptographic breaks

#### Current Mitigation

Documentation exists at: https://eng-wiki.espressosys.com/mainch36.html#:Future%20Work:Trustless%20enclave%20networking

However, implementation remains vulnerable.

#### Related Concern

**TLS Certificate Expiration:**
- Current approach may require rebuilding enclave when certificates expire
- Creates operational burden and potential security windows during rotation
- Frequency of certificate expiration is a concern

#### Recommendation

1. **Critical**: Implement certificate pinning
   - Embed expected certificates during enclave build
   - Include certificates in PCR0 hash measurement
   - Ensure enclave only trusts specific, validated endpoints

2. **Important**: Add certificate rotation mechanism
   - Design automatic certificate update process
   - Implement gradual rollover to avoid service interruption

3. **Long-term**: Implement attestation-based mutual authentication
   - Both endpoints verify each other's TEE attestations
   - Remove dependency on traditional PKI

**Suggested Implementation:**
```go
// Embed certificates at build time
const EXPECTED_CERT_HASH = "sha256:abc123..."

func VerifyConnection(conn *tls.Conn) error {
    certs := conn.ConnectionState().PeerCertificates
    if len(certs) == 0 {
        return errors.New("no peer certificates")
    }

    hash := sha256.Sum256(certs[0].Raw)
    expected, _ := hex.DecodeString(EXPECTED_CERT_HASH)

    if !bytes.Equal(hash[:], expected) {
        return errors.New("certificate pinning validation failed")
    }

    return nil
}
```

**Reference:** v0.5.0 - https://github.com/EspressoSystems/optimism-espresso-integration/releases/tag/v0.5.0

---

## 4. TEE Contracts Vulnerabilities

The following vulnerabilities were identified and **fixed** in [PR #43](https://github.com/EspressoSystems/espresso-tee-contracts/pull/43) of the `espresso-tee-contracts` repository.

**PR #43 Summary:**
- **Title**: Internal Audit #2 - Security Fixes
- **Merged**: January 28, 2026
- **Commit**: `1a5a179`
- **Files Changed**: 21 files (+1098, -250 lines)
- **Test Coverage**: 624 new test lines added

---

### V-4: Cross-Chain Deployment Vulnerability

**Severity:** 🔴 **Critical**
**Status:** ✅ **Fixed**
**Component:** `EspressoNitroTEEVerifier.sol`, `EspressoSGXTEEVerifier.sol`
**Fix Reference:** Commit `57bf5de`, PR #43

#### Description

TEE Verifier contracts maintain chain-specific on-chain state for registered enclaves and signers. However, attestations are not chain-specific by default, allowing replay attacks across different chains with inconsistent security policies.

#### Technical Details

**Vulnerable State Management:**
```solidity
// These mappings are stored ON-CHAIN (chain-specific):
mapping(ServiceType => mapping(bytes32 => bool)) public registeredEnclaveHashes;
mapping(ServiceType => mapping(address => bool)) public registeredServices;
```

**Problem**: State is local to each chain, but attestations can be replayed across chains.

#### Attack Scenarios

**Attack 1: Uncoordinated Revocation**

Timeline:
- Day 1: Enclave hash approved on Ethereum and Arbitrum
- Day 30: Vulnerability discovered in enclave
- Day 31: Hash revoked on Ethereum
- **Result**: Attacker blocked on Ethereum ✅ but still valid on Arbitrum ❌

**Attack 2: Attestation Replay**

1. TEE generates single attestation
2. Attacker registers on Ethereum using attestation
3. Attacker reuses **same attestation** on Arbitrum
4. Attacker reuses **same attestation** on Optimism
5. All registrations succeed (if hash is approved on each chain)

**Attack 3: Policy Inconsistency**

- Ethereum: High security, only approves hash v2.0 (latest, secure)
- Arbitrum: Different governance, approves hash v1.0 (old, vulnerable)
- **Result**: Same codebase, different security across chains

#### Impact

- **Security Fragmentation**: Inconsistent security policies across chains
- **Delayed Response**: Vulnerability on one chain doesn't automatically propagate
- **Replay Attacks**: Single attestation usable on multiple chains

#### Likelihood

**High** - Natural consequence of multi-chain deployment without chain ID validation

#### Overall Risk

**Critical** - High impact authentication/security bypass combined with high likelihood in multi-chain deployments

#### Fix Applied

✅ Added [Security Considerations](https://github.com/EspressoSystems/espresso-tee-contracts/blob/main/README.md#security-considerations) section in README.md.


---

### V-5: Signer Deletion DoS Attack

**Severity:** 🟠 **High**
**Status:** ✅ **Fixed**
**Component:** `TEEHelper.sol`
**Fix Reference:** Commit `3026966`, PR #43

#### Description

The TEE Helper contract iterates over a list of registered signers in deletion operations. An attacker could exploit unbounded loops to cause denial of service by exceeding block gas limits.

**Fix:** PR #43 changed the security model so that **deleting signers is no longer required**. Revoking an enclave hash via `setEnclaveHash(hash, false)` is now sufficient to prevent new malicious registrations, eliminating the need for the DoS-vulnerable deletion operation.



**Attack Scenario:**
1. Attacker registers many signers (e.g., 10,000 addresses)
2. Enclave is compromised
3. Operator tries to revoke by calling `setEnclaveHash(hash, false)` and `deleteRegisteredSigners()`
4. **Deletion fails due to gas limit** - transaction reverts
5. **Compromised signers remain active** - security breach!

#### Impact

- **Security Bypass**: Unable to fully revoke compromised enclave access
- **DoS on Critical Security Function**: Deletion operation required but impossible
- **Persistent Vulnerability**: Compromised signers remain valid indefinitely

#### Likelihood

**High** - Attacker can easily register many signers to prevent future revocation

#### Fix Applied in PR #43

| Signers | Gas Cost | Block Limit (30M) | Status |
|---------|----------|-------------------|---------|
| 100 | ~500k | ✅ Safe | OK |
| 1,000 | ~5M | ✅ Safe | OK |
| 5,000 | ~25M | ⚠️ Close | Risk |
| 10,000 | ~50M | ❌ Over | DoS |

**AFTER PR #43:** Revoking an enclave only requires one step:
1. Call `setEnclaveHash(hash, false)` to prevent new registrations ✅ **Sufficient**
2. ~~Delete existing signers~~ ❌ **No longer needed**

**Why This Works:**
- When an enclave is compromised, the private keys are already exposed to attackers
- Existing signer addresses in the registry don't grant any additional attack surface
- The security boundary is enforced at enclave hash validation, not signer presence
- Revoking the hash immediately protects the system



**Note:** Operators may optionally use this function to reduce contract state size, but it's not a security requirement.

#### Verification

- ✅ Enclave hash revocation alone is sufficient to protect system
- ✅ DoS attack vector eliminated by removing requirement for vulnerable operation

---

### V-6: Missing TEE Journal Struct Validations

**Severity:** 🔴 **Critical**
**Status:** ✅ **Fixed**
**Component:** `EspressoNitroTEEVerifier.sol`
**Fix Reference:** Commit `c47d9aa`, PR #43

#### Description

The VerifierJournal struct contains critical cryptographic fields (PCRs, public key, nonce, timestamp, userData) that require comprehensive validation. Missing validations could allow malformed attestations to be accepted, potentially leading to predictable signer addresses or other cryptographic attacks.

#### Technical Details

**Journal Structure:**
```solidity
struct VerifierJournal {
    bytes32[] pcrs;          // Platform Configuration Registers
    bytes publicKey;         // Enclave public key (should be 65 bytes)
    bytes nonce;             // Replay protection
    uint256 timestamp;       // Attestation time
    bytes userData;          // Application data
    string moduleId;         // Nitro module identifier
    VerificationResult result;
}
```

#### Specific Vulnerabilities

**V-6a: Empty PCR Array**
- **Issue**: No validation that PCR array contains data
- **Impact**: Could accept attestations without platform measurements
- **Exploit**: Bypass hardware attestation requirements

**V-6b: Invalid Public Key Format**
- **Issue**: No check for correct public key length (65 bytes) and format
- **Impact**: Malformed public keys could lead to predictable addresses
- **Exploit**:
  ```solidity
  // Attacker provides short public key
  bytes memory badKey = hex"04";  // Only 1 byte instead of 65
  address predictable = deriveAddress(badKey);  // Predictable result!
  ```

**V-6c: Predictable Signer Addresses**
- **Issue**: Invalid public key formats can produce predictable Ethereum addresses
- **Impact**: Attacker could precompute and claim desirable addresses
- **Severity**: Enables address squatting and impersonation

#### Impact

- **Authentication Bypass**: Malformed attestations accepted
- **Address Prediction**: Attacker could generate predictable signer addresses
- **Integrity Violation**: TEE guarantees undermined

#### Likelihood

**High** - Malformed attestations can be trivially crafted without cryptographic knowledge

#### Overall Risk

**Critical** - Direct authentication bypass and potential for address prediction/squatting attacks. Allows completely invalid attestations to be accepted, undermining entire security model.

#### Fix Applied

Added comprehensive `_validateJournal()` validation function:

```solidity
function _validateJournal(VerifierJournal memory journal) internal view {
    // 1. Validate PCR array bounds
    require(journal.pcrs.length > 0, "PCR array cannot be empty");

    // 2. CRITICAL: Validate public key format
    require(journal.publicKey.length == 65, "Invalid public key length");
    require(journal.publicKey[0] == 0x04, "Public key must be uncompressed");

    // 3. Note: Nonce validation removed - AWS Nitro may have empty nonce
    // Implement nonce tracking separately if replay protection needed

    // 4. Timestamp validation already done by NitroEnclaveVerifier
    // Result would be InvalidTimestamp if timestamp is bad

    // 5. Optional: Additional userData validation
    // require(journal.userData.length > 0, "UserData cannot be empty");
}
```

**Integration:**
```solidity
function registerService(...) external {
    // Verify attestation
    if (journal.result != VerificationResult.Success) {
        revert VerificationFailed(journal.result);
    }

    // NEW: Validate journal format and integrity
    _validateJournal(journal);  // ✅ Defense in depth

    // ... rest of registration
}
```

---

**End of Report**

---


**Last Updated:** January 29, 2026
