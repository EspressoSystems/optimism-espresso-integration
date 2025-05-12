package environment_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/setuputils"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

const MOCK_AUTHENTICATOR_ADDR = "0xBAD8ae8e955D24D64D0b21b839ceF63075757Bc7"

// TestE2eDevNetWithInvalidAttestation verifies that the batcher correctly fails to register
// when provided with an invalid attestation. This test ensures that the batch inbox contract
// properly validates attestations
func TestE2eDevNetWithoutAuthenticatingBatches(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	launcher := new(env.EspressoDevNodeLauncherDocker)

	system, _, err :=
		launcher.StartDevNet(ctx, t, 0,
			env.Config(func(cfg *e2esys.SystemConfig) {
				cfg.L1Allocs[common.HexToAddress(MOCK_AUTHENTICATOR_ADDR)] = types.Account{
					Code: hexutil.MustDecode("0x6080604052348015600e575f5ffd5b00fea26469706673582212202ea613f61d0ebf0addf2d722c3c06a3a613dd4932a28ed1c2e9aca8ca656808964736f6c634300081e0033"),
				}
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
	// if they don't succeed
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

	_, err = geth.WaitForBlockToBeSafe(new(big.Int).SetUint64(1), system.NodeClient(e2esys.RoleVerif), time.Minute*1)
	require.Error(t, err)
}

// This ETHBackend wraps a real ETHBackend and forwards all
// calls to it, except EstimateGas and CallContract calls, which always succeeds
// Wrapping Txmgr's backend with it ensures that Txmgr will always send
// transactions, even if they would be reverted.
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
