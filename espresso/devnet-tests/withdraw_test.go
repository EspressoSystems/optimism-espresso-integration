package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	espressobindings "github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain"
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

func TestWithdrawal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up(NON_TEE))
	defer func() { require.NoError(t, d.Down()) }()

	alice := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)

	// Get contract addresses
	systemConfig, _, err := d.SystemConfig(ctx)
	require.NoError(t, err)
	factoryAddr, _ := systemConfig.DisputeGameFactory(&bind.CallOpts{})
	portalAddr, _ := systemConfig.OptimismPortal(&bind.CallOpts{})

	factory, _ := nodebindings.NewDisputeGameFactoryCaller(factoryAddr, d.L1)
	portal2, _ := nodepreview.NewOptimismPortal2(portalAddr, d.L1)
	portal2Caller, _ := nodepreview.NewOptimismPortal2Caller(portalAddr, d.L1)
	gameType, _ := portal2Caller.RespectedGameType(&bind.CallOpts{})
	l1ChainID, _ := d.L1.ChainID(ctx)

	// Step 1: Wait for proposer to start (just need 1 game)
	t.Log("Waiting for proposer to create first game...")
	require.Eventually(t, func() bool {
		count, _ := factory.GameCount(&bind.CallOpts{})
		if count != nil && count.Cmp(common.Big0) > 0 {
			t.Logf("Proposer started: %d games", count)
			return true
		}
		return false
	}, 3*time.Minute, 2*time.Second, "proposer didn't start")

	// Step 2: Initiate withdrawal
	t.Log("Initiating withdrawal on L2...")
	withdrawalAmount := big.NewInt(1e18) // 1 ETH

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

	// Deposit ETH to L1 bridge to fund withdrawals
	t.Log("Depositing ETH to L1 bridge...")
	rollupConfig, _ := d.RollupConfig(ctx)
	depositContract, _ := bindings.NewOptimismPortal(rollupConfig.DepositContractAddress, d.L1)
	l1Opts, _ := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	l1Opts.Value = new(big.Int).Mul(withdrawalAmount, big.NewInt(2))
	l1Opts.GasLimit = 500000
	depositTx, err := depositContract.DepositTransaction(l1Opts, common.Address{}, l1Opts.Value, 21000, false, nil)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, depositTx.Hash())
	require.NoError(t, err)
	t.Log("Deposit complete!")

	// Step 3: Wait for a dispute game covering the withdrawal block
	t.Logf("Waiting for dispute game covering L2 block %d...", withdrawReceipt.BlockNumber)
	var gameIndex *big.Int
	var gameProxy common.Address
	require.Eventually(t, func() bool {
		count, _ := factory.GameCount(&bind.CallOpts{})
		if count == nil || count.Cmp(common.Big0) == 0 {
			return false
		}
		games, _ := factory.FindLatestGames(&bind.CallOpts{}, gameType,
			new(big.Int).Sub(count, common.Big1), big.NewInt(1))
		if len(games) > 0 {
			latestL2Block := new(big.Int).SetBytes(games[0].ExtraData[0:32])
			t.Logf("Latest game: index=%d, L2Block=%d (need >= %d)",
				games[0].Index, latestL2Block, withdrawReceipt.BlockNumber)
			if latestL2Block.Cmp(withdrawReceipt.BlockNumber) >= 0 {
				gameIndex = games[0].Index
				info, _ := factory.GameAtIndex(&bind.CallOpts{}, games[0].Index)
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

	l1Opts, _ = bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	proveTx, err := disputeGame.Prove(l1Opts, []byte{}) // empty proof for mock verifier
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, proveTx.Hash())
	require.NoError(t, err)
	t.Log("Dispute game proven!")

	// Step 5: Prove withdrawal transaction
	t.Log("Proving withdrawal transaction...")
	games, _ := factory.FindLatestGames(&bind.CallOpts{}, gameType, gameIndex, big.NewInt(1))
	l2BlockNum := new(big.Int).SetBytes(games[0].ExtraData[0:32])
	l2Header, _ := d.L2Seq.HeaderByNumber(ctx, l2BlockNum)

	proofCl := gethclient.New(d.L2Seq.Client())
	params, err := withdrawals.ProveWithdrawalParametersForBlock(ctx, proofCl, d.L2Seq, withdrawTx.Hash(), l2Header, gameIndex)
	require.NoError(t, err)

	wd := crossdomain.NewWithdrawal(params.Nonce, &params.Sender, &params.Target, params.Value, params.GasLimit, params.Data)
	withdrawalTxStruct := nodepreview.TypesWithdrawalTransaction{
		Nonce: wd.Nonce, Sender: *wd.Sender, Target: *wd.Target,
		Value: wd.Value, GasLimit: wd.GasLimit, Data: wd.Data,
	}
	outputProof := nodepreview.TypesOutputRootProof{
		Version:                  params.OutputRootProof.Version,
		StateRoot:                params.OutputRootProof.StateRoot,
		MessagePasserStorageRoot: params.OutputRootProof.MessagePasserStorageRoot,
		LatestBlockhash:          params.OutputRootProof.LatestBlockhash,
	}

	l1Opts, _ = bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	l1Opts.GasLimit = 500000
	proveWithdrawTx, err := portal2.ProveWithdrawalTransaction(l1Opts, withdrawalTxStruct, gameIndex, outputProof, params.WithdrawalProof)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(ctx, d.L1, proveWithdrawTx.Hash())
	require.NoError(t, err)
	t.Log("Withdrawal proven!")

	// Step 6: Wait for proof maturity + game resolution
	t.Log("Waiting for proof maturity and game resolution...")
	maturityDelay, _ := portal2Caller.ProofMaturityDelaySeconds(&bind.CallOpts{})
	time.Sleep(time.Duration(maturityDelay.Int64()+5) * time.Second)

	require.Eventually(t, func() bool {
		status, _ := disputeGame.Status(&bind.CallOpts{})
		if status != 0 {
			t.Logf("Game resolved with status: %d", status)
			return true
		}
		// Try manual resolution (ignore errors, we'll retry)
		resolveOpts, _ := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
		_, _ = disputeGame.Resolve(resolveOpts)
		return false
	}, 5*time.Minute, 5*time.Second, "game not resolved")

	// Wait for dispute game finality delay (game must be resolved for this long before finalization)
	t.Log("Waiting for dispute game finality delay...")
	time.Sleep(10 * time.Second) // DISPUTE_GAME_FINALITY_DELAY_SECONDS is 6, add buffer

	// Step 7: Finalize withdrawal
	t.Log("Finalizing withdrawal...")
	balanceBefore, _ := d.L1.BalanceAt(ctx, alice, nil)

	l1Opts, _ = bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	l1Opts.GasLimit = 300000
	finalizeTx, err := portal2.FinalizeWithdrawalTransaction(l1Opts, withdrawalTxStruct)
	require.NoError(t, err)
	finalizeReceipt, err := wait.ForReceiptOK(ctx, d.L1, finalizeTx.Hash())
	require.NoError(t, err)

	// Verify balance increased
	balanceAfter, _ := d.L1.BalanceAt(ctx, alice, nil)
	fees := new(big.Int).Mul(new(big.Int).SetUint64(finalizeReceipt.GasUsed), finalizeReceipt.EffectiveGasPrice)
	expectedChange := new(big.Int).Sub(withdrawalAmount, fees)
	actualChange := new(big.Int).Sub(balanceAfter, balanceBefore)
	require.Equal(t, actualChange, expectedChange, "balance mismatch")

	t.Log("Withdrawal completed successfully!")
}
