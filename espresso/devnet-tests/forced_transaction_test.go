package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
<<<<<<< HEAD
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
=======
	"github.com/ethereum-optimism/optimism/op-core/predeploys"
>>>>>>> celo-integration-rebase-16
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// Time to wait for a transaction to be enforced after submission.
const WAIT_FORCED_TXN_TIME = 25 * time.Second

// ForcedTransaction attempts to verify that the forced transaction mechanism works for the
// current Docker Compose devnet
func TestForcedTransaction(t *testing.T) {
	// Set up the test timeout condition.
	// Extended timeout to accommodate slower processing in test environments
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	// Launch docker compose devnet
	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	l2Verif := d.L2Verif
	l1Client := d.L1

	// Alice address
	aliceAddress := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)

	// Initial L2 balance
	initialBalance, err := l2Verif.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err, "Failed to get initial balance")
	require.True(t, initialBalance.Cmp(big.NewInt(0)) > 0, "Alice should have positive L2 balance")
	t.Logf("Alice initial L2 balance: %s", initialBalance.String())

	// Get OptimismPortal address from SystemConfig BEFORE stopping the sequencer
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err, "Failed to get systemConfig")
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get optimismPortalAddr")

	// Stop the sequencer to force inclusion via L1 deposits
	// This is required to test the forced transaction mechanism
	t.Log("Stopping sequencer to test forced inclusion...")
	err = d.ServiceDown("op-node-sequencer")
	require.NoError(t, err, "Failed to stop sequencer")

	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, l1Client)
	require.NoError(t, err, "Failed to create Optimism portal")

	// L1 signer
	l1ChainID, err := l1Client.ChainID(ctx)
	require.NoError(t, err, "Failed to get l1ChainID")

	opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err, "Failed to create withdrawal transaction options")

	withdrawalAmount := new(big.Int).SetUint64(1000)
	// Don't set opts.Value - we want to use Alice's existing L2 balance, not mint new ETH
	opts.GasLimit = 500000 // Set explicit gas limit to avoid gas estimation issues in ResourceMetering

	// Forced transaction via L1 deposit
	tx, err := portal.DepositTransaction(
		opts,
		common.HexToAddress(predeploys.L2ToL1MessagePasser),
		withdrawalAmount,
		uint64(100_000), // L2 gas limit - reduced since we just need the deposit to go through
		false,
		nil,
	)
	require.NoError(t, err, "Failed to create transaction")
	t.Logf("Deposit transaction submitted: %s", tx.Hash().Hex())
	receipt, err := bind.WaitMined(ctx, l1Client, tx)
	require.NoError(t, err, "Transaction not minted")

	// Note: we only check the transaction goes through, we don't wait for the sequencer window to expire and the deposit to be processed by the verifier.
	// This is because it would take too long to test on the local devnet as reducing the sequencer window size breaks the other tests.
	t.Logf("Deposit transaction mined in block %d", receipt.BlockNumber.Uint64())

}
