package devnet_tests

import (
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/require"
)

func TestBatcherRestart(t *testing.T) {
	profile := ProfileFromEnv(t)
	t.Run(string(profile), func(t *testing.T) {
		ctx := t.Context()

		d := NewDevnet(ctx, t, profile)
		require.NoError(t, d.Up())
		defer func() {
			require.NoError(t, d.Down())
		}()

		require.NoError(t, d.WaitForBatcher(ctx))

		// Send a transaction just to check that everything has started up ok.
		require.NoError(t, d.RunSimpleL2Burn())

		// Shut down the batcher and have another transaction submitted while it is down.
		require.NoError(t, d.ServiceDown(OpBatcher))
		d.SleepOutageDuration()

		receipt, err := d.SubmitSimpleL2Burn()
		require.NoError(t, err)

		// Check that while the batcher is down, the verifier does NOT process submitted transactions.
		d.SleepOutageDuration()
		_, err = d.L2Verif.TransactionReceipt(ctx, receipt.Receipt.TxHash)
		require.ErrorIs(t, err, ethereum.NotFound)

		// Bring the batcher back up and check that it processes the transaction which was submitted
		// while it was down.
		require.NoError(t, d.ServiceUp(OpBatcher))
		require.NoError(t, d.WaitForBatcher(ctx))
		require.NoError(t, d.VerifySimpleL2Burn(receipt))

		// Submit another transaction at the end just to check that things stay working.
		d.SleepRecoveryDuration()
		require.NoError(t, d.RunSimpleL2Burn())
	})
}
