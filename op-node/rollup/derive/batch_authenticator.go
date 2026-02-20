package derive

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// BatchAuthLookbackWindow defines how many L1 blocks before the batch submission
// to scan for a BatchInfoAuthenticated event. The authentication transaction must
// land in this window (or in the same block as the batch submission) for the batch
// to be considered valid.
//
// At ~12s per L1 block, 100 blocks ≈ 20 minutes. This gives the batcher ample time
// to land the batch data transaction on L1 after the authentication transaction,
// even under moderate L1 congestion or batcher restarts. The window is intentionally
// generous: a tighter window risks rejecting valid batches during congestion spikes,
// while a wider window only increases the receipt scan range (mitigated by the
// CachingReceiptsProvider LRU cache).
const BatchAuthLookbackWindow uint64 = 100

var (
	// BatchInfoAuthenticatedABI is the event signature for BatchInfoAuthenticated(bytes32 indexed commitment, address indexed signer).
	BatchInfoAuthenticatedABI     = "BatchInfoAuthenticated(bytes32,address)"
	BatchInfoAuthenticatedABIHash = crypto.Keccak256Hash([]byte(BatchInfoAuthenticatedABI))
)

// ComputeCalldataBatchHash computes keccak256(calldata), matching the BatchAuthenticator
// contract's calldata batch validation path.
func ComputeCalldataBatchHash(data []byte) common.Hash {
	return crypto.Keccak256Hash(data)
}

// ComputeBlobBatchHash computes keccak256(concat(blobHashes)), matching the BatchAuthenticator
// contract's blob batch validation path.
func ComputeBlobBatchHash(blobHashes []common.Hash) common.Hash {
	concatenated := make([]byte, 32*len(blobHashes))
	for i, h := range blobHashes {
		copy(concatenated[i*32:(i+1)*32], h[:])
	}
	return crypto.Keccak256Hash(concatenated)
}

// FindBatchAuthEvent scans the given receipts for a BatchInfoAuthenticated event
// emitted by authenticatorAddr with a commitment matching batchHash.
// Returns true if such an event is found.
func FindBatchAuthEvent(receipts types.Receipts, authenticatorAddr common.Address, batchHash common.Hash) bool {
	for _, receipt := range receipts {
		if receipt.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for _, lg := range receipt.Logs {
			if lg.Address != authenticatorAddr {
				continue
			}
			// BatchInfoAuthenticated has 3 topics: event sig, indexed commitment, indexed signer
			if len(lg.Topics) >= 2 &&
				lg.Topics[0] == BatchInfoAuthenticatedABIHash &&
				lg.Topics[1] == batchHash {
				return true
			}
		}
	}
	return false
}

// collectAuthEventsFromReceipts extracts all authenticated batch hashes from the given receipts.
// It returns the set of commitment hashes that have been authenticated by the given authenticator.
func collectAuthEventsFromReceipts(receipts types.Receipts, authenticatorAddr common.Address) map[common.Hash]bool {
	result := make(map[common.Hash]bool)
	for _, receipt := range receipts {
		if receipt.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for _, lg := range receipt.Logs {
			if lg.Address != authenticatorAddr {
				continue
			}
			if len(lg.Topics) >= 2 && lg.Topics[0] == BatchInfoAuthenticatedABIHash {
				result[lg.Topics[1]] = true
			}
		}
	}
	return result
}

// CollectAuthenticatedBatches scans L1 receipts in the range
// [ref.Number - BatchAuthLookbackWindow, ref.Number] and returns the set of all
// batch commitment hashes that were authenticated via BatchInfoAuthenticated events.
//
// This is called once per L1 block by the data source, and the returned set is checked
// against each candidate batch transaction. This avoids rescanning the lookback window
// for every individual batch transaction.
//
// Using event scanning (rather than L1 contract state reads) keeps the derivation
// pipeline compatible with the op-program fault proof environment, which can only
// access L1 block headers, transactions, receipts, and blobs.
func CollectAuthenticatedBatches(
	ctx context.Context,
	fetcher L1Fetcher,
	ref eth.L1BlockRef,
	authenticatorAddr common.Address,
	logger log.Logger,
) (map[common.Hash]bool, error) {
	startBlock := ref.Number
	if startBlock > BatchAuthLookbackWindow {
		startBlock = ref.Number - BatchAuthLookbackWindow
	} else {
		startBlock = 0
	}

	allAuthenticated := make(map[common.Hash]bool)
	for blockNum := startBlock; blockNum <= ref.Number; blockNum++ {
		blockRef, err := fetcher.L1BlockRefByNumber(ctx, blockNum)
		if err != nil {
			return nil, NewTemporaryError(fmt.Errorf("batch auth: failed to fetch L1 block ref %d: %w", blockNum, err))
		}
		_, receipts, err := fetcher.FetchReceipts(ctx, blockRef.Hash)
		if err != nil {
			return nil, NewTemporaryError(fmt.Errorf("batch auth: failed to fetch receipts for block %d: %w", blockNum, err))
		}
		for h := range collectAuthEventsFromReceipts(receipts, authenticatorAddr) {
			allAuthenticated[h] = true
		}
	}

	if len(allAuthenticated) > 0 {
		logger.Debug("collected authenticated batches from lookback window",
			"count", len(allAuthenticated), "startBlock", startBlock, "endBlock", ref.Number)
	}
	return allAuthenticated, nil
}
