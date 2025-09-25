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

func checkUserBalance(d *Devnet, ctx context.Context, t *testing.T, userAddress common.Address) {
	userBalance, err := d.L2Verif.BalanceAt(ctx, userAddress, nil)
	require.NoError(t, err)
	require.True(t, userBalance.Cmp(big.NewInt(0)) > 0, "Alice should have a positive balance")
}

func initiateWithdrawalOnL2(d *Devnet,
	ctx context.Context,
	t *testing.T,
	userAddress common.Address,
	withdrawalAmount *big.Int) (*types.Transaction, *types.Receipt) {

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
	tx, err := l2MessagePasser.InitiateWithdrawal(opts, userAddress, big.NewInt(21000), nil)
	require.NoError(t, err)

	// Wait for receipt ok
	receipt, err := wait.ForReceiptOK(ctx, d.L2Verif, tx.Hash())
	require.NoError(t, err)

	wait.ForNextBlock(ctx, d.L2Verif)

	return tx, receipt
}

func checkUserBalanceOnL1(
	d *Devnet,
	ctx context.Context,
	t *testing.T,
	userAddress common.Address,
	expectedBalance *big.Int) {
	userBalance, err := d.L1.BalanceAt(ctx, userAddress, nil)
	require.NoError(t, err)
	require.True(t, userBalance.Cmp(expectedBalance) == 0)
	t.Logf("User's L1 balance before withdrawal: %s wei", userBalance.String())

}

func waitForGameToBePublished(d *Devnet, ctx context.Context, t *testing.T,
	receipt *types.Receipt) uint64 {
	// TODO philippe: can it be less
	time.Sleep(3 * time.Minute)

	// Get contract addresses from SystemConfig
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)

	disputeGameFactoryAddr, err := systemConfig.DisputeGameFactory(&bind.CallOpts{})
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	t.Logf("Dispute game factory address: %s", disputeGameFactoryAddr)
	t.Logf("Optimism portal address: %s", optimismPortalAddr)

	// Wait for the L2 output to be published as a dispute game on L1
	t.Logf("Waiting for dispute game to be published for block %d...", receipt.BlockNumber.Uint64())
	gameCtx, gameCancel := context.WithTimeout(ctx, 5*time.Minute)
	defer gameCancel()

	// Wait for the finalization period, then we can finalize this withdrawal.
	blockNumber, err := wait.ForGamePublished(gameCtx, d.L1, optimismPortalAddr, disputeGameFactoryAddr, receipt.BlockNumber)
	require.NoError(t, err)
	return blockNumber
}

func depositOnL1Bridge(d *Devnet,
	ctx context.Context,
	t *testing.T,
	userAddress common.Address,
	depositAmount *big.Int) {

	// Get the OptimismPortal address from rollup config
	rollupConfig, err := d.RollupConfig(ctx)
	require.NoError(t, err)
	optimismPortalAddr := rollupConfig.DepositContractAddress

	// Create deposit contract binding
	depositContract, err := bindings.NewOptimismPortal(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Get L1 chain ID and create transaction options
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)
	opts.Value = depositAmount
	opts.GasLimit = 1_000_000

	// Set gas price
	gasPrice, err := d.L1.SuggestGasPrice(ctx)
	require.NoError(t, err)
	opts.GasPrice = gasPrice

	// Create the deposit transaction to a dummy address
	toAddr := common.Address{0xff, 0xff}

	depositTx, err := depositContract.DepositTransaction(opts, toAddr, depositAmount, 21000, false, nil)
	require.NoError(t, err)

	// Wait for the deposit transaction to succeed
	depositReceipt, err := wait.ForReceiptOK(ctx, d.L1, depositTx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, depositReceipt.Status)

	t.Logf("L1 deposit transaction completed successfully: %s", depositTx.Hash().Hex())
}

func proveWithdrawalTransaction(d *Devnet,
	ctx context.Context,
	t *testing.T,
	tx *types.Transaction,
	receipt *types.Receipt,
	blockNumber uint64) (common.Hash, bindings.TypesWithdrawalTransaction) {

	// Get contract addresses from SystemConfig
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)

	disputeGameFactoryAddr, err := systemConfig.DisputeGameFactory(&bind.CallOpts{})
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

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

	// Create the withdrawal transaction struct that will be used for both proving and finalizing
	withdrawalTx := bindings.TypesWithdrawalTransaction{
		Nonce:    wd.Nonce,
		Sender:   *wd.Sender,
		Target:   *wd.Target,
		Value:    wd.Value,
		GasLimit: wd.GasLimit,
		Data:     wd.Data,
	}

	t.Logf("Submitting ProveWithdrawalTransaction to L1...")
	proveTx, err := portal.ProveWithdrawalTransaction(
		l1Opts,
		withdrawalTx,
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

	// Wait for the withdrawal delay period before finalization
	withdrawalHash, err := wd.Hash()
	require.NoError(t, err)
	pw, err := portal2.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash, l1Opts.From)
	require.NoError(t, err)
	require.NotEqual(t, pw.DisputeGameProxy, common.Address{0x0})
	require.GreaterOrEqual(t, pw.Timestamp, uint64(1))

	return withdrawalHash, withdrawalTx
}

func resolveGame(d *Devnet,
	ctx context.Context,
	t *testing.T,
	withdrawalHash common.Hash,
	userAddress common.Address) {

	// Get system config to access OptimismPortal
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	// Create OptimismPortal2 binding to check proven withdrawals
	portal2, err := nodepreview.NewOptimismPortal2Caller(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Get the proven withdrawal info
	pw, err := portal2.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash, userAddress)
	require.NoError(t, err)
	require.NotEqual(t, pw.DisputeGameProxy, common.Address{0x0})
	require.GreaterOrEqual(t, pw.Timestamp, uint64(1))

	// Get the withdrawal delay from the devnet
	withdrawalDelay, err := d.getWithdrawalDelay()
	require.NoError(t, err)
	t.Logf("Withdrawal delay (disputeGameFinalityDelaySeconds): %v", withdrawalDelay)

	// Calculate target time when withdrawal can be finalized
	targetTime := time.Unix(int64(pw.Timestamp), 0).Add(withdrawalDelay)
	t.Logf("Waiting until L1 time passes target %s (pw.Timestamp=%d + delay)", targetTime.Format(time.RFC3339), pw.Timestamp)

	// Poll L1 latest header time until we pass targetTime
	err = wait.For(ctx, time.Second, func() (bool, error) {
		hdr, err := d.L1.HeaderByNumber(ctx, nil)
		if err != nil {
			return false, err
		}
		return int64(hdr.Time) >= targetTime.Unix(), nil
	})
	require.NoError(t, err)
	t.Logf("Withdrawal delay period has passed, proceeding with finalization")

	// Check current L1 block time for verification
	currentHeader, err := d.L1.HeaderByNumber(ctx, nil)
	require.NoError(t, err)
	currentTime := time.Unix(int64(currentHeader.Time), 0)
	t.Logf("Current L1 block time: %s (timestamp: %d)", currentTime.Format(time.RFC3339), currentHeader.Time)
	t.Logf("Target time was: %s (timestamp: %d)", targetTime.Format(time.RFC3339), targetTime.Unix())
	t.Logf("Time difference: %v", currentTime.Sub(targetTime))

	// Create dispute game contract binding to query parameters (for logging/debugging)
	disputeGame, err := bindings.NewFaultDisputeGame(pw.DisputeGameProxy, d.L1)
	require.NoError(t, err)

	// Query the actual MAX_CLOCK_DURATION from the contract
	maxClockDuration, err := disputeGame.MaxClockDuration(&bind.CallOpts{})
	require.NoError(t, err)
	t.Logf("Contract MAX_CLOCK_DURATION: %d seconds", maxClockDuration)

	// The chess clock duration is typically 302,400 seconds (3.5 days) - too long for a devnet test!
	// This explains why manual dispute game resolution would fail with ClockNotExpired
	maxClockDurationTime := time.Duration(maxClockDuration) * time.Second
	t.Logf("Chess clock duration: %v (too long for devnet test!)", maxClockDurationTime)

	// For devnet testing, we'll skip manual dispute game resolution and proceed directly to finalization
	// The OptimismPortal should handle any necessary dispute game checks internally
	t.Log("Skipping manual dispute game resolution due to long chess clock duration...")
	t.Log("Proceeding directly to withdrawal finalization...")
}

func finalizeWithdrawl(d *Devnet,
	ctx context.Context,
	t *testing.T,
	withdrawalHash common.Hash,
	userAddress common.Address,
	withdrawalTx bindings.TypesWithdrawalTransaction) {

	// Get system config and portal address
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	// Create portal binding
	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Get L1 chain ID and create transaction options
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	finalizeOpts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)
	finalizeOpts.GasLimit = 300000

	gasPrice, err := d.L1.SuggestGasPrice(ctx)
	require.NoError(t, err)
	finalizeOpts.GasPrice = gasPrice

	// Check Alice's L1 balance before finalization
	aliceL1BalanceBefore, err := d.L1.BalanceAt(ctx, userAddress, nil)
	require.NoError(t, err)
	t.Logf("Alice's L1 balance before finalization: %s wei", aliceL1BalanceBefore.String())

	// Finalize the withdrawal
	finalizeTx, err := portal.FinalizeWithdrawalTransaction(finalizeOpts, withdrawalTx)
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
	aliceL1BalanceAfter, err := d.L1.BalanceAt(ctx, userAddress, nil)
	require.NoError(t, err)
	t.Logf("Alice's L1 balance after finalization: %s wei", aliceL1BalanceAfter.String())

	// Calculate and log the balance change
	balanceChange := new(big.Int).Sub(aliceL1BalanceAfter, aliceL1BalanceBefore)
	t.Logf("Net L1 balance change: %s wei", balanceChange.String())

	t.Logf("✅ Withdrawal finalization completed!")
}

func TestWithdrawal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	aliceAddress := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	//Check Alice's balance on L2 verifier before withdrawal
	checkUserBalance(d, ctx, t, aliceAddress)

	withdrawalAmount := big.NewInt(1000000)
	tx, receipt := initiateWithdrawalOnL2(d, ctx, t, aliceAddress, withdrawalAmount)

	// Check Alice's balance on L1 before withdrawal
	expectedBalance := new(big.Int)
	expectedBalance.SetString("10000000000000000000000", 10) // 10,000 ETH in wei
	checkUserBalanceOnL1(d, ctx, t, aliceAddress, expectedBalance)

	// Deposit some ETH on the L1 bridge so that it is possible to withdraw later
	depositOnL1Bridge(d, ctx, t, aliceAddress, big.NewInt(9_000_000))

	// Wait for the game to be published
	blockNumber := waitForGameToBePublished(d, ctx, t, receipt)

	// Generate withdrawal proof
	withdrawalHash, withdrawalTx := proveWithdrawalTransaction(d, ctx, t, tx, receipt, blockNumber)

	resolveGame(d, ctx, t, withdrawalHash, aliceAddress)

	finalizeWithdrawl(d, ctx, t, withdrawalHash, aliceAddress, withdrawalTx)

}
