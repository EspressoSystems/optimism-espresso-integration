# Espresso Integration Security Audit Plan

This document provides a comprehensive audit plan for the Celo/Espresso integration code in the Optimism stack. The plan is organized by component priority and criticality.

---

## Overview

The Espresso integration introduces several critical components that ensure consistency between the state derived by the Caff node and a normal OP verifier node. Any vulnerabilities could result in users losing funds or break safety guarantees from an Espresso app perspective.

### Risk Assessment

| Component | Risk Level | Impact |
|-----------|------------|--------|
| L1 Contracts (BatchAuthenticator, BatchInbox) | **Critical** | Fund loss, state inconsistency |
| Streamer | **High** | Safety violations, liveness failures |
| Batcher Logic | **High** | Data integrity, transaction ordering |
| Derivation Pipeline | **Medium-High** | Fundamental behavior changes |

---

## 1. L1 Contracts Audit

### 1.1 BatchAuthenticator.sol
**File:** `packages/contracts-bedrock/src/L1/BatchAuthenticator.sol` (93 lines)

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| L1-AUTH-01 | Review ECDSA signature verification logic in `authenticateBatchInfo()` - ensure proper v-value normalization (lines 66-87) | Critical | 4 |
| L1-AUTH-02 | Audit `registerSigner()` function and its interaction with `IEspressoTEEVerifier` (lines 89-92) | Critical | 3 |
| L1-AUTH-03 | Verify access control: `Ownable` pattern and `switchBatcher()` function security (lines 61-63) | High | 2 |
| L1-AUTH-04 | Audit immutable variables: `teeBatcher`, `nonTeeBatcher`, `preRegisteredBatcher` - verify proper initialization and zero-address checks | High | 2 |
| L1-AUTH-05 | Review `validBatchInfo` mapping - check for replay attack vectors | Critical | 3 |
| L1-AUTH-06 | Analyze trust model: `preRegisteredBatcher` bypass of TEE verification (lines 78-83) | High | 2 |
| L1-AUTH-07 | Verify constructor parameter validation and ownership transfer | Medium | 1 |

#### Key Security Concerns
- Signature malleability in ECDSA recovery
- State manipulation through batch info mapping
- TEE batcher vs non-TEE batcher mode switching attacks
- Pre-registered batcher bypass mechanism

---

### 1.2 BatchInbox.sol
**File:** `packages/contracts-bedrock/src/L1/BatchInbox.sol` (77 lines)

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| L1-INBOX-01 | Audit `fallback()` function - ensure proper authentication flow for TEE mode (lines 27-60) | Critical | 4 |
| L1-INBOX-02 | Review blob hash verification logic using `blobhash()` opcode (lines 43-53) | Critical | 4 |
| L1-INBOX-03 | Verify calldata batch hash verification (lines 54-59) | Critical | 2 |
| L1-INBOX-04 | Analyze batcher authorization checks for both TEE and fallback modes | High | 3 |
| L1-INBOX-05 | Review error handling and revert message construction | Medium | 1 |
| L1-INBOX-06 | Check for DoS vectors in blob iteration loop (lines 45-49) | High | 2 |
| L1-INBOX-07 | Verify integration with `IBatchAuthenticator` interface | High | 2 |

#### Key Security Concerns
- Blob hash manipulation or spoofing
- Race conditions between `authenticateBatchInfo` and batch submission
- Gas griefing through excessive blob counts
- Mode switching attacks (TEE vs fallback)

---

### 1.3 External Dependency: espresso-tee-contracts
**Repository:** `github.com/EspressoSystems/espresso-tee-contracts`

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| L1-TEE-01 | Review `IEspressoTEEVerifier` interface implementation | Critical | 4 |
| L1-TEE-02 | Audit ZK attestation verification (PR #29) | Critical | 8 |
| L1-TEE-03 | Verify Nitro TEE verifier `registeredSigners` mapping | High | 3 |
| L1-TEE-04 | Review signer registration flow and trust model | High | 4 |

#### Note
These contracts are part of global infrastructure shared across Nitro and OP stack and have been used in production. Lower priority but should still be reviewed.

---

## 2. Streamer Audit

### 2.1 Core Streamer Logic
**File:** `espresso/streamer.go` (605 lines)

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| STR-01 | Audit `BatchStreamer` struct initialization and state management (lines 71-142) | High | 4 |
| STR-02 | Review `Reset()` function - verify proper state restoration (lines 145-150) | Critical | 2 |
| STR-03 | Audit `Refresh()` function - analyze sync status handling and reset conditions (lines 163-184) | Critical | 4 |
| STR-04 | Review `CheckBatch()` function - verify L1 origin finality checks (lines 188-229) | Critical | 4 |
| STR-05 | Audit `Update()` method - analyze block fetching and batch processing (lines 273-347) | High | 6 |
| STR-06 | Review `fetchHotShotRange()` - verify transaction processing and position tracking (lines 355-386) | High | 4 |
| STR-07 | Audit `streamHotShotRange()` - analyze streaming API implementation (lines 394-448) | High | 4 |
| STR-08 | Review `processRemainingBatches()` - verify batch recovery logic (lines 452-497) | High | 3 |
| STR-09 | Audit `processEspressoTransaction()` - verify batch validation flow (lines 501-537) | High | 3 |
| STR-10 | Review `confirmEspressoBlockHeight()` - L1 reorg handling (lines 569-599) | Critical | 4 |
| STR-11 | Analyze `BatchBuffer` operations - verify ordering and deduplication | High | 3 |
| STR-12 | Review `RemainingBatches` map - check for memory leaks and race conditions | Medium | 2 |

#### Key Security Concerns
- Race conditions between L1/L2/Espresso state updates
- Batch ordering violations leading to safety issues
- L1 reorg handling edge cases
- Memory exhaustion through unprocessed batches
- Off-by-one errors in position tracking

---

### 2.2 Supporting Files

#### 2.2.1 Interface Definition
**File:** `espresso/interface.go` (66 lines)

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| STR-INT-01 | Review `EspressoStreamer` interface contract | Medium | 1 |

#### 2.2.2 Batch Buffer
**File:** `espresso/batch_buffer.go` (91 lines)

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| STR-BUF-01 | Audit `BatchBuffer` generic implementation | Medium | 2 |
| STR-BUF-02 | Verify `TryInsert()` binary search correctness | High | 2 |
| STR-BUF-03 | Review `Pop()` and `Peek()` operations for thread safety | Medium | 1 |

---

## 3. Batcher Logic Audit

### 3.1 Driver (Espresso Mode)
**File:** `op-batcher/batcher/driver.go` (1159 lines)

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| BAT-DRV-01 | Audit `DriverSetup` struct - verify Espresso-specific fields (lines 94-111) | High | 2 |
| BAT-DRV-02 | Review `StartBatchSubmitting()` Espresso mode initialization (lines 215-243) | Critical | 4 |
| BAT-DRV-03 | Audit `registerBatcher()` call and error handling (line 217-219) | Critical | 2 |
| BAT-DRV-04 | Review `espressoSubmitter` initialization and worker spawning | High | 3 |
| BAT-DRV-05 | Audit `teeAuthGroup` limiter (128 concurrent goroutines) | Medium | 2 |
| BAT-DRV-06 | Review `sendTx()` function - Espresso path vs standard path (lines 1036-1048) | Critical | 3 |
| BAT-DRV-07 | Audit `clearState()` - verify Espresso streamer reset (lines 825-868) | High | 2 |

---

### 3.2 Espresso-Specific Logic
**File:** `op-batcher/batcher/espresso.go` (1153 lines)

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| BAT-ESP-01 | Audit `espressoTransactionSubmitter` architecture (lines 91-104) | High | 4 |
| BAT-ESP-02 | Review job submission/retry logic - `evaluateSubmission()` (lines 223-242) | Critical | 3 |
| BAT-ESP-03 | Audit receipt verification - `evaluateVerification()` (lines 321-345) | Critical | 3 |
| BAT-ESP-04 | Review worker queue pattern (lines 395-480) | High | 4 |
| BAT-ESP-05 | Audit `espressoSubmitTransactionWorker()` (lines 495-550) | High | 3 |
| BAT-ESP-06 | Review `espressoVerifyTransactionWorker()` (lines 558-614) | High | 3 |
| BAT-ESP-07 | Audit `queueBlockToEspresso()` - batch conversion and signing (lines 655-675) | Critical | 3 |
| BAT-ESP-08 | Review `espressoSyncAndRefresh()` function (lines 677-698) | High | 2 |
| BAT-ESP-09 | Audit `espressoBatchLoadingLoop()` - main Espresso polling loop (lines 700-773) | Critical | 6 |
| BAT-ESP-10 | Review `BlockLoader` and `nextBlockRange()` logic (lines 775-900) | High | 4 |
| BAT-ESP-11 | Audit `espressoBatchQueueingLoop()` (lines 905-944) | High | 3 |
| BAT-ESP-12 | Review `registerBatcher()` ZK proof generation and submission (lines 963-1025) | Critical | 6 |
| BAT-ESP-13 | Audit `GenerateZKProof()` HTTP interaction (lines 1027-1063) | High | 3 |
| BAT-ESP-14 | Review `sendTxWithEspresso()` - commitment signing and authentication (lines 1067-1145) | Critical | 6 |

#### Key Security Concerns
- Transaction replay attacks
- Race conditions in concurrent transaction submission
- Job queue overflow/underflow
- Improper error handling leading to transaction loss
- ZK proof generation failures
- Signature verification bypass

---

## 4. Derivation Pipeline Audit

### 4.1 Data Source
**File:** `op-node/rollup/derive/data_source.go` (122 lines)

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| DER-DS-01 | Review modified `isValidBatchTx()` function - removed sender verification (lines 98-121) | Critical | 4 |
| DER-DS-02 | Verify integration with BatchInbox contract authentication model | High | 3 |
| DER-DS-03 | Analyze impact of removing L1Signer verification | Critical | 4 |

#### Key Security Concerns
- The comment at line 114-117 indicates sender verification is now delegated to BatchInbox contract
- Ensure this change doesn't introduce any bypass opportunities
- Verify consistency with upstream OP behavior when Espresso is disabled

---

### 4.2 Espresso Batch
**File:** `op-node/rollup/derive/espresso_batch.go` (133 lines)

#### Tasks

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| DER-EB-01 | Audit `EspressoBatch` struct and its conversion methods | High | 3 |
| DER-EB-02 | Review `ToEspressoTransaction()` - RLP encoding and signing (lines 44-61) | Critical | 4 |
| DER-EB-03 | Audit `BlockToEspressoBatch()` - L1 info deposit extraction (lines 63-83) | High | 3 |
| DER-EB-04 | Review `UnmarshalEspressoTransaction()` - signature verification (lines 95-113) | Critical | 4 |
| DER-EB-05 | Audit `ToBlock()` - verify transaction reconstruction (lines 118-132) | Critical | 4 |

#### Key Security Concerns
- RLP encoding/decoding consistency
- Signature verification soundness
- L1 info deposit manipulation
- Transaction ordering preservation

---

## 5. Cross-Cutting Concerns

### 5.1 Integration Testing

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| INT-01 | Review existing streamer tests (`espresso/streamer_test.go`) for coverage gaps | High | 4 |
| INT-02 | Identify missing edge case tests for L1 reorg scenarios | Critical | 4 |
| INT-03 | Review batcher<->streamer integration scenarios | High | 4 |
| INT-04 | Verify contract deployment and initialization sequences | High | 3 |

### 5.2 Configuration and Deployment

| Task ID | Description | Priority | Estimated Hours |
|---------|-------------|----------|-----------------|
| CFG-01 | Review rollup configuration parameters for Espresso integration | Medium | 2 |
| CFG-02 | Verify default values and parameter validation | Medium | 2 |
| CFG-03 | Audit attestation service configuration | High | 2 |

---

## 6. Summary

### Total Estimated Hours by Component

| Component | Estimated Hours |
|-----------|-----------------|
| L1 Contracts (BatchAuthenticator, BatchInbox) | ~40 hours |
| External TEE Contracts | ~19 hours |
| Streamer | ~45 hours |
| Batcher Logic | ~64 hours |
| Derivation Pipeline | ~22 hours |
| Cross-Cutting Concerns | ~21 hours |
| **Total** | **~211 hours** |

### Priority Order

1. **Immediate (Critical Path)**
   - L1-AUTH-01 through L1-AUTH-06
   - L1-INBOX-01 through L1-INBOX-04
   - STR-02, STR-03, STR-04, STR-10
   - BAT-ESP-07, BAT-ESP-12, BAT-ESP-14
   - DER-DS-01, DER-EB-02, DER-EB-04, DER-EB-05

2. **High Priority**
   - Remaining Streamer tasks
   - Remaining Batcher tasks
   - Integration testing

3. **Medium Priority**
   - Configuration review
   - Code quality improvements

---

## 7. Previous Audit Status

### Completed Audits

| Date | Auditor | Status | Notes |
|------|---------|--------|-------|
| May 29, 2025 | External | ✅ Complete | Version 0.1.0 - All fixes implemented |
| November 7, 2025 | Olympix | ⚠️ Pending fixes | Minor findings - Fixes due before January 9, 2026 |
| November 12, 2025 | Alysia Huggins | ✅ Mostly complete | One exception pending discussion (PCR0 Hash Generation) |

### Outstanding Items

1. **Olympix Findings** - Implementation pending, deadline January 9, 2026
2. **PCR0 Hash Generation** - Requires further discussion for both OP and Nitro chains

---

## 8. External Resources

### TEE Security References
- [Trail of Bits: Top TEE Bugs Before Audit](https://watch.getcontrast.io/register/trail-of-bits-top-tee-bugs-you-should-fix-before-your-audit)
- [Trail of Bits: TEE Infrastructure Changes](https://watch.getcontrast.io/register/trail-of-bits-after-wiretap-and-battering-ram-what-changes-for-tee-based-blockchain-infrastructure)
- [Trail of Bits TEE Blog](https://blog.trailofbits.com/categories/trusted-execution-environment/)

---

## 9. Appendix: File References

| File | Lines | Component |
|------|-------|-----------|
| `packages/contracts-bedrock/src/L1/BatchAuthenticator.sol` | 93 | L1 Contracts |
| `packages/contracts-bedrock/src/L1/BatchInbox.sol` | 77 | L1 Contracts |
| `espresso/streamer.go` | 605 | Streamer |
| `espresso/interface.go` | 66 | Streamer |
| `espresso/batch_buffer.go` | 91 | Streamer |
| `op-batcher/batcher/driver.go` | 1159 | Batcher |
| `op-batcher/batcher/espresso.go` | 1153 | Batcher |
| `op-node/rollup/derive/data_source.go` | 122 | Derivation |
| `op-node/rollup/derive/espresso_batch.go` | 133 | Derivation |

---

*Document generated: January 12, 2026*
*Based on commit: 15c5f3f8f1465159eb9f3ec803a529ce37a8e624 (referenced in source document)*










