//go:build mips64

package derive

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/log"
)

// espressoAttributesQueue is the MIPS64/fault-proof stub.
// Espresso derivation is not available in the fault-proof program.
type espressoAttributesQueue struct {
	isCaffNode           bool
	caffeinationHeightL2 uint64
}

func newEspressoAttributesQueue(_ log.Logger, cfg *rollup.Config) espressoAttributesQueue {
	return espressoAttributesQueue{
		isCaffNode:           cfg.CaffNodeConfig.Enabled,
		caffeinationHeightL2: cfg.CaffNodeConfig.CaffeinationHeightL2,
	}
}

// nextBatch always falls through to the regular L1-based derivation path.
func (e *espressoAttributesQueue) nextBatch(ctx context.Context, parent eth.L2BlockRef, _ uint64, _ L1Fetcher, prev SingularBatchProvider, _ log.Logger) (*SingularBatch, bool, error) {
	return prev.NextBatch(ctx, parent)
}
