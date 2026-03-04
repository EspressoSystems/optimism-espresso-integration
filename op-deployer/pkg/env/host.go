package env

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script/forking"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/broadcaster"

	"github.com/ethereum-optimism/optimism/op-chain-ops/foundry"
	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func DefaultScriptHost(
	bcaster broadcaster.Broadcaster,
	lgr log.Logger,
	deployer common.Address,
	artifacts foundry.StatDirFs,
	additionalOpts ...script.HostOption,
) (*script.Host, error) {
	scriptCtx := script.DefaultContext
	scriptCtx.Sender = deployer
	scriptCtx.Origin = deployer
	h := script.NewHost(
		lgr,
		&foundry.ArtifactsFS{FS: artifacts},
		nil,
		scriptCtx,
		append([]script.HostOption{
			script.WithBroadcastHook(bcaster.Hook),
			script.WithIsolatedBroadcasts(),
			script.WithCreate2Deployer(),
		}, additionalOpts...)...,
	)

	if err := h.EnableCheats(); err != nil {
		return nil, fmt.Errorf("failed to enable cheats: %w", err)
	}

	return h, nil
}

func DefaultForkedScriptHost(
	ctx context.Context,
	bcaster broadcaster.Broadcaster,
	lgr log.Logger,
	deployer common.Address,
	artifacts foundry.StatDirFs,
	forkRPC *rpc.Client,
	additionalOpts ...script.HostOption,
) (*script.Host, error) {
	client := ethclient.NewClient(forkRPC)

	latest, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	return ForkedScriptHost(
		bcaster,
		lgr,
		deployer,
		artifacts,
		forkRPC,
		latest.Number,
		additionalOpts...,
	)
}

// ForkedScriptHostForBootstrap creates a script host for bootstrap superchain (and similar)
// without WithIsolatedBroadcasts, so that CREATEs inside vm.broadcast() persist in the
// simulation and scripts like DeploySuperchain.assertValidOutput succeed.
// Uses WithCreate2Deployer so the handleCaller hook overrides the CREATE2 caller to
// DeterministicDeployerAddress, matching what computeCreate2Address expects.
func ForkedScriptHostForBootstrap(
	ctx context.Context,
	bcaster broadcaster.Broadcaster,
	lgr log.Logger,
	deployer common.Address,
	artifacts foundry.StatDirFs,
	forkRPC *rpc.Client,
) (*script.Host, error) {
	client := ethclient.NewClient(forkRPC)
	latest, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}
	// Same as DefaultScriptHost but omit WithIsolatedBroadcasts so broadcast CREATEs persist.
	scriptCtx := script.DefaultContext
	scriptCtx.Sender = deployer
	scriptCtx.Origin = deployer
	opts := []script.HostOption{
		script.WithBroadcastHook(bcaster.Hook),
		script.WithCreate2Deployer(),
		script.WithForkHook(func(cfg *script.ForkConfig) (forking.ForkSource, error) {
			src, err := forking.RPCSourceByNumber(cfg.URLOrAlias, forkRPC, *cfg.BlockNumber)
			if err != nil {
				return nil, fmt.Errorf("failed to create RPC fork source: %w", err)
			}
			return forking.Cache(src), nil
		}),
	}
	h := script.NewHost(
		lgr,
		&foundry.ArtifactsFS{FS: artifacts},
		nil,
		scriptCtx,
		opts...,
	)
	if err := h.EnableCheats(); err != nil {
		return nil, fmt.Errorf("failed to enable cheats: %w", err)
	}
	if _, err := h.CreateSelectFork(
		script.ForkWithURLOrAlias("main"),
		script.ForkWithBlockNumberU256(latest.Number),
	); err != nil {
		return nil, fmt.Errorf("failed to select fork: %w", err)
	}
	return h, nil
}

func ForkedScriptHost(
	bcaster broadcaster.Broadcaster,
	lgr log.Logger,
	deployer common.Address,
	artifacts foundry.StatDirFs,
	forkRPC *rpc.Client,
	blockNumber *big.Int,
	additionalOpts ...script.HostOption,
) (*script.Host, error) {
	opts := append([]script.HostOption{
		script.WithForkHook(func(cfg *script.ForkConfig) (forking.ForkSource, error) {
			src, err := forking.RPCSourceByNumber(cfg.URLOrAlias, forkRPC, *cfg.BlockNumber)
			if err != nil {
				return nil, fmt.Errorf("failed to create RPC fork source: %w", err)
			}
			return forking.Cache(src), nil
		}),
	}, additionalOpts...)
	h, err := DefaultScriptHost(
		bcaster,
		lgr,
		deployer,
		artifacts,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create default script host: %w", err)
	}

	if _, err := h.CreateSelectFork(
		script.ForkWithURLOrAlias("main"),
		script.ForkWithBlockNumberU256(blockNumber),
	); err != nil {
		return nil, fmt.Errorf("failed to select fork: %w", err)
	}

	return h, nil
}
