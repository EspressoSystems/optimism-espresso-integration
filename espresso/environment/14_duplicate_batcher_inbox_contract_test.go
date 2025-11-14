package environment_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// Test private key for PreApprovedBatcher (TEE batcher)
const preApprovedBatcherPrivateKey = "5fede428b9506dee864b0d85aefb2409f4728313eb41da4121409299c487f816"

func setupBatchInboxEnv(ctx context.Context, t *testing.T) (*e2esys.System, *bindings.BatchInbox, *big.Int) {
	t.Helper()
	launcher := &env.EspressoDevNodeLauncherDocker{
		EnclaveBatcher: false, // Explicitly set to use non-enclave mode
	}
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
	tx2, err := inbox.Fallback(nonTeeAuth, []byte("hello"))
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
	tx2, err := inbox.Fallback(teeAuth, []byte("unauth"))
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
	tx, err := inbox.Fallback(teeAuth, []byte("needs-auth"))
	require.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, system.NodeClient(e2esys.RoleL1), tx)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusFailed, receipt.Status)
}
