package proofs

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// Test_EspressoCeloIntegrationActivation tests that the EspressoCeloIntegration
// activation works correctly by verifying that Espresso integration is enabled after the activation timestamp.
// In a real deployment, this would enable the batcher's isEspressoEnabled() method and initialize the
// espressoStreamer for dual submission to L1 and Espresso, therefore allowing the caff node to receive batches from Espresso.
//
// This test sets up a full Espresso environment with dev node, starts the
// system with EspressoCeloIntegration activation configured, and verifies
// that Espresso integration is enabled by checking caff node is making progress
// after the activation time.
func Test_EspressoCeloIntegrationActivation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Sishan TODO: Remove this and initialize the e2esys with UseEspresso = False
	// Launch Espresso dev node environment
	launcher := new(env.EspressoDevNodeLauncherDocker)
	system, espressoDevNode, err := launcher.StartDevNet(
		ctx,
		t,
		env.WithSequencerUseFinalized(true),
		env.WithL1BlockTime(12*time.Second),
		env.WithL2BlockTime(2*time.Second),
	)
a
	require.NoError(t, err, "failed to start dev environment with espresso dev node")

	defer env.Stop(t, system)
	defer env.Stop(t, espressoDevNode)

	// Launch Caff Node for Espresso integration
	caffNode, err := env.LaunchCaffNode(t, system, espressoDevNode)
	require.NoError(t, err, "failed to start caff node")
	defer env.Stop(t, caffNode)

	// Configure EspressoCeloIntegration activation time for testing
	// Set it to activate 10 seconds from now
	currentTime := uint64(time.Now().Unix())
	activationTime := currentTime + 10
	system.RollupConfig.EspressoCeloIntegrationTime = &activationTime

	t.Logf("EspressoCeloIntegration activation time: %d (current: %d)", activationTime, currentTime)

	// Wait for system to stabilize
	time.Sleep(2 * time.Second)

	// Verify activation state before activation time
	preActivationTime := activationTime - 1
	postActivationTime := activationTime + 1

	require.False(t, system.RollupConfig.IsEspressoCeloIntegration(preActivationTime), "EspressoCeloIntegration should not be active before activation time")
	require.True(t, system.RollupConfig.IsEspressoCeloIntegration(activationTime), "EspressoCeloIntegration should be active at activation time")
	require.True(t, system.RollupConfig.IsEspressoCeloIntegration(postActivationTime), "EspressoCeloIntegration should remain active after activation")

	// Set up transaction submitters to generate activity
	keys := system.Cfg.Secrets
	addresses := keys.Addresses()
	signer := geth_types.LatestSignerForChainID(system.Cfg.L2ChainIDBig())

	// Submit some transactions before activation
	seqClient := system.NodeClient(e2esys.RoleSeq)

	// Submit a few transactions before activation
	for i := 0; i < 3; i++ {
		tx := &geth_types.DynamicFeeTx{
			ChainID:   system.RollupConfig.L2ChainID,
			Nonce:     uint64(i),
			To:        &addresses.Bob,
			Value:     big.NewInt(1),
			Gas:       21000,
			GasFeeCap: big.NewInt(1000000000),
			GasTipCap: big.NewInt(1000000000),
		}

		signedTx, err := geth_types.SignTx(geth_types.NewTx(tx), signer, keys.Alice)
		require.NoError(t, err, "failed to sign transaction")

		err = seqClient.SendTransaction(ctx, signedTx)
		require.NoError(t, err, "failed to send transaction")

		t.Logf("Submitted pre-activation transaction %d", i)
		time.Sleep(1 * time.Second)
	}

	// Verify that the Caff node is not receiving data (indicates Espresso integration is not started)
	caffClient := system.NodeClient(env.RoleCaffNode)
	caffHeader, err := caffClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err, "failed to get latest header from Caff node")

	// Verify that the Caff node is not processing blocks
	require.Equal(t, uint64(0), caffHeader.Number.Uint64(), "Caff node should NOT have processed blocks")

	// Wait for activation time to pass
	t.Log("Waiting for EspressoCeloIntegration activation...")
	timeToWait := time.Until(time.Unix(int64(activationTime), 0)) + 2*time.Second
	if timeToWait > 0 {
		time.Sleep(timeToWait)
	}
	// Verify we're now past activation
	nowTime := uint64(time.Now().Unix())
	t.Logf("Current time: %d, activation time: %d", nowTime, activationTime)
	require.True(t, nowTime >= activationTime, "Should be past activation time")
	require.True(t, system.RollupConfig.IsEspressoCeloIntegration(nowTime), "EspressoCeloIntegration should be active now")

	// Submit transactions after activation to test batcher behavior
	for i := 3; i < 8; i++ {
		tx := &geth_types.DynamicFeeTx{
			ChainID:   system.RollupConfig.L2ChainID,
			Nonce:     uint64(i),
			To:        &addresses.Bob,
			Value:     big.NewInt(1),
			Gas:       21000,
			GasFeeCap: big.NewInt(1000000000),
			GasTipCap: big.NewInt(1000000000),
		}

		signedTx, err := geth_types.SignTx(geth_types.NewTx(tx), signer, keys.Alice)
		require.NoError(t, err, "failed to sign transaction")

		err = seqClient.SendTransaction(ctx, signedTx)
		require.NoError(t, err, "failed to send transaction")

		t.Logf("Submitted post-activation transaction %d", i)
		time.Sleep(1 * time.Second)
	}

	// Wait for blocks to be processed and batched
	time.Sleep(5 * time.Second)

	// Verify that the system is functioning correctly after activation
	// Check that we have processed blocks beyond the activation time
	header, err := seqClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err, "failed to get latest header")

	t.Logf("Latest block time: %d, activation time: %d", header.Time, activationTime)

	// Verify that the activation logic works correctly
	require.True(t, system.RollupConfig.IsEspressoCeloIntegration(header.Time), "EspressoCeloIntegration should be active for current block")

	// Check if we're at the activation block or past it
	if system.RollupConfig.IsEspressoCeloIntegrationActivationBlock(header.Time) {
		t.Log("Current block is the EspressoCeloIntegration activation block")
	} else if system.RollupConfig.IsEspressoCeloIntegration(header.Time - system.RollupConfig.BlockTime) {
		t.Log("EspressoCeloIntegration was activated in a previous block")
	}

	t.Log("EspressoCeloIntegration activation verified successfully")

	// Verify that the Caff node is receiving data (indicates Espresso integration is working)
	caffHeader, err = caffClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err, "failed to get latest header from Caff node")

	// Verify that both sequencer and caff node are progressing
	require.Greater(t, header.Number.Uint64(), uint64(0), "Sequencer should have processed blocks")
	require.Greater(t, caffHeader.Number.Uint64(), uint64(0), "Caff node should have processed blocks")

	// Final verification: The key test is that EspressoCeloIntegration activation
	// time-based logic is working correctly.
	finalTime := uint64(time.Now().Unix())
	require.True(t, system.RollupConfig.IsEspressoCeloIntegration(finalTime),
		"EspressoCeloIntegration should remain active at test completion")

}
