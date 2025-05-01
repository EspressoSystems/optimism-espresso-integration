package environment_test

import (
	"context"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/stretchr/testify/require"
)

// TestBatcherWaitForFinality is a test that attempts to make sure that the batcher waits for the
// derived L1 block to be finalized before submitting a new block.
//
// This tests is designed to evaluate Test 8.1.1 as outlined within the Espresso Celo Integration
// plan. It has stated task definition as follows:
//
//	Arrange:
//		Run the sequencer and the batcher in Espresso mode.
//	Act:
//		Wait until a new block is finalized.
//	Assert:
//		The batcher doesn't submit a block without finalized L1 origin to the L1.
//		After the L1 origin is finalized, the batcher submits the block.
func TestBatcherWaitForFinality(t *testing.T) {
	// Basic test setup.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}
	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	caffNode, err := env.LaunchDecaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}
	defer env.Stop(t, caffNode)

	rollupClient := system.RollupClient("verifier")

	// Verify that the batcher waits for the L1 origin to be finalized before submitting a new
	// block to the L1.
	initialSeqStatus, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)
	initialFinalizedL1Number := initialSeqStatus.FinalizedL1.Number
	initialSafeL1Number := initialSeqStatus.SafeL1.Number
	require.LessOrEqual(t, initialSafeL1Number, initialFinalizedL1Number + 1, "Safe L1 number too large")

	// Verify that eventually a new block will be finalized, which will enable the batcher to
	// submit another block to the L1.
	tickerFinality := time.NewTicker(1 * time.Second)
	defer tickerFinality.Stop()

	for {
		select {
		case <-ctx.Done():
			require.FailNow(t, "Timeout: Finalized L1 number not increased")
		case <-tickerFinality.C:
			seqStatusAfterWait, err := rollupClient.SyncStatus(context.Background())
			require.NoError(t, err)

			// Wait for a new block to be finalized.
			if seqStatusAfterWait.FinalizedL1.Number > initialFinalizedL1Number {
				tickerSubmission := time.NewTicker(1 * time.Second)
				defer tickerSubmission.Stop()

				for {
					select {
					case <-ctx.Done():
						require.FailNow(t, "Timeout: Safe L1 number not increased")
					case <-tickerSubmission.C:
						seqStatusAfterWait, err := rollupClient.SyncStatus(context.Background())
						require.NoError(t, err)

						if seqStatusAfterWait.SafeL1.Number > initialSafeL1Number {
							return
						}
					}
				}
			}
		}
	}
}

// TestCaffNodeWaitForFinality is a test that attempts to make sure that the Caff node waits for
// the derived L1 block to be finalized before updating its record.
//
// This tests is designed to evaluate Test 8.2.1 as outlined within the Espresso Celo Integration
// plan. It has stated task definition as follows:
//
//	Arrange:
//		Run the sequencer and the Caff node in Espresso mode.
//	Act:
//		Wait the Caff node's batch buffer is empty.
//	Assert:
//		The Caff node doesn't update its batch buffer without finalized L1 origin to the L1.
//		After the L1 origin is finalized, the Caff node updates the batch buffer.
func TestCaffNodeWaitForFinality(t *testing.T) {
	// Basic test setup.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}
	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	caffNode, err := env.LaunchDecaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}
	defer env.Stop(t, caffNode)

	rollupClient := system.RollupClient("verifier")
	streamer := caffNode.OpNode.EspressoStreamer()

	// Wait for the batch buffer to be empty which will trigger the Caff node to sync the status.
	tickerBufferClear := time.NewTicker(1 * time.Second)
	defer tickerBufferClear.Stop()

	for {
		select {
		case <-ctx.Done():
			require.FailNow(t, "Timeout: Batch buffer not processed")
		case <-tickerBufferClear.C:
			if streamer.BatchBuffer.Len() == 0 {
				initialSeqStatus, err := rollupClient.SyncStatus(context.Background())
				require.NoError(t, err)
				initialFinalizedL1Number := initialSeqStatus.FinalizedL1.Number

				// Wait for a new block to be finalized on the L1.
				tickerFinality := time.NewTicker(1 * time.Second)
				defer tickerFinality.Stop()

				for {
					select {
					case <-ctx.Done():
						require.FailNow(t, "Timeout: Finalized L1 number not increased")
					case <-tickerFinality.C:
						seqStatusAfterWait, err := rollupClient.SyncStatus(context.Background())
						require.NoError(t, err)

						if seqStatusAfterWait.SafeL1.Number > initialFinalizedL1Number {
							// Verify that eventually the Caff node will update its batch buffer
							// after the L1 origin update.
							tickerBufferInsert := time.NewTicker(1 * time.Second)
							defer tickerBufferInsert.Stop()

							for {
								select {
								case <-ctx.Done():
									require.FailNow(t, "Timeout: Batch buffer not inserted")
								case <-tickerBufferInsert.C:
									if streamer.BatchBuffer.Len() > 0 {
										return
									}
								}
							}
						}

						// Verify that the Caff node doesn't update its batch buffer because the L1
						// origin isn't updated.
						require.Equal(t, streamer.BatchBuffer.Len(), 0, "Batch buffer not empty")
					}
				}
			}
		}
	}
}

