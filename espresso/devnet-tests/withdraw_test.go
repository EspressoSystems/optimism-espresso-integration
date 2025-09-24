package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain"
	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	nodebindings "github.com/ethereum-optimism/optimism/op-node/bindings"
	nodepreview "github.com/ethereum-optimism/optimism/op-node/bindings/preview"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/stretchr/testify/require"
)

func TestWithdrawal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	//Check Alice's balance on L2 verifier before withdrawal
	aliceAddress := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)
	aliceBalance, err := d.L2Verif.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	require.True(t, aliceBalance.Cmp(big.NewInt(0)) > 0, "Alice should have a positive balance")

	// Initiate withdrawal on L2
	withdrawalAmount := big.NewInt(1000000) // Withdraw 1 000 000 wei

	// Bind to L2ToL1MessagePasser contract
	l2ToL1MessagePasserAddr := common.HexToAddress("0x4200000000000000000000000000000000000016") // L2ToL1MessagePasser predeploy
	l2MessagePasser, err := bindings.NewL2ToL1MessagePasser(l2ToL1MessagePasserAddr, d.L2Seq)
	require.NoError(t, err)

	// Get the correct L2 chain ID from the sequencer
	chainID, err := d.L2Seq.ChainID(ctx)
	require.NoError(t, err)

	// Create transaction options for Alice
	opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, chainID)
	require.NoError(t, err)
	opts.Value = withdrawalAmount

	// Initiate withdrawal - this sends ETH to L2ToL1MessagePasser and emits an event
	tx, err := l2MessagePasser.InitiateWithdrawal(opts, aliceAddress, big.NewInt(21000), nil)
	require.NoError(t, err)

	// Wait for receipt ok
	receipt, err := wait.ForReceiptOK(ctx, d.L2Verif, tx.Hash())
	require.NoError(t, err)

	// Check Alice's balance on L1 before withdrawal
	aliceL1Balance, err := d.L1.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	expectedBalance := new(big.Int)
	expectedBalance.SetString("10000000000000000000000", 10) // 10,000 ETH in wei
	require.True(t, aliceL1Balance.Cmp(expectedBalance) == 0, "Alice should have exactly 10,000 ETH")
	t.Logf("Alice's L1 balance before withdrawal: %s wei", aliceL1Balance.String())

	wait.ForNextBlock(ctx, d.L2Verif)

	// // Check proposer account balance and permissions
	// // The proposer uses the same mnemonic as batcher: "test test test test test test test test test test test junk"
	// proposerPrivKey := d.secrets.Alice // This should match the mnemonic account
	// proposerAddr := crypto.PubkeyToAddress(proposerPrivKey.PublicKey)
	// proposerBalance, err := d.L1.BalanceAt(ctx, proposerAddr, nil)
	// require.NoError(t, err)
	// t.Logf("Proposer account %s L1 balance: %s ETH", proposerAddr.Hex(), new(big.Int).Div(proposerBalance, big.NewInt(1e18)).String())

	// Get contract addresses from SystemConfig
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)

	disputeGameFactoryAddr, err := systemConfig.DisputeGameFactory(&bind.CallOpts{})
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	// Wait for the finalization period, then we can finalize this withdrawal.
	blockNumber, err := wait.ForGamePublished(ctx, d.L1, optimismPortalAddr, disputeGameFactoryAddr, receipt.BlockNumber)
	require.Nil(t, err)

	// Generate withdrawal proof using fault proofs
	t.Logf("Generating withdrawal proof for transaction %s", tx.Hash().Hex())

	// Set up clients for proof generation
	receiptCl := d.L2Seq
	headerCl := d.L2Seq
	proofCl := gethclient.New(receiptCl.Client())

	// Set up contract bindings for proof generation
	factory, err := nodebindings.NewDisputeGameFactoryCaller(disputeGameFactoryAddr, d.L1)
	require.NoError(t, err)

	portal2, err := nodepreview.NewOptimismPortal2Caller(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Get the block header for proof generation
	header, err := receiptCl.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	require.NoError(t, err)
	t.Logf("Using block %d (hash: %s) for withdrawal proof", blockNumber, header.Hash().Hex())

	// Generate withdrawal proof parameters using fault proofs
	t.Logf("Generating withdrawal proof using fault proofs...")
	params, err := withdrawals.ProveWithdrawalParametersFaultProofs(ctx, proofCl, receiptCl, headerCl, tx.Hash(), factory, portal2)
	require.NoError(t, err)
	t.Logf("Fault proofs withdrawal parameters generated successfully")

	// Bind to OptimismPortal contract on L1
	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Create transaction options for Alice on L1
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	l1Opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)

	// Set proper gas configuration
	l1Opts.GasLimit = 500000
	gasPrice, err := d.L1.SuggestGasPrice(ctx)
	require.NoError(t, err)
	l1Opts.GasPrice = gasPrice

	// Build the withdrawal message with correct field types
	wd := crossdomain.NewWithdrawal(
		params.Nonce,
		&params.Sender,
		&params.Target,
		params.Value,
		params.GasLimit,
		params.Data,
	)

	t.Logf("Submitting ProveWithdrawalTransaction to L1...")
	proveTx, err := portal.ProveWithdrawalTransaction(
		l1Opts,
		bindings.TypesWithdrawalTransaction{
			Nonce:    wd.Nonce,
			Sender:   *wd.Sender,
			Target:   *wd.Target,
			Value:    wd.Value,
			GasLimit: wd.GasLimit,
			Data:     wd.Data,
		},
		params.L2OutputIndex,
		bindings.TypesOutputRootProof{
			Version:                  params.OutputRootProof.Version,
			StateRoot:                params.OutputRootProof.StateRoot,
			MessagePasserStorageRoot: params.OutputRootProof.MessagePasserStorageRoot,
			LatestBlockhash:          params.OutputRootProof.LatestBlockhash,
		},
		params.WithdrawalProof,
	)
	require.NoError(t, err)

	// Wait for the proof transaction to be mined
	proveReceipt, err := bind.WaitMined(ctx, d.L1, proveTx)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, proveReceipt.Status)

	t.Logf("Withdrawal proof transaction successful: %s", proveTx.Hash().Hex())
	t.Logf("Gas used: %d", proveReceipt.GasUsed)
	t.Logf("Withdrawal can now be finalized after the finalization period")

	// Resolve the dispute game before checking maturity to ensure the root claim is accepted.
	wdHash, err := wd.Hash()
	require.NoError(t, err)
	pwBeforeResolve, err := portal2.ProvenWithdrawals(&bind.CallOpts{}, wdHash, l1Opts.From)
	require.NoError(t, err)
	require.NotEqual(t, pwBeforeResolve.DisputeGameProxy, common.Address{0x0})

	// Check Alice's L1 balance before finalization
	aliceL1BalanceBefore, err := d.L1.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	t.Logf("Alice's L1 balance before finalization: %s ETH", new(big.Int).Div(aliceL1BalanceBefore, big.NewInt(1e18)).String())

	// Finalize the withdrawal
	t.Logf("Finalizing withdrawal transaction...")

	// Create new transaction options for finalization
	finalizeOpts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)
	finalizeOpts.GasLimit = 300000
	finalizeOpts.GasPrice = gasPrice

	finalizeTx, err := portal.FinalizeWithdrawalTransaction(
		finalizeOpts,
		bindings.TypesWithdrawalTransaction{
			Nonce:    params.Nonce,
			Sender:   params.Sender,
			Target:   params.Target,
			Value:    params.Value,
			GasLimit: params.GasLimit,
			Data:     params.Data,
		},
	)
	require.NoError(t, err)

	// Wait for finalization transaction to be mined
	finalizeCtx, finalizeCancel := context.WithTimeout(ctx, 5*time.Minute)
	defer finalizeCancel()
	finalizeReceipt, err := wait.ForReceiptOK(finalizeCtx, d.L1, finalizeTx.Hash())
	require.NoError(t, err, "finalize withdrawal")
	require.Equal(t, types.ReceiptStatusSuccessful, finalizeReceipt.Status)

	t.Logf("Withdrawal finalization successful: %s", finalizeTx.Hash().Hex())
	t.Logf("Finalization gas used: %d", finalizeReceipt.GasUsed)

	// Check Alice's L1 balance after finalization
	aliceL1BalanceAfter, err := d.L1.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	t.Logf("Alice's L1 balance after finalization: %s ETH", new(big.Int).Div(aliceL1BalanceAfter, big.NewInt(1e18)).String())

	// Require that finalization was successful
	require.Equal(t, types.ReceiptStatusSuccessful, finalizeReceipt.Status, "Withdrawal finalization must succeed")

	// Calculate the net change (accounting for gas costs)
	balanceChange := new(big.Int).Sub(aliceL1BalanceAfter, aliceL1BalanceBefore)
	t.Logf("Net L1 balance change: %s wei", balanceChange.String())
	t.Logf("Net L1 balance change: %s ETH", new(big.Int).Div(balanceChange, big.NewInt(1e18)).String())

	// Calculate expected gas costs for both prove and finalize transactions
	totalGasCost := new(big.Int).Mul(big.NewInt(int64(proveReceipt.GasUsed+finalizeReceipt.GasUsed)), gasPrice)
	t.Logf("Total gas cost: %s wei", totalGasCost.String())
	t.Logf("Withdrawal amount: %s wei", withdrawalAmount.String())

	// Calculate expected balance change: withdrawal amount minus gas costs
	expectedChange := new(big.Int).Sub(withdrawalAmount, totalGasCost)
	t.Logf("Expected balance change: %s wei", expectedChange.String())

	// Verify the balance change matches expectations
	// The balance should increase by exactly the withdrawal amount minus gas costs
	require.Equal(t, expectedChange, balanceChange,
		"Balance change should equal withdrawal amount (%s wei) minus gas costs (%s wei)",
		withdrawalAmount.String(), totalGasCost.String())

	// Additional verification: ensure Alice actually received the withdrawal amount
	// by checking that the balance increased by at least the withdrawal amount minus reasonable gas costs
	minExpectedIncrease := new(big.Int).Sub(withdrawalAmount, big.NewInt(1e15)) // Allow up to 0.001 ETH for gas
	require.True(t, balanceChange.Cmp(minExpectedIncrease) >= 0,
		"Balance should increase by at least %s wei (withdrawal minus reasonable gas), but only increased by %s wei",
		minExpectedIncrease.String(), balanceChange.String())

	t.Logf("✅ Withdrawal verification successful!")
	t.Logf("✅ Alice's L1 balance increased by %s wei as expected", balanceChange.String())

	// Note: After finalization, the ProvenWithdrawals mapping entry may be cleared
	// so we don't query it again to avoid "execution reverted" errors.
	// The successful finalization transaction receipt and balance verification confirm completion.

}
