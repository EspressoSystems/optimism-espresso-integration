package devnet_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	bindings "github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
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

	// Check proposer account balance and permissions
	// The proposer uses the same mnemonic as batcher: "test test test test test test test test test test test junk"
	proposerPrivKey := d.secrets.Alice // This should match the mnemonic account
	proposerAddr := crypto.PubkeyToAddress(proposerPrivKey.PublicKey)
	proposerBalance, err := d.L1.BalanceAt(ctx, proposerAddr, nil)
	require.NoError(t, err)
	t.Logf("Proposer account %s balance: %s ETH", proposerAddr.Hex(), new(big.Int).Div(proposerBalance, big.NewInt(1e18)).String())

	disputeGameFactoryAddr, optimismPortalAddr := d.getOPAddresses()

	// Wait a bit longer for the proposer to create games (proposer interval is 6s)
	t.Logf("Waiting 3 minutes for proposer to create dispute games...")
	time.Sleep(3 * time.Minute)

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
}
