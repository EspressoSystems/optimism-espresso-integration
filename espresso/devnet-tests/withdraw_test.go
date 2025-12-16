package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	espressobindings "github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	nodebindings "github.com/ethereum-optimism/optimism/op-node/bindings"
	nodepreview "github.com/ethereum-optimism/optimism/op-node/bindings/preview"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/stretchr/testify/require"
)

// l2BlockFromExtraData extracts the L2 block number from a dispute game's ExtraData.
// The first 32 bytes of ExtraData contain the L2 block number as a big-endian uint256.
func l2BlockFromExtraData(extraData []byte) *big.Int {
	return new(big.Int).SetBytes(extraData[:32])
}

func TestWithdrawal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() { require.NoError(t, d.Down()) }()

	alice := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)
	callOpts := &bind.CallOpts{}
	l1ChainID, _ := d.L1.ChainID(ctx)
	l1Transactor := func() *bind.TransactOpts {
		opts, _ := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
		return opts
	}

	// Get contract addresses
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)
	factoryAddr, _ := systemConfig.DisputeGameFactory(callOpts)
	portalAddr, _ := systemConfig.OptimismPortal(callOpts)

	factory, _ := nodebindings.NewDisputeGameFactoryCaller(factoryAddr, d.L1)
	portal2, _ := nodepreview.NewOptimismPortal2(portalAddr, d.L1)
	portal2Caller, _ := nodepreview.NewOptimismPortal2Caller(portalAddr, d.L1)
	gameType, _ := portal2Caller.RespectedGameType(callOpts)

	// Step 1: Wait for proposer to start (just need 1 game)
	t.Log("Waiting for proposer to create first game...")
	require.Eventually(t, func() bool {
		count, _ := factory.GameCount(callOpts)
		if count != nil && count.Cmp(common.Big0) > 0 {
			t.Logf("Proposer started: %d games", count)
			return true
		}
		return false
	}, 3*time.Minute, 2*time.Second, "proposer didn't start")

	// Step 2: Initiate withdrawal

	t.Log("Initiating withdrawal on L2...")
	withdrawalAmount := big.NewInt(1e18) // 1 ETH

	// Deposit ETH to L1 bridge to fund withdrawals
	t.Log("Depositing ETH to L1 bridge...")
	rollupConfig, _ := d.RollupConfig(ctx)
	depositContract, _ := bindings.NewOptimismPortal(rollupConfig.DepositContractAddress, d.L1)
	depositOpts := l1Transactor()
	depositOpts.Value = new(big.Int).Mul(withdrawalAmount, big.NewInt(2))
	depositOpts.GasLimit = 500000
	depositTx, err := depositContract.DepositTransaction(depositOpts, common.Address{}, depositOpts.Value, 21000, false, nil)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, depositTx.Hash())
	require.NoError(t, err)
	t.Log("Deposit complete!")

	l2MessagePasser, _ := bindings.NewL2ToL1MessagePasser(
		common.HexToAddress("0x4200000000000000000000000000000000000016"), d.L2Seq)
	l2ChainID, _ := d.L2Seq.ChainID(ctx)
	l2Opts, _ := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l2ChainID)
	l2Opts.Value = withdrawalAmount

	withdrawTx, err := l2MessagePasser.InitiateWithdrawal(l2Opts, alice, big.NewInt(21000), nil)
	require.NoError(t, err)
	withdrawReceipt, err := wait.ForReceiptOK(ctx, d.L2Verif, withdrawTx.Hash())
	require.NoError(t, err)
	t.Logf("Withdrawal initiated at L2 block %d", withdrawReceipt.BlockNumber)

	// Step 3: Wait for a dispute game covering the withdrawal block
	t.Logf("Waiting for dispute game covering L2 block %d...", withdrawReceipt.BlockNumber)
	var gameIndex *big.Int
	var gameProxy common.Address
	var gameL2Block *big.Int
	require.Eventually(t, func() bool {
		count, _ := factory.GameCount(callOpts)
		if count == nil || count.Cmp(common.Big0) == 0 {
			return false
		}
		games, _ := factory.FindLatestGames(callOpts, gameType, new(big.Int).Sub(count, common.Big1), big.NewInt(1))
		if len(games) > 0 {
			gameL2Block = l2BlockFromExtraData(games[0].ExtraData)
			t.Logf("Latest game: index=%d, L2Block=%d (need >= %d)", games[0].Index, gameL2Block, withdrawReceipt.BlockNumber)
			if gameL2Block.Cmp(withdrawReceipt.BlockNumber) >= 0 {
				gameIndex = games[0].Index
				info, _ := factory.GameAtIndex(callOpts, games[0].Index)
				gameProxy = info.Proxy
				return true
			}
		}
		return false
	}, 10*time.Minute, 5*time.Second, "no game covering withdrawal block")

	// Step 4: Prove the dispute game (mock verifier accepts empty proof)
	t.Log("Proving dispute game...")
	disputeGame, err := espressobindings.NewOPSuccinctFaultDisputeGame(gameProxy, d.L1)
	require.NoError(t, err)

	proveTx, err := disputeGame.Prove(l1Transactor(), []byte{})
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, proveTx.Hash())
	require.NoError(t, err)
	t.Log("Dispute game proven!")

	// Step 5: Prove withdrawal transaction
	t.Log("Proving withdrawal transaction...")
	l2Header, _ := d.L2Seq.HeaderByNumber(ctx, gameL2Block)
	proofCl := gethclient.New(d.L2Seq.Client())
	params, err := withdrawals.ProveWithdrawalParametersForBlock(ctx, proofCl, d.L2Seq, withdrawTx.Hash(), l2Header, gameIndex)
	require.NoError(t, err)

	withdrawalTxStruct := nodepreview.TypesWithdrawalTransaction{
		Nonce: params.Nonce, Sender: params.Sender, Target: params.Target,
		Value: params.Value, GasLimit: params.GasLimit, Data: params.Data,
	}
	outputProof := nodepreview.TypesOutputRootProof{
		Version:                  params.OutputRootProof.Version,
		StateRoot:                params.OutputRootProof.StateRoot,
		MessagePasserStorageRoot: params.OutputRootProof.MessagePasserStorageRoot,
		LatestBlockhash:          params.OutputRootProof.LatestBlockhash,
	}

	proveWithdrawOpts := l1Transactor()
	proveWithdrawOpts.GasLimit = 500000
	proveWithdrawTx, err := portal2.ProveWithdrawalTransaction(proveWithdrawOpts, withdrawalTxStruct, gameIndex, outputProof, params.WithdrawalProof)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, proveWithdrawTx.Hash())
	require.NoError(t, err)
	t.Log("Withdrawal proven!")

	// Step 6: Wait for proof maturity + game resolution
	t.Log("Waiting for proof maturity and game resolution...")
	maturityDelay, _ := portal2Caller.ProofMaturityDelaySeconds(callOpts)
	time.Sleep(time.Duration(maturityDelay.Int64()+5) * time.Second)

	require.Eventually(t, func() bool {
		status, _ := disputeGame.Status(callOpts)
		if status != 0 {
			t.Logf("Game resolved with status: %d", status)
			return true
		}
		_, _ = disputeGame.Resolve(l1Transactor()) // try manual resolution
		return false
	}, 5*time.Minute, 5*time.Second, "game not resolved")

	t.Log("Waiting for dispute game finality delay...")
	time.Sleep(10 * time.Second) // DISPUTE_GAME_FINALITY_DELAY_SECONDS is 6 + 4seconds to be safe

	// Step 7: Finalize withdrawal
	t.Log("Finalizing withdrawal...")
	balanceBefore, _ := d.L1.BalanceAt(ctx, alice, nil)

	finalizeOpts := l1Transactor()
	finalizeOpts.GasLimit = 300000
	finalizeTx, err := portal2.FinalizeWithdrawalTransaction(finalizeOpts, withdrawalTxStruct)
	require.NoError(t, err)
	finalizeReceipt, err := wait.ForReceiptOK(ctx, d.L1, finalizeTx.Hash())
	require.NoError(t, err)

	// Verify balance increased (withdrawal amount minus gas fees)
	balanceAfter, _ := d.L1.BalanceAt(ctx, alice, nil)
	gasCost := new(big.Int).Mul(big.NewInt(int64(finalizeReceipt.GasUsed)), finalizeReceipt.EffectiveGasPrice)
	expectedBalance := new(big.Int).Add(balanceBefore, new(big.Int).Sub(withdrawalAmount, gasCost))
	require.Equal(t, expectedBalance, balanceAfter, "balance mismatch")

	t.Log("Withdrawal completed successfully!")
}
