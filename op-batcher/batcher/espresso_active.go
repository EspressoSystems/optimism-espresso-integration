package batcher

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
)

// isBatcherActive checks if the current batcher is the active one by querying
// the BatchAuthenticator contract. Returns true if this batcher instance should
// be publishing batches, false if it should stay idle.
//
// The active batcher is determined by the contract's activeIsTee flag:
//   - If activeIsTee is true, the TEE batcher address is active
//   - If activeIsTee is false, the non-TEE (fallback) batcher address is active
//
// This method compares the batcher's own address (from TxMgr) against the
// contract's registered TEE and non-TEE batcher addresses.
func (l *BatchSubmitter) isBatcherActive(ctx context.Context) (bool, error) {
	// Check if contract code exists at the address
	code, err := l.L1Client.CodeAt(ctx, l.RollupConfig.BatchAuthenticatorAddress, nil)
	if err != nil {
		return false, fmt.Errorf("failed to check code at BatchAuthenticator address: %w", err)
	}
	if len(code) == 0 {
		return false, fmt.Errorf("no contract code at BatchAuthenticator address %s", l.RollupConfig.BatchAuthenticatorAddress.Hex())
	}

	batchAuthenticator, err := bindings.NewBatchAuthenticator(l.RollupConfig.BatchAuthenticatorAddress, l.L1Client)
	if err != nil {
		return false, fmt.Errorf("failed to create BatchAuthenticator binding: %w", err)
	}

	cCtx, cancel := context.WithTimeout(ctx, l.Config.NetworkTimeout)
	defer cancel()

	callOpts := &bind.CallOpts{Context: cCtx}

	activeIsTee, err := batchAuthenticator.ActiveIsTee(callOpts)
	if err != nil {
		return false, fmt.Errorf("failed to check activeIsTee: %w", err)
	}

	teeBatcherAddr, err := batchAuthenticator.TeeBatcher(callOpts)
	if err != nil {
		return false, fmt.Errorf("failed to get TEE batcher address: %w", err)
	}

	nonTeeBatcherAddr, err := batchAuthenticator.NonTeeBatcher(callOpts)
	if err != nil {
		return false, fmt.Errorf("failed to get non-TEE batcher address: %w", err)
	}

	batcherAddr := l.Txmgr.From()

	isActive := (activeIsTee && batcherAddr == teeBatcherAddr) ||
		(!activeIsTee && batcherAddr == nonTeeBatcherAddr)

	if !isActive {
		l.Log.Info("Batcher is not the active batcher, skipping publish",
			"batcherAddr", batcherAddr,
			"activeIsTee", activeIsTee,
			"teeBatcher", teeBatcherAddr,
			"nonTeeBatcher", nonTeeBatcherAddr,
		)
	}

	return isActive, nil
}

// hasBatchAuthenticator returns true if the rollup config has a non-zero
// BatchAuthenticatorAddress, indicating that batcher active/idle checking
// should be performed before publishing.
func (l *BatchSubmitter) hasBatchAuthenticator() bool {
	return l.RollupConfig.BatchAuthenticatorAddress != (common.Address{})
}
