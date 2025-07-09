package proofs

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// Test_EspressoCeloIntegrationActivation tests that the EspressoCeloIntegration
// activation works correctly by verifying that Espresso integration is enabled after the activation timestamp.
// In a real deployment, this would enable the batcher's isEspressoEnabled() method and initialize the
// espressoStreamer for dual submission to L1 and Espresso, therefore allowing the caff node to receive batches from Espresso.
//
// This test starts with a normal e2esys (without Espresso), sets an activation time,
// verifies caff node cannot make progress, then starts the espressoDevNode,
// verifies caff node still cannot progress, then waits until activation time passes
// and verifies that caff node can make progress.
func Test_EspressoCeloIntegrationActivation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var system *e2esys.System
	var espressoDevNode env.EspressoDevNode
	var caffNode *env.CaffNodeInstance

	// Ensure cleanup happens even if test fails
	defer func() {
		if caffNode != nil {
			env.Stop(t, caffNode)
		}
		if espressoDevNode != nil {
			env.Stop(t, espressoDevNode)
		}
		if system != nil {
			env.Stop(t, system)
		}
		// Give time for goroutines to finish
		time.Sleep(1 * time.Second)
	}()

	// Step 1: Start with normal e2esys (without Espresso initially)
	sysConfig := e2esys.DefaultSystemConfig(t, e2esys.WithAllocType(config.AllocTypeStandard))
	system, err := sysConfig.Start(t)
	require.NoError(t, err, "failed to launch without Espresso dev node")

	// Step 2: Configure EspressoCeloIntegration activation time for testing
	// Set it to activate 15 seconds from now

	currentTime := uint64(time.Now().Unix())
	activationTime := currentTime + 15
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

	// Submit some transactions to generate activity before starting Espresso
	seqClient := system.NodeClient(e2esys.RoleSeq)

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

		t.Logf("Submitted transaction %d (before Espresso)", i)
		time.Sleep(1 * time.Second)
	}

	// Verify normal sequencer is working but we haven't started Espresso integration yet
	seqHeader, err := seqClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err, "failed to get sequencer header")
	require.Greater(t, seqHeader.Number.Uint64(), uint64(0), "Sequencer should process blocks normally")
	t.Logf("Sequencer at block %d before Espresso", seqHeader.Number.Uint64())

	// Step 3: Start the Espresso dev node
	t.Log("Starting Espresso dev node...")
	launcher := new(env.EspressoDevNodeLauncherDocker)
	_, espressoDevNode, err = launcher.StartDevNet(ctx, t, env.WithL1FinalizedDistance(0), env.WithSequencerUseFinalized(true))
	require.NoError(t, err, "failed to start espresso dev node")

	// Launch Caff Node with the espresso dev node
	caffNode, err = env.LaunchCaffNode(t, system, espressoDevNode)
	require.NoError(t, err, "failed to start caff node with espresso dev node")

	// Submit more transactions to test with Espresso dev node running but before activation
	for i := 3; i < 6; i++ {
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

		t.Logf("Submitted transaction %d (with Espresso dev node)", i)
		time.Sleep(1 * time.Second)
	}

	// Wait a bit for any potential processing
	time.Sleep(3 * time.Second)

	// Verify that the Caff node still cannot make progress (Espresso integration not activated yet)
	caffClient := system.NodeClient(env.RoleCaffNode)
	caffHeader, err := caffClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err, "failed to get latest header from Caff node")

	// Get current sequencer progress to compare
	currentSeqHeader, err := seqClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err, "failed to get current sequencer header")

	t.Logf("Before activation - Sequencer: block %d, Caff: block %d",
		currentSeqHeader.Number.Uint64(), caffHeader.Number.Uint64())

	// Verify that the Caff node is behind the sequencer (cannot process blocks before activation)
	// The caff node should be at genesis because Espresso integration is not active
	require.Less(t, caffHeader.Number.Uint64(), currentSeqHeader.Number.Uint64(),
		"Caff node should be behind sequencer before activation time")
	require.Equal(t, uint64(0), caffHeader.Number.Uint64(), "Caff node should be at genesis before activation time")

	// Step 4: Wait for activation time to pass
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
	for i := 6; i < 10; i++ {
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

	// Verify that the Caff node can now make progress (indicates Espresso integration is working)
	time.Sleep(10 * time.Second)
	caffHeader, err = caffClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err, "failed to get latest header from Caff node")

	// Verify that both sequencer and caff node are progressing
	require.Greater(t, header.Number.Uint64(), uint64(0), "Sequencer should have processed blocks")
	require.Greater(t, caffHeader.Number.Uint64(), uint64(0), "Caff node should now be able to process blocks after activation")

	t.Logf("Sequencer at block %d, Caff node at block %d", header.Number.Uint64(), caffHeader.Number.Uint64())

	// Final verification: The key test is that EspressoCeloIntegration activation
	// time-based logic is working correctly.
	finalTime := uint64(time.Now().Unix())
	require.True(t, system.RollupConfig.IsEspressoCeloIntegration(finalTime),
		"EspressoCeloIntegration should remain active at test completion")

	t.Log("Test completed successfully: Caff node can make progress after activation time")
}
