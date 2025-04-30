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

func TestE2eDevNetWithL1Reorg(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, _, err := launcher.StartDevNet(ctx, t, 8)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l1Client := system.NodeClient(e2esys.RoleL1)

	// Wait for batcher to start advancing L2 head
	_, err = geth.WaitForBlockToBeSafe(big.NewInt(2), l2Seq, 2*time.Minute)
	if have, want := err, error(nil); have != want {
		t.Fatalf("L2 isn't progressing:\nhave:\n\t%v\nwant:\n\t%v", have, want)
	}

	t.Log("L2 is progressing")

	// Wait for L2 head to be based off non-genesis block
	headL1Info := &derive.L1BlockInfo{}
	var unsafeL2Height uint64
	for headL1Info.Number == 0 {
		unsafeL2Height, err = l2Seq.BlockNumber(ctx)
		require.NoError(t, err)

		l2Head, err := l2Seq.BlockByNumber(ctx, new(big.Int).SetUint64(unsafeL2Height))
		require.NoError(t, err)

		_, headL1Info, err = derive.BlockToSingularBatch(system.RollupCfg(), l2Head)
		require.NoError(t, err)
	}

	l1Origin, err := l1Client.BlockByNumber(ctx, new(big.Int).SetUint64(headL1Info.Number))
	require.NoError(t, err)

	// Introduce a reorg at L1
	t.Logf("Introducing reorg at L1Origin %d, l2Head %d", l1Origin.Number(), unsafeL2Height)
	err = system.ForkL1(l1Origin.ParentHash())
	require.NoError(t, err)

	// Wait for SafeL2 to advance despite the reorg
	_, err = geth.WaitForBlockToBeSafe(new(big.Int).SetUint64(unsafeL2Height+1), l2Seq, 2*time.Minute)
	require.NoError(t, err)

	cancel()
}
