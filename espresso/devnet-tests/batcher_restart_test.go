package devnet_tests

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum"
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
	// TODO this delay shouldn't really be necessary. It just takes the verifier a little while to
	// receive the transaction after the batcher starts up; a few seconds longer than the time limit
	// that is built into `VerifySimpleL2Burn`. This can likely be fixed by changing some
	// configuration, such as how quickly the batcher posts or the L1 block or epoch time.
	d.SleepRecoveryDuration()
	require.NoError(t, d.VerifySimpleL2Burn(receipt))

	// Submit another transaction at the end just to check that things stay working.
	d.SleepRecoveryDuration()
	require.NoError(t, d.RunSimpleL2Burn())
}
