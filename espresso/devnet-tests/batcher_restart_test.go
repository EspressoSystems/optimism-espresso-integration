package devnet_tests

import (
	"context"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/require"
)

func TestBatcherRestart(t *testing.T) {
	t.Run("non-tee", func(t *testing.T) {
		runTest(t, ComposeProfileNonTee)
	})

	t.Run("tee", func(t *testing.T) {
		env.RunOnlyWithEnclave(t)
		runTest(t, ComposeProfileTee)
	})
}

func runTest(t *testing.T, profile ComposeProfile) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t, profile)
	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	// Shut down the batcher and have another transaction submitted while it is down.
	require.NoError(t, d.ServiceDown(ComposeServiceBatcher))
	d.SleepOutageDuration()

	receipt, err := d.SubmitSimpleL2Burn()
	require.NoError(t, err)

	// Check that while the batcher is down, the verifier does NOT process submitted transactions.
	d.SleepOutageDuration()
	_, err = d.L2Verif.TransactionReceipt(ctx, receipt.Receipt.TxHash)
	require.ErrorIs(t, err, ethereum.NotFound)

	// Bring the batcher back up and check that it processes the transaction which was submitted
	// while it was down.
	require.NoError(t, d.ServiceUp(ComposeServiceBatcher))
	require.NoError(t, d.VerifySimpleL2Burn(receipt))

	// Submit another transaction at the end just to check that things stay working.
	d.SleepRecoveryDuration()
	require.NoError(t, d.RunSimpleL2Burn())
}
