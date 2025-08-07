package devnet_tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBatcherRestart(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx)
	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	// Shut down the batcher and have another transaction submitted while it is down.
	require.NoError(t, d.ServiceDown("op-batcher"))
	d.WaitOutage()

	receipt, err := d.SubmitSimpleL2Burn()
	require.NoError(t, err)

	// Check that while the batcher is down, the verifier does NOT process submitted transactions.
	// TODO: currently this fails because the sequencer and verifier use the same op-geth instance.
	// So the verifier gets the transaction from the sequencer. Re-enable this check when the
	// verifier is using its own separate op-geth, so that it is forced to get the transaction from
	// the batcher via the inbox.
	// _, err = d.L2Verif.TransactionReceipt(ctx, receipt.Receipt.TxHash)
	// require.ErrorIs(t, err, ethereum.NotFound)

	// Bring the batcher back up and check that it processes the transaction which was submitted
	// while it was down.
	require.NoError(t, d.ServiceUp("op-batcher"))
	require.NoError(t, d.VerifySimpleL2Burn(receipt))

	// Submit another transaction at the end just to check that things stay working.
	d.WaitSuccess()
	require.NoError(t, d.RunSimpleL2Burn())
}
