package batcher

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/espresso/bindings"
)

// isBatcherActive checks if the current batcher is the active one by querying
// the BatchAuthenticator contract. Returns true if this batcher instance should
// be publishing batches, false if it should stay idle.
//
// The active batcher is determined by the contract's activeIsEspresso flag:
//   - If activeIsEspresso is true, the Espresso batcher address is active
//   - If activeIsEspresso is false, the fallback batcher address is active
//
// This method compares the batcher's own address (from TxMgr) against the
// contract's registered Espresso batcher address and the SystemConfig batcher address.
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

	activeIsEspresso, err := batchAuthenticator.ActiveIsEspresso(callOpts)
	if err != nil {
		return false, fmt.Errorf("failed to check activeIsEspresso: %w", err)
	}

	batcherAddr := l.Txmgr.From()

	isActive := (activeIsEspresso && l.Config.UseEspresso) ||
		(!activeIsEspresso && !l.Config.UseEspresso)

	if !isActive {
		l.Log.Info("Batcher is not the active batcher, skipping publish",
			"batcherAddr", batcherAddr,
			"activeIsEspresso", activeIsEspresso,
			"UseEspresso", l.Config.UseEspresso,
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

// isEspressoEnforcementActive returns true when the EspressoEnforcement hardfork
// is active for the current L1 tip time.
func (l *BatchSubmitter) isEspressoEnforcementActive(ctx context.Context) (bool, error) {
	tip, err := l.l1Tip(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to fetch L1 tip for EspressoEnforcement gate: %w", err)
	}
	return l.RollupConfig.IsEspressoEnforcement(tip.Time), nil
}
