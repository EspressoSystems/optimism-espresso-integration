package espresso_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DummyBatch is a test implementation of the Batch interface
type DummyBatch struct {
	number   uint64
	l1Origin eth.BlockID
}

func (b DummyBatch) Number() uint64 {
	return b.number
}

func (b DummyBatch) L1Origin() eth.BlockID {
	return b.l1Origin
}

func (b DummyBatch) Hash() common.Hash {
	return common.Hash{}
}

func (b DummyBatch) Header() *types.Header {
	return nil
}

// Basic batch buffer test
func FuzzBatchBufferBasic(f *testing.F) {
	// Generate corpus for batch buffer
	f.Add([]byte{0})
	f.Add([]byte{1})
	f.Add([]byte{0, 0, 0, 0, 0, 0, 0, 0})
	f.Add([]byte{1, 0, 0, 0, 0, 0, 0, 0})

	f.Fuzz(func(t *testing.T, data []byte) {
		b := espresso.NewBatchBuffer[DummyBatch]()
		var num uint64
		err := binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &num)

		if err != nil {
			b.Insert(DummyBatch{number: num, l1Origin: eth.BlockID{Number: num}}, 0)
		} else {
			b.Peek()
			b.Pop()
		}
	})
}

// FuzzBatchBuffer tests the BatchBuffer implementation using Go's built-in fuzzer
func FuzzBatchBufferSimple(f *testing.F) {
	// Add some seed corpus
	f.Add(uint64(0))
	f.Add(uint64(1))
	f.Add(uint64(100))
	f.Add(uint64(1000))

	// Fuzz test
	f.Fuzz(func(t *testing.T, num uint64) {
		b := espresso.NewBatchBuffer[DummyBatch]()

		// Test insertion
		b.Insert(DummyBatch{number: num, l1Origin: eth.BlockID{Number: num}}, 0)
		if b.Len() != 1 {
			t.Errorf("Expected buffer length 1, got %d", b.Len())
		}

		// Test peek
		batch := b.Peek()
		if batch == nil {
			t.Fatal("Expected non-nil batch from Peek")
		}
		if batch.Number() != num {
			t.Errorf("Expected batch number %d, got %d", num, batch.Number())
		}

		// Test pop
		batch = b.Pop()
		if batch == nil {
			t.Fatal("Expected non-nil batch from Pop")
		}
		if batch.Number() != num {
			t.Errorf("Expected batch number %d, got %d", num, batch.Number())
		}
		if b.Len() != 0 {
			t.Errorf("Expected empty buffer after Pop, got length %d", b.Len())
		}
	})
}

// FuzzBatchBufferInsertMultiple tests inserting multiple batches into the buffer
func FuzzBatchBufferInsertMultiple(f *testing.F) {
	// Add some seed corpus
	f.Add(uint64(1), uint64(2), uint64(3))
	f.Add(uint64(10), uint64(5), uint64(15))

	// Fuzz test with multiple insertions
	f.Fuzz(func(t *testing.T, num1, num2, num3 uint64) {
		b := espresso.NewBatchBuffer[DummyBatch]()

		// Insert batches
		b.Insert(DummyBatch{number: num1, l1Origin: eth.BlockID{Number: num1}}, 0)
		b.Insert(DummyBatch{number: num2, l1Origin: eth.BlockID{Number: num2}}, 0)
		b.Insert(DummyBatch{number: num3, l1Origin: eth.BlockID{Number: num3}}, 0)

		// Test length
		if b.Len() != 3 {
			t.Errorf("Expected buffer length 3, got %d", b.Len())
		}

		// Test clear
		b.Clear()
		if b.Len() != 0 {
			t.Errorf("Expected empty buffer after Clear, got length %d", b.Len())
		}
	})
}

// FuzzBatchBufferTryInsert tests the TryInsert method
func FuzzBatchBufferTryInsert(f *testing.F) {
	f.Add(uint64(1), uint64(2))
	f.Add(uint64(100), uint64(100))

	f.Fuzz(func(t *testing.T, num1, num2 uint64) {
		b := espresso.NewBatchBuffer[DummyBatch]()

		// Insert first batch
		batch1 := DummyBatch{number: num1, l1Origin: eth.BlockID{Number: num1}}
		pos, _ := b.TryInsert(batch1)
		b.Insert(batch1, pos)

		// Try inserting second batch
		batch2 := DummyBatch{number: num2, l1Origin: eth.BlockID{Number: num2}}
		pos, exists := b.TryInsert(batch2)

		// If numbers are the same, it should detect as already existing
		if num1 == num2 && !exists {
			t.Errorf("Expected duplicate batch to be detected")
		}

		// If not duplicate, insert it
		if !exists {
			b.Insert(batch2, pos)

			// Get the batch at position
			gotBatch := b.Get(pos)
			if gotBatch == nil {
				t.Fatal("Expected non-nil batch from Get")
			}
			if gotBatch.Number() != num2 {
				t.Errorf("Expected batch number %d, got %d", num2, gotBatch.Number())
			}
		}
	})
}
