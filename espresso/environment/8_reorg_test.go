package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func runL1Reorg(ctx context.Context, t *testing.T, system *e2esys.System) {
	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l1Client := system.NodeClient(e2esys.RoleL1)
	caffClient := system.NodeClient(env.RoleCaffNode)

	// Wait for batcher to start advancing L2 head
	_, err := geth.WaitForBlockToBeSafe(big.NewInt(2), l2Seq, 2*time.Minute)
	if have, want := err, error(nil); have != want {
		t.Fatalf("L2 isn't progressing:\nhave:\n\t%v\nwant:\n\t%v", have, want)
	}

	t.Log("L2 is progressing")

	// Wait for L2 head to be based off non-genesis unfinalized block
	l2HeadL1Info := &derive.L1BlockInfo{}
	var l2Head *types.Block
	var unsafeL2Height uint64
	var l1Height uint64
	for l2HeadL1Info.Number == 0 || (l1Height-l2HeadL1Info.Number) >= system.Cfg.L1FinalizedDistance {
		unsafeL2Height, err = l2Seq.BlockNumber(ctx)
		require.NoError(t, err)

		l2Head, err = l2Seq.BlockByNumber(ctx, new(big.Int).SetUint64(unsafeL2Height))
		require.NoError(t, err)

		_, l2HeadL1Info, err = derive.BlockToSingularBatch(system.RollupCfg(), l2Head)
		require.NoError(t, err)

		l1Height, err = l1Client.BlockNumber(ctx)
		require.NoError(t, err)
	}

	l1Origin, err := l1Client.BlockByNumber(ctx, new(big.Int).SetUint64(l2HeadL1Info.Number))
	require.NoError(t, err)

	// Introduce a reorg at L1
	t.Logf("Introducing reorg at L1Origin %d, L1Head %d, l2Head %d", l1Origin.Number(), l1Height, unsafeL2Height)
	err = system.ForkL1(l1Origin.ParentHash())
	require.NoError(t, err)

	// Wait for SafeL2 to advance despite the reorg
	_, err = geth.WaitForBlockToBeSafe(new(big.Int).SetUint64(unsafeL2Height+1), l2Seq, 2*time.Minute)
	require.NoError(t, err)

	// Check that safe chain doesn't contain the forked block
	newL2Head, err := l2Seq.BlockByNumber(ctx, new(big.Int).SetUint64(unsafeL2Height))
	require.NoError(t, err)
	require.NotEqual(t, newL2Head.Hash(), l2Head.Hash())

	// Check that Caff node came to the same conclusion
	caffL2Head, err := caffClient.BlockByNumber(ctx, new(big.Int).SetUint64(unsafeL2Height))
	require.NoError(t, err)
	require.Equal(t, caffL2Head.Hash(), newL2Head.Hash())
}

// TestE2eDevNetWithL1Reorg tests how the batcher and Caff node handle an L1 reorg.
// Specifically, it focuses on cases where unsafe L2 chain contains blocks that
// reference unfinalized L1 blocks as their origin.
//
// The test is defined as follows
// Arrange:
//
//	Running Sequencer, Batcher in Espresso mode, Caff node & OP node.
//
// Act:
//
//	Wait for sequencer to propose an unsafe L2 block with unfinalized L1 origin
//	Simulate L1 reorg at that block's origin
//
// Assert:
//
//	Assert that derivation pipeline still progresses
//	Assert that Caff and OP node report a new block at the target L2 height
func TestE2eDevNetWithL1Reorg(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, devNode, err := launcher.StartDevNet(ctx, t, 16)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	caffNode, err := env.LaunchDecaffNode(t, system, devNode)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start caff node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Shut down the Caff Node
	defer env.Stop(t, caffNode)

	runL1Reorg(ctx, t, system)
}
