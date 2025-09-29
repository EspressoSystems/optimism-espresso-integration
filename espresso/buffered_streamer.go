package espresso

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// BufferedEspressoStreamer is a wrapper around EspressoStreamerIFace that
// buffers batches to avoid repeated calls to the underlying streamer.
//
// This structure is meant to help the underlying streamer avoid getting
// reset too frequently.  This has primarily been added as an in-between
// layer for the Batch, which seems to need to rewind constantly, which is
// not great for the EspressoStreamer which wants to only progress forward
// and not rewind.
//
// The general idea is to take advantage that we should have a safe starting
// position for the batches being reported to the streamer that is being
// updated frequently.
//
// We can use this safe starting position to store a buffer as needed to store
// all batches  from the safe position to whatever the current latest batch is.
// This allows us to avoid needing to rewind the streamer, and instead just
// adjust the read position of the buffered streamer.
type BufferedEspressoStreamer[B Batch] struct {
	streamer EspressoStreamer[B]

	batches []*B

	// local offset
	readPos uint64

	startingBatchPos    uint64
	currentSafeL1Origin eth.BlockID
}

// Compile time assertion to ensure BufferedEspressoStreamer implements
// EspressoStreamerIFace
var _ EspressoStreamer[Batch] = (*BufferedEspressoStreamer[Batch])(nil)

// NewBufferedEspressoStreamer creates a new BufferedEspressoStreamer instance.
func NewBufferedEspressoStreamer[B Batch](streamer EspressoStreamer[B]) *BufferedEspressoStreamer[B] {
	return &BufferedEspressoStreamer[B]{
		streamer: streamer,
	}
}

// Update implements EspressoStreamerIFace
func (b *BufferedEspressoStreamer[B]) Update(ctx context.Context) error {
	return b.streamer.Update(ctx)
}

// handleL2PositionUpdate handles the update of the L2 position for the
// buffered streamer.
//
// There are three conditions to consider:
//  1. If the next position is before the starting batch position, we need to
//     reset the underlying streamer, and dump our local buffer, as this
//     indicates a need to move backwards before our earliest known batch.
//  2. If the next position is after our starting batch position, then we
//     can drop all earlier stored batches in our buffer, and adjust our
//     read position accordingly.  This should appear to the consumer as nothing
//     has changed progression-wise, but it allows us to reclaim memory.
//  3. If the next position is the same as our starting batch position, then
//     we do nothing, as we are already at the correct position.
func (b *BufferedEspressoStreamer[B]) handleL2PositionUpdate(nextPosition uint64) {
	if nextPosition < b.startingBatchPos {
		// If the next position is before the starting batch position,
		// we need to reset the buffered streamer to ensure we don't
		// miss any batches.
		b.readPos = 0
		b.startingBatchPos = nextPosition
		b.batches = make([]*B, 0)
		b.streamer.Reset()
		return
	}

	if nextPosition > b.startingBatchPos {
		// We want to advance the read position, and we are indicating that
		// we no longer will need to refer to older batches.  So instead, we
		// will want to adjust the buffer, and read position based on the
		// new nextPosition.

		positionAdjustment := nextPosition - b.startingBatchPos
		if positionAdjustment <= uint64(len(b.batches)) {
			// If the adjustment is within the bounds of the current buffer,
			// we can simply adjust the read position and starting batch position.
			b.batches = b.batches[positionAdjustment:]
			b.readPos -= positionAdjustment
		} else {
			b.batches = make([]*B, 0)
			b.readPos = 0
		}
		b.startingBatchPos = nextPosition
		return
	}
}

// RefreshSafeL1Origin updates the safe L1 origin for the buffered streamer.
// This method attempts to safely handle the adjustment of the safeL1Origin
// without needing to defer to the underlying streamer unless necessary.
func (b *BufferedEspressoStreamer[B]) RefreshSafeL1Origin(safeL1Origin eth.BlockID) error {
	if safeL1Origin.Number < b.currentSafeL1Origin.Number {
		// If the safeL1Origin is before the starting batch position, we need to
		// reset the buffered streamer to ensure we don't miss any batches.
		b.currentSafeL1Origin = safeL1Origin
		b.startingBatchPos = 0
		b.readPos = 0
		b.batches = make([]*B, 0)
		if cast, castOk := b.streamer.(interface{ RefreshSafeL1Origin(eth.BlockID) error }); castOk {
			// If the underlying streamer has a method to refresh the safe L1 origin,
			// we call it to ensure it is aware of the new safe L1 origin.
			return cast.RefreshSafeL1Origin(safeL1Origin)
		}
		return nil
	}

	b.currentSafeL1Origin = safeL1Origin
	return nil
}

// Refresh implements EspressoStreamerIFace
func (b *BufferedEspressoStreamer[B]) Refresh(ctx context.Context, finalizedL1 eth.L1BlockRef, safeBatchNumber uint64, safeL1Origin eth.BlockID) error {
	b.handleL2PositionUpdate(safeBatchNumber)
	if err := b.RefreshSafeL1Origin(safeL1Origin); err != nil {
		return err
	}

	return b.streamer.Refresh(ctx, finalizedL1, safeBatchNumber, safeL1Origin)
}

// Reset resets the buffered streamer state to the last known good
// safe batch position.
func (b *BufferedEspressoStreamer[B]) Reset() {
	// Reset the buffered streamer state
	b.readPos = 0
}

// HasNext implements EspressoStreamerIFace
//
// It checks to see if there are any batches left to read in its local buffer.
// If there are no batches left in the buffer, it defers to the underlying
// streamer to determine if there are more batches available.
func (b *BufferedEspressoStreamer[B]) HasNext(ctx context.Context) bool {
	if b.readPos < uint64(len(b.batches)) {
		return true
	}

	return b.streamer.HasNext(ctx)
}

// Next implements EspressoStreamerIFace
//
// It returns the next batch from the local buffer if available, or fetches
// it from the underlying streamer if not, appending to its local underlying
// buffer in the process.
func (b *BufferedEspressoStreamer[B]) Next(ctx context.Context) *B {
	if b.readPos < uint64(len(b.batches)) {
		// If we have a batch in the buffer, return it
		batch := b.batches[b.readPos]
		b.readPos++
		return batch
	}

	// If we don't have a batch in the buffer, fetch the next one from the streamer
	batch := b.streamer.Next(ctx)

	// No more batches available at the moment
	if batch == nil {
		return nil
	}

	number := (*batch).Number()
	if number < b.startingBatchPos {
		// If the batch number is before the starting batch position, we ignore
		// it, and want to fetch the next one
		return b.Next(ctx)
	}

	b.batches = append(b.batches, batch)
	b.readPos++
	return batch

}

// UnmarshalBatch implements EspressoStreamerIFace
func (b *BufferedEspressoStreamer[B]) UnmarshalBatch(data []byte) (*B, error) {
	// Delegate the unmarshalling to the underlying streamer
	return b.streamer.UnmarshalBatch(data)
}
