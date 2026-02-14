package devnet_tests

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/require"
)

func TestBatcherRestart(t *testing.T) {
	// Use a timeout so the test fails with a clear error before the runner's 30m limit.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	// Shut down the batcher and have another transaction submitted while it is down.
	require.NoError(t, d.ServiceDown("op-batcher"))
	d.SleepOutageDuration()

	receipt, err := d.SubmitSimpleL2Burn()
	require.NoError(t, err)

	// Check that while the batcher is down, the verifier does NOT process submitted transactions.
	d.SleepOutageDuration()
	_, err = d.L2Verif.TransactionReceipt(ctx, receipt.Receipt.TxHash)
	require.ErrorIs(t, err, ethereum.NotFound)

	// Bring the batcher back up and check that it processes the transaction which was submitted
	// while it was down.
	require.NoError(t, d.ServiceUp("op-batcher"))
	require.NoError(t, d.VerifySimpleL2Burn(receipt))

	// Submit another transaction at the end just to check that things stay working.
	d.SleepRecoveryDuration()
	require.NoError(t, d.RunSimpleL2Burn())
}
