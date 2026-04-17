package devnet_tests

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	espressobindings "github.com/ethereum-optimism/optimism/espresso/bindings"
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
func l2BlockFromExtraData(extraData []byte) (*big.Int, error) {
	if len(extraData) < 32 {
		return nil, fmt.Errorf("extraData too short: got %d bytes, need at least 32", len(extraData))
	}
	return new(big.Int).SetBytes(extraData[:32]), nil
}

func TestWithdrawal(t *testing.T) {
	profile := ProfileFromEnv(t)
	t.Run(string(profile), func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		d := NewDevnet(ctx, t, profile)
		require.NoError(t, d.Up())
		defer func() { require.NoError(t, d.Down()) }()

		require.NoError(t, d.WaitForBatcher(ctx))

		alice := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)
		callOpts := &bind.CallOpts{Context: ctx}

		l1ChainID, err := d.L1.ChainID(ctx)
		require.NoError(t, err, "failed to get L1 chain ID")

		l1Transactor := func() *bind.TransactOpts {
			opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
			require.NoError(t, err, "failed to create L1 transactor")
			return opts
		}

		// Get contract addresses
		systemConfig, _, err := d.SystemConfig(ctx)
		require.NoError(t, err)

		factoryAddr, err := systemConfig.DisputeGameFactory(callOpts)
		require.NoError(t, err, "failed to get DisputeGameFactory address")

		portalAddr, err := systemConfig.OptimismPortal(callOpts)
		require.NoError(t, err, "failed to get OptimismPortal address")

		factory, err := nodebindings.NewDisputeGameFactoryCaller(factoryAddr, d.L1)
		require.NoError(t, err, "failed to bind DisputeGameFactory")

		portal2, err := nodepreview.NewOptimismPortal2(portalAddr, d.L1)
		require.NoError(t, err, "failed to bind OptimismPortal2")

		portal2Caller, err := nodepreview.NewOptimismPortal2Caller(portalAddr, d.L1)
		require.NoError(t, err, "failed to bind OptimismPortal2Caller")

		gameType, err := portal2Caller.RespectedGameType(callOpts)
		require.NoError(t, err, "failed to get respected game type")

		// Step 1: Wait for proposer to start (just need 1 game)
		t.Log("Waiting for proposer to create first game...")
		require.Eventually(t, func() bool {
			count, err := factory.GameCount(callOpts)
			if err != nil {
				t.Logf("Error getting game count: %v", err)
				return false
			}
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
		rollupConfig, err := d.RollupConfig(ctx)
		require.NoError(t, err, "failed to get rollup config")

		depositContract, err := bindings.NewOptimismPortal(rollupConfig.DepositContractAddress, d.L1)
		require.NoError(t, err, "failed to bind deposit contract")

		depositOpts := l1Transactor()
		depositOpts.Value = new(big.Int).Mul(withdrawalAmount, big.NewInt(2))
		depositOpts.GasLimit = 500000
		depositTx, err := depositContract.DepositTransaction(depositOpts, common.Address{}, depositOpts.Value, 21000, false, nil)
		require.NoError(t, err)
		_, err = wait.ForReceiptOK(ctx, d.L1, depositTx.Hash())
		require.NoError(t, err)
		t.Log("Deposit complete!")

		l2MessagePasser, err := bindings.NewL2ToL1MessagePasser(
			common.HexToAddress("0x4200000000000000000000000000000000000016"), d.L2Seq)
		require.NoError(t, err, "failed to bind L2ToL1MessagePasser")

		l2ChainID, err := d.L2Seq.ChainID(ctx)
		require.NoError(t, err, "failed to get L2 chain ID")

		l2Opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l2ChainID)
		require.NoError(t, err, "failed to create L2 transactor")
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
			count, err := factory.GameCount(callOpts)
			if err != nil || count == nil || count.Cmp(common.Big0) == 0 {
				return false
			}
			games, err := factory.FindLatestGames(callOpts, gameType, new(big.Int).Sub(count, common.Big1), big.NewInt(1))
			if err != nil || len(games) == 0 {
				return false
			}
			var parseErr error
			gameL2Block, parseErr = l2BlockFromExtraData(games[0].ExtraData)
			if parseErr != nil {
				t.Logf("Error parsing extraData: %v", parseErr)
				return false
			}
			t.Logf("Latest game: index=%d, L2Block=%d (need >= %d)", games[0].Index, gameL2Block, withdrawReceipt.BlockNumber)
			if gameL2Block.Cmp(withdrawReceipt.BlockNumber) >= 0 {
				gameIndex = games[0].Index
				info, err := factory.GameAtIndex(callOpts, games[0].Index)
				if err != nil {
					t.Logf("Error getting game info: %v", err)
					return false
				}
				gameProxy = info.Proxy
				return true
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
		l2Header, err := d.L2Seq.HeaderByNumber(ctx, gameL2Block)
		require.NoError(t, err, "failed to get L2 header")

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
		proveReceipt, err := wait.ForReceiptOK(ctx, d.L1, proveWithdrawTx.Hash())
		require.NoError(t, err)
		t.Log("Withdrawal proven!")

		// Step 6: Wait for proof maturity + game resolution
		t.Log("Waiting for proof maturity and game resolution...")

		maturityDelay, err := portal2Caller.ProofMaturityDelaySeconds(callOpts)
		require.NoError(t, err, "failed to get proof maturity delay")

		disputeGameFinalityDelay, err := portal2Caller.DisputeGameFinalityDelaySeconds(callOpts)
		require.NoError(t, err, "failed to get dispute game finality delay")

		// Wait for game to be resolved
		require.Eventually(t, func() bool {
			status, err := disputeGame.Status(callOpts)
			if err != nil {
				t.Logf("Error getting game status: %v", err)
				return false
			}
			if status != 0 {
				t.Logf("Game resolved with status: %d", status)
				return true
			}
			// Try manual resolution
			resolveTx, err := disputeGame.Resolve(l1Transactor())
			if err != nil {
				t.Logf("Error calling Resolve (may be expected): %v", err)
				return false
			}
			_, err = wait.ForReceiptOK(ctx, d.L1, resolveTx.Hash())
			if err != nil {
				t.Logf("Resolve tx failed: %v", err)
				return false
			}
			return false // Recheck status on next iteration
		}, 5*time.Minute, 5*time.Second, "game not resolved")

		// Wait for proof maturity + finality delays by polling L1 block time
		t.Log("Waiting for proof maturity and dispute game finality delays...")
		proveBlock, err := d.L1.HeaderByNumber(ctx, proveReceipt.BlockNumber)
		require.NoError(t, err, "failed to get prove block header")
		targetTime := proveBlock.Time + maturityDelay.Uint64() + disputeGameFinalityDelay.Uint64() + 10

		require.Eventually(t, func() bool {
			header, err := d.L1.HeaderByNumber(ctx, nil)
			if err != nil {
				return false
			}
			if header.Time >= targetTime {
				return true
			}
			t.Logf("Waiting for delays: %ds remaining", targetTime-header.Time)
			return false
		}, 3*time.Minute, 2*time.Second, "timeout waiting for delays")

		// Step 7: Finalize withdrawal
		t.Log("Finalizing withdrawal...")
		balanceBefore, err := d.L1.BalanceAt(ctx, alice, nil)
		require.NoError(t, err, "failed to get balance before finalize")

		finalizeOpts := l1Transactor()
		finalizeOpts.GasLimit = 300000
		finalizeTx, err := portal2.FinalizeWithdrawalTransaction(finalizeOpts, withdrawalTxStruct)
		require.NoError(t, err)
		finalizeReceipt, err := wait.ForReceiptOK(ctx, d.L1, finalizeTx.Hash())
		require.NoError(t, err)

		// Verify balance increased (withdrawal amount minus gas fees)
		balanceAfter, err := d.L1.BalanceAt(ctx, alice, nil)
		require.NoError(t, err, "failed to get balance after finalize")

		gasCost := new(big.Int).Mul(big.NewInt(int64(finalizeReceipt.GasUsed)), finalizeReceipt.EffectiveGasPrice)
		balanceChange := new(big.Int).Sub(balanceAfter, balanceBefore)
		expectedChange := new(big.Int).Sub(withdrawalAmount, gasCost)

		// Use GreaterOrEqual to account for any minor discrepancies (e.g., rounding)
		// The balance should increase by at least (withdrawalAmount - gasCost)
		require.True(t, balanceChange.Cmp(expectedChange) >= 0,
			"balance didn't increase as expected: got change %s, expected at least %s", balanceChange, expectedChange)

		t.Log("Withdrawal completed successfully!")
	})
}
