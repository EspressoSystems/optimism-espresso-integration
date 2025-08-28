package devnet_tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/bindings"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	l1Client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)
	defer l1Client.Close()

	aliceL1Balance, err := l1Client.BalanceAt(ctx, aliceAddress, nil)
	require.NoError(t, err)
	expectedBalance := new(big.Int)
	expectedBalance.SetString("10000000000000000000000", 10) // 10,000 ETH in wei
	require.True(t, aliceL1Balance.Cmp(expectedBalance) == 0, "Alice should have exactly 10,000 ETH")
	t.Logf("Alice's L1 balance: %s wei", aliceL1Balance.String())

}
