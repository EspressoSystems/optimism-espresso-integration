package environment_test

import (
	"context"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

func TestBatcherSwitching(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, espressoDevNode, err := launcher.StartE2eDevnet(ctx, t)
	require.NoError(t, err)

	l1Client := system.NodeClient(e2esys.RoleL1)

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
	options, err := bind.NewKeyedTransactorWithChainID(system.Config().Secrets.Deployer, system.Cfg.L1ChainIDBig())
	require.NoError(t, err)

	batchAuthenticator, err := bindings.NewBatchAuthenticator(system.RollupConfig.BatchAuthenticatorAddress, l1Client)
	require.NoError(t, err)

	tx, err := batchAuthenticator.SwitchBatcher(options)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, l1Client, tx.Hash())
	require.NoError(t, err)

	// Start the fallback batcher
	err = system.FallbackBatchSubmitter.TestDriver().StartBatchSubmitting()
	require.NoError(t, err)

	// Everything should still work
	env.RunSimpleL2Burn(ctx, t, system)
}
