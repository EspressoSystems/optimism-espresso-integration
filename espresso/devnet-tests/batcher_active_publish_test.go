package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// hasBatchTransactions checks if any transactions were sent to the BatchInbox from the given sender in the block range.
func hasBatchTransactions(ctx context.Context, client *ethclient.Client, batchInboxAddr, senderAddr common.Address, startBlock, endBlock uint64) (bool, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	chainID, err := client.ChainID(timeoutCtx)
	if err != nil {
		return false, err
	}
	signer := types.NewCancunSigner(chainID)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlock)),
		ToBlock:   big.NewInt(int64(endBlock)),
		Addresses: []common.Address{batchInboxAddr},
	}

	logs, err := client.FilterLogs(timeoutCtx, query)
	if err != nil {
		return false, err
	}

	maxChecks := 50
	checked := 0
	for _, log := range logs {
		if checked >= maxChecks {
			break
		}
		checked++

		tx, _, err := client.TransactionByHash(timeoutCtx, log.TxHash)
		if err != nil {
			continue
		}

		txSender, err := types.Sender(signer, tx)
		if err != nil {
			continue
		}

		if txSender == senderAddr {
			return true, nil
		}
	}

	return false, nil
}

// TestBatcherActivePublishOnly tests that only the active batcher publishes to L1.
func TestBatcherActivePublishOnly(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Minute)
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer require.NoError(t, d.Down())

	// op-batcher-fallback starts stopped in NON_TEE profile, so start it first
	require.NoError(t, d.ServiceUp("op-batcher-fallback"))
	require.NoError(t, d.StartBatcherSubmitting("op-batcher-fallback"))
	time.Sleep(5 * time.Second) // Let batchers initialize

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

	// Ensure verifier services are running before using them
	require.NoError(t, d.ServiceUp("op-geth-verifier"), "op-geth-verifier should be running")
	require.NoError(t, d.ServiceUp("op-node-verifier"), "op-node-verifier should be running")

	startBlock, err := d.L1.BlockNumber(ctx)
	require.NoError(t, err)
	t.Logf("Starting from block %d", startBlock)

	// Retry RunSimpleL2Burn in case services need time to be ready
	var burnErr error
	for i := 0; i < 3; i++ {
		burnErr = d.RunSimpleL2Burn()
		if burnErr == nil {
			break
		}
		if i < 2 {
			t.Logf("RunSimpleL2Burn attempt %d failed, retrying: %v", i+1, burnErr)
			time.Sleep(5 * time.Second)
		}
	}
	require.NoError(t, burnErr)
	t.Logf("Generated L2 transactions")

	// Wait for batcher to publish
	time.Sleep(30 * time.Second)
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

	// Ensure verifier services are still running after switch
	require.NoError(t, d.ServiceUp("op-geth-verifier"), "op-geth-verifier should be running after switch")
	require.NoError(t, d.ServiceUp("op-node-verifier"), "op-node-verifier should be running after switch")

	// Wait for services to stabilize and process backlog
	time.Sleep(10 * time.Second)

	startBlockAfter, err := d.L1.BlockNumber(ctx)
	require.NoError(t, err)
	t.Logf("After switch, starting from block %d", startBlockAfter)

	time.Sleep(30 * time.Second)
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
