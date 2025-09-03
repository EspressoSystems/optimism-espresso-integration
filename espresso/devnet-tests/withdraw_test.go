package devnet_tests

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	nodebindings "github.com/ethereum-optimism/optimism/op-node/bindings"
	nodepreview "github.com/ethereum-optimism/optimism/op-node/bindings/preview"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/stretchr/testify/require"
)

func TestWithdraw(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := NewDevnet(ctx, t)
	require.NoError(t, d.Up())
	defer func() {
		require.NoError(t, d.Down())
	}()

	// Send a transaction just to check that everything has started up ok.
	require.NoError(t, d.RunSimpleL2Burn())

	// Check Alice's balance on L2 verifier before withdrawal
	aliceAddress := crypto.PubkeyToAddress(d.secrets.Alice.PublicKey)
	aliceBalance, err := d.L2Verif.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	require.True(t, aliceBalance.Cmp(big.NewInt(0)) > 0, "Alice should have a positive balance")

	// Initiate withdrawal on L2
	withdrawalAmount := big.NewInt(1000000) // Withdraw 1 000 000 wei

	// Bind to L2ToL1MessagePasser contract
	l2MessagePasser, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, d.L2Seq)
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

	// Read actual deployed contract addresses from deployment state
	// This matches what the op-proposer service does in docker-compose.yml
	deploymentStateFile := "../deployment/deployer/state.json"

	// Check if deployment state file exists
	if _, err := os.Stat(deploymentStateFile); os.IsNotExist(err) {
		t.Skipf("Deployment state file not found: %s. Make sure devnet is properly deployed.", deploymentStateFile)
		return
	}

	// Read and parse the deployment state
	stateData, err := os.ReadFile(deploymentStateFile)
	if err != nil {
		t.Fatalf("Failed to read deployment state: %v", err)
	}

	var deploymentState struct {
		OpChainDeployments []struct {
			DisputeGameFactoryProxyAddress string `json:"disputeGameFactoryProxyAddress"`
			OptimismPortalProxyAddress     string `json:"optimismPortalProxyAddress"`
		} `json:"opChainDeployments"`
	}

	if err := json.Unmarshal(stateData, &deploymentState); err != nil {
		t.Fatalf("Failed to parse deployment state: %v", err)
	}

	if len(deploymentState.OpChainDeployments) == 0 {
		t.Fatal("No OP chain deployments found in state file")
	}

	disputeGameFactoryAddr := common.HexToAddress(deploymentState.OpChainDeployments[0].DisputeGameFactoryProxyAddress)
	optimismPortalAddr := common.HexToAddress(deploymentState.OpChainDeployments[0].OptimismPortalProxyAddress)

	t.Logf("Using actual deployed contracts from devnet")
	t.Logf("OptimismPortalProxy: %s", optimismPortalAddr.Hex())
	t.Logf("DisputeGameFactoryProxy: %s", disputeGameFactoryAddr.Hex())

	// Check if the proposer publishes games
	// Check if any dispute games exist
	factory, err := bindings.NewDisputeGameFactoryCaller(disputeGameFactoryAddr, d.L1)
	require.NoError(t, err)

	// Get the total number of games created
	gameCount, err := factory.GameCount(&bind.CallOpts{})
	if err != nil {
		t.Logf("Error getting game count: %v", err)
		return
	}

	t.Logf("Total dispute games created: %s", gameCount.String())

	// Check recent games
	if gameCount.Uint64() > 0 {
		// Get the latest game
		latestGameIndex := new(big.Int).Sub(gameCount, big.NewInt(1))
		gameInfo, err := factory.GameAtIndex(&bind.CallOpts{}, latestGameIndex)
		if err == nil {
			t.Logf("Latest game: proxy=%s, type=%d, timestamp=%d",
				gameInfo.Proxy.Hex(), gameInfo.GameType, gameInfo.Timestamp)
		}
	}

	// Wait for the L2 output to be published as a dispute game on L1
	t.Logf("Waiting for dispute game to be published for block %d", receipt.BlockNumber)
	var blockNumber uint64

	// For now, assume we're using fault proofs (dispute games)
	t.Logf("Waiting for dispute game to be published...")
	blockNumber, err = wait.ForGamePublished(ctx, d.L1, optimismPortalAddr, disputeGameFactoryAddr, receipt.BlockNumber)
	require.NoError(t, err)
	t.Logf("Dispute game published for block %d", blockNumber)

	// Wait for another block to be mined to ensure timestamp increases
	err = wait.ForNextBlock(ctx, d.L1)
	require.NoError(t, err)

	// Set up clients for proof generation
	receiptCl := d.L2Seq
	headerCl := d.L2Seq
	proofCl := gethclient.New(receiptCl.Client())

	// Set up contract bindings for proof generation
	factory2, err := nodebindings.NewDisputeGameFactoryCaller(disputeGameFactoryAddr, d.L1)
	require.NoError(t, err)

	portal2, err := nodepreview.NewOptimismPortal2Caller(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Generate withdrawal proof parameters using fault proofs
	params, err := withdrawals.ProveWithdrawalParametersFaultProofs(ctx, proofCl, receiptCl, headerCl, tx.Hash(), factory2, portal2)
	require.NoError(t, err)

	// Bind to OptimismPortal contract on L1
	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Create transaction options for Alice on L1
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	l1Opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)

	// Submit the withdrawal proof transaction
	proveTx, err := portal.ProveWithdrawalTransaction(
		l1Opts,
		bindings.TypesWithdrawalTransaction{
			Nonce:    params.Nonce,
			Sender:   params.Sender,
			Target:   params.Target,
			Value:    params.Value,
			GasLimit: params.GasLimit,
			Data:     params.Data,
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
	t.Logf("Withdrawal can now be finalized after the finalization period")

}
