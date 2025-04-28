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
//		Running the sequencer and the batcher in Espresso mode.
//	Act:
//		Send a single transaction derived from an unfinalized L1 block.
//	Assert:
//		The batcher doesn't submit the transaction to the L1 immediatly.
//		After the derived L1 block is finalized, the batcher submits the transaciton.
func TestBatcherWaitForFinality(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t, 0)
	// Signal the testnet to shut down
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	caffNode, err := env.LaunchDecaffNode(t, system, espressoDevNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer env.Stop(t, caffNode)

	l2Seq := system.NodeClient("sequencer")
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	rollupClient := system.RollupClient("verifier")

	intialSeqStatus, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)

    privateKey := system.Cfg.Secrets.Bob
    if err != nil {
        t.Fatalf("failed to create transaction options for Bob: %v", err)
    }

	_ = helpers.SendL2Tx(t, system.Cfg, l2Seq, privateKey, func(opts *helpers.TxOpts) {
		opts.Value = big.NewInt(1)
		opts.Nonce = 1 // Already have deposit
		opts.ToAddr = &common.Address{0xff, 0xff}
		opts.VerifyOnClients(l2Verif)
	})

	_ = helpers.SendL2Tx(t, system.Cfg, l2Seq, privateKey, func(opts *helpers.TxOpts) {
		opts.Value = big.NewInt(1)
		opts.Nonce = 2
		opts.ToAddr = &common.Address{0xff, 0xff}
		opts.VerifyOnClients(l2Verif)
	})

	// Verify that no block is finalized.
	SeqStatusBeforeWait, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)
	require.Equal(t, SeqStatusBeforeWait.FinalizedL1.Number, intialSeqStatus.FinalizedL1.Number, "Finalized L1 number increased")

	// TODO (Keyao) Find a proper time or a better way to handle the wait
	time.Sleep(2 * time.Second)
	SeqStatusAfterWait1, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)
	require.Equal(t, SeqStatusAfterWait1.FinalizedL1.Number, intialSeqStatus.FinalizedL1.Number + 2, "Finalized L1 number increased")

	// Verify that both blocks are finalized.
	time.Sleep(5 * time.Second)
	SeqStatusAfterWait, err := rollupClient.SyncStatus(context.Background())
	require.NoError(t, err)
	require.Greater(t, SeqStatusAfterWait.FinalizedL1.Number, intialSeqStatus.FinalizedL1.Number + 2, "Finalized L1 number not increased")
}
