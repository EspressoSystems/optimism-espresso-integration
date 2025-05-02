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
	"github.com/stretchr/testify/require"
)

func runL1Reorg(ctx context.Context, t *testing.T, system *e2esys.System) {
	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l1Client := system.NodeClient(e2esys.RoleL1)

	// Wait for batcher to start advancing L2 head
	_, err := geth.WaitForBlockToBeSafe(big.NewInt(2), l2Seq, 2*time.Minute)
	if have, want := err, error(nil); have != want {
		t.Fatalf("L2 isn't progressing:\nhave:\n\t%v\nwant:\n\t%v", have, want)
	}

	t.Log("L2 is progressing")

	// Wait for L2 head to be based off non-genesis unfinalized block
	l2HeadL1Info := &derive.L1BlockInfo{}
	var unsafeL2Height uint64
	var l1Height uint64
	for l2HeadL1Info.Number == 0 && (l1Height-l2HeadL1Info.Number) < system.Cfg.L1FinalizedDistance {
		unsafeL2Height, err = l2Seq.BlockNumber(ctx)
		require.NoError(t, err)

		l2Head, err := l2Seq.BlockByNumber(ctx, new(big.Int).SetUint64(unsafeL2Height))
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
}

func TestE2eDevNetWithL1Reorg(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, _, err := launcher.StartDevNet(ctx, t, 16)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	runL1Reorg(ctx, t, system)
}
