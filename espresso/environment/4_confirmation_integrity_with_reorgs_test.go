package environment_test

import (
	"context"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/assert"
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
//		Store the (unfinalized) head of the L1 in variable h.
//		Wait for the first n batches to be posted on possibly unfinalized L1 blocks.
//		Collect the batches of the corresponding batches and store them in list L.
//		Reorg the L1 to height h.
//		Wait for the L2 to reach safe height n again as the corresponding batches are submitted again to L1.
//		Store these n batches in L'
//	Assert:
//		L == L'

// TODO function to compute the hash of a batch
func BatchHash(b *derive.SingularBatch) string {
	hashes := ""
	for _, tx := range b.Transactions {
		hashes = hashes + tx.String()
	}

	return hashes
}

func getBatchesPublishedOnUnfinalizedL1Blocks(ctx context.Context, t *testing.T, system *e2esys.System, l1Client *ethclient.Client, l2Seq *ethclient.Client, l2Verif *ethclient.Client) ([]string, uint64, uint64) {
	var batchInfo []string

	l1Height, err := l1Client.BlockNumber(ctx)
	require.NoError(t, err)
	l1HeightStart := l1Height
	log.Info("L1 height to reorg to", "height", l1HeightStart)
	// Keep monitoring L2 blocks while L1 is producing unfinalized blocks
	i := int64(0)
	// Fetch height of the most recent block which is unsafe and which batch will be sent to L1 at a later stage
	unsafeL2BlockNumber, err := l2Seq.BlockNumber(ctx)

	require.NoError(t, err)
	for (l1Height - l1HeightStart) < system.Cfg.L1FinalizedDistance {
		height := uint64(i) + unsafeL2BlockNumber

		l2Head, err := l2Seq.BlockByNumber(ctx, new(big.Int).SetUint64(height))
		time.Sleep(500 * time.Millisecond)
		if err != nil {
			continue
		} else {
			i++
		}
		batch, l2HeadL1Info, err := derive.BlockToSingularBatch(system.RollupCfg(), l2Head)
		log.Info("l2HeadL1Info", "value", l2HeadL1Info)

		batchHash := BatchHash(batch)
		t.Log("New element", "value", batchHash, "list length", len(batchInfo))
		// Insert new elements only
		if !slices.Contains(batchInfo, batchHash) {
			batchInfo = append(batchInfo, batchHash)
		}

		l1Height, err = l1Client.BlockNumber(ctx)
		require.NoError(t, err)

	}

	return batchInfo, l1HeightStart, unsafeL2BlockNumber
}

func getFirstNL2SafeBlocks(ctx context.Context, t *testing.T, system *e2esys.System, l2Verif *ethclient.Client, n int, startIndex uint64) []string {
	var batches []string

	for i := 0; i < n; i++ {
		height := startIndex + uint64(i)
		_, err := geth.WaitForBlockToBeSafe(big.NewInt(int64(height)), l2Verif, 2*time.Minute)
		require.NoError(t, err)

		l2Head, err := l2Verif.BlockByNumber(ctx, new(big.Int).SetUint64(height))
		require.NoError(t, err)

		//hash := l2Head.Hash().String()
		batch, _, err := derive.BlockToSingularBatch(system.RollupCfg(), l2Head)
		batchHash := BatchHash(batch)
		batches = append(batches, batchHash)

	}
	return batches
}

func run(ctx context.Context, t *testing.T, system *e2esys.System) {
	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l2Verif := system.NodeClient(e2esys.RoleVerif)
	l1Client := system.NodeClient(e2esys.RoleL1)

	var unsafeL2Height uint64
	var l1Height uint64

	// Wait for batcher to start advancing L2 head
	_, err := geth.WaitForBlockToBeSafe(big.NewInt(2), l2Seq, 2*time.Minute)
	if have, want := err, error(nil); have != want {
		t.Fatalf("L2 isn't progressing:\nhave:\n\t%v\nwant:\n\t%v", have, want)
	}

	t.Log("L2 is progressing")

	// Fetch batches before reorg
	batchesBefore, L1BlockHeightToReorgTo, startIndex := getBatchesPublishedOnUnfinalizedL1Blocks(ctx, t, system, l1Client, l2Seq, l2Verif)

	l1Origin, err := l1Client.BlockByNumber(ctx, new(big.Int).SetUint64(L1BlockHeightToReorgTo))
	require.NoError(t, err)

	log.Info("+++ L2 blocks before reorg", "value", batchesBefore)
	// Introduce a reorg at L1
	l1Height, err = l1Client.BlockNumber(ctx)
	require.NoError(t, err)
	t.Logf("Introducing reorg at L1Origin %d, L1Head %d, l2Head %d", l1Origin.Number(), l1Height, unsafeL2Height)
	err = system.ForkL1(l1Origin.ParentHash())
	require.NoError(t, err)

	n := len(batchesBefore)
	batchesAfter := getFirstNL2SafeBlocks(ctx, t, system, l2Verif, n, startIndex)

	log.Info("+++ L2 blocks after reorg", "value", batchesAfter)

	assert.Equal(t, batchesAfter, batchesBefore)

}

func TestConfirmationIntegrityWithReorgs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, _, err := launcher.StartDevNet(ctx, t, 12)
	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	run(ctx, t, system)

}
