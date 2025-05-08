package environment_test

import (
	"context"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
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
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Minute)
	defer cancel()
	launcher := new(env.EspressoDevNodeLauncherDocker)

	// Set NonFinalizedProposals to true and SequencerUseFinalized to false, to make sure we are
	// testing how the batcher handles the finality.
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

	initialStatus, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)
	initialSafeL1Number := initialStatus.SafeL1.Number

	// Wait for new blocks to be finalized, which will enable the batcher to submit more blocks to
	// to the L1.
	tickerFinality := time.NewTicker(1 * time.Second)
	defer tickerFinality.Stop()

	for {
		select {
		case <-ctx.Done():
			require.FailNow(t, "Timeout: Finalized L1 number not increased by 10")
		case <-tickerFinality.C:
			// Verify that the batcher waits for the L1 origin to be finalized before submitting a new
			// block to the L1.
			statusAfterWait, err := rollupClient.SyncStatus(context.Background())
			require.NoError(t, err)
			require.LessOrEqual(t, statusAfterWait.SafeL2.L1Origin.Number, statusAfterWait.FinalizedL1.Number, "L1 origin not finalized before submission")

			// Exit the test if there are 10 new safe blocks on the L1.
			if statusAfterWait.SafeL1.Number >= initialSafeL1Number + 10 {
				return
			}
		}
	}
}
