package environment_test

import (
	"context"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
	"math/big"
	"slices"
	"testing"
	"time"
)

// TestConfirmationIntegrityWithReorgs
// Post batches to both Espresso and the L1 then force the L1 to reorg back to an earlier state in which those batches have not been posted.
// Wait for some time and check that the batches are eventually posted to the L1 again, in the same order as they were originally sequenced,
// as if the reorg did not happen.
// More specifically the test is defined as follows
//	Arrange:
//		Running Sequencer, Batcher in Espresso mode, OP node.
//	Act:
//		Wait for 10 L2 blocks to be posted on L1.
//		Collect the hashes of the corresponding blocks
//		Reorg the chain 10 blocks earlier.
//	Assert:
//		Wait for 10 L2 blocks to be posted on L1 again
//		Collect the hashes of the 10 next blocks from the OP node
//		Check that these hashes are the same and in the same order as the ones collected earlier

func getNextNBatches(ctx context.Context, t *testing.T, system *e2esys.System, l1Client *ethclient.Client, l2Seq *ethclient.Client, numberOfBatches int) []string {
	var l2HeadHashes []string

	l2HeadL1Info := &derive.L1BlockInfo{}

	//l2HeadL1Info.Number == 0 && (l1Height-l2HeadL1Info.Number) < system.Cfg.L1FinalizedDistance && (len(l2HeadHashes) < numberOfBatches)
	for len(l2HeadHashes) < numberOfBatches {
		unsafeL2Height, err := l2Seq.BlockNumber(ctx)
		require.NoError(t, err)

		l2Head, err := l2Seq.BlockByNumber(ctx, new(big.Int).SetUint64(unsafeL2Height))
		require.NoError(t, err)

		hash := l2Head.Hash().String()
		if !slices.Contains(l2HeadHashes, hash) {
			l2HeadHashes = append(l2HeadHashes, hash)
			t.Log("New element", "value", hash, "list length", len(l2HeadHashes))
		}

		_, l2HeadL1Info, err = derive.BlockToSingularBatch(system.RollupCfg(), l2Head)
		log.Info("l2HeadL1Info", "value", l2HeadL1Info)

		//l1Height, err := l1Client.BlockNumber(ctx)
		require.NoError(t, err)
	}

	return l2HeadHashes
}

// TODO Should be merged with https://github.com/EspressoSystems/optimism-espresso-integration/pull/119
func runL1Reorg2(ctx context.Context, t *testing.T, system *e2esys.System) {
	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l1Client := system.NodeClient(e2esys.RoleL1)

	// Wait for batcher to start advancing L2 head
	_, err := geth.WaitForBlockToBeSafe(big.NewInt(2), l2Seq, 2*time.Minute)
	if have, want := err, error(nil); have != want {
		t.Fatalf("L2 isn't progressing:\nhave:\n\t%v\nwant:\n\t%v", have, want)
	}

	t.Log("L2 is progressing")

	// Wait for L2 head to be based off non-genesis unfinalized block

	batches := getNextNBatches(ctx, t, system, l1Client, l2Seq, 10)

	// Wait for these blocks to be final TODO

	log.Info("+++ L2 blocks before reorg", "value", batches)

	var unsafeL2Height uint64
	var l1Height uint64

	l1Height, err = l1Client.BlockNumber(ctx)
	l1Origin, err := l1Client.BlockByNumber(ctx, new(big.Int).SetUint64(l1Height))
	require.NoError(t, err)

	// Introduce a reorg at L1
	t.Logf("Introducing reorg at L1Origin %d, L1Head %d, l2Head %d", l1Origin.Number(), l1Height, unsafeL2Height)
	err = system.ForkL1(l1Origin.ParentHash())
	require.NoError(t, err)

	// Wait for SafeL2 to advance despite the reorg
	//_, err = geth.WaitForBlockToBeSafe(new(big.Int).SetUint64(unsafeL2Height+1), l2Seq, 2*time.Minute)
	//require.NoError(t, err)

}

func TestConfirmationIntegrityWithReorgs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, _, err := launcher.StartDevNet(ctx, t, 16)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	runL1Reorg2(ctx, t, system)

}
