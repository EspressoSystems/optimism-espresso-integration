package espresso

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum-optimism/optimism/espresso/bindings"
)

// FetchEspressoBatcherAddress reads the Espresso batcher address from the BatchAuthenticator
// contract on L1. This is used by the caff node to determine which address signed
// Espresso batches, since the Espresso batcher may use a different key than the
// SystemConfig batcher (fallback batcher).
func FetchEspressoBatcherAddress(ctx context.Context, l1Client *ethclient.Client, batchAuthenticatorAddr common.Address) (common.Address, error) {
	caller, err := bindings.NewBatchAuthenticatorCaller(batchAuthenticatorAddr, l1Client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to bind BatchAuthenticator at %s: %w", batchAuthenticatorAddr, err)
	}
	addr, err := caller.EspressoBatcher(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call BatchAuthenticator.espressoBatcher(): %w", err)
	}
	return addr, nil
}
