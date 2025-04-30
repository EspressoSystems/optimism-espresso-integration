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
// This tests is designed to evaluate Test 8.2 as outlined within the Espresso Celo Integration
// plan. It has stated task definition as follows:
//
//	Arrange:
//		Run the sequencer and the batcher in Espresso mode.
//	Act:
//		Send two transactions.
//	Assert:
//		The batcher doesn't submit the second transaction to the L1 immediatly.
//		After the first block is finalized the L1, the batcher submits the second transaction.
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
