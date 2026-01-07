package devnet_tests

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// hasBatchTransactions checks if any transactions were sent to the BatchInbox from the given sender in the block range by inspecting transactions directly (since BatchInbox doesn't emit logs).
func hasBatchTransactions(ctx context.Context, client *ethclient.Client, batchInboxAddr, senderAddr common.Address, startBlock, endBlock uint64) (bool, error) {
	for i := startBlock; i <= endBlock; i++ {
		timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		block, err := client.BlockByNumber(timeoutCtx, new(big.Int).SetUint64(i))
		cancel()
		if err != nil {
			return false, fmt.Errorf("failed to get block %d: %w", i, err)
		}

		for _, tx := range block.Transactions() {
			if tx.To() != nil && *tx.To() == batchInboxAddr {
				signer := types.LatestSignerForChainID(tx.ChainId())
				sender, err := types.Sender(signer, tx)
				if err != nil {
					continue
				}
				if sender == senderAddr {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

// TestBatcherActivePublishOnly tests that only the active batcher publishes to L1.
func TestBatcherActivePublishOnly(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Initialize devnet with NON_TEE profile (starts both batchers)
	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send initial transaction to verify everything has started up ok
	require.NoError(t, d.RunSimpleL2Burn())
	config, err := d.RollupConfig(ctx)
	require.NoError(t, err)

	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	deployerOpts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Deployer, l1ChainID)
	require.NoError(t, err)

	batchAuthenticator, err := bindings.NewBatchAuthenticator(config.BatchAuthenticatorAddress, d.L1)
	require.NoError(t, err)

	teeBatcherAddr, err := batchAuthenticator.TeeBatcher(&bind.CallOpts{})
	require.NoError(t, err)
	nonTeeBatcherAddr, err := batchAuthenticator.NonTeeBatcher(&bind.CallOpts{})
	require.NoError(t, err)

	activeIsTee, err := batchAuthenticator.ActiveIsTee(&bind.CallOpts{})
	require.NoError(t, err)
	t.Logf("Initial state: activeIsTee = %v", activeIsTee)

	// Switch to non-TEE if TEE is active initially
	if activeIsTee {
		t.Logf("Switching to non-TEE batcher...")
		switchTx, err := batchAuthenticator.SwitchBatcher(deployerOpts)
		require.NoError(t, err)
		receipt, err := wait.ForReceiptOK(ctx, d.L1, switchTx.Hash())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		activeIsTee = false
		t.Logf("Switched to non-TEE batcher")
	}

	// Test non-TEE batcher
	// Wait a bit for services to stabilize after switch
	time.Sleep(5 * time.Second)

	startBlock, err := d.L1.BlockNumber(ctx)
	require.NoError(t, err)
	t.Logf("Starting from block %d", startBlock)

	// Generate L2 traffic for non-TEE batcher
	burnErr := d.RunSimpleL2Burn()
	require.NoError(t, burnErr)
	t.Logf("Generated L2 transaction for non-TEE batcher")

	// Wait for batcher to publish
	time.Sleep(30 * time.Second)
	t.Logf("Waited 30s for L1 confirmation")

	endBlock, err := d.L1.BlockNumber(ctx)
	require.NoError(t, err)
	t.Logf("Checking blocks %d-%d", startBlock, endBlock)

	teePublished, err := hasBatchTransactions(ctx, d.L1, config.BatchInboxAddress, teeBatcherAddr, startBlock, endBlock)
	require.NoError(t, err)
	nonTeePublished, err := hasBatchTransactions(ctx, d.L1, config.BatchInboxAddress, nonTeeBatcherAddr, startBlock, endBlock)
	require.NoError(t, err)

	t.Logf("TEE batcher published: %v, non-TEE batcher published: %v", teePublished, nonTeePublished)

	require.True(t, nonTeePublished, "non-TEE batcher should publish when active")
	require.False(t, teePublished, "TEE batcher should NOT publish when inactive")

	// Switch to TEE and test
	t.Logf("Switching to TEE batcher...")
	switchTx, err := batchAuthenticator.SwitchBatcher(deployerOpts)
	require.NoError(t, err)
	receipt, err := wait.ForReceiptOK(ctx, d.L1, switchTx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	t.Logf("Switched to TEE batcher")

	// Wait for services to stabilize and process backlog
	time.Sleep(10 * time.Second)

	startBlockAfter, err := d.L1.BlockNumber(ctx)
	require.NoError(t, err)
	t.Logf("After switch, starting from block %d", startBlockAfter)

	// Generate L2 traffic for TEE batcher
	burnReceiptAfter, err := d.SubmitSimpleL2Burn()
	require.NoError(t, err)
	t.Logf("Generated L2 transaction for TEE batcher: %s (L2 block %d)", burnReceiptAfter.Receipt.TxHash, burnReceiptAfter.Receipt.BlockNumber)

	// Wait for batcher to publish
	time.Sleep(30 * time.Second)
	t.Logf("Waited 30s for L1 confirmation (after switch)")

	endBlockAfter, err := d.L1.BlockNumber(ctx)
	require.NoError(t, err)
	t.Logf("Checking blocks %d-%d after switch", startBlockAfter, endBlockAfter)

	teePublishedAfter, err := hasBatchTransactions(ctx, d.L1, config.BatchInboxAddress, teeBatcherAddr, startBlockAfter, endBlockAfter)
	require.NoError(t, err)
	nonTeePublishedAfter, err := hasBatchTransactions(ctx, d.L1, config.BatchInboxAddress, nonTeeBatcherAddr, startBlockAfter, endBlockAfter)
	require.NoError(t, err)

	t.Logf("After switch - TEE batcher published: %v, non-TEE batcher published: %v", teePublishedAfter, nonTeePublishedAfter)

	require.True(t, teePublishedAfter, "TEE batcher should publish after becoming active")
	require.False(t, nonTeePublishedAfter, "non-TEE batcher should NOT publish after becoming inactive")
}
