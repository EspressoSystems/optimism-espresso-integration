package devnet_tests

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestChangeBatchInboxOwner(t *testing.T) {
	// Load environment variables from .env file
	err := LoadDevnetEnv()
	require.NoError(t, err, "Failed to load .env file")

	ctx, cancel := context.WithCancel(context.Background())
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

	// Use batch authenticator owner key to sign the transaction
	batchAuthenticatorPrivateKeyHex := os.Getenv("BATCH_AUTHENTICATOR_OWNER_PRIVATE_KEY")
	require.NotEmpty(t, batchAuthenticatorPrivateKeyHex, "BATCH_AUTHENTICATOR_OWNER_PRIVATE_KEY must be set")
	t.Logf("Using BATCH_AUTHENTICATOR_OWNER_PRIVATE_KEY from environment: %s...", batchAuthenticatorPrivateKeyHex[:10])

	batchAuthenticatorKey, err := crypto.HexToECDSA(strings.TrimPrefix(batchAuthenticatorPrivateKeyHex, "0x"))
	require.NoError(t, err)

	batchAuthenticatorOwnerOpts, err := bind.NewKeyedTransactorWithChainID(batchAuthenticatorKey, l1ChainID)
	require.NoError(t, err)

	// Call TransferOwnership
	tx, err := batchAuthenticator.TransferOwnership(batchAuthenticatorOwnerOpts, bobAddress)
	require.NoError(t, err)

	// Wait for transaction receipt and check if it succeeded
	_, err = wait.ForReceiptOK(ctx, d.L1, tx.Hash())
	require.NoError(t, err)

	// Ensure the owner has been changed
	newOwner, err := batchAuthenticator.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, newOwner, bobAddress)

	// Check that everything still functions
	require.NoError(t, d.RunSimpleL2Burn())
}
