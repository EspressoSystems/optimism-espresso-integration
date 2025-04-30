package derive

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
)

// L1BlockRefClient is a wrapper around eth.L1BlockRef that implements the espresso.L1Client interface
type L1BlockRefClient struct {
	L1FinalizedBlock   func() (eth.L1BlockRef, error)
	L1BlockRefByNumber func(ctx context.Context, num uint64) (eth.L1BlockRef, error)
}

// NewL1BlockRefClient creates a new L1BlockRefClient
func NewL1BlockRefClient(L1FinalizedBlock func() (eth.L1BlockRef, error), L1BlockRefByNumber func(ctx context.Context, num uint64) (eth.L1BlockRef, error)) *L1BlockRefClient {
	return &L1BlockRefClient{
		L1FinalizedBlock:   L1FinalizedBlock,
		L1BlockRefByNumber: L1BlockRefByNumber,
	}
}

// HeaderHashByNumber implements the espresso.L1Client interface
func (c *L1BlockRefClient) HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error) {
	expectedL1BlockRef, err := c.L1BlockRefByNumber(ctx, number.Uint64())
	if err != nil {
		return common.Hash{}, err
	}

	return expectedL1BlockRef.Hash, nil
}
