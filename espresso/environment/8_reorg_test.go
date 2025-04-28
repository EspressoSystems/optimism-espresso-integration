package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/common"
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

	l2Seq := system.NodeClient("sequencer")
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	rollupClient := system.RollupClient("verifier")

	// Record the initial sync status.
	initialSeqStatus, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)
	initialFinalizedL1Number := initialSeqStatus.FinalizedL1.Number
	initialCurrentL1Number := initialSeqStatus.CurrentL1.Number

	// Send two transactions.
    privateKey := system.Cfg.Secrets.Bob
    if err != nil {
        t.Fatalf("failed to create transaction options for Bob: %v", err)
    }
	env.SendL2TxNoReceipt(t, system.Cfg, l2Seq, privateKey, func(opts *helpers.TxOpts) {
		opts.Value = big.NewInt(1)
		opts.Nonce = 1 // Already have deposit
		opts.ToAddr = &common.Address{0xff, 0xff}
		opts.VerifyOnClients(l2Verif)
	})
	env.SendL2TxNoReceipt(t, system.Cfg, l2Seq, privateKey, func(opts *helpers.TxOpts) {
		opts.Value = big.NewInt(1)
		opts.Nonce = 2
		opts.ToAddr = &common.Address{0xff, 0xff}
		opts.VerifyOnClients(l2Verif)
	})

	// Verify that the second block is not submitted to the L1 before the first block is finalized.
	SeqStatusBeforeWait, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)
	require.Equal(t, SeqStatusBeforeWait.FinalizedL1.Number, initialSeqStatus.FinalizedL1.Number, "Finalized L1 number not expected to increase")
	require.Less(t, SeqStatusBeforeWait.CurrentL1.Number, initialSeqStatus.CurrentL1.Number + 2, "Current L1 number not expected to increase by more than 1")

	// Verify that the second block is submitted to the L1 after the first block is finalized.
	tickerFinality := time.NewTicker(1 * time.Second)
	defer tickerFinality.Stop()

	for {
		select {
		case <-ctx.Done():
			require.FailNow(t, "Timeout: Finalized L1 number not increased")
		case <-tickerFinality.C:
			seqStatusAfterWait, err := rollupClient.SyncStatus(context.Background())
			require.NoError(t, err)

			// Wait for the first block to be finalized.
			if seqStatusAfterWait.FinalizedL1.Number > initialFinalizedL1Number {
				tickerSubmission := time.NewTicker(1 * time.Second)
				defer tickerSubmission.Stop()

				for {
					select {
					case <-ctx.Done():
						require.FailNow(t, "Timeout: Current L1 number not increased by 2")
					case <-tickerSubmission.C:
						seqStatusAfterWait, err := rollupClient.SyncStatus(context.Background())
						require.NoError(t, err)

						if seqStatusAfterWait.CurrentL1.Number > initialCurrentL1Number + 1 {
							return
						}
					}
				}
			}
		}
	}
}
