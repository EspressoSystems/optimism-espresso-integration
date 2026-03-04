package derive

import (
	"context"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
)

func TestComputeCalldataBatchHash(t *testing.T) {
	data := []byte("hello world")
	hash := ComputeCalldataBatchHash(data)
	expected := crypto.Keccak256Hash(data)
	require.Equal(t, expected, hash)
}

func TestComputeCalldataBatchHashEmpty(t *testing.T) {
	hash := ComputeCalldataBatchHash([]byte{})
	expected := crypto.Keccak256Hash([]byte{})
	require.Equal(t, expected, hash)
}

func TestComputeBlobBatchHash(t *testing.T) {
	h1 := common.HexToHash("0x0100000000000000000000000000000000000000000000000000000000000001")
	h2 := common.HexToHash("0x0100000000000000000000000000000000000000000000000000000000000002")

	hash := ComputeBlobBatchHash([]common.Hash{h1, h2})

	// Manually compute expected: keccak256(h1 ++ h2)
	concatenated := make([]byte, 64)
	copy(concatenated[0:32], h1[:])
	copy(concatenated[32:64], h2[:])
	expected := crypto.Keccak256Hash(concatenated)
	require.Equal(t, expected, hash)
}

func TestComputeBlobBatchHashSingle(t *testing.T) {
	h := common.HexToHash("0xabcdef")
	hash := ComputeBlobBatchHash([]common.Hash{h})
	expected := crypto.Keccak256Hash(h[:])
	require.Equal(t, expected, hash)
}

func TestFindBatchAuthEvent(t *testing.T) {
	authenticatorAddr := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	batchHash := crypto.Keccak256Hash([]byte("test batch data"))

	t.Run("event found", func(t *testing.T) {
		receipts := types.Receipts{
			{
				Status: types.ReceiptStatusSuccessful,
				Logs: []*types.Log{
					{
						Address: authenticatorAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							batchHash,
						},
					},
				},
			},
		}
		require.True(t, FindBatchAuthEvent(receipts, authenticatorAddr, batchHash))
	})

	t.Run("event not found - wrong hash", func(t *testing.T) {
		wrongHash := crypto.Keccak256Hash([]byte("wrong data"))
		receipts := types.Receipts{
			{
				Status: types.ReceiptStatusSuccessful,
				Logs: []*types.Log{
					{
						Address: authenticatorAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							wrongHash,
						},
					},
				},
			},
		}
		require.False(t, FindBatchAuthEvent(receipts, authenticatorAddr, batchHash))
	})

	t.Run("event not found - wrong address", func(t *testing.T) {
		wrongAddr := common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")
		receipts := types.Receipts{
			{
				Status: types.ReceiptStatusSuccessful,
				Logs: []*types.Log{
					{
						Address: wrongAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							batchHash,
						},
					},
				},
			},
		}
		require.False(t, FindBatchAuthEvent(receipts, authenticatorAddr, batchHash))
	})

	t.Run("event not found - reverted receipt", func(t *testing.T) {
		receipts := types.Receipts{
			{
				Status: types.ReceiptStatusFailed,
				Logs: []*types.Log{
					{
						Address: authenticatorAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							batchHash,
						},
					},
				},
			},
		}
		require.False(t, FindBatchAuthEvent(receipts, authenticatorAddr, batchHash))
	})

	t.Run("event not found - empty receipts", func(t *testing.T) {
		require.False(t, FindBatchAuthEvent(types.Receipts{}, authenticatorAddr, batchHash))
	})

	t.Run("event found among multiple receipts", func(t *testing.T) {
		receipts := types.Receipts{
			{
				Status: types.ReceiptStatusSuccessful,
				Logs: []*types.Log{
					{
						Address: common.HexToAddress("0x1111"),
						Topics:  []common.Hash{common.HexToHash("0xdead")},
					},
				},
			},
			{
				Status: types.ReceiptStatusSuccessful,
				Logs: []*types.Log{
					{
						Address: authenticatorAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							batchHash,
						},
					},
				},
			},
		}
		require.True(t, FindBatchAuthEvent(receipts, authenticatorAddr, batchHash))
	})
}

// buildL1Chain creates a chain of L1BlockRef values with proper parent-hash linkage.
// The chain goes from block number `start` to `end` (inclusive).
// Returns a slice indexed by block number (relative to start), and the full map by number.
func buildL1Chain(rng *rand.Rand, start, end uint64) map[uint64]eth.L1BlockRef {
	chain := make(map[uint64]eth.L1BlockRef)
	for num := start; num <= end; num++ {
		ref := eth.L1BlockRef{
			Number: num,
			Hash:   testutils.RandomHash(rng),
		}
		if num > start {
			ref.ParentHash = chain[num-1].Hash
		}
		chain[num] = ref
	}
	return chain
}

func TestCollectAuthenticatedBatches(t *testing.T) {
	logger := testlog.Logger(t, log.LevelDebug)
	ctx := context.Background()
	rng := rand.New(rand.NewSource(1234))

	authenticatorAddr := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	batchHash := crypto.Keccak256Hash([]byte("test batch data"))

	// Build a matching receipt
	matchingReceipts := types.Receipts{
		{
			Status: types.ReceiptStatusSuccessful,
			Logs: []*types.Log{
				{
					Address: authenticatorAddr,
					Topics: []common.Hash{
						BatchInfoAuthenticatedABIHash,
						batchHash,
					},
				},
			},
		},
	}
	emptyReceipts := types.Receipts{}

	// expectChainTraversal sets up mock expectations for a backward parent-hash
	// traversal from chain[end] down to chain[start]. For each block it expects
	// FetchReceipts (by hash), and for all blocks except the first (end) it
	// expects L1BlockRefByHash to resolve the parent hash.
	// receiptsByBlock allows overriding receipts for specific block numbers.
	expectChainTraversal := func(l1F *testutils.MockL1Source, chain map[uint64]eth.L1BlockRef, start, end uint64, receiptsByBlock map[uint64]types.Receipts) {
		for num := end; num >= start; num-- {
			ref := chain[num]
			receipts := emptyReceipts
			if r, ok := receiptsByBlock[num]; ok {
				receipts = r
			}
			l1F.ExpectFetchReceipts(ref.Hash, nil, receipts, nil)
			// L1BlockRefByHash is called for every block except the first one (ref itself)
			if num > start {
				l1F.ExpectL1BlockRefByHash(chain[num-1].Hash, chain[num-1], nil)
			}
			if num == 0 {
				break // avoid underflow
			}
		}
	}

	t.Run("found in same block", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		chain := buildL1Chain(rng, 100, 200)
		ref := chain[200]

		// Auth event is in block 200 (same block as ref). Traversal goes 200 -> 100.
		expectChainTraversal(l1F, chain, 100, 200, map[uint64]types.Receipts{
			200: matchingReceipts,
		})

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.True(t, result[batchHash])
		require.Len(t, result, 1)
		l1F.AssertExpectations(t)
	})

	t.Run("found in earliest block of window", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		chain := buildL1Chain(rng, 100, 200)
		ref := chain[200]

		// Auth event is in block 100 (last block of the lookback window).
		expectChainTraversal(l1F, chain, 100, 200, map[uint64]types.Receipts{
			100: matchingReceipts,
		})

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.True(t, result[batchHash])
		require.Len(t, result, 1)
		l1F.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		chain := buildL1Chain(rng, 100, 200)
		ref := chain[200]

		// No auth event in any block in the window
		expectChainTraversal(l1F, chain, 100, 200, nil)

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.Len(t, result, 0)
		l1F.AssertExpectations(t)
	})

	t.Run("low block number - window clamps to 0", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		chain := buildL1Chain(rng, 0, 10)
		ref := chain[10]

		// Window should clamp to [0, 10]. Auth event is in block 10.
		expectChainTraversal(l1F, chain, 0, 10, map[uint64]types.Receipts{
			10: matchingReceipts,
		})

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.True(t, result[batchHash])
		require.Len(t, result, 1)
		l1F.AssertExpectations(t)
	})

	t.Run("multiple hashes collected", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		chain := buildL1Chain(rng, 0, 10)
		ref := chain[10]

		batchHash2 := crypto.Keccak256Hash([]byte("second batch"))
		multiReceipts := types.Receipts{
			{
				Status: types.ReceiptStatusSuccessful,
				Logs: []*types.Log{
					{
						Address: authenticatorAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							batchHash,
						},
					},
					{
						Address: authenticatorAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							batchHash2,
						},
					},
				},
			},
		}

		// Both auth events are in block 10
		expectChainTraversal(l1F, chain, 0, 10, map[uint64]types.Receipts{
			10: multiReceipts,
		})

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.Len(t, result, 2)
		require.True(t, result[batchHash])
		require.True(t, result[batchHash2])
		l1F.AssertExpectations(t)
	})
}

func TestBatchInfoAuthenticatedABIHash(t *testing.T) {
	// Verify the ABI hash matches what Solidity would compute
	expected := crypto.Keccak256Hash([]byte("BatchInfoAuthenticated(bytes32)"))
	require.Equal(t, expected, BatchInfoAuthenticatedABIHash)
}
