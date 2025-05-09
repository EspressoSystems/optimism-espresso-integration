package environment_test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/espresso"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/stretchr/testify/require"
)

// VerifyL1OriginFinalized checks whether every batch in the batch buffer has a finalized L1
// origin.
func VerifyL1OriginFinalized(rollupClient *sources.RollupClient, streamer *espresso.EspressoStreamer[derive.EspressoBatch]) bool {
	batch := streamer.BatchBuffer.Pop()
	for batch != nil{
		origin := (batch).L1Origin()
		status, err := rollupClient.SyncStatus(context.Background())
		if err != nil || origin.Number > status.FinalizedL1.Number {
			return false
		}
		batch = streamer.BatchBuffer.Pop()
	}
	return true
}

// VerifyBatchBufferUpdated checks whether the batch buffer is updated before the timeout.
func VerifyBatchBufferUpdated(ctx context.Context, streamer *espresso.EspressoStreamer[derive.EspressoBatch]) bool {
	tickerBufferInsert := time.NewTicker(100 * time.Millisecond)
	defer tickerBufferInsert.Stop()
	for {
		select {
		case <-ctx.Done():
			return false
		case <-tickerBufferInsert.C:
			if streamer.BatchBuffer.Len() > 0 {
				return true
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
//		Wait until the Caff node's batch buffer is empty.
//	Assert:
//		The Caff node doesn't insert a batch without finalized L1 origin to the batch buffer.
//		After the L1 origin is finalized, the Caff node inserts the batch.
func TestCaffNodeWaitForFinality(t *testing.T) {
	// Basic test setup.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	launcher := new(env.EspressoDevNodeLauncherDocker)

	// Set L1FinalizedDistance to nonzero, NonFinalizedProposals to true, and SequencerUseFinalized
	// to false, to make sure we are testing how the Caff node handles the finality.
	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, env.WithL1FinalizedDistance(4), env.WithNonFinalizedProposals(true), env.WithSequencerUseFinalized(false))
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

	rollupClient := system.RollupClient(e2esys.RoleVerif)
	streamer := caffNode.OpNode.EspressoStreamer()

	initialStatus, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)

	// Wait for the batch buffer to be empty which will trigger the Caff node to sync the status
	// and insert more batches to the buffer.
	for {
		if streamer.BatchBuffer.Len() == 0 {
			// Wait for the finalized L1 number and the batch buffer to be updated.
			for {
				if streamer.BatchBuffer.Len() > 0 {
					// Verify that any batch inserted into the batch buffer has a finalized L1
					// origin.
					if !VerifyL1OriginFinalized(rollupClient, streamer) {
						require.FailNow(t, "Timeout: L1 origin not finalized")
					}
				} else {
					statusAfterWait, err := rollupClient.SyncStatus(context.Background())
					require.NoError(t, err)
					if statusAfterWait.FinalizedL1.Number > initialStatus.FinalizedL1.Number {
						// Verify that eventually the batch buffer will be updated.
						if !VerifyBatchBufferUpdated(ctx, streamer) {
							require.FailNow(t, "Timeout: Batch buffer not updated")
						}
						return
					}
				}
			}
		}

		time.Sleep(10 * time.Millisecond)
	}
}
