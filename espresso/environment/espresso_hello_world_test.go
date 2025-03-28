package environment_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"

	op_e2e "github.com/ethereum-optimism/optimism/op-e2e"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// TestSystemE2E sets up a L1 Geth node, a rollup node, and a L2 geth node and then confirms that L1 deposits are reflected on L2.
// All nodes are run in process (but are the full nodes, not mocked or stubbed).
// How to run this test:
// > (cd packages/contracts-bedrock && just build-dev)
// > cd espresso/environment
// > go test espresso_hello_world_test.go
func TestHelloWorldEspresso(t *testing.T) {
	op_e2e.InitParallel(t)

	// ESPRESSO: Maybe we need to tweak the config here
	cfg := e2esys.DefaultSystemConfig(t)

	sys, err := cfg.Start(t)

	log := testlog.Logger(t, log.LevelInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.NodeClient("l1")
	l2Seq := sys.NodeClient("sequencer")
	l2Verif := sys.NodeClient("verifier")

	// ESPRESSO
	// Here we should spin up a Caff node, or maybe change the e2sys logic so that the Caff node is already predefined and can be accessed via
	// l2VerifCaff:= sys.NodeClient("caff_verifier")

	// Transactor Account
	ethPrivKey := sys.Cfg.Secrets.Alice

	// Send Transaction & wait for success
	fromAddr := sys.Cfg.Secrets.Addresses().Alice

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	startBalance, err := l2Verif.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)

	// Send deposit transaction
	// Philippe: I understand that this tx is sent on L1 with effect on L2 (bridge transaction)
	opts, err := bind.NewKeyedTransactorWithChainID(ethPrivKey, sys.Cfg.L1ChainIDBig())
	require.Nil(t, err)
	mintAmount := big.NewInt(1_000_000_000_000)
	opts.Value = mintAmount
	helpers.SendDepositTx(t, sys.Cfg, l1Client, l2Verif, opts, nil)

	// Confirm balance
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	endBalance, err := wait.ForBalanceChange(ctx, l2Verif, fromAddr, startBalance)
	require.Nil(t, err)

	diff := new(big.Int)
	diff = diff.Sub(endBalance, startBalance)
	require.Equal(t, mintAmount, diff, "Did not get expected balance change")

	// Submit TX to L2 sequencer node
	// This tx burns the tokens

	amountBurnt := big.NewInt(500_000_000)
	burnAddress := common.Address{0xff, 0xff}
	helpers.SendL2Tx(t, sys.Cfg, l2Seq, ethPrivKey, func(opts *helpers.TxOpts) {
		opts.Value = amountBurnt
		opts.Nonce = 1 // Already have deposit
		opts.ToAddr = &burnAddress
		opts.VerifyOnClients(l2Verif)
	})

	// Philippe: Check the money has been burnt using normal OP Node
	balanceBurnAddress, err := wait.ForBalanceChange(ctx, l2Verif, burnAddress, big.NewInt(0))
	require.Nil(t, err)

	// All the money received has been burnt
	require.Equal(t, balanceBurnAddress, amountBurnt)

	// ESPRESSO
	// We should also check that the money has been burnt using the Caff node
	//....

}
