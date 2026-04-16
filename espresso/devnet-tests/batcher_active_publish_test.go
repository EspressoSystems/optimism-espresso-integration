package devnet_tests

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/espresso/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// hasBatchTransactions checks if any transactions were sent to the BatchInbox from the given sender.
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
	profile := ProfileFromEnv(t)
	t.Run(string(profile), func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
		defer cancel()

		d := NewDevnet(ctx, t, profile)
		require.NoError(t, d.Up())
		defer func() {
			require.NoError(t, d.Down())
		}()

		require.NoError(t, d.WaitForBatcher(ctx))

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

		espressoBatcherAddr, err := batchAuthenticator.EspressoBatcher(&bind.CallOpts{})
		require.NoError(t, err)
		fallbackBatcherAddr := config.Genesis.SystemConfig.BatcherAddr

		activeIsEspresso, err := batchAuthenticator.ActiveIsEspresso(&bind.CallOpts{})
		require.NoError(t, err)
		t.Logf("Initial state: activeIsEspresso = %v", activeIsEspresso)

		// verifyPublishing helper function
		verifyPublishing := func(expectEspressoActive bool) {
			t.Logf("Verifying publishing for state: expectEspressoActive=%v", expectEspressoActive)

			startBlock, err := d.L1.BlockNumber(ctx)
			require.NoError(t, err)
			t.Logf("Starting from block %d", startBlock)

			// Generate L2 traffic
			burnReceipt, err := d.SubmitSimpleL2Burn()
			require.NoError(t, err)
			t.Logf("Generated L2 transaction: %s (L2 block %d)", burnReceipt.Receipt.TxHash, burnReceipt.Receipt.BlockNumber)

			// Wait for batcher to publish
			// We wait long enough for the active batcher to publish, but not so long that we timeout the test
			// The idle batcher check inside the driver should prevent it from publishing
			time.Sleep(60 * time.Second)
			t.Logf("Waited 60s for L1 confirmation")

			endBlock, err := d.L1.BlockNumber(ctx)
			require.NoError(t, err)
			t.Logf("Checking blocks %d-%d", startBlock, endBlock)

			espressoPublished, err := hasBatchTransactions(ctx, d.L1, config.BatchInboxAddress, espressoBatcherAddr, startBlock, endBlock)
			require.NoError(t, err)
			fallbackPublished, err := hasBatchTransactions(ctx, d.L1, config.BatchInboxAddress, fallbackBatcherAddr, startBlock, endBlock)
			require.NoError(t, err)

			t.Logf("Espresso batcher published: %v, fallback batcher published: %v", espressoPublished, fallbackPublished)

			if expectEspressoActive {
				require.True(t, espressoPublished, "Espresso batcher should publish when active")
				require.False(t, fallbackPublished, "fallback batcher should NOT publish when inactive")
			} else {
				require.True(t, fallbackPublished, "fallback batcher should publish when active")
				require.False(t, espressoPublished, "Espresso batcher should NOT publish when inactive")
			}
		}

		// 1. Verify initial state
		verifyPublishing(activeIsEspresso)

		// 2. Switch state
		t.Logf("Switching batcher state...")
		switchTx, err := batchAuthenticator.SwitchBatcher(deployerOpts)
		require.NoError(t, err)
		receipt, err := wait.ForReceiptOK(ctx, d.L1, switchTx.Hash())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

		// Update expected state
		activeIsEspresso = !activeIsEspresso
		t.Logf("Switched state to: activeIsEspresso=%v", activeIsEspresso)

		// Wait for services to stabilize after switch. In-flight sendTxWithEspresso goroutines
		// spawned before deactivation can take ~25s to drain their queued Txmgr.Send calls,
		// so we wait long enough for all residual transactions to land on L1 before capturing
		// startBlock in the next verifyPublishing call.
		time.Sleep(60 * time.Second)

		// 3. Verify new state
		verifyPublishing(activeIsEspresso)
	})
}
