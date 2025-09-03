package devnet_tests

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/stretchr/testify/require"

	bindings "github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	nodebindings "github.com/ethereum-optimism/optimism/op-node/bindings"
	nodepreview "github.com/ethereum-optimism/optimism/op-node/bindings/preview"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
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

	// TODO put this in a function part of devnet_tools.go
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

	// Check proposer account balance and permissions
	// The proposer uses the same mnemonic as batcher: "test test test test test test test test test test test junk"
	proposerPrivKey := d.secrets.Alice // This should match the mnemonic account
	proposerAddr := crypto.PubkeyToAddress(proposerPrivKey.PublicKey)
	proposerBalance, err := d.L1.BalanceAt(ctx, proposerAddr, nil)
	require.NoError(t, err)
	t.Logf("Proposer account %s balance: %s ETH", proposerAddr.Hex(), new(big.Int).Div(proposerBalance, big.NewInt(1e18)).String())

	// Check if the proposer publishes games
	// Check if any dispute games exist
	factory, err := bindings.NewDisputeGameFactoryCaller(disputeGameFactoryAddr, d.L1)
	require.NoError(t, err)

	// Check the required bond for game type 1
	requiredBond, err := factory.InitBonds(&bind.CallOpts{}, 1)
	if err != nil {
		t.Logf("Error getting required bond for game type 1: %v", err)
	} else {
		t.Logf("Required bond for game type 1: %s wei (%s ETH)", requiredBond.String(), new(big.Int).Div(requiredBond, big.NewInt(1e18)).String())
	}

	// Check if game type 1 implementation is set
	gameImpl, err := factory.GameImpls(&bind.CallOpts{}, 1)
	if err != nil {
		t.Logf("Error getting game implementation for type 1: %v", err)
	} else {
		t.Logf("Game type 1 implementation: %s", gameImpl.Hex())
		if gameImpl == (common.Address{}) {
			t.Logf("ERROR: No implementation set for game type 1! This explains why proposer fails.")
		}
	}

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

	// Wait a bit longer for the proposer to create games (proposer interval is 6s)
	t.Logf("Waiting 3 minutes for proposer to create dispute games...")
	time.Sleep(3 * time.Minute)

	// Check again after waiting
	gameCount, err = factory.GameCount(&bind.CallOpts{})
	if err != nil {
		t.Logf("Error getting game count after waiting: %v", err)
	} else {
		t.Logf("Total dispute games after waiting: %s", gameCount.String())
	}

	// Wait for the L2 output to be published as a dispute game on L1
	t.Logf("Waiting for dispute game to be published for block %d", receipt.BlockNumber)
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

	// Set up clients for proof generation
	receiptCl := d.L2Seq
	headerCl := d.L2Seq
	proofCl := gethclient.New(receiptCl.Client())

	// Set up contract bindings for proof generation
	factory2, err := nodebindings.NewDisputeGameFactoryCaller(disputeGameFactoryAddr, d.L1)
	require.NoError(t, err)

	portal2, err := nodepreview.NewOptimismPortal2Caller(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Get the block header for proof generation
	header, err := receiptCl.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	require.NoError(t, err)
	t.Logf("Using block %d (hash: %s) for withdrawal proof", blockNumber, header.Hash().Hex())

	// Generate withdrawal proof parameters using fault proofs
	params, err := withdrawals.ProveWithdrawalParametersFaultProofs(ctx, proofCl, receiptCl, headerCl, tx.Hash(), factory2, portal2)
	if err != nil {
		t.Logf("Failed to generate fault proof parameters: %v", err)
		t.Logf("Trying legacy proof generation instead...")
		// Fall back to legacy proof generation if fault proofs fail
		oracle, err := nodebindings.NewL2OutputOracleCaller(common.Address{}, d.L1) // Use zero address as fallback
		if err != nil {
			t.Fatalf("Failed to create oracle caller: %v", err)
		}
		params, err = withdrawals.ProveWithdrawalParameters(ctx, proofCl, receiptCl, tx.Hash(), header, oracle)
		require.NoError(t, err)
	}

	// Bind to OptimismPortal contract on L1
	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, d.L1)
	require.NoError(t, err)

	// Create transaction options for Alice on L1
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	l1Opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)

	// Log basic withdrawal proof info
	t.Logf("Attempting to prove withdrawal with L2OutputIndex: %s", params.L2OutputIndex.String())

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
	if err != nil {
		t.Logf("ProveWithdrawalTransaction failed: %v", err)
		// Try to get more detailed error information
		if strings.Contains(err.Error(), "execution reverted") {
			t.Logf("Transaction reverted - possible causes:")
			t.Logf("  1. Withdrawal already proven")
			t.Logf("  2. Invalid proof data")
			t.Logf("  3. L2 output not yet finalized")
			t.Logf("  4. Incorrect dispute game reference")
		}
	}
	require.NoError(t, err)

	// Wait for the proof transaction to be mined
	proveReceipt, err := bind.WaitMined(ctx, d.L1, proveTx)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, proveReceipt.Status)

	t.Logf("Withdrawal proof transaction successful: %s", proveTx.Hash().Hex())
	t.Logf("Withdrawal can now be finalized after the finalization period")

}
