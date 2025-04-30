package environment

import (
	"context"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
	"time"
)

// runSimpleL2Transfer runs a simple L2 burn transaction and verifies it on the
// L2 Verifier.
func RunSimpleL2Transfer(ctx context.Context, t *testing.T, system *e2esys.System, nonce uint64, amount big.Int, l2Seq *ethclient.Client, l2Verif *ethclient.Client) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	privateKey := system.Cfg.Secrets.Bob

	t.Log("Sending tx", "nonce", nonce)

	destAddress := system.Cfg.Secrets.Addresses().Alice
	receipt := helpers.SendL2Tx(
		t,
		system.Cfg,
		l2Seq,
		privateKey,
		L2TxWithOptions(
			L2TxWithAmount(&amount),
			L2TxWithNonce(nonce),
			L2TxWithToAddress(&destAddress),
			L2TxWithVerifyOnClients(l2Verif),
		),
	)
	t.Log("Receipt", receipt)

	cancel()
}

// runSimpleL1TransferAndVerifier runs a simple L1 transfer and verifies it on
// the L2 Verifier.
func RunSimpleL1TransferAndVerifier(ctx context.Context, t *testing.T, system *e2esys.System) {
	privateKey := system.Cfg.Secrets.Bob

	l1Client := system.NodeClient(e2esys.RoleL1)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	fromAddress := system.Cfg.Secrets.Addresses().Bob

	// Send Transaction on L1, and wait for verification on the L2 Verifier
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	// Get the Starting Balance of the Address
	startBalance, err := l2Verif.BalanceAt(ctx, fromAddress, nil)
	if have, want := err, error(nil); have != want {
		t.Errorf("attempt to get starting balance for %s failed:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", fromAddress, have, want)
	}

	// Create a new Keyed Transaction
	options, err := bind.NewKeyedTransactorWithChainID(privateKey, system.Cfg.L1ChainIDBig())
	if have, want := err, error(nil); have != want {
		t.Errorf("attempt to get keyed transaction with chain ID %d failed:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", system.Cfg.L1ChainIDBig(), have, want)
	}

	if err == nil {
		// We can only continue with these tests if the error above was nil

		// Send a Deposit Transaction
		mintAmount := big.NewInt(1_000_000_000_000)
		options.Value = mintAmount
		_ = helpers.SendDepositTx(t, system.Cfg, l1Client, l2Verif, options, nil)

		endBalance, err := wait.ForBalanceChange(ctx, l2Verif, fromAddress, startBalance)
		if have, want := err, error(nil); have != want {
			t.Errorf("waiting for balance change returned with error:\nhave:\n\t\"%v\"\nwant:\t\n\"%v\"\n", have, want)
		}

		diff := new(big.Int).Sub(endBalance, startBalance)
		if have, want := diff, mintAmount; have.Cmp(want) != 0 {
			t.Errorf("balance change does not match mint amount:\nhave;\n\t\"%s\"\nwant:\n\t\"%s\"\n", have, want)
		}
	}

	cancel()
}

// runSimpleL2Burn runs a simple L2 burn transaction and verifies it on the
// L2 Verifier.
func RunSimpleL2Burn(ctx context.Context, t *testing.T, system *e2esys.System) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	privateKey := system.Cfg.Secrets.Bob

	l2Seq := system.NodeClient(e2esys.RoleSeq)
	l2Verif := system.NodeClient(e2esys.RoleVerif)

	amountToBurn := big.NewInt(500_000_000)
	burnAddress := common.Address{0xff, 0xff}
	_ = helpers.SendL2Tx(
		t,
		system.Cfg,
		l2Seq,
		privateKey,
		L2TxWithOptions(
			L2TxWithAmount(amountToBurn),
			L2TxWithNonce(1), // Already have deposit
			L2TxWithToAddress(&burnAddress),
			L2TxWithVerifyOnClients(l2Verif),
		),
	)

	// Check the balance of hte burn address using the L2 Verifier
	balanceBurned, err := wait.ForBalanceChange(ctx, l2Verif, burnAddress, big.NewInt(0))
	if have, want := err, error(nil); have != want {
		t.Errorf("wait for balance change for burn address %s failed:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", burnAddress, have, want)
	}

	// Make sure that these match
	if have, want := balanceBurned, amountToBurn; have.Cmp(want) != 0 {
		t.Errorf("balance of burn address does not match amount burned:\nhave:\n\t\"%s\"\nwant:\n\t\"%s\"\n", have, want)
	}

	cancel()
}
