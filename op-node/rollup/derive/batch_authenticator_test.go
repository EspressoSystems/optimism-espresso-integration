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
	signerAddr := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")

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
							common.BytesToHash(signerAddr.Bytes()),
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
							common.BytesToHash(signerAddr.Bytes()),
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
							common.BytesToHash(signerAddr.Bytes()),
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
							common.BytesToHash(signerAddr.Bytes()),
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
							common.BytesToHash(signerAddr.Bytes()),
						},
					},
				},
			},
		}
		require.True(t, FindBatchAuthEvent(receipts, authenticatorAddr, batchHash))
	})
}

func TestCollectAuthenticatedBatches(t *testing.T) {
	logger := testlog.Logger(t, log.LevelDebug)
	ctx := context.Background()
	rng := rand.New(rand.NewSource(1234))

	authenticatorAddr := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	batchHash := crypto.Keccak256Hash([]byte("test batch data"))
	signerAddr := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")

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
						common.BytesToHash(signerAddr.Bytes()),
					},
				},
			},
		},
	}
	emptyReceipts := types.Receipts{}

	t.Run("found in same block", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		ref := eth.L1BlockRef{Number: 200, Hash: common.HexToHash("0xaa")}

		// We scan from block 100 (200 - 100) to 200. Auth event is in block 200 (same block).
		for blockNum := uint64(100); blockNum < 200; blockNum++ {
			blockRef := eth.L1BlockRef{Number: blockNum, Hash: testutils.RandomHash(rng)}
			l1F.ExpectL1BlockRefByNumber(blockNum, blockRef, nil)
			l1F.ExpectFetchReceipts(blockRef.Hash, nil, emptyReceipts, nil)
		}
		// Block 200 has the matching event
		l1F.ExpectL1BlockRefByNumber(200, ref, nil)
		l1F.ExpectFetchReceipts(ref.Hash, nil, matchingReceipts, nil)

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.True(t, result[batchHash])
		require.Len(t, result, 1)
		l1F.AssertExpectations(t)
	})

	t.Run("found in earlier block", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		ref := eth.L1BlockRef{Number: 200, Hash: common.HexToHash("0xaa")}

		// Auth event is in block 100 (first block of window). CollectAuthenticatedBatches
		// always scans the full window, so we need expectations for all blocks.
		for blockNum := uint64(100); blockNum <= 200; blockNum++ {
			var blockRef eth.L1BlockRef
			if blockNum == 100 {
				blockRef = eth.L1BlockRef{Number: 100, Hash: testutils.RandomHash(rng)}
				l1F.ExpectL1BlockRefByNumber(blockNum, blockRef, nil)
				l1F.ExpectFetchReceipts(blockRef.Hash, nil, matchingReceipts, nil)
			} else if blockNum == 200 {
				l1F.ExpectL1BlockRefByNumber(blockNum, ref, nil)
				l1F.ExpectFetchReceipts(ref.Hash, nil, emptyReceipts, nil)
			} else {
				blockRef = eth.L1BlockRef{Number: blockNum, Hash: testutils.RandomHash(rng)}
				l1F.ExpectL1BlockRefByNumber(blockNum, blockRef, nil)
				l1F.ExpectFetchReceipts(blockRef.Hash, nil, emptyReceipts, nil)
			}
		}

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.True(t, result[batchHash])
		require.Len(t, result, 1)
		l1F.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		ref := eth.L1BlockRef{Number: 200, Hash: common.HexToHash("0xaa")}

		// No auth event in any block in the window
		for blockNum := uint64(100); blockNum <= 200; blockNum++ {
			blockRef := eth.L1BlockRef{Number: blockNum, Hash: testutils.RandomHash(rng)}
			l1F.ExpectL1BlockRefByNumber(blockNum, blockRef, nil)
			l1F.ExpectFetchReceipts(blockRef.Hash, nil, emptyReceipts, nil)
		}

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.Len(t, result, 0)
		l1F.AssertExpectations(t)
	})

	t.Run("low block number - window clamps to 0", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		ref := eth.L1BlockRef{Number: 10, Hash: common.HexToHash("0xaa")}

		// Window should be [0, 10]
		for blockNum := uint64(0); blockNum < 10; blockNum++ {
			blockRef := eth.L1BlockRef{Number: blockNum, Hash: testutils.RandomHash(rng)}
			l1F.ExpectL1BlockRefByNumber(blockNum, blockRef, nil)
			l1F.ExpectFetchReceipts(blockRef.Hash, nil, emptyReceipts, nil)
		}
		l1F.ExpectL1BlockRefByNumber(10, ref, nil)
		l1F.ExpectFetchReceipts(ref.Hash, nil, matchingReceipts, nil)

		result, err := CollectAuthenticatedBatches(ctx, l1F, ref, authenticatorAddr, logger)
		require.NoError(t, err)
		require.True(t, result[batchHash])
		require.Len(t, result, 1)
		l1F.AssertExpectations(t)
	})

	t.Run("multiple hashes collected", func(t *testing.T) {
		l1F := &testutils.MockL1Source{}
		ref := eth.L1BlockRef{Number: 10, Hash: common.HexToHash("0xaa")}

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
							common.BytesToHash(signerAddr.Bytes()),
						},
					},
					{
						Address: authenticatorAddr,
						Topics: []common.Hash{
							BatchInfoAuthenticatedABIHash,
							batchHash2,
							common.BytesToHash(signerAddr.Bytes()),
						},
					},
				},
			},
		}

		for blockNum := uint64(0); blockNum < 10; blockNum++ {
			blockRef := eth.L1BlockRef{Number: blockNum, Hash: testutils.RandomHash(rng)}
			l1F.ExpectL1BlockRefByNumber(blockNum, blockRef, nil)
			l1F.ExpectFetchReceipts(blockRef.Hash, nil, emptyReceipts, nil)
		}
		l1F.ExpectL1BlockRefByNumber(10, ref, nil)
		l1F.ExpectFetchReceipts(ref.Hash, nil, multiReceipts, nil)

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
	expected := crypto.Keccak256Hash([]byte("BatchInfoAuthenticated(bytes32,address)"))
	require.Equal(t, expected, BatchInfoAuthenticatedABIHash)
}
