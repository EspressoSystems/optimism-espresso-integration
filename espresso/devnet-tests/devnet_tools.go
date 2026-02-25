package devnet_tests

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/helpers"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	opclient "github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/joho/godotenv"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/config/secrets"
)

type Devnet struct {
	ctx           context.Context
	secrets       secrets.Secrets
	outageTime    time.Duration
	successTime   time.Duration
	L1            *ethclient.Client
	L2Seq         *ethclient.Client
	L2SeqRollup   *sources.RollupClient
	L2Verif       *ethclient.Client
	L2VerifRollup *sources.RollupClient
}

// LoadEnvFile loads environment variables from a .env file
func LoadEnvFile(filename string) error {
	return godotenv.Load(filename)
}

// LoadDevnetEnv loads the espresso/.env file for devnet tests
func LoadDevnetEnv() error {
	// Get the path to the espresso/.env file relative to the test directory
	envPath := filepath.Join("..", ".env")
	return LoadEnvFile(envPath)
}

func NewDevnet(ctx context.Context, t *testing.T) *Devnet {

	if testing.Short() {
		t.Skip("skipping devnet test in short mode")
	}

	d := new(Devnet)
	d.ctx = ctx

	mnemonics := *secrets.DefaultMnemonicConfig
	mnemonics.Batcher = "m/44'/60'/0'/0/0"
	secrets, err := mnemonics.Secrets()
	if err != nil {
		panic(fmt.Sprintf("failed to create default secrets: %e", err))
	}
	d.secrets = *secrets

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

func (d *Devnet) isRunning() bool {
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "ps", "-q",
	)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		log.Error("failed to check if devnet is running", "error", err)
		return false
	}
	out := strings.TrimSpace(buf.String())
	return len(out) > 0
}

// The setting for `COMPOES_PROFILES` when running the Docker Compose.
type ComposeProfile string

const (
	TEE     ComposeProfile = "tee"
	NON_TEE ComposeProfile = "default"
)

func (d *Devnet) Up(profile ComposeProfile) (err error) {
	// If devnet is already running (e.g. leftover from a previous test in the same process,
	// or CI run), tear it down and then start fresh so this test can proceed.
	if d.isRunning() {
		log.Info("devnet already running, tearing down before starting fresh")
		if err := d.Down(); err != nil {
			return fmt.Errorf("tearing down existing devnet: %w", err)
		}
		// Brief wait so docker compose has released resources before we start again.
		for i := 0; i < 30; i++ {
			if d.ctx.Err() != nil {
				return fmt.Errorf("context cancelled while waiting for devnet to stop: %w", d.ctx.Err())
			}
			if !d.isRunning() {
				break
			}
			sleepContext(d.ctx, time.Second)
		}
		if d.isRunning() {
			return fmt.Errorf("devnet still running after Down(), shut it down manually")
		}
	}

	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "up", "-d",
		// Blockscout is a heavy monitoring stack not needed for devnet tests.
		// Scaling it to 0 reduces memory pressure and avoids OOM kills.
		"--scale", "blockscout=0",
		"--scale", "blockscout-frontend=0",
		"--scale", "blockscout-db=0",
	)
	cmd.Env = append(os.Environ(),
		"COMPOSE_PROFILES="+string(profile),
		fmt.Sprintf("OP_BATCHER_PRIVATE_KEY=%s", hex.EncodeToString(crypto.FromECDSA(d.secrets.Batcher))),
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

	if testing.Verbose() {
		// Stream logs to stdout while the test runs. This goroutine will automatically exit when
		// the context is cancelled.
		go func() {
			cmd = exec.CommandContext(d.ctx, "docker", "compose", "logs", "-f")
			cmd.Stdout = os.Stdout
			// We don't care about the error return of this command, since it's always going to be
			// killed by the context cancellation.
			_ = cmd.Run()
		}()
	}

	// Open RPC clients for the different nodes. Retry until containers are up and listening.
	if err := d.connectClientsWithRetry(); err != nil {
		return err
	}

	// Wait for nodes to respond to RPC so we don't return and then hang on first test RPC.
	// Use a bounded timeout so a bad rebase or broken devnet fails fast instead of timing out the test.
	waitCtx, waitCancel := context.WithTimeout(d.ctx, 5*time.Minute)
	defer waitCancel()
	if err := wait.ForNodeUp(waitCtx, d.L1, log.Root()); err != nil {
		return fmt.Errorf("L1 not ready: %w", err)
	}
	if err := wait.ForNodeUp(waitCtx, d.L2Seq, log.Root()); err != nil {
		return fmt.Errorf("L2 sequencer not ready: %w", err)
	}
	if err := wait.ForNodeUp(waitCtx, d.L2Verif, log.Root()); err != nil {
		return fmt.Errorf("L2 verifier not ready: %w", err)
	}

	return nil
}

// connectClientsWithRetry opens RPC clients, retrying until containers are up or timeout/context cancel.
func (d *Devnet) connectClientsWithRetry() error {
	const retryInterval = 5 * time.Second
	const retryTimeout = 2 * time.Minute
	deadline := time.Now().Add(retryTimeout)
	var err error
	for time.Now().Before(deadline) {
		if d.ctx.Err() != nil {
			return fmt.Errorf("context cancelled while connecting to devnet: %w", d.ctx.Err())
		}
		d.L2Seq, err = d.serviceClient("op-geth-sequencer", 8546)
		if err != nil {
			log.Debug("waiting for op-geth-sequencer", "err", err)
			sleepContext(d.ctx, retryInterval)
			continue
		}
		d.L2SeqRollup, err = d.rollupClient("op-node-sequencer", 9545)
		if err != nil {
			d.L2Seq.Close()
			d.L2Seq = nil
			log.Debug("waiting for op-node-sequencer", "err", err)
			sleepContext(d.ctx, retryInterval)
			continue
		}
		d.L2Verif, err = d.serviceClient("op-geth-verifier", 8546)
		if err != nil {
			d.L2Seq.Close()
			d.L2SeqRollup.Close()
			d.L2Seq, d.L2SeqRollup = nil, nil
			log.Debug("waiting for op-geth-verifier", "err", err)
			sleepContext(d.ctx, retryInterval)
			continue
		}
		d.L2VerifRollup, err = d.rollupClient("op-node-verifier", 9546)
		if err != nil {
			d.L2Verif.Close()
			d.L2Seq.Close()
			d.L2SeqRollup.Close()
			d.L2Verif, d.L2Seq, d.L2SeqRollup = nil, nil, nil
			log.Debug("waiting for op-node-verifier", "err", err)
			sleepContext(d.ctx, retryInterval)
			continue
		}
		d.L1, err = d.serviceClient("l1-geth", 8545)
		if err != nil {
			d.L2VerifRollup.Close()
			d.L2Verif.Close()
			d.L2Seq.Close()
			d.L2SeqRollup.Close()
			d.L2VerifRollup, d.L2Verif, d.L2Seq, d.L2SeqRollup = nil, nil, nil, nil
			log.Debug("waiting for l1-geth", "err", err)
			sleepContext(d.ctx, retryInterval)
			continue
		}
		return nil
	}
	return fmt.Errorf("devnet services did not become reachable within %v (last err: %w)", retryTimeout, err)
}

// sleepContext sleeps for d or until ctx is cancelled.
func sleepContext(ctx context.Context, d time.Duration) {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return
	case <-t.C:
		return
	}
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

// callBatcherRPC calls a batcher RPC method on a running batcher service
func (d *Devnet) callBatcherRPC(service, method string) error {
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "exec", "-T", service,
		"sh", "-c",
		fmt.Sprintf("wget -q -O- --header='Content-Type: application/json' --post-data='{\"jsonrpc\":\"2.0\",\"method\":\"%s\",\"params\":[],\"id\":1}' http://localhost:8545", method),
	)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to call %s (%w): %s", method, err, buf.String())
	}
	log.Info("RPC call successful", "service", service, "method", method, "response", buf.String())
	return nil
}

// StartBatcherSubmitting starts batch submission on a running batcher service
func (d *Devnet) StartBatcherSubmitting(service string) error {
	log.Info("starting batch submission", "service", service)
	return d.callBatcherRPC(service, "admin_startBatcher")
}

// StopBatcherSubmitting stops batch submission on a running batcher service
func (d *Devnet) StopBatcherSubmitting(service string) error {
	log.Info("stopping batch submission", "service", service)
	return d.callBatcherRPC(service, "admin_stopBatcher")
}

func (d *Devnet) RollupConfig(ctx context.Context) (*rollup.Config, error) {
	return d.L2SeqRollup.RollupConfig(ctx)
}

func (d *Devnet) SystemConfig(ctx context.Context) (*bindings.SystemConfig, *bind.TransactOpts, error) {
	config, err := d.RollupConfig(ctx)
	if err != nil {
		return nil, nil, err
	}
	contract, err := bindings.NewSystemConfig(config.L1SystemConfigAddress, d.L1)
	if err != nil {
		return nil, nil, err
	}

	owner, err := bind.NewKeyedTransactorWithChainID(d.secrets.Deployer, config.L1ChainID)
	if err != nil {
		return nil, nil, err
	}

	return contract, owner, nil
}

// Submits a transaction and waits until it is confirmed by the sequencer (but not necessarily the verifier).
func (d *Devnet) SubmitL2Tx(applyTxOpts helpers.TxOptsFn) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(d.ctx, 3*time.Minute)
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

	log.Info("submitted transaction to sequencer", "hash", tx.Hash(), "receipt", receipt)

	return receipt, nil
}

// Waits for a previously submitted transaction to be confirmed by the verifier.
func (d *Devnet) VerifyL2Tx(receipt *types.Receipt) error {
	timeout := 2 * time.Minute
	// Use longer timeout in CI environments due to Espresso processing delays
	if os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "" {
		timeout = 3 * time.Minute
		log.Info("CI environment detected, using extended timeout for transaction verification", "hash", receipt.TxHash, "timeout", timeout)
	}
	return d.VerifyL2TxWithTimeout(receipt, timeout)
}

// VerifyL2TxWithTimeout waits for the verifier to confirm the tx, using the given timeout.
func (d *Devnet) VerifyL2TxWithTimeout(receipt *types.Receipt, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(d.ctx, timeout)
	defer cancel()
	return d.verifyL2TxWithContext(ctx, receipt)
}

func (d *Devnet) verifyL2TxWithContext(ctx context.Context, receipt *types.Receipt) error {
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

func (d *Devnet) SendL1Tx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	err := d.L1.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return wait.ForReceiptOK(ctx, d.L1, tx.Hash())
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
	return d.VerifySimpleL2BurnWithTimeout(receipt, 2*time.Minute)
}

// VerifySimpleL2BurnWithTimeout waits for the verifier to confirm the burn, using the given timeout.
func (d *Devnet) VerifySimpleL2BurnWithTimeout(receipt *BurnReceipt, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(d.ctx, timeout)
	defer cancel()

	if err := d.verifyL2TxWithContext(ctx, receipt.Receipt); err != nil {
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

	if d.L1 != nil {
		d.L1.Close()
	}
	if d.L2Seq != nil {
		d.L2Seq.Close()
	}
	if d.L2SeqRollup != nil {
		d.L2SeqRollup.Close()
	}
	if d.L2Verif != nil {
		d.L2Verif.Close()
	}
	if d.L2VerifRollup != nil {
		d.L2VerifRollup.Close()
	}

	// Use timeout flag for faster Docker shutdown
	cmd := exec.CommandContext(
		d.ctx,
		"docker", "compose", "down", "-v", "--remove-orphans", "--timeout", "10",
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to shut down docker: %w", err)
	}

	outBatcher, _ := exec.Command("docker", "ps", "-q", "--filter", "ancestor=op-batcher-tee:espresso").Output()
	batcherContainers := strings.Fields(string(outBatcher))
	if len(batcherContainers) > 0 {
		cmd = exec.Command("docker", append([]string{"stop"}, batcherContainers...)...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to stop the batcher container: %w", err)
		}
		cmd = exec.Command("docker", append([]string{"rm"}, batcherContainers...)...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to remove the batcher container: %w", err)
		}
	}

	outEnclave, _ := exec.Command("docker", "ps", "-aq", "--filter", "name=batcher-enclaver-").Output()
	enclaveContainers := strings.Fields(string(outEnclave))
	if len(enclaveContainers) > 0 {
		cmd = exec.Command("docker", append([]string{"stop"}, enclaveContainers...)...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to stop the enclave container: %w", err)
		}
		cmd = exec.Command("docker", append([]string{"rm"}, enclaveContainers...)...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to remove the enclave container: %w", err)
		}
	}

	return nil
}

type TaggedWriter struct {
	inner   io.Writer
	tag     string
	newline bool
}

func NewTaggedWriter(tag string, inner io.Writer) *TaggedWriter {
	return &TaggedWriter{
		inner:   inner,
		tag:     tag,
		newline: true,
	}
}

// Implementation of io.Write interface for TaggedWriter.
// Allows to prepend a tag to each line of output.
// The `p` parameter is the tag to add at the beginning of each line.
func (w *TaggedWriter) Write(p []byte) (int, error) {
	if w.newline {
		if _, err := fmt.Fprintf(w.inner, "%s | ", w.tag); err != nil {
			return 0, err
		}
		w.newline = false
	}

	written := 0
	for i := range len(p) {
		// Buffer bytes until we hit a newline.
		if p[i] == '\n' {
			// Print everything we've buffered up to and including the newline.
			line := p[written : i+1]
			n, err := w.inner.Write(line)
			written += n
			if err != nil || n < len(line) {
				return written, err
			}

			// If that's the end of the output, return, but make a note that the buffer ended with a
			// newline and we need to print the tag before the next message.
			if written == len(p) {
				w.newline = true
				return written, nil
			}

			// Otherwise print the tag now before proceeding with the next line in `p`.
			if _, err := fmt.Fprintf(w.inner, "%s | ", w.tag); err != nil {
				return written, err
			}
		}
	}

	// Print anything that was buffered after the final newline.
	if written < len(p) {
		line := p[written:]
		n, err := w.inner.Write(line)
		written += n
		if err != nil || n < len(line) {
			return written, err
		}
	}

	return written, nil
}

func (d *Devnet) OpChallenger(opts ...string) error {
	return d.opChallengerCmd(opts...).Run()
}

type ChallengeGame struct {
	Index      uint64
	Address    common.Address
	OutputRoot []byte
	Claims     uint64
}

func ParseChallengeGame(s string) (ChallengeGame, error) {
	fields := strings.Fields(s)
	if len(fields) < 8 {
		return ChallengeGame{}, fmt.Errorf("challenge game is missing fields; expected at least 8 but got only %v", len(fields))
	}

	index, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		return ChallengeGame{}, fmt.Errorf("index invalid: %w", err)
	}

	address := common.HexToAddress(fields[1])

	outputRoot := common.Hex2Bytes(fields[6])

	claims, err := strconv.ParseUint(fields[7], 10, 64)
	if err != nil {
		return ChallengeGame{}, fmt.Errorf("claims count invalid: %w", err)
	}

	return ChallengeGame{
		Index:      index,
		Address:    address,
		OutputRoot: outputRoot,
		Claims:     claims,
	}, nil
}

func (d *Devnet) ListChallengeGames() ([]ChallengeGame, error) {
	// Succinct only supports contract-based query
	games, err := d.ListChallengeGamesFromContract()
	if err == nil && len(games) > 0 {
		return games, nil
	}

	return nil, fmt.Errorf("failed to list challenge games: %w", err)
}

// ListChallengeGamesFromContract queries games directly from the DisputeGameFactory contract
// Only supports OPSuccinctFaultDisputeGame (game type 42)
func (d *Devnet) ListChallengeGamesFromContract() ([]ChallengeGame, error) {
	ctx := d.ctx

	// Get SystemConfig to find DisputeGameFactory address
	systemConfig, _, err := d.SystemConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get system config: %w", err)
	}

	factoryAddr, err := systemConfig.DisputeGameFactory(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get dispute game factory address: %w", err)
	}

	// Bind to DisputeGameFactory
	factory, err := bindings.NewDisputeGameFactory(factoryAddr, d.L1)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to dispute game factory: %w", err)
	}

	// Get game count
	gameCount, err := factory.GameCount(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get game count: %w", err)
	}

	var games []ChallengeGame
	for i := uint64(0); i < gameCount.Uint64(); i++ {
		// Get game at index
		gameInfo, err := factory.GameAtIndex(&bind.CallOpts{}, new(big.Int).SetUint64(i))
		if err != nil {
			log.Warn("failed to get game at index", "index", i, "error", err)
			continue
		}

		// Only include game type 42 (OPSuccinctFaultDisputeGame)
		if gameInfo.GameType != 42 {
			continue
		}

		gameProxy := gameInfo.Proxy

		// Get root claim from the game contract
		// OPSuccinctFaultDisputeGame only has root claim, no claim tree
		disputeGame, err := bindings.NewFaultDisputeGame(gameProxy, d.L1)
		if err != nil {
			log.Warn("failed to bind to dispute game", "address", gameProxy, "error", err)
			continue
		}

		rootClaim, err := disputeGame.RootClaim(&bind.CallOpts{})
		if err != nil {
			log.Warn("failed to get root claim", "address", gameProxy, "error", err)
			continue
		}

		games = append(games, ChallengeGame{
			Index:      i,
			Address:    gameProxy,
			OutputRoot: rootClaim[:],
			Claims:     1, // OPSuccinctFaultDisputeGame only has root claim
		})
	}

	return games, nil
}

func (d *Devnet) OpChallengerOutput(opts ...string) (string, error) {
	cmd := d.opChallengerCmd(opts...)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (d *Devnet) opChallengerCmd(opts ...string) *exec.Cmd {
	opts = append([]string{"compose", "exec", "op-challenger", "entrypoint.sh", "op-challenger"}, opts...)
	cmd := exec.CommandContext(
		d.ctx,
		"docker",
		opts...,
	)
	if testing.Verbose() {
		cmd.Stdout = NewTaggedWriter("op-challenger-cmd", os.Stdout)
		cmd.Stderr = NewTaggedWriter("op-challenger-cmd", os.Stderr)
	}
	log.Info("invoking op-challenger", "cmd", cmd)
	return cmd
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
	client, err := ethclient.DialContext(d.ctx, fmt.Sprintf("http://127.0.0.1:%d", port))
	if err != nil {
		return nil, fmt.Errorf("could not open %s RPC client: %w", service, err)
	}
	return client, nil
}

func (d *Devnet) rollupClient(service string, port uint16) (*sources.RollupClient, error) {
	port, err := d.hostPort(service, port)
	if err != nil {
		return nil, fmt.Errorf("could not get %s port: %w", service, err)
	}
	rpc, err := opclient.NewRPC(d.ctx, log.Root(), fmt.Sprintf("http://127.0.0.1:%d", port), opclient.WithDialAttempts(10))
	if err != nil {
		return nil, fmt.Errorf("could not open %s RPC client: %w", service, err)
	}

	client := sources.NewRollupClient(rpc)
	return client, nil
}
