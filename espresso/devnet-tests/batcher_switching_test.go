package devnet_tests

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

// TestBatcherSwitching tests that the batcher can be switched from the TEE-enabled
// batcher to a fallback non-TEE batcher using the BatchAuthenticator contract.
//
// This is the devnet equivalent of TestBatcherSwitching from the E2E tests.
// The test runs two batchers in parallel:
// - op-batcher: The primary batcher with Espresso enabled (initially active)
// - op-batcher-fallback: The fallback batcher without Espresso (initially stopped)
func TestBatcherSwitching(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize devnet with NON_TEE profile (starts both batchers)
	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send initial transaction to verify everything has started up ok
	require.NoError(t, d.RunSimpleL2Burn())

	// Get rollup config to access BatchAuthenticator address
	config, err := d.RollupConfig(ctx)
	require.NoError(t, err)

	// Get L1 chain ID for transaction signing
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	// Create transactor options using the deployer key (owner of BatchAuthenticator)
	deployerOpts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Deployer, l1ChainID)
	require.NoError(t, err)

	// Bind to BatchAuthenticator contract
	batchAuthenticator, err := bindings.NewBatchAuthenticator(config.BatchAuthenticatorAddress, d.L1)
	require.NoError(t, err)

	// Check current active batcher state before switching
	activeIsTee, err := batchAuthenticator.ActiveIsTee(&bind.CallOpts{})
	require.NoError(t, err)
	t.Logf("Before switch: activeIsTee = %v", activeIsTee)

	// Stop the primary "TEE" batcher (op-batcher with Espresso enabled)
	require.NoError(t, d.StopBatcherSubmitting("op-batcher"))
	t.Logf("Stopped op-batcher batch submission")

	// Switch active batcher via BatchAuthenticator contract
	tx, err := batchAuthenticator.SwitchBatcher(deployerOpts)
	require.NoError(t, err)
	t.Logf("Submitted switchBatcher transaction: %s", tx.Hash().Hex())

	// Wait for transaction receipt
	receipt, err := wait.ForReceiptOK(ctx, d.L1, tx.Hash())
	require.NoError(t, err)
	t.Logf("SwitchBatcher transaction confirmed in block %d", receipt.BlockNumber.Uint64())

	// Verify the switch happened
	activeIsTeeAfter, err := batchAuthenticator.ActiveIsTee(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotEqual(t, activeIsTee, activeIsTeeAfter, "activeIsTee should have toggled")
	t.Logf("After switch: activeIsTee = %v", activeIsTeeAfter)

	// Start the fallback batcher
	require.NoError(t, d.StartBatcherSubmitting("op-batcher-fallback"))
	t.Logf("Started op-batcher-fallback batch submission")

	// Verify everything still works with the fallback batcher
	require.NoError(t, d.RunSimpleL2Burn())
	t.Logf("Transaction verified with fallback batcher")

	// Submit another transaction and verify system continues to work
	d.SleepRecoveryDuration()
	require.NoError(t, d.RunSimpleL2Burn())
	t.Logf("System continues to work after batcher switch")
}
