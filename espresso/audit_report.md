# OP Streamer Component Audit

## General Notes

- We could simplify the name of `PollingHotShotPollingInterval` to something like `HotshotPollingInterval`.
- The value `HOTSHOT_BLOCK_STREAM_LIMIT` is never used in the code.
- Only real issues found would happen in the case of race conditions.

## File Dive: `espresso/batch_buffer.go`

### `Insert`
```go
func (b *BatchBuffer[B]) Insert(batch B, i int)
```

- Calling this function twice with the same batch will cause the batch to be inserted twice. This method does not check for duplicates; it unconditionally inserts at index `i`.

### `TryInsert`
```go
func (b *BatchBuffer[B]) TryInsert(batch B) (int, bool)
```

- This function assumes the list is already sorted but it should always be the case.

---

## File Dive: `espresso/streamer.go`

### `GetFinalizedL1`
```go
func GetFinalizedL1(header *espressoCommon.HeaderImpl) espressoCommon.L1BlockInfo
```

- Could we not create an espresso-network go sdk function for this instead if this information is present in all the headers?

### `Refresh`
```go
func (s *BatchStreamer[B]) Refresh(ctx context.Context, finalizedL1 eth.L1BlockRef, safeBatchNumber uint64, safeL1Origin eth.BlockID) error
```

- Line 173 compares `fallbackBatchPos` (Batch Index) with `hotShotPos` (Espresso Block Height). How would that be possible?

### `processEspressoTransaction`
```go
func (s *BatchStreamer[B]) processEspressoTransaction(ctx context.Context, transaction espressoCommon.Bytes)
```

- The `Debug` log in `Update` (line 304) is redundant because `fetchHotShotRange` immediately logs `Trace`. We could remove the `Debug` log in `Update` and rely on `fetchHotShotRange`'s `Trace` log for low-level debugging and its `Info` log (line 344) for successful fetches. This clarifies the logs and reduces noise.
- The log "Batch already in buffer" (line 435) is misleading. It refers to `RemainingBatches` (the pending map), NOT `BatchBuffer` (the main sorted slice).  "Batch already in remaining list" would be more accurate!
- If the batch is found in the map, it warns but then immediately overwrites it with the new copy (`s.RemainingBatches[hash] = *batch`). Since the hash is the key, the content should be identical, making this overwrite benign but redundant.

### `confirmEspressoBlockHeight`
```go
func (s *BatchStreamer[B]) confirmEspressoBlockHeight(safeL1Origin eth.BlockID) (shouldReset bool)
```

- The function returns false when FinalizedState() fails meaning "do not reset the streamer", treating a network/RPC failure the same way as "no reorg happened".

## Deeper Dive on Component Flow

### 1. The All At Once RPC Calls"
- **Background**: The function `CheckBatch` makes a synchronous L1 RPC call (`HeaderHashByNumber`) if the batch's L1 origin is said to be finalized:
- **Scenario**:
  1.  The node accumulates 500 batches in `RemainingBatches` while waiting for L1 finality.
  2.  L1 finalizes a new state.
  3.  `processRemainingBatches` runs and iterates all 500 batches.
  4.  The "Finalized" check now passes for all of them.
  5.  The node executes 500 sequential synchronous RPC calls to the L1 node inside the `Update` loop.
- **Consequence**: This leads to the streamer freezing for seconds or minutes and stops fetching new Espresso blocks.

### 2. Infinite Buffer Growth
- **Background**: The function `HasNext` only returns `true` if `BatchBuffer.Peek() == BatchPos`.
- **Scenario**:
  1. The node is expecting Batch #100.
  2. Espresso delivers Batch #101, #102, ... #50,000.
  3. Batch #100 is still missing.
- **Consequence**: `BatchBuffer` has no size limit. It will accept and store Batches #101 through #50,000 in memory, waiting forever for #100. The node will run out of memory and crash. There is no mechanism to invalidate the stream if a batch is permanently lost.

### 3. TEE Networking Layer MitM Attack (Asana Task #1212819560549006)
- **Background**: The TEE enclave networking layer could potentially be abused to feed the enclave with arbitrary input, such as maliciously crafted HotShot blocks.
- **Issue**: Despite documentation about trustless enclave networking (https://eng-wiki.espressosys.com/mainch36.html#:Future%20Work:Trustless%20enclave%20networking), the system remains vulnerable to Man-in-the-Middle attacks.
- **Root Cause**: The enclave networking layer does not have sufficient protection against network-layer attacks that could inject malicious data into the TEE.
- **Related Concern**: TLS certificate management may require rebuilding the enclave whenever certificates expire, which could create operational challenges and potential security windows.
- **Reference**: https://github.com/EspressoSystems/optimism-espresso-integration/releases/tag/v0.5.0
- **Recommendation**: Implement certificate pinning or embed certificates during the build process so that PCR0 hash includes these certificates, ensuring the enclave only trusts specific, validated endpoints.

## TEE Contracts Vulnerabilities (PR #43: EspressoSystems/espresso-tee-contracts)

The following vulnerabilities were identified and fixed in PR #43 of the espresso-tee-contracts repository.

### 4. Cross-Chain Deployment Vulnerability (Documentation Gap)
- **Background**: TEE Verifier contracts maintain on-chain state for registered enclaves and signers using chain-local mappings.
- **Issue**: State is not synchronized across chains, but attestations are NOT chain-specific by default and can be replayed across multiple chains.
- **Attack Scenarios**:
  1. **Uncoordinated Revocation**: A compromised enclave hash can be revoked on one chain (e.g., Ethereum) but remain valid on another (e.g., Arbitrum), allowing continued attacks on the unrevoked chain.
  2. **Attestation Replay**: A single TEE attestation can be reused to register on multiple chains (Ethereum, Arbitrum, Optimism) if the hash is approved on all chains.
  3. **Different Security Policies**: Each chain may have different approved hashes and governance, creating inconsistent security across deployments.
- **Root Cause**: Contracts do not validate that attestations are intended for a specific chain (no chain ID verification).
- **Fix Applied**:
  - Added comprehensive documentation in `CROSS_CHAIN_SECURITY.md` explaining the risks
  - Updated `VULNERABILITY_REPORT.md` with deployment guidance
  - Recommended TEE encode target chain ID in attestation userData and validate it: `require(attestedChainId == block.chainid, "Attestation for wrong chain")`
- **Status**: Fixed via documentation and recommended implementation pattern
- **Reference**: Commit 57bf5de, PR #43

### 5. Signers List DoS Attack
- **Background**: The TEE Helper contract iterates over a list of registered signers in certain operations.
- **Issue**: An attacker could exploit unbounded loops over the signers list to cause denial of service, making the contract unusable due to gas limits.
- **Attack Scenario**:
  1. Attacker registers many signers (within allowed limits)
  2. Operations that iterate over all signers hit gas limits
  3. Contract becomes unable to process legitimate requests
- **Fix Applied**:
  - Added `MAX_BATCH_DELETE_SIZE` constant to limit batch operations
  - Implemented batched deletion logic with size checks
  - Added validation to prevent operations that exceed gas limits
  - Created comprehensive test suite in `test/TEEHelper_DoSFix.t.sol` (248 lines)
  - Added signer validation tests in `test/SignerValidation.t.sol` (164 lines)
- **Status**: Fixed and tested
- **Reference**: Commit 3026966, PR #43

### 6. Missing TEE Journal Struct Validations
- **Background**: The VerifierJournal struct contains critical fields (PCRs, public key, nonce, timestamp, userData) that must be properly validated.
- **Issue**: Insufficient validation of journal fields could allow malformed or malicious attestations to be accepted.
- **Specific Vulnerabilities**:
  1. **Empty PCR array**: No validation that PCR array contains data
  2. **Invalid public key format**: Missing check for correct public key length (65 bytes) and format (uncompressed, starting with 0x04)
  3. **Predictable addresses**: Malformed public keys could lead to predictable or exploitable signer addresses
- **Fix Applied**:
  - Added `_validateJournal()` internal function with comprehensive checks:
    - `require(journal.pcrs.length > 0, "PCR array cannot be empty")`
    - `require(journal.publicKey.length == 65, "Invalid public key length")`
    - `require(journal.publicKey[0] == 0x04, "Public key must be uncompressed")`
  - Created test suite in `test/JournalValidation.t.sol` (212 lines) testing:
    - Valid attestation passes all validations
    - Old attestations are rejected (timestamp validation)
    - Future attestations are rejected (timestamp validation)
- **Status**: Fixed and tested
- **Reference**: Commits c47d9aa, PR #43

**Overall Status**: All three vulnerabilities in PR #43 have been addressed with fixes and comprehensive test coverage.
