package espresso

import (
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// mockBatch is a simple implementation of the Batch interface for testing
type mockBatch struct {
	number   uint64
	hash     common.Hash
	l1Origin eth.BlockID
}

func (m mockBatch) Number() uint64 {
	return m.number
}

func (m mockBatch) L1Origin() eth.BlockID {
	return m.l1Origin
}

func (m mockBatch) Header() *types.Header {
	return &types.Header{
		Number: big.NewInt(int64(m.number)),
	}
}

func (m mockBatch) Hash() common.Hash {
	return m.hash
}

// newMockBatch creates a mock batch with the given number and a hash derived from the number
func newMockBatch(number uint64) mockBatch {
	return mockBatch{
		number: number,
		hash:   common.BigToHash(big.NewInt(int64(number))),
		l1Origin: eth.BlockID{
			Number: number,
			Hash:   common.BigToHash(big.NewInt(int64(number))),
		},
	}
}

// newMockBatchWithHash creates a mock batch with a specific number and hash
func newMockBatchWithHash(number uint64, hash common.Hash) mockBatch {
	return mockBatch{
		number: number,
		hash:   hash,
		l1Origin: eth.BlockID{
			Number: number,
			Hash:   common.BigToHash(big.NewInt(int64(number))),
		},
	}
}

// TestBatchBufferInsertAndRetrieve verifies basic insert and retrieval behavior.
// Note: the buffer no longer enforces a capacity limit internally. Bounding is
// done by the streamer, which drops batches too far ahead of the current position
// (see MaxBatchOutOfOrder). The buffer itself accepts any number of inserts.
func TestBatchBufferInsertAndRetrieve(t *testing.T) {
	buffer := NewBatchBuffer[mockBatch](0)

	// Verify buffer starts empty
	require.Equal(t, 0, buffer.Len())

	// Insert batches
	batch1 := newMockBatch(1)
	batch2 := newMockBatch(2)
	batch3 := newMockBatch(3)
	batch4 := newMockBatch(4)

	err := buffer.Insert(batch1)
	require.NoError(t, err)
	require.Equal(t, 1, buffer.Len())

	err = buffer.Insert(batch2)
	require.NoError(t, err)
	require.Equal(t, 2, buffer.Len())

	err = buffer.Insert(batch3)
	require.NoError(t, err)
	require.Equal(t, 3, buffer.Len())

	// Inserting a fourth batch succeeds (no capacity limit)
	err = buffer.Insert(batch4)
	require.NoError(t, err)
	require.Equal(t, 4, buffer.Len())

	// Verify all batches are accessible and in sorted order
	for i := 0; i < 4; i++ {
		got := buffer.Get(i)
		require.NotNil(t, got)
		require.Equal(t, uint64(i+1), got.Number())
	}

	// Verify Get returns nil for out of bounds
	require.Nil(t, buffer.Get(4))
}

// TestBatchBufferInsertDuplicateHandling verifies that:
// - Inserting the exact same batch (same number AND same hash) does not create a duplicate
// - Inserting a batch with the same number but different hash IS allowed
func TestBatchBufferInsertDuplicateHandling(t *testing.T) {
	const batchNumberN uint64 = 42

	buffer := NewBatchBuffer[mockBatch](0)

	// Create first batch with number N and hash H1
	hashH1 := common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	batchH1 := newMockBatchWithHash(batchNumberN, hashH1)

	// Insert first batch
	err := buffer.Insert(batchH1)
	require.NoError(t, err)
	require.Equal(t, 1, buffer.Len())

	// Insert the exact same batch again (same number N, same hash H1)
	// This should return ErrDuplicateBatch and not create a duplicate
	err = buffer.Insert(batchH1)
	require.ErrorIs(t, err, ErrDuplicateBatch)
	require.Equal(t, 1, buffer.Len(), "duplicate batch with same number and hash should not be inserted")

	// Create a different batch with same number N but different hash H2
	hashH2 := common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222")
	batchH2 := newMockBatchWithHash(batchNumberN, hashH2)

	// Insert batch with same number but different hash - should be allowed
	err = buffer.Insert(batchH2)
	require.NoError(t, err)
	require.Equal(t, 2, buffer.Len(), "batch with same number but different hash should be inserted")

	// Verify both batches can be retrieved
	first := buffer.Get(0)
	require.NotNil(t, first)

	second := buffer.Get(1)
	require.NotNil(t, second)

	// Verify they both have the same batch number
	require.Equal(t, batchNumberN, first.Number())
	require.Equal(t, batchNumberN, second.Number())

	// Verify they have different hashes
	require.NotEqual(t, first.Hash(), second.Hash())

	// Verify insertion order is preserved (H1 first, H2 second)
	require.Equal(t, hashH1, first.Hash())
	require.Equal(t, hashH2, second.Hash())
}

// TestBatchBufferPeekAndPop verifies Peek returns without removing and Pop removes
func TestBatchBufferPeekAndPop(t *testing.T) {
	buffer := NewBatchBuffer[mockBatch](10)

	// Verify Peek on empty buffer returns nil
	require.Nil(t, buffer.Peek())

	// Verify Pop on empty buffer returns nil
	require.Nil(t, buffer.Pop())

	// Insert a batch
	batch1 := newMockBatch(1)
	err := buffer.Insert(batch1)
	require.NoError(t, err)

	// Peek should return the batch without removing
	peeked := buffer.Peek()
	require.NotNil(t, peeked)
	require.Equal(t, uint64(1), peeked.Number())
	require.Equal(t, 1, buffer.Len())

	// Peek again should return the same batch
	peekedAgain := buffer.Peek()
	require.Equal(t, peeked.Number(), peekedAgain.Number())
	require.Equal(t, peeked.Hash(), peekedAgain.Hash())

	// Pop should return and remove the batch
	popped := buffer.Pop()
	require.NotNil(t, popped)
	require.Equal(t, uint64(1), popped.Number())
	require.Equal(t, 0, buffer.Len())

	// Pop on now-empty buffer should return nil
	require.Nil(t, buffer.Pop())
}

// TestBatchBufferSortedOrder verifies batches are stored in sorted order by batch number
func TestBatchBufferSortedOrder(t *testing.T) {
	buffer := NewBatchBuffer[mockBatch](10)

	// Insert batches out of order
	err := buffer.Insert(newMockBatch(5))
	require.NoError(t, err)
	err = buffer.Insert(newMockBatch(2))
	require.NoError(t, err)
	err = buffer.Insert(newMockBatch(8))
	require.NoError(t, err)
	err = buffer.Insert(newMockBatch(1))
	require.NoError(t, err)

	require.Equal(t, 4, buffer.Len())

	// Verify Get returns them in sorted order
	require.Equal(t, uint64(1), buffer.Get(0).Number())
	require.Equal(t, uint64(2), buffer.Get(1).Number())
	require.Equal(t, uint64(5), buffer.Get(2).Number())
	require.Equal(t, uint64(8), buffer.Get(3).Number())

	// Verify Pop returns them in sorted order
	require.Equal(t, uint64(1), buffer.Pop().Number())
	require.Equal(t, uint64(2), buffer.Pop().Number())
	require.Equal(t, uint64(5), buffer.Pop().Number())
	require.Equal(t, uint64(8), buffer.Pop().Number())

	// Buffer should be empty now
	require.Equal(t, 0, buffer.Len())
}

// TestBatchBufferClear verifies Clear removes all batches
func TestBatchBufferClear(t *testing.T) {
	buffer := NewBatchBuffer[mockBatch](10)

	// Insert some batches
	err := buffer.Insert(newMockBatch(1))
	require.NoError(t, err)
	err = buffer.Insert(newMockBatch(2))
	require.NoError(t, err)
	err = buffer.Insert(newMockBatch(3))
	require.NoError(t, err)
	require.Equal(t, 3, buffer.Len())

	// Clear the buffer
	buffer.Clear()

	// Verify buffer is empty
	require.Equal(t, 0, buffer.Len())
	require.Nil(t, buffer.Peek())
	require.Nil(t, buffer.Pop())
	require.Nil(t, buffer.Get(0))

	// Verify we can insert again after clear
	err = buffer.Insert(newMockBatch(1))
	require.NoError(t, err)
	require.Equal(t, 1, buffer.Len())
}

// TestBatchBufferGetOutOfBounds verifies Get returns nil for invalid indices
func TestBatchBufferGetOutOfBounds(t *testing.T) {
	buffer := NewBatchBuffer[mockBatch](10)

	// Empty buffer - all indices should return nil
	require.Nil(t, buffer.Get(0))
	require.Nil(t, buffer.Get(1))

	// Insert one batch
	err := buffer.Insert(newMockBatch(1))
	require.NoError(t, err)

	// Valid index
	require.NotNil(t, buffer.Get(0))

	// Invalid indices
	require.Nil(t, buffer.Get(1))
	require.Nil(t, buffer.Get(100))
}
