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

// initiateWithdrawalOnL2 initiates a withdrawal on L2 and returns the transaction and receipt
func initiateWithdrawalOnL2(d *Devnet, ctx context.Context, t *testing.T, userAddress common.Address, withdrawalAmount *big.Int) (*types.Transaction, *types.Receipt) {
	// Bind to L2ToL1MessagePasser contract
	l2ToL1MessagePasserAddr := common.HexToAddress("0x4200000000000000000000000000000000000016")
	l2MessagePasser, err := bindings.NewL2ToL1MessagePasser(l2ToL1MessagePasserAddr, d.L2Seq)
	require.NoError(t, err)

	// Create transaction options
	chainID, err := d.L2Seq.ChainID(ctx)
	require.NoError(t, err)
	opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, chainID)
	require.NoError(t, err)
	opts.Value = withdrawalAmount

	// Initiate withdrawal
	tx, err := l2MessagePasser.InitiateWithdrawal(opts, userAddress, big.NewInt(21000), nil)
	require.NoError(t, err)

	// Wait for confirmation
	receipt, err := wait.ForReceiptOK(ctx, d.L2Verif, tx.Hash())
	require.NoError(t, err)
	err = wait.ForNextBlock(ctx, d.L2Verif)
	require.NoError(t, err)

	return tx, receipt
}

// waitForGameToBePublished waits for the dispute game to be published on L1
func waitForGameToBePublished(d *Devnet, ctx context.Context, t *testing.T, receipt *types.Receipt) {
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)

	disputeGameFactoryAddr, err := systemConfig.DisputeGameFactory(&bind.CallOpts{})
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	// Retry up to 5 times with 1-minute timeout each
	for i := 0; i < 5; i++ {
		gameCtx, gameCancel := context.WithTimeout(ctx, 1*time.Minute)
		_, err = wait.ForGamePublished(gameCtx, d.L1, optimismPortalAddr, disputeGameFactoryAddr, receipt.BlockNumber)
		gameCancel()
		if err == nil {
			return
		}
	}
	require.NoError(t, err)
}

// depositOnL1Bridge deposits ETH on L1 to fund withdrawals
func depositOnL1Bridge(d *Devnet, ctx context.Context, t *testing.T, depositAmount *big.Int) {
	rollupConfig, err := d.RollupConfig(ctx)
	require.NoError(t, err)

	depositContract, err := bindings.NewOptimismPortal(rollupConfig.DepositContractAddress, d.L1)
	require.NoError(t, err)

	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)
	opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)
	opts.Value = depositAmount
	opts.GasLimit = 1_000_000
	gasPrice, err := d.L1.SuggestGasPrice(ctx)
	require.NoError(t, err)
	opts.GasPrice = gasPrice

	// Deposit to dummy address
	depositTx, err := depositContract.DepositTransaction(opts, common.Address{0xff, 0xff}, depositAmount, 21000, false, nil)
	require.NoError(t, err)

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

// waitForResolvedGame waits for the dispute game to resolve and withdrawal to be ready for finalization
func waitForResolvedGame(d *Devnet, ctx context.Context, t *testing.T, withdrawalHash common.Hash, userAddress common.Address) {
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	portal2, err := nodepreview.NewOptimismPortal2Caller(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Get proven withdrawal info
	pw, err := portal2.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash, userAddress)
	require.NoError(t, err)
	require.NotEqual(t, pw.DisputeGameProxy, common.Address{0x0})

	// Wait for proof maturity delay
	withdrawalDelay, err := portal2.ProofMaturityDelaySeconds(&bind.CallOpts{})
	require.NoError(t, err)
	withdrawalDelayDuration := time.Duration(withdrawalDelay.Int64()+1) * time.Second // +1 for safety
	targetTime := time.Unix(int64(pw.Timestamp), 0).Add(withdrawalDelayDuration)

	err = wait.For(ctx, time.Second, func() (bool, error) {
		hdr, err := d.L1.HeaderByNumber(ctx, nil)
		if err != nil {
			return false, err
		}
		return int64(hdr.Time) >= targetTime.Unix(), nil
	})
	require.NoError(t, err)

	// Wait for dispute game to auto-resolve (10 seconds + 30 second buffer)
	disputeGame, err := bindings.NewFaultDisputeGame(pw.DisputeGameProxy, d.L1)
	require.NoError(t, err)

	maxClockDuration, err := disputeGame.MaxClockDuration(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, uint64(10), maxClockDuration, "Expected 10-second dispute game clock for devnet")

	// Wait for game resolution
	totalWaitTime := time.Duration(maxClockDuration)*time.Second + 30*time.Second
	time.Sleep(totalWaitTime)

	disputeCtx, disputeCancel := context.WithTimeout(ctx, totalWaitTime)
	defer disputeCancel() // Ensure context resources are released
	err = wait.For(disputeCtx, time.Second, func() (bool, error) {
		gameStatus, err := disputeGame.Status(&bind.CallOpts{})
		require.NoError(t, err)
		require.NotEqual(t, gameStatus, 0, "Dispute game should have resolved automatically")
		return true, nil
	})
	require.NoError(t, err)

}

// finalizeWithdrawal finalizes the withdrawal on L1 and verifies balance change
func finalizeWithdrawal(d *Devnet, ctx context.Context, t *testing.T, userAddress common.Address, withdrawalTx bindings.TypesWithdrawalTransaction, withdrawalAmount *big.Int) {
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)
	optimismPortalAddr, err := systemConfig.OptimismPortal(&bind.CallOpts{})
	require.NoError(t, err)

	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)
	finalizeOpts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)
	finalizeOpts.GasLimit = 300000
	gasPrice, err := d.L1.SuggestGasPrice(ctx)
	require.NoError(t, err)
	finalizeOpts.GasPrice = gasPrice

	// Get balance before finalization
	balanceBefore, err := d.L1.BalanceAt(ctx, userAddress, nil)
	require.NoError(t, err)

	// Finalize withdrawal
	finalizeTx, err := portal.FinalizeWithdrawalTransaction(finalizeOpts, withdrawalTx)
	require.NoError(t, err)

	finalizeReceipt, err := wait.ForReceiptOK(ctx, d.L1, finalizeTx.Hash())
	require.NoError(t, err)

	// Verify balance change
	_, err = wait.ForBalanceChange(ctx, d.L1, userAddress, balanceBefore)
	require.NoError(t, err)
	balanceAfter, err := d.L1.BalanceAt(ctx, userAddress, nil)
	require.NoError(t, err)

	balanceChange := new(big.Int).Sub(balanceAfter, balanceBefore)
	fees := new(big.Int).Mul(new(big.Int).SetUint64(finalizeReceipt.GasUsed), finalizeReceipt.EffectiveGasPrice)
	expectedChange := new(big.Int).Sub(withdrawalAmount, fees)
	require.Equal(t, 0, balanceChange.Cmp(expectedChange), "Balance change should match withdrawal amount minus fees")
}

func TestWithdrawal(t *testing.T) {
	t.Skip("Temporarily skipped: Re-enable once Succinct Integration is investigated.")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() {
		require.NoError(t, d.Down())
	}()

	aliceAddress := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)

	// Verify devnet is running
	require.NoError(t, d.RunSimpleL2Burn())

	// Verify Alice has L2 balance
	userBalance, err := d.L2Verif.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	require.True(t, userBalance.Cmp(big.NewInt(0)) > 0, "Alice should have positive L2 balance")

	withdrawalAmount := big.NewInt(1000000000000000000) // 1 ETH
	tx, receipt := initiateWithdrawalOnL2(d, ctx, t, aliceAddress, withdrawalAmount)

	// Deposit ETH on L1 bridge to fund withdrawals
	depositAmount := new(big.Int).Mul(withdrawalAmount, big.NewInt(2))
	depositOnL1Bridge(d, ctx, t, depositAmount)

	// Wait for dispute game publication
	waitForGameToBePublished(d, ctx, t, receipt)

	// Prove withdrawal transaction
	withdrawalHash, withdrawalTx := proveWithdrawalTransaction(d, ctx, t, tx)

	// Wait for game resolution
	waitForResolvedGame(d, ctx, t, withdrawalHash, aliceAddress)

	// Finalize withdrawal
	finalizeWithdrawal(d, ctx, t, aliceAddress, withdrawalTx, withdrawalAmount)
}
