package devnet_tests

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	e2ebindings "github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestChangeBatchInboxOwner(t *testing.T) {
	// Load environment variables from .env file
	err := LoadDevnetEnv()
	require.NoError(t, err, "Failed to load .env file")

	// 25 min: group 0 runs after TestChallengeGame; devnet bring-up can be slow under CI load.
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	config, err := d.RollupConfig(ctx)
	require.NoError(t, err)

	batchAuthenticator, err := bindings.NewBatchAuthenticator(config.BatchAuthenticatorAddress, d.L1)
	require.NoError(t, err)

	// Get L1 chain ID for transaction signing
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	// Check current owner first
	currentOwner, err := batchAuthenticator.Owner(&bind.CallOpts{})
	require.NoError(t, err)

	// Check that the new owner is different from the current one
	bobAddress := d.secrets.Addresses().Bob
	require.NotEqual(t, currentOwner, bobAddress)

	// Get the ProxyAdmin address from the BatchAuthenticator proxy
	proxyContract, err := e2ebindings.NewProxy(config.BatchAuthenticatorAddress, d.L1)
	require.NoError(t, err)

	var result []interface{}
	proxyRaw := &e2ebindings.ProxyRaw{Contract: proxyContract}
	err = proxyRaw.Call(&bind.CallOpts{}, &result, "admin")
	require.NoError(t, err)
	require.Len(t, result, 1, "admin() should return one value")
	proxyAdminAddress := result[0].(common.Address)
	require.NotEqual(t, proxyAdminAddress, common.Address{}, "ProxyAdmin address should not be zero")

	// Get ProxyAdmin contract binding
	proxyAdmin, err := e2ebindings.NewProxyAdmin(proxyAdminAddress, d.L1)
	require.NoError(t, err)

	// Verify current owner matches initially (they're set to the same address during deployment)
	proxyAdminOwner, err := proxyAdmin.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, currentOwner, proxyAdminOwner, "BatchAuthenticator owner should initially match ProxyAdmin owner")

	// Use batch authenticator owner key to sign the transaction
	batchAuthenticatorPrivateKeyHex := os.Getenv("BATCH_AUTHENTICATOR_OWNER_PRIVATE_KEY")
	require.NotEmpty(t, batchAuthenticatorPrivateKeyHex, "BATCH_AUTHENTICATOR_OWNER_PRIVATE_KEY must be set")
	t.Logf("Using BATCH_AUTHENTICATOR_OWNER_PRIVATE_KEY from environment: %s...", batchAuthenticatorPrivateKeyHex[:10])

	batchAuthenticatorKey, err := crypto.HexToECDSA(strings.TrimPrefix(batchAuthenticatorPrivateKeyHex, "0x"))
	require.NoError(t, err)

	batchAuthenticatorOwnerOpts, err := bind.NewKeyedTransactorWithChainID(batchAuthenticatorKey, l1ChainID)
	require.NoError(t, err)

	// Transfer ownership of both ProxyAdmin and BatchAuthenticator
	// Note: BatchAuthenticator and ProxyAdmin have independent ownership since the migration
	// to OwnableWithGuardiansUpgradeable, so we need to transfer both.

	// 1. Transfer ProxyAdmin ownership
	tx, err := proxyAdmin.TransferOwnership(batchAuthenticatorOwnerOpts, bobAddress)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, tx.Hash())
	require.NoError(t, err)

	// 2. Transfer BatchAuthenticator ownership (2-step process with Ownable2StepUpgradeable)
	// Step 2a: Current owner initiates transfer
	tx, err = batchAuthenticator.TransferOwnership(batchAuthenticatorOwnerOpts, bobAddress)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, tx.Hash())
	require.NoError(t, err)

	// Step 2b: New owner (Bob) accepts ownership
	bobOpts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Bob, l1ChainID)
	require.NoError(t, err)
	tx, err = batchAuthenticator.AcceptOwnership(bobOpts)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, tx.Hash())
	require.NoError(t, err)

	// Verify ProxyAdmin owner has been changed
	newProxyAdminOwner, err := proxyAdmin.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, bobAddress, newProxyAdminOwner, "ProxyAdmin owner should be updated to Bob")

	// Verify BatchAuthenticator owner has been changed
	newOwner, err := batchAuthenticator.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, bobAddress, newOwner, "BatchAuthenticator owner should be updated to Bob")

	// Check that everything still functions
	require.NoError(t, d.RunSimpleL2Burn())
}
