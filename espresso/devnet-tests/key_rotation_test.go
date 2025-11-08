package devnet_tests

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

func TestRotateBatcherKey(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)

	// We're going to change batcher key to Bob's, verify that it won't be a no-op
	require.NotEqual(t, d.secrets.Batcher, d.secrets.Bob)

	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	// Shut down the batcher
	require.NoError(t, d.ServiceDown("op-batcher"))
	d.SleepOutageDuration()

	// Change the batch sender key to Bob
	contract, owner, err := d.SystemConfig(ctx)
	require.NoError(t, err)

	tx, err := contract.SetBatcherHash(owner, eth.AddressAsLeftPaddedHash(d.secrets.Addresses().Bob))
	require.NoError(t, err)

	_, err = d.SendL1Tx(ctx, tx)
	require.NoError(t, err)

	d.secrets.Batcher = d.secrets.Bob

	// Restart the batcher
	require.NoError(t, d.ServiceUp("op-batcher"))
	d.SleepOutageDuration()

	// Send a transaction to check the L2 still runs
	require.NoError(t, d.RunSimpleL2Burn())
}

func TestChangeBatchInboxOwner(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)

	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	config, err := d.RollupConfig(ctx)
	require.NoError(t, err)

	// Change the BatchAuthenticator's owner
	batchAuthenticator, err := bindings.NewBatchAuthenticator(config.BatchAuthenticatorAddress, d.L1)
	require.NoError(t, err)

	// Get L1 chain ID for transaction signing
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	// Check current owner first
	currentOwner, err := batchAuthenticator.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	t.Logf("Current BatchAuthenticator owner: %s", currentOwner.Hex())
	t.Logf("Deployer address: %s", d.secrets.Addresses().Deployer.Hex())
	t.Logf("Bob address: %s", d.secrets.Addresses().Bob.Hex())

	// Use deployer key to sign the transaction
	deployerOpts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Deployer, l1ChainID)
	require.NoError(t, err)

	// Verify we're transferring to the right address
	bobAddress := d.secrets.Addresses().Bob
	t.Logf("Attempting to transfer ownership from %s to %s", currentOwner.Hex(), bobAddress.Hex())
	
	// Call TransferOwnership
	tx, err := batchAuthenticator.TransferOwnership(deployerOpts, bobAddress)
	require.NoError(t, err)
	t.Logf("TransferOwnership transaction hash: %s", tx.Hash().Hex())
	
	// Wait for transaction receipt and check if it succeeded
	receipt, err := wait.ForReceiptOK(ctx, d.L1, tx.Hash())
	require.NoError(t, err)
	t.Logf("TransferOwnership transaction receipt status: %d", receipt.Status)
	t.Logf("TransferOwnership transaction gas used: %d", receipt.GasUsed)
	t.Logf("TransferOwnership transaction completed successfully")

	// Ensure the owner has been changed
	newOwner, err := batchAuthenticator.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	t.Logf("New owner after transfer: %s", newOwner.Hex())
	t.Logf("Expected Bob address: %s", d.secrets.Addresses().Bob.Hex())
	require.Equal(t, newOwner, d.secrets.Addresses().Bob)

	// Check that everything still functions
	require.NoError(t, d.RunSimpleL2Burn())
}
