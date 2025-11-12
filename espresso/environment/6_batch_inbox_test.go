package environment_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/setuputils"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// Test private key for PreApprovedBatcher (TEE batcher)
const preApprovedBatcherPrivateKey = "5fede428b9506dee864b0d85aefb2409f4728313eb41da4121409299c487f816"

func setupBatchInboxEnv(ctx context.Context, t *testing.T) (*e2esys.System, *bindings.BatchInbox, *big.Int) {
    t.Helper()
    launcher := new(env.EspressoDevNodeLauncherDocker)
    system, _, err := launcher.StartE2eDevnet(ctx, t,
        env.Config(func(cfg *e2esys.SystemConfig) {
            cfg.DisableBatcher = true
        }),
    )
    require.NoError(t, err)

    l1 := system.NodeClient(e2esys.RoleL1)
    chainID, err := l1.ChainID(ctx)
    require.NoError(t, err)

    inbox, err := bindings.NewBatchInbox(system.RollupConfig.BatchInboxAddress, l1)
    require.NoError(t, err)

    // Fund the PreApprovedBatcher account if needed
    pk, _ := crypto.HexToECDSA(preApprovedBatcherPrivateKey)
    addr := crypto.PubkeyToAddress(pk.PublicKey)
    if balance, _ := l1.BalanceAt(ctx, addr, nil); balance.Sign() == 0 {
        nonce, _ := l1.PendingNonceAt(ctx, crypto.PubkeyToAddress(system.Cfg.Secrets.Deployer.PublicKey))
        gasPrice, _ := l1.SuggestGasPrice(ctx)
        tx := types.NewTx(&types.LegacyTx{
            Nonce:    nonce,
            To:       &addr,
            Value:    big.NewInt(1e18),
            Gas:      21000,
            GasPrice: gasPrice,
        })
        signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), system.Cfg.Secrets.Deployer)
        l1.SendTransaction(ctx, signedTx)
        bind.WaitMined(ctx, l1, signedTx)
    }

    return system, inbox, chainID
}

func authForAddress(t *testing.T, system *e2esys.System, chainID *big.Int, addr common.Address) *bind.TransactOpts {
	t.Helper()
	for _, secret := range []*ecdsa.PrivateKey{
		system.Cfg.Secrets.Deployer,
		system.Cfg.Secrets.Batcher,
		system.Cfg.Secrets.Bob,
	} {
		if crypto.PubkeyToAddress(secret.PublicKey) == addr {
			auth, _ := bind.NewKeyedTransactorWithChainID(secret, chainID)
			return auth
		}
	}
	// Check PreApprovedBatcher
	if pk, _ := crypto.HexToECDSA(preApprovedBatcherPrivateKey); crypto.PubkeyToAddress(pk.PublicKey) == addr {
		auth, _ := bind.NewKeyedTransactorWithChainID(pk, chainID)
		return auth
	}
	t.Fatalf("no auth available for address %s", addr)
	return nil
}

// TestE2eDevnetWithoutAuthenticatingBatches verifies BatchInboxContract behaviour when batches
// aren't attested before being posted to batch inbox. To do this, we substitute BatchAuthenticatorAddress
// in batcher config with a zero address, which will never revert as it has no contract deployed.
// This way we trick batcher into posting unauthenticated batches to batch inbox.
// We then verify that these batches aren't accepted by the batch inbox contract and derivation pipeline.
//
// The test is defined as follows
// Arrange:
//
//	Deploy a mock BatchAuthenticator.
//	Configure batcher to use said authenticator instead of the real one.
//	Start sequencer, batcher in Espresso mode and OP node.
//
// Assert:
//
//	Assert that transaction submitting the batch was reverted by
//	batch inbox contract
//	Assert that derivation pipeline doesn't progress
func TestE2eDevnetWithoutAuthenticatingBatches(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, _, err :=
		launcher.StartE2eDevnet(ctx, t,
			env.Config(func(cfg *e2esys.SystemConfig) {
				cfg.DisableBatcher = true
			}),
		)

	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to start dev environment with espresso dev node:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	batchDriver := system.BatchSubmitter.TestDriver()
	// Set mock batcher authenticator address
	batchDriver.BatchSubmitter.RollupConfig.BatchAuthenticatorAddress = common.Address{}

	// Substitute batcher's transaction manager with one that always sends transactions, even
	// if they won't succeed. Otherwise batcher wouldn't submit transactions that would revert to
	// batch inbox
	txMgrCliConfig := setuputils.NewTxMgrConfig(system.NodeEndpoint(e2esys.RoleL1), system.Cfg.Secrets.Batcher)
	txMgrConfig, err := txmgr.NewConfig(txMgrCliConfig, log.Root())
	require.NoError(t, err)
	txMgrConfig.Backend = AlwaysSendingETHBackend{
		inner: txMgrConfig.Backend,
	}
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig("always-sending", log.Root(), &metrics.NoopTxMetrics{}, txMgrConfig)
	require.NoError(t, err)
	batchDriver.Txmgr = txMgr

	// Start the batcher
	err = batchDriver.StartBatchSubmitting()
	require.NoError(t, err, "Couldn't start batcher")
	l1Client := system.NodeClient(e2esys.RoleL1)

	// Wait for batcher to submit a transaction to BatchInbox
	var batchInboxTxHash common.Hash
	for {
		l1Height, err := l1Client.BlockNumber(ctx)
		require.NoError(t, err)
		_, err = geth.FindBlock(l1Client,
			0,
			int(l1Height),
			time.Minute*2,
			func(block *types.Block) (bool, error) {
				for _, tx := range block.Transactions() {
					if *tx.To() == system.RollupConfig.BatchInboxAddress {
						batchInboxTxHash = tx.Hash()
						return true, nil
					}
				}
				return false, nil
			})
		if err == nil {
			break
		}
	}

	receipt, err := l1Client.TransactionReceipt(ctx, batchInboxTxHash)
	require.NoError(t, err)

	require.Equal(t, receipt.Status, types.ReceiptStatusFailed, "transaction should've been rejected by BatchInbox contract")

	_, err = geth.WaitForBlockToBeSafe(new(big.Int).SetUint64(1), system.NodeClient(e2esys.RoleVerif), time.Minute)
	require.Error(t, err)
}

func TestBatchInbox_SwitchActiveBatcher(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    system, inbox, chainID := setupBatchInboxEnv(ctx, t)
    deployerAuth, _ := bind.NewKeyedTransactorWithChainID(system.Cfg.Secrets.Deployer, chainID)
    tx, err := inbox.SwitchBatcher(deployerAuth)
    require.NoError(t, err)
    _, err = bind.WaitMined(ctx, system.NodeClient(e2esys.RoleL1), tx)
    require.NoError(t, err)
}

func TestBatchInbox_ActiveNonTeeBatcherAllowsPosting(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    system, inbox, chainID := setupBatchInboxEnv(ctx, t)
    deployerAuth, _ := bind.NewKeyedTransactorWithChainID(system.Cfg.Secrets.Deployer, chainID)
    tx, err := inbox.SwitchBatcher(deployerAuth)
    require.NoError(t, err)
    _, err = bind.WaitMined(ctx, system.NodeClient(e2esys.RoleL1), tx)
    require.NoError(t, err)
    // Determine non-TEE batcher from contract and post with its key
    nonTeeAddr, err := inbox.NonTeeBatcher(&bind.CallOpts{Context: ctx})
    require.NoError(t, err)
    nonTeeAuth := authForAddress(t, system, chainID, nonTeeAddr)
    tx2, err := inbox.PostCalldata(nonTeeAuth, []byte("hello"))
    require.NoError(t, err)
    receipt, err := bind.WaitMined(ctx, system.NodeClient(e2esys.RoleL1), tx2)
    require.NoError(t, err)
    require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
}

func TestBatchInbox_InactiveBatcherReverts(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    system, inbox, chainID := setupBatchInboxEnv(ctx, t)
    deployerAuth, _ := bind.NewKeyedTransactorWithChainID(system.Cfg.Secrets.Deployer, chainID)
    tx, err := inbox.SwitchBatcher(deployerAuth)
    require.NoError(t, err)
    _, err = bind.WaitMined(ctx, system.NodeClient(e2esys.RoleL1), tx)
    require.NoError(t, err)
    teeAddr, err := inbox.TeeBatcher(&bind.CallOpts{Context: ctx})
    require.NoError(t, err)
    teeAuth := authForAddress(t, system, chainID, teeAddr)
    teeAuth.GasLimit = 100000 // Bypass gas estimation
    tx2, err := inbox.PostCalldata(teeAuth, []byte("unauth"))
    require.NoError(t, err)
    receipt, err := bind.WaitMined(ctx, system.NodeClient(e2esys.RoleL1), tx2)
    require.NoError(t, err)
    require.Equal(t, types.ReceiptStatusFailed, receipt.Status)
}

func TestBatchInbox_TEEBatcherRequiresAuthentication(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    system, inbox, chainID := setupBatchInboxEnv(ctx, t)
    teeAddr, err := inbox.TeeBatcher(&bind.CallOpts{Context: ctx})
    require.NoError(t, err)
    teeAuth := authForAddress(t, system, chainID, teeAddr)
    // Disable gas estimation to force sending a transaction that will revert
    teeAuth.GasLimit = 100000
    teeAuth.NoSend = false
    tx, err := inbox.PostCalldata(teeAuth, []byte("needs-auth"))
    require.NoError(t, err)
    receipt, err := bind.WaitMined(ctx, system.NodeClient(e2esys.RoleL1), tx)
    require.NoError(t, err)
    require.Equal(t, types.ReceiptStatusFailed, receipt.Status)
}

// A wrapper for testing that proxies all calls to ETHBackend unchanged,
// except EstimateGas and CallContract calls, which always "succeed"
// without making any actual RPC calls.
//
// Wrapping SimpleTxManager's backend with it ensures that SimpleTxManager will always send
// transactions, even if they would be reverted. The reason for this behaviour is
// that SimpleTxManager will check whether transaction will be executed successfully
// before submitting it, either by calling CallContract if transaction request had
// set the gas cap, or by checking EstimateGas return value if transaction request
// doesn't have the gas cap set. Mocking these two methods to always succeed thus
// makes SimpleTxManager submit even invalid transactions, which it wouldn't normally do.
type AlwaysSendingETHBackend struct {
	inner txmgr.ETHBackend
}

// BlockNumber implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) BlockNumber(ctx context.Context) (uint64, error) {
	return m.inner.BlockNumber(ctx)
}

// CallContract implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return []byte{}, nil
}

// Close implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) Close() {
	m.inner.Close()
}

// EstimateGas implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 1_000_000, nil
}

// HeaderByNumber implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return m.inner.HeaderByNumber(ctx, number)
}

// NonceAt implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return m.inner.NonceAt(ctx, account, blockNumber)
}

// PendingNonceAt implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return m.inner.PendingNonceAt(ctx, account)
}

// SendTransaction implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return m.inner.SendTransaction(ctx, tx)
}

// SuggestGasTipCap implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return m.inner.SuggestGasTipCap(ctx)
}

// TransactionReceipt implements txmgr.ETHBackend.
func (m AlwaysSendingETHBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return m.inner.TransactionReceipt(ctx, txHash)
}

// Ensure conformance to ETHBackend
var _ txmgr.ETHBackend = AlwaysSendingETHBackend{}
