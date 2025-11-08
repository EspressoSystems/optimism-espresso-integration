package devnet_tests

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
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

	batchAuthenticator, err := bindings.NewBatchAuthenticator(config.BatchAuthenticatorAddress, d.L1)
	require.NoError(t, err)
	currentOwner, err := batchAuthenticator.Owner(&bind.CallOpts{})
	require.NoError(t, err)

	// The BatchAuthenticator should be owned by the deployer
	deployerAddress := d.secrets.Addresses().Deployer
	aliceAddress := d.secrets.Addresses().Alice

	t.Logf("Current owner: %s", currentOwner.Hex())
	t.Logf("Deployer address: %s", deployerAddress.Hex())

	// Verify the contract is owned by the deployer (as expected from deployment)
	require.Equal(t, currentOwner, deployerAddress,
		"BatchAuthenticator should be owned by deployer %s, but is owned by %s",
		deployerAddress.Hex(), currentOwner.Hex())

	require.NotEqual(t, deployerAddress, aliceAddress, "Alice should not be the current owner")

	// Create transaction options using the deployer's private key
	chainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	ownerAuth, err := bind.NewKeyedTransactorWithChainID(d.secrets.Deployer, chainID)
	require.NoError(t, err)

	tx, err := batchAuthenticator.TransferOwnership(ownerAuth, aliceAddress)
	require.NoError(t, err, "Ownership transfer transaction building failed.")

	// Send transaction using the devnet's SendL1Tx method
	receipt, err := d.SendL1Tx(ctx, tx)
	require.NoError(t, err, "Failed to send ownership transfer transaction.")
	require.Equal(t, receipt.Status, uint64(1), "Transaction failed")

	// Ensure the owner has been changed
	newOwner, err := batchAuthenticator.Owner(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get new owner.")
	require.Equal(t, newOwner, aliceAddress, "New Owner is not Alice")

	// Check that everything still functions
	require.NoError(t, d.RunSimpleL2Burn())
}
