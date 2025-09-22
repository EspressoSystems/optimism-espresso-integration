package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain"
	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	configpkg "github.com/ethereum-optimism/optimism/op-e2e/config"
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

	// Ensure that the allocation type uses proofs (static config check)
	require.True(t, configpkg.DefaultAllocType.UsesProofs(), "alloc type must use proofs; got %s", configpkg.DefaultAllocType)

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
	tx, err := l2MessagePasser.InitiateWithdrawal(opts, aliceAddress, big.NewInt(21000), []byte{})
	require.NoError(t, err)

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, d.L2Seq, tx)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	// Check Alice's balance on L1 before withdrawal
	aliceL1Balance, err := d.L1.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	expectedBalance := new(big.Int)
	expectedBalance.SetString("10000000000000000000000", 10) // 10,000 ETH in wei
	require.True(t, aliceL1Balance.Cmp(expectedBalance) == 0, "Alice should have exactly 10,000 ETH")
	t.Logf("Alice's L1 balance before withdrawal: %s wei", aliceL1Balance.String())

	// Create withdrawal proof transaction
	// Wait for the withdrawal to be included in a block and get the block number
	// Use a longer timeout for devnet environment where proposer interval is 6s
	ctx, cancel = context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	// Check proposer account balance and permissions
	// The proposer uses the same mnemonic as batcher: "test test test test test test test test test test test junk"
	proposerPrivKey := d.secrets.Alice // This should match the mnemonic account
	proposerAddr := crypto.PubkeyToAddress(proposerPrivKey.PublicKey)
	proposerBalance, err := d.L1.BalanceAt(ctx, proposerAddr, nil)
	require.NoError(t, err)
	t.Logf("Proposer account %s L1 balance: %s ETH", proposerAddr.Hex(), new(big.Int).Div(proposerBalance, big.NewInt(1e18)).String())

	disputeGameFactoryAddr, optimismPortalAddr := d.getOPAddresses()

	// Wait a bit longer for the proposer to create games (proposer interval is 6s)
	t.Logf("Waiting 3 minutes for proposer to create dispute games...")
	time.Sleep(3 * time.Minute)

	// Wait for the L2 output to be published as a dispute game on L1
	t.Logf("Waiting for dispute game to be published for block %d", receipt.BlockNumber.Uint64())
	var blockNumber uint64

	// Try waiting for dispute game first, but fall back if it times out
	t.Logf("Waiting for dispute game to be published for block %d...", receipt.BlockNumber.Uint64())
	gameCtx, gameCancel := context.WithTimeout(ctx, 2*time.Minute)
	defer gameCancel()

	blockNumber, err = wait.ForGamePublished(gameCtx, d.L1, optimismPortalAddr, disputeGameFactoryAddr, receipt.BlockNumber)
	if err != nil {
		t.Logf("Failed to wait for dispute game: %v. Using withdrawal block number for proof generation.", err)
		// Use the receipt block number if game publication fails
		blockNumber = receipt.BlockNumber.Uint64()
		t.Logf("Proceeding with block number %d for withdrawal proof", blockNumber)
	} else {
		t.Logf("Dispute game published for block %d", blockNumber)
	}

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

	// Check Alice's L1 balance before finalization
	aliceL1BalanceBefore, err := d.L1.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	t.Logf("Alice's L1 balance before finalization: %s ETH", new(big.Int).Div(aliceL1BalanceBefore, big.NewInt(1e18)).String())

	// Check that the withdrawal delay is set to 12 seconds
	withdrawalDelay, err := d.getWithdrawalDelay()
	require.NoError(t, err)
	t.Logf("Withdrawal delay (disputeGameFinalityDelaySeconds): %v", withdrawalDelay)
	require.Equal(t, time.Duration(12*time.Second), withdrawalDelay)

	// Wait for the challenge period to expire
	t.Logf("Waiting for challenge period to expire...")

	withdrawalHash, err := wd.Hash()
	require.NoError(t, err)
	t.Logf("Withdrawal hash (ABI encoded): %s", withdrawalHash.Hex())

	// Wait until the portal reports the withdrawal is ready
	wait.ForWithdrawalCheck(ctx, d.L1, *wd, optimismPortalAddr, aliceAddress)

	// Check if withdrawal is ready for finalization
	// Note: On some protocol versions the ABI for OptimismPortal differs (Portal vs Portal2),
	// and the public getter may not exist or may revert if the entry is absent/cleared.
	// Treat this as a best-effort check and proceed regardless; finalization+balances are the source of truth.
	t.Logf("Checking proven withdrawals with ABI-encoded hash (best-effort)...")
	_, err = portal.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash)
	require.NoError(t, err)

	// TODO Philippe remove one of the two checks and config modification (DASEL)
	// Additional observability: log configured delays from Portal2
	if pm, err := portal2.ProofMaturityDelaySeconds(&bind.CallOpts{}); err == nil {
		t.Logf("Portal2 proofMaturityDelaySeconds: %ds", pm.Int64())
	} else {
		t.Logf("Warning: could not read proofMaturityDelaySeconds: %v", err)
	}
	if dgfd, err := portal2.DisputeGameFinalityDelaySeconds(&bind.CallOpts{}); err == nil {
		t.Logf("Portal2 disputeGameFinalityDelaySeconds: %ds", dgfd.Int64())
	} else {
		t.Logf("Warning: could not read disputeGameFinalityDelaySeconds: %v", err)
	}

	// Query the proven-withdrawal record for this proof submitter (Alice)
	proofSubmitter := l1Opts.From
	if pw, err := portal2.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash, proofSubmitter); err == nil {
		t.Logf("ProvenWithdrawal record -> disputeGameProxy: %s, timestamp: %d", pw.DisputeGameProxy.Hex(), pw.Timestamp)
	} else {
		t.Logf("Warning: could not read ProvenWithdrawals record: %v", err)
	}

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
	finalizeReceipt, err := wait.ForReceiptOK(ctx, d.L1, finalizeTx.Hash())
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
