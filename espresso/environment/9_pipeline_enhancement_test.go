package environment_test

import (
	"context"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	"io"
	"log/slog"
	"math/big"
	"testing"
)

// TestPipelineEnhancement is a test that ensures the derivation pipeline does not include batches from reverted L1 transactions submitted to the inbox contract.
// Attempt to post a batch that fails in the batch inbox contract. The revert transactions should not be included by the derivation pipeline.
//		Arrange:
//			Running Sequencer, Batcher in Espresso mode
//		Act:
//			Send a transaction not signed by the batcher (which means it will revert) to the inbox contract with a specific sequence of bytes e.g. 0x42424242424242424242
//          Fetch N, the block number where the transaction was included (even though it is reverted)
//		Assert:
//			Instantiate a CalldataSource object with block N. CalldataSource is the object that reads the calldata from the inbox contract and filters it using the isValidBatchTx function.
//			Verify that the concatenated bytes of all batch transactions does not include 0x42424242424242424242

func TestPipelineEnhancement(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartDevNet(ctx, t)
	require.NoError(t, err, "failed to start dev environment with espresso dev node")

	// Stop the batcher to ensure no valid batch is posted to L1.
	driver := system.BatchSubmitter.TestDriver()
	driver.StopBatchSubmitting(ctx)

	l1Client := system.NodeClient(e2esys.RoleL1)

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	// Send a transaction not signed by the batcher to the inbox contract
	// Create the transaction
	txData := []byte("42424242424242424242")

	tx := gethTypes.MustSignNewTx(system.Cfg.Secrets.Bob, system.RollupConfig.L1Signer(), &gethTypes.DynamicFeeTx{
		ChainID:   system.Cfg.L1ChainIDBig(),
		Nonce:     0,
		GasTipCap: big.NewInt(1 * params.GWei),
		GasFeeCap: big.NewInt(10 * params.GWei),
		Gas:       5_000_000,
		To:        &system.RollupConfig.BatchInboxAddress,
		Value:     big.NewInt(0),
		Data:      txData,
	})

	l := log.NewLogger(slog.Default().Handler())
	err = l1Client.SendTransaction(ctx, tx)
	require.NoError(t, err)

	receipt, err := wait.ForReceiptFail(ctx, l1Client, tx.Hash())
	require.Equal(t, receipt.Status, gethTypes.ReceiptStatusFailed)
	require.NoError(t, err, "Waiting for receipt on transaction", tx)

	l1ClientFetching, _ := client.NewRPC(ctx, nil, system.NodeEndpoint(e2esys.RoleL1).RPC())
	l1RefClient, err := sources.NewL1Client(l1ClientFetching, l, nil, sources.L1ClientDefaultConfig(system.RollupConfig, true, sources.RPCKindStandard))

	system.RollupConfig.EcotoneTime = nil
	factory := derive.NewDataSourceFactory(l, system.RollupConfig, l1RefClient, nil, nil)

	batcherAddress := crypto.PubkeyToAddress(*system.BatchSubmitter.BatcherPublicKey)
	l1Block, err := l1Client.BlockByNumber(ctx, receipt.BlockNumber)
	require.NoError(t, err)
	l1BlockRef := eth.L1BlockRef{
		Hash:       l1Block.Hash(),
		Number:     l1Block.NumberU64(),
		ParentHash: l1Block.ParentHash(),
		Time:       l1Block.Time(),
	}
	datas, err := factory.OpenData(ctx, l1BlockRef, batcherAddress)
	require.NoError(t, err)

	data, err := datas.Next(ctx)

	// The L1 data collected by the derivation pipeline is empty because the batch information has been discarded
	require.Equal(t, data, eth.Data(nil))
	require.Equal(t, err, io.EOF)
}
