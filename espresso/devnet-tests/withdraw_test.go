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

	// Try legacy withdrawal proof first (more compatible)
	t.Logf("Attempting legacy withdrawal proof generation...")
	
	// For legacy proofs, we need an L2OutputOracle instead of dispute game factory
	// Check if we can use legacy method
	params, err := withdrawals.ProveWithdrawalParameters(ctx, proofCl, receiptCl, tx.Hash(), header, nil)
	if err != nil {
		t.Logf("Legacy withdrawal proof failed: %v", err)
		t.Logf("Falling back to fault proofs...")
		
		// Fallback to fault proofs
		params, err = withdrawals.ProveWithdrawalParametersFaultProofs(ctx, proofCl, receiptCl, headerCl, tx.Hash(), factory2, portal2)
		require.NoError(t, err)
		t.Logf("Using fault proofs for withdrawal")
	} else {
		t.Logf("Using legacy withdrawal proof")
	}

	// Bind to OptimismPortal contract on L1
	portal, err := bindings.NewOptimismPortal(optimismPortalAddr, d.L1)
	require.NoError(t, err)
	
	// Validate contract deployment
	t.Logf("=== Contract Validation ===")
	contractCode, err := d.L1.CodeAt(ctx, optimismPortalAddr, nil)
	if err != nil {
		t.Fatalf("Failed to get contract code: %v", err)
	}
	if len(contractCode) == 0 {
		t.Fatalf("No contract deployed at OptimismPortal address %s", optimismPortalAddr.Hex())
	}
	t.Logf("OptimismPortal contract verified at %s (code length: %d)", optimismPortalAddr.Hex(), len(contractCode))
	
	// Try to get contract version to verify it's the right contract
	version, err := portal.Version(&bind.CallOpts{})
	if err != nil {
		t.Logf("Warning: Could not get contract version: %v", err)
	} else {
		t.Logf("OptimismPortal version: %s", version)
	}
	t.Logf("=== End Contract Validation ===")

	// Create transaction options for Alice on L1
	l1ChainID, err := d.L1.ChainID(ctx)
	require.NoError(t, err)

	l1Opts, err := bind.NewKeyedTransactorWithChainID(d.secrets.Alice, l1ChainID)
	require.NoError(t, err)
	
	// Set proper gas configuration
	l1Opts.GasLimit = 500000 // Set a reasonable gas limit
	gasPrice, err := d.L1.SuggestGasPrice(ctx)
	if err != nil {
		t.Logf("Warning: Could not get suggested gas price: %v", err)
		l1Opts.GasPrice = big.NewInt(20000000000) // 20 gwei fallback
	} else {
		l1Opts.GasPrice = gasPrice
	}
	t.Logf("Set gas limit: %d, gas price: %s", l1Opts.GasLimit, l1Opts.GasPrice.String())

	// Log basic withdrawal proof info
	t.Logf("Attempting to prove withdrawal with L2OutputIndex: %s", params.L2OutputIndex.String())

	// Pre-flight validation checks
	t.Logf("=== Pre-flight Validation ===")

	// 1. Check if withdrawal is already proven
	withdrawalHash := crypto.Keccak256Hash(
		params.Nonce.Bytes(),
		params.Sender.Bytes(),
		params.Target.Bytes(),
		params.Value.Bytes(),
		params.GasLimit.Bytes(),
		params.Data,
	)
	t.Logf("Withdrawal hash: %s", withdrawalHash.Hex())

	// Check if this is a fault proofs vs legacy contract issue
	t.Logf("Checking contract compatibility...")
	
	provenWithdrawal, err := portal.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash)
	if err != nil {
		t.Logf("Warning: Could not check proven withdrawals: %v", err)
		t.Logf("This may indicate a contract version mismatch or fault proofs not enabled")
		
		// Try to check if this is a legacy vs fault proofs issue
		// In fault proofs, the contract interface may be different
		if strings.Contains(err.Error(), "execution reverted") {
			t.Logf("CRITICAL: Contract calls are reverting - possible causes:")
			t.Logf("  1. Wrong contract address")
			t.Logf("  2. Contract not properly initialized")
			t.Logf("  3. Fault proofs not enabled/configured")
			t.Logf("  4. Contract ABI mismatch")
		}
	} else {
		t.Logf("Proven withdrawal status - OutputRoot: %s, Timestamp: %s, L2OutputIndex: %s",
			common.Bytes2Hex(provenWithdrawal.OutputRoot[:]),
			provenWithdrawal.Timestamp.String(),
			provenWithdrawal.L2OutputIndex.String())

		if provenWithdrawal.Timestamp.Cmp(big.NewInt(0)) > 0 {
			t.Logf("WARNING: Withdrawal may already be proven (timestamp > 0)")
		}
	}

	// 2. Check if output is finalized
	isFinalized, err := portal.IsOutputFinalized(&bind.CallOpts{}, params.L2OutputIndex)
	if err != nil {
		t.Logf("Warning: Could not check if output is finalized: %v", err)
	} else {
		t.Logf("L2 Output %s finalized: %t", params.L2OutputIndex.String(), isFinalized)
		if isFinalized {
			t.Logf("WARNING: Output is already finalized - this may cause issues")
		}
	}

	// 3. Validate proof parameters
	t.Logf("Withdrawal transaction details:")
	t.Logf("  Nonce: %s", params.Nonce.String())
	t.Logf("  Sender: %s", params.Sender.Hex())
	t.Logf("  Target: %s", params.Target.Hex())
	t.Logf("  Value: %s", params.Value.String())
	t.Logf("  GasLimit: %s", params.GasLimit.String())
	t.Logf("  Data length: %d", len(params.Data))
	t.Logf("  WithdrawalProof length: %d", len(params.WithdrawalProof))

	// 4. Check account balance and gas estimation
	balance, err := d.L1.BalanceAt(ctx, crypto.PubkeyToAddress(d.secrets.Alice.PublicKey), nil)
	if err != nil {
		t.Logf("Warning: Could not check account balance: %v", err)
	} else {
		t.Logf("Account balance: %s wei", balance.String())
	}

	// 5. Try to estimate gas for the transaction
	l1Opts.NoSend = true // Don't actually send, just estimate
	_, err = portal.ProveWithdrawalTransaction(
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
	l1Opts.NoSend = false // Reset for actual transaction

	if err != nil {
		t.Logf("Gas estimation failed: %v", err)
		t.Logf("This indicates the transaction will likely fail")
	} else {
		t.Logf("Gas estimation successful")
	}

	t.Logf("=== End Pre-flight Validation ===")

	// Contract state inspection
	t.Logf("=== Contract State Inspection ===")

	// Check portal contract state
	paused, err := portal.Paused(&bind.CallOpts{})
	if err != nil {
		t.Logf("Warning: Could not check if portal is paused: %v", err)
	} else {
		t.Logf("OptimismPortal paused: %t", paused)
		if paused {
			t.Logf("ERROR: Portal is paused - transactions will fail")
		}
	}

	// Check L2 Oracle address
	l2Oracle, err := portal.L2Oracle(&bind.CallOpts{})
	if err != nil {
		t.Logf("Warning: Could not get L2Oracle address: %v", err)
	} else {
		t.Logf("L2Oracle address: %s", l2Oracle.Hex())
	}

	// Check guardian address
	guardian, err := portal.Guardian(&bind.CallOpts{})
	if err != nil {
		t.Logf("Warning: Could not get guardian address: %v", err)
	} else {
		t.Logf("Guardian address: %s", guardian.Hex())
	}

	t.Logf("=== End Contract State Inspection ===")

	// Submit the withdrawal proof transaction
	t.Logf("=== Submitting ProveWithdrawalTransaction ===")
	t.Logf("Transaction will be sent from: %s", crypto.PubkeyToAddress(d.secrets.Alice.PublicKey).Hex())
	t.Logf("Gas limit: %d", l1Opts.GasLimit)
	t.Logf("Gas price: %s", l1Opts.GasPrice.String())

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

		// Enhanced error analysis
		errorStr := err.Error()
		t.Logf("Full error string: %s", errorStr)

		if strings.Contains(errorStr, "execution reverted") {
			t.Logf("Transaction reverted - analyzing revert reason...")

			// Common revert reasons for proveWithdrawalTransaction:
			if strings.Contains(errorStr, "OptimismPortal: withdrawal hash has already been proven") {
				t.Logf("ERROR: Withdrawal already proven")
			} else if strings.Contains(errorStr, "OptimismPortal: invalid output root proof") {
				t.Logf("ERROR: Invalid output root proof")
			} else if strings.Contains(errorStr, "OptimismPortal: output root proof is not valid") {
				t.Logf("ERROR: Output root proof validation failed")
			} else if strings.Contains(errorStr, "OptimismPortal: cannot prove a withdrawal with a finalized output") {
				t.Logf("ERROR: Output already finalized")
			} else {
				t.Logf("Generic revert - possible causes:")
				t.Logf("  1. Withdrawal already proven")
				t.Logf("  2. Invalid proof data")
				t.Logf("  3. L2 output not yet finalized")
				t.Logf("  4. Incorrect dispute game reference")
				t.Logf("  5. Invalid withdrawal proof")
			}
		} else if strings.Contains(errorStr, "insufficient funds") {
			t.Logf("ERROR: Insufficient gas funds")
		} else if strings.Contains(errorStr, "nonce too low") {
			t.Logf("ERROR: Nonce issue")
		}

		// Run comprehensive debugging
		debugWithdrawalTransaction(t, ctx, portal, params, tx.Hash())
	}
	require.NoError(t, err)

	// Wait for the proof transaction to be mined
	proveReceipt, err := bind.WaitMined(ctx, d.L1, proveTx)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, proveReceipt.Status)

	t.Logf("Withdrawal proof transaction successful: %s", proveTx.Hash().Hex())
	t.Logf("Withdrawal can now be finalized after the finalization period")

}

// debugWithdrawalTransaction provides comprehensive debugging for withdrawal proof failures
func debugWithdrawalTransaction(t *testing.T, ctx context.Context, portal *bindings.OptimismPortal, params withdrawals.ProvenWithdrawalParameters, txHash common.Hash) {
	t.Logf("=== Comprehensive Withdrawal Debug ===")

	// 1. Validate withdrawal hash calculation
	withdrawalHash := crypto.Keccak256Hash(
		params.Nonce.Bytes(),
		params.Sender.Bytes(),
		params.Target.Bytes(),
		params.Value.Bytes(),
		params.GasLimit.Bytes(),
		params.Data,
	)
	t.Logf("Calculated withdrawal hash: %s", withdrawalHash.Hex())

	// 2. Check proven withdrawals mapping
	provenWithdrawal, err := portal.ProvenWithdrawals(&bind.CallOpts{}, withdrawalHash)
	if err != nil {
		t.Logf("ERROR: Could not query proven withdrawals: %v", err)
	} else {
		t.Logf("Proven withdrawal details:")
		t.Logf("  OutputRoot: %s", common.Bytes2Hex(provenWithdrawal.OutputRoot[:]))
		t.Logf("  Timestamp: %s", provenWithdrawal.Timestamp.String())
		t.Logf("  L2OutputIndex: %s", provenWithdrawal.L2OutputIndex.String())

		if provenWithdrawal.Timestamp.Cmp(big.NewInt(0)) > 0 {
			t.Logf("ISSUE: Withdrawal already proven at timestamp %s", provenWithdrawal.Timestamp.String())
		}
	}

	// 3. Validate output root proof components
	t.Logf("Output root proof validation:")
	t.Logf("  Version: %s", common.Bytes2Hex(params.OutputRootProof.Version[:]))
	t.Logf("  StateRoot: %s", common.Bytes2Hex(params.OutputRootProof.StateRoot[:]))
	t.Logf("  MessagePasserStorageRoot: %s", common.Bytes2Hex(params.OutputRootProof.MessagePasserStorageRoot[:]))
	t.Logf("  LatestBlockhash: %s", common.Bytes2Hex(params.OutputRootProof.LatestBlockhash[:]))

	// 4. Check withdrawal proof structure
	t.Logf("Withdrawal proof structure:")
	t.Logf("  Proof elements: %d", len(params.WithdrawalProof))
	for i, proof := range params.WithdrawalProof {
		t.Logf("  Proof[%d]: %s (length: %d)", i, common.Bytes2Hex(proof), len(proof))
	}

	// 5. Check L2 output finalization status
	isFinalized, err := portal.IsOutputFinalized(&bind.CallOpts{}, params.L2OutputIndex)
	if err != nil {
		t.Logf("ERROR: Could not check output finalization: %v", err)
	} else {
		t.Logf("L2 Output %s finalized: %t", params.L2OutputIndex.String(), isFinalized)
	}

	// 6. Check contract pause status
	paused, err := portal.Paused(&bind.CallOpts{})
	if err != nil {
		t.Logf("ERROR: Could not check pause status: %v", err)
	} else {
		t.Logf("Portal paused: %t", paused)
		if paused {
			t.Logf("ISSUE: Portal is paused - all transactions will fail")
		}
	}

	t.Logf("=== End Comprehensive Debug ===")
}
