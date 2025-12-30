package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

func TestBatcherSwitching(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	// We will need this config to start a new instance of "TEE" batcher
	// with parameters tweaked.
	batcherConfig := &batcher.CLIConfig{}
	system, espressoDevNode, err := launcher.StartE2eDevnet(ctx, t, env.WithSequencerUseFinalized(true), env.GetBatcherConfig(batcherConfig))
	require.NoError(t, err)

	l1Client := system.NodeClient(e2esys.RoleL1)
	verifClient := system.NodeClient(e2esys.RoleVerif)
	espClient := espressoClient.NewClient(espressoDevNode.EspressoUrls()[0])

	deployerTransactor, err := bind.NewKeyedTransactorWithChainID(system.Config().Secrets.Deployer, system.Cfg.L1ChainIDBig())
	require.NoError(t, err)

	batchAuthenticator, err := bindings.NewBatchAuthenticator(system.RollupConfig.BatchAuthenticatorAddress, l1Client)
	require.NoError(t, err)

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	// Send Transaction on L1, and wait for verification on the L2 Verifier
	env.RunSimpleL1TransferAndVerifier(ctx, t, system)

	// Verify everything works
	env.RunSimpleL2Burn(ctx, t, system)

	// Stop the "TEE" batcher
	err = system.BatchSubmitter.TestDriver().StopBatchSubmitting(ctx)
	require.NoError(t, err)

	// Switch active batcher
	tx, err := batchAuthenticator.SwitchBatcher(deployerTransactor)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, l1Client, tx.Hash())
	require.NoError(t, err)

	// Start the fallback batcher
	err = system.FallbackBatchSubmitter.TestDriver().StartBatchSubmitting()
	require.NoError(t, err)

	// Everything should still work
	env.RunSimpleL2Burn(ctx, t, system)

	// Stop the fallback batcher
	err = system.FallbackBatchSubmitter.TestDriver().StopBatchSubmitting(ctx)
	require.NoError(t, err)

	// Switch batcher back to the "TEE" batcher
	tx, err = batchAuthenticator.SwitchBatcher(deployerTransactor)
	require.NoError(t, err)
	switchReceipt, err := wait.ForReceiptOK(ctx, l1Client, tx.Hash())
	require.NoError(t, err)

	// Give things time to settle
	var l2Height uint64

	ticker := time.NewTicker(100 * time.Millisecond)
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

Loop:
	for {
		select {
		case <-timeoutCtx.Done():
			panic("Timeout waiting for verifier derivation pipeline to advance past the fallback batcher switchoff point")
		case <-ticker.C:
			status, err := system.RollupClient(e2esys.RoleVerif).SyncStatus(ctx)
			require.NoError(t, err)
			if status.CurrentL1.Number > switchReceipt.BlockNumber.Uint64() {
				l2Height = status.LocalSafeL2.Number
				break Loop
			}
		}
	}

	espHeight, err := espClient.FetchLatestBlockHeight(ctx)
	require.NoError(t, err)

	// Start a new "TEE" batcher
	batcherConfig.Espresso.CaffeinationHeightEspresso = espHeight
	batcherConfig.Espresso.CaffeinationHeightL2 = l2Height
	newBatcher, err := batcher.BatcherServiceFromCLIConfig(ctx, "0.0.1", batcherConfig, system.BatchSubmitter.Log)
	require.NoError(t, err)
	err = newBatcher.Start(ctx)
	require.NoError(t, err)

	// Everything should still work
	env.RunSimpleL2Burn(ctx, t, system)

	caffNode, err := env.LaunchCaffNode(t, system, espressoDevNode, func(c *config.Config) {
		c.Rollup.CaffNodeConfig.CaffeinationHeightEspresso = espHeight
		c.Rollup.CaffNodeConfig.CaffeinationHeightL2 = l2Height
	})
	require.NoError(t, err)
	defer env.Stop(t, caffNode)

	caffClient := system.NodeClient(env.RoleCaffNode)

	verifHeight, err := verifClient.BlockNumber(ctx)
	require.NoError(t, err)
	verifBlock, err := verifClient.BlockByNumber(ctx, new(big.Int).SetUint64(verifHeight))
	require.NoError(t, err)

	err = wait.ForBlock(ctx, caffClient, verifHeight)
	require.NoError(t, err)

	caffBlock, err := caffClient.BlockByNumber(ctx, new(big.Int).SetUint64(verifHeight))
	require.NoError(t, err)

	require.Equal(t, verifBlock.Hash(), caffBlock.Hash())
}
