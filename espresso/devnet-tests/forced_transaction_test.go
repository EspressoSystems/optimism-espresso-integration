package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
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

	// Get OptimismPortal address from SystemConfig
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err, "Failed to get systemConfig")
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get optimismPortalAddr")

	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, l1Client)
	require.NoError(t, err, "Failed to create Optimism portal")

	// L1 signer
	l1ChainID, err := l1Client.ChainID(ctx)
	require.NoError(t, err, "Failed to get l1ChainID")

	opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err, "Failed to create withdrawal transaction options")

	withdrawalAmount := new(big.Int).SetUint64(1000)

	// Forced transaction via L1 deposit
	tx, err := portal.DepositTransaction(
		opts,
		common.HexToAddress(predeploys.L2ToL1MessagePasser),
		withdrawalAmount,
		uint64(300_000),
		false,
		nil,
	)
	require.NoError(t, err, "Failed to create transaction")
	_, err = bind.WaitMined(ctx, l1Client, tx)
	require.NoError(t, err, "Transaction not minted")
	// Wait for forced inclusion
	time.Sleep(WAIT_FORCED_TXN_TIME)

	newBalance, err := wait.ForBalanceChange(ctx, l2Verif, aliceAddress, initialBalance)
	require.NoError(t, err, "Failed to get newBalance")

	require.LessOrEqualf(
		t,
		newBalance.Uint64(),
		initialBalance.Uint64()-withdrawalAmount.Uint64(),
		"Balance not decreased",
	)
}
