package devnet_tests

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/config/secrets"
)

type Devnet struct {
	ctx         context.Context
	secrets     secrets.Secrets
	outageTime  time.Duration
	successTime time.Duration

	L2Seq   *ethclient.Client
	L2Verif *ethclient.Client
}

func NewDevnet(ctx context.Context, t *testing.T) *Devnet {
	if testing.Short() {
		t.Skip("skipping devnet test in short mode")
	}

	d := new(Devnet)
	d.ctx = ctx
	d.secrets = *secrets.DefaultSecrets

	var err error
	if outageTime, ok := os.LookupEnv("ESPRESSO_DEVNET_TESTS_OUTAGE_PERIOD"); ok {
		d.outageTime, err = time.ParseDuration(outageTime)
		if err != nil {
			panic(fmt.Sprintf("invalid value for ESPRESSO_DEVNET_TESTS_OUTAGE_PERIOD: %e", err))
		}
	} else {
		d.outageTime = 10 * time.Second
	}
	if successTime, ok := os.LookupEnv("ESPRESSO_DEVNET_TESTS_LIVENESS_PERIOD"); ok {
		d.successTime, err = time.ParseDuration(successTime)
		if err != nil {
			panic(fmt.Sprintf("invalid value for ESPRESSO_DEVNET_TESTS_LIVENESS_PERIOD: %e", err))
		}
	} else {
		d.successTime = 10 * time.Second
	}

	return d
}

func (d *Devnet) Up() (err error) {
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "up", "-d",
	)
	buf := new(bytes.Buffer)
	cmd.Stderr = buf
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start docker compose (%w): %s", err, buf.String())
	}

	// Shut down the now-running devnet if we exit this function with an error (in which case the
	// caller expects the devnet not to be running and will not be responsible for shutting it down
	// themselves).
	defer func() {
		if err != nil {
			if downErr := d.Down(); downErr != nil {
				log.Error("error shutting down devnet after encountering another error", "error", downErr)
			}
		}
	}()

	// Shut down the devnet automatically when the lifetime of the context ends.
	go func() {
		<-d.ctx.Done()
		if err := d.Down(); err != nil {
			log.Error("error shutting down devnet asynchronously", "error", err)
		}
	}()

	// Open RPC clients for the different nodes.
	d.L2Seq, err = d.serviceClient("op-geth-sequencer", 8546)
	if err != nil {
		return err
	}
	d.L2Verif, err = d.serviceClient("op-geth-verifier", 8546)
	if err != nil {
		return err
	}

	return nil
}

func (d *Devnet) ServiceUp(service string) error {
	log.Info("bringing up service", "service", service)
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "up", "-d", service,
	)
	return cmd.Run()
}

func (d *Devnet) ServiceDown(service string) error {
	log.Info("shutting down service", "service", service)
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "down", service,
	)
	return cmd.Run()
}

func (d *Devnet) ServiceRestart(service string) error {
	if err := d.ServiceDown(service); err != nil {
		return err
	}
	if err := d.ServiceUp(service); err != nil {
		return err
	}
	return nil
}

// Submits a transaction and waits until it is confirmed by the sequencer (but not necessarily the verifier).
func (d *Devnet) SubmitL2Tx(applyTxOpts helpers.TxOptsFn) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(d.ctx, 2*time.Minute)
	defer cancel()

	chainID, err := d.L2Seq.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	privKey := d.secrets.Alice
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	balance, err := d.L2Seq.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("getting initial sender balance: %w", err)
	}
	if balance.Cmp(big.NewInt(0)) <= 0 {
		return nil, fmt.Errorf("sender account empty")
	}
	nonce, err := d.L2Seq.NonceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting nonce: %w", err)
	}
	log.Debug("sender wallet", "private key", privKey, "address", address, "balance", balance, "nonce", nonce)

	opts := &helpers.TxOpts{
		ToAddr:         nil,
		Nonce:          nonce,
		Value:          common.Big0,
		GasTipCap:      big.NewInt(10),
		GasFeeCap:      big.NewInt(1000000000),
		Gas:            21_000,
		Data:           nil,
		ExpectedStatus: types.ReceiptStatusSuccessful,
	}
	applyTxOpts(opts)

	tx := types.MustSignNewTx(privKey, types.LatestSignerForChainID(chainID), &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     opts.Nonce,
		To:        opts.ToAddr,
		Value:     opts.Value,
		GasTipCap: opts.GasTipCap,
		GasFeeCap: opts.GasFeeCap,
		Gas:       opts.Gas,
		Data:      opts.Data,
	})
	log.Info("send transaction", "from", address, "hash", tx.Hash())
	if err := d.L2Seq.SendTransaction(ctx, tx); err != nil {
		return nil, fmt.Errorf("sending L2 tx: %w", err)
	}

	receipt, err := wait.ForReceiptOK(ctx, d.L2Seq, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("waiting for L2 tx: %w", err)
	}
	if opts.ExpectedStatus != receipt.Status {
		return nil, fmt.Errorf("wrong status: have %d, want %d", receipt.Status, opts.ExpectedStatus)
	}

	return receipt, nil
}

// Waits for a previously submitted transaction to be confirmed by the verifier.
func (d *Devnet) VerifyL2Tx(receipt *types.Receipt) error {
	ctx, cancel := context.WithTimeout(d.ctx, 2*time.Minute)
	defer cancel()

	log.Info("waiting for transaction verification", "hash", receipt.TxHash)
	verified, err := wait.ForReceiptOK(ctx, d.L2Verif, receipt.TxHash)
	if err != nil {
		return fmt.Errorf("waiting for L2 tx on verification client: %w", err)
	}
	if !reflect.DeepEqual(receipt, verified) {
		return fmt.Errorf("verification client returned incorrect receipt\nSeq:  %v\nVerif: %v", receipt, verified)
	}
	return nil
}

// Submits a transaction and waits for it to be verified.
func (d *Devnet) RunL2Tx(applyTxOpts helpers.TxOptsFn) error {
	receipt, err := d.SubmitL2Tx(applyTxOpts)
	if err != nil {
		return err
	}
	return d.VerifyL2Tx(receipt)
}

type BurnReceipt struct {
	InitialBurnBalance *big.Int
	BurnAmount         *big.Int
	BurnAddress        common.Address
	Receipt            *types.Receipt
}

// Submits a burn transaction and waits until it is confirmed by the sequencer (but not necessarily the verifier).
func (d *Devnet) SubmitSimpleL2Burn() (*BurnReceipt, error) {
	var err error

	receipt := new(BurnReceipt)
	receipt.BurnAddress = common.Address{0xff, 0xff}
	receipt.BurnAmount = big.NewInt(1)

	receipt.InitialBurnBalance, err = d.L2Verif.BalanceAt(d.ctx, receipt.BurnAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("getting initial burn address balance: %w", err)
	}

	tx := env.L2TxWithOptions(
		env.L2TxWithAmount(receipt.BurnAmount),
		env.L2TxWithToAddress(&receipt.BurnAddress),
		env.L2TxWithVerifyOnClients(d.L2Verif),
	)
	if receipt.Receipt, err = d.SubmitL2Tx(tx); err != nil {
		return nil, err
	}
	return receipt, nil
}

// Waits for a previously submitted burn transaction to be confirmed by the verifier.
func (d *Devnet) VerifySimpleL2Burn(receipt *BurnReceipt) error {
	ctx, cancel := context.WithTimeout(d.ctx, 2*time.Minute)
	defer cancel()

	if err := d.VerifyL2Tx(receipt.Receipt); err != nil {
		return err
	}

	// Check the balance of the burn address using the L2 Verifier
	final, err := wait.ForBalanceChange(ctx, d.L2Verif, receipt.BurnAddress, receipt.InitialBurnBalance)
	if err != nil {
		return fmt.Errorf("waiting for balance change for burn address %s: %w", receipt.BurnAddress, err)
	}
	balanceBurned := new(big.Int).Sub(final, receipt.InitialBurnBalance)
	if balanceBurned.Cmp(receipt.BurnAmount) != 0 {
		return fmt.Errorf("incorrect amount burned (have %s, want %s)", balanceBurned, receipt.BurnAmount)
	}

	return nil
}

// RunSimpleL2Burn runs a simple L2 burn transaction and verifies it on the L2 Verifier.
func (d *Devnet) RunSimpleL2Burn() error {
	receipt, err := d.SubmitSimpleL2Burn()
	if err != nil {
		return err
	}
	return d.VerifySimpleL2Burn(receipt)
}

// Wait for a configurable amount of time while simulating an outage.
func (d *Devnet) SleepOutageDuration() {
	log.Info("sleeping during simulated outage", "duration", d.outageTime)
	time.Sleep(d.outageTime)
}

// Wait for a configurable amount of time before considering a run a success.
func (d *Devnet) SleepRecoveryDuration() {
	log.Info("sleeping to check that things stay working", "duration", d.successTime)
	time.Sleep(d.successTime)
}

func (d *Devnet) Down() error {
	log.Info("devnet shutting down")
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "down", "-v", "--remove-orphans",
	)
	return cmd.Run()
}

// Get the host port mapped to `privatePort` for the given Docker service.
func (d *Devnet) hostPort(service string, privatePort uint16) (uint16, error) {
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "port", service, fmt.Sprint(privatePort),
	)
	cmd.Stdout = buf
	cmd.Stderr = errBuf

	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("command failed (%w)\nStdout: %s\nStderr: %s", err, buf.String(), errBuf.String())
	}
	out := strings.TrimSpace(buf.String())
	_, portStr, found := strings.Cut(out, ":")
	if !found {
		return 0, fmt.Errorf("invalid output from docker port: %s (missing : separator)", out)
	}

	port, err := strconv.ParseInt(portStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid output from docker port: %s (%w)", out, err)
	}
	return uint16(port), nil
}

// Open an Ethereum RPC client for a Docker service running an RPC server on the given port.
func (d *Devnet) serviceClient(service string, port uint16) (*ethclient.Client, error) {
	port, err := d.hostPort(service, port)
	if err != nil {
		return nil, fmt.Errorf("could not get %s port: %w", service, err)
	}
	client, err := ethclient.DialContext(d.ctx, fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		return nil, fmt.Errorf("could not open %s RPC client: %w", service, err)
	}
	return client, nil
}
