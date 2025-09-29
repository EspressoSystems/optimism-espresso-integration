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

	err = wait.ForNextBlock(ctx, d.L2Verif)
	require.NoError(t, err)

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
	receipt *types.Receipt) {
	// Get contract addresses from SystemConfig
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)

	disputeGameFactoryAddr, err := systemConfig.DisputeGameFactory(&bind.CallOpts{})
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	// Wait for the L2 output to be published as a dispute game on L1
	gameCtx, gameCancel := context.WithTimeout(ctx, 5*time.Minute)
	defer gameCancel()

	_, err = wait.ForGamePublished(gameCtx, d.L1, optimismPortalAddr, disputeGameFactoryAddr, receipt.BlockNumber)
	require.NoError(t, err)

}

func depositOnL1Bridge(d *Devnet,
	ctx context.Context,
	t *testing.T,
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

}

func proveWithdrawalTransaction(d *Devnet,
	ctx context.Context,
	t *testing.T,
	tx *types.Transaction,

) (common.Hash, bindings.TypesWithdrawalTransaction) {

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

	// Generate withdrawal proof parameters using fault proofs
	params, err := withdrawals.ProveWithdrawalParametersFaultProofs(ctx, proofCl, receiptCl, headerCl, tx.Hash(), factory, portal2)
	require.NoError(t, err)

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

	// Wait for the withdrawal delay period before finalization
	withdrawalHash, err := wd.Hash()
	require.NoError(t, err)
	pw, err := portal2.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash, l1Opts.From)
	require.NoError(t, err)
	require.NotEqual(t, pw.DisputeGameProxy, common.Address{0x0})
	require.GreaterOrEqual(t, pw.Timestamp, uint64(1))

	return withdrawalHash, withdrawalTx
}

func waitForResolvedGame(d *Devnet,
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

	// Get the withdrawal delay from the devnet (this is the PROOF_MATURITY_DELAY_SECONDS)
	withdrawalDelay, err := portal2.ProofMaturityDelaySeconds(&bind.CallOpts{})
	require.NoError(t, err)

	// Calculate target time when withdrawal can be finalized
	// The contract requires: block.timestamp provenWithdrawal.timestamp > PROOF_MATURITY_DELAY_SECONDS
	// So we need to wait for PROOF_MATURITY_DELAY_SECONDS + 1 second to be safe
	withdrawalDelayDuration := time.Duration(withdrawalDelay.Int64()) * time.Second
	targetTime := time.Unix(int64(pw.Timestamp), 0).Add(withdrawalDelayDuration).Add(1 * time.Second)

	// Poll L1 latest header time until we pass targetTime
	err = wait.For(ctx, time.Second, func() (bool, error) {
		hdr, err := d.L1.HeaderByNumber(ctx, nil)
		if err != nil {
			return false, err
		}
		return int64(hdr.Time) >= targetTime.Unix(), nil
	})
	require.NoError(t, err)

	// Create dispute game contract binding to query parameters (for logging/debugging)
	disputeGame, err := bindings.NewFaultDisputeGame(pw.DisputeGameProxy, d.L1)
	require.NoError(t, err)

	// Query the actual MAX_CLOCK_DURATION from the contract
	maxClockDuration, err := disputeGame.MaxClockDuration(&bind.CallOpts{})
	require.NoError(t, err)

	// Verify that the dispute game clock duration is set correctly for devnet testing
	expectedClockDuration := uint64(10) // 10 seconds - ultra-fast testing configuration
	require.Equal(t, expectedClockDuration, maxClockDuration,
		"Dispute game clock duration must be %d seconds for devnet testing. "+
			"Current value is %d seconds. Please check that prepare-allocs.sh correctly sets "+
			"faultGameMaxClockDuration to %d in the chain configuration.",
		expectedClockDuration, maxClockDuration, expectedClockDuration)

	maxClockDurationTime := time.Duration(maxClockDuration) * time.Second

	// Wait for the dispute game clock to expire plus additional finality delay
	// The game should resolve automatically after the clock expires
	disputeGameFinalityDelay := 30 * time.Second // Additional safety buffer for game finality
	totalWaitTime := maxClockDurationTime + disputeGameFinalityDelay

	time.Sleep(totalWaitTime)

	// Check final game status (should be auto-resolved by now)
	gameStatus, err := disputeGame.Status(&bind.CallOpts{})
	require.NoError(t, err)
	require.NotEqual(t, gameStatus, 0, "Dispute game should have resolved automatically")

}

func finalizeWithdrawl(d *Devnet,
	ctx context.Context,
	t *testing.T,
	withdrawalHash common.Hash,
	userAddress common.Address,
	withdrawalTx bindings.TypesWithdrawalTransaction,
	withdrawalAmount *big.Int) {

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

	// Finalize the withdrawal
	finalizeTx, err := portal.FinalizeWithdrawalTransaction(finalizeOpts, withdrawalTx)
	require.NoError(t, err, "Finalize withdrawal transaction")

	// Wait for finalization transaction to be mined
	finalizeCtx, finalizeCancel := context.WithTimeout(ctx, 5*time.Minute)
	defer finalizeCancel()
	finalizeReceipt, err := wait.ForReceiptOK(finalizeCtx, d.L1, finalizeTx.Hash())
	require.NoError(t, err, "Finalize withdrawal")

	// Check Alice's L1 balance after finalization
	_, err = wait.ForBalanceChange(ctx, d.L1, userAddress, withdrawalAmount)
	require.NoError(t, err)
	aliceL1BalanceAfter, err := d.L1.BalanceAt(ctx, userAddress, nil)
	require.NoError(t, err)

	// Calculate and verify the balance change
	balanceChange := new(big.Int).Sub(aliceL1BalanceAfter, aliceL1BalanceBefore)
	fees := new(big.Int).Mul(new(big.Int).SetUint64(finalizeReceipt.GasUsed), finalizeReceipt.EffectiveGasPrice)
	expectedBalanceChange := new(big.Int).Sub(withdrawalAmount, fees)
	require.True(t, balanceChange.Cmp(expectedBalanceChange) == 0)
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

	// Check Alice's balance on L2 verifier before withdrawal
	checkUserBalance(d, ctx, t, aliceAddress)

	withdrawalAmount := new(big.Int)
	withdrawalAmount.SetString("1000000000000000000", 10) // 1 ETH in wei
	tx, receipt := initiateWithdrawalOnL2(d, ctx, t, aliceAddress, withdrawalAmount)

	// Check Alice's balance on L1 before withdrawal
	expectedBalance := new(big.Int)
	expectedBalance.SetString("10000000000000000000000", 10) // 10,000 ETH in wei
	checkUserBalanceOnL1(d, ctx, t, aliceAddress, expectedBalance)

	// Deposit some ETH on the L1 bridge so that it is possible to withdraw later
	depositAmount := new(big.Int).Mul(withdrawalAmount, big.NewInt(2))
	depositOnL1Bridge(d, ctx, t, depositAmount)

	// Wait for the game to be published
	waitForGameToBePublished(d, ctx, t, receipt)

	// Generate withdrawal proof
	withdrawalHash, withdrawalTx := proveWithdrawalTransaction(d, ctx, t, tx)

	// Wait for the game to be resolved
	waitForResolvedGame(d, ctx, t, withdrawalHash, aliceAddress)

	// Transfer the funds to Alice on L1
	finalizeWithdrawl(d, ctx, t, withdrawalHash, aliceAddress, withdrawalTx, withdrawalAmount)

}
