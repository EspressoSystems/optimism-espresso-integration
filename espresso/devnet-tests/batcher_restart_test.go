package devnet_tests

import (
	"context"
	"testing"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/require"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
)

func TestBatcherRestart(t *testing.T) {
	testRestart(t, false)
}

func TestEnclaveRestart(t *testing.T) {
	env.RunOnlyWithEnclave(t)
	testRestart(t, true)
}

func testRestart(t *testing.T, tee bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	profile := DevnetProfileNonTee
	if tee {
		profile = DevnetProfileTee
	}

	fmt.Printf("profile: %v, tee: %v\n", profile, tee)

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(profile))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// STOP HERE
    fmt.Println("FLAG: Sleep START 10 min")
    time.Sleep(10 * time.Minute)
    fmt.Println("FLAG: Sleep FINISHED")

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
