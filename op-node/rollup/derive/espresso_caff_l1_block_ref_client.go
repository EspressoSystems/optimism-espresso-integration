package derive

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// L1BlockRefClient is a wrapper around eth.L1BlockRef that implements the espresso.L1Client interface
type L1BlockRefClient struct {
	L1Fetcher L1Fetcher
}

// NewL1BlockRefClient creates a new L1BlockRefClient
func NewL1BlockRefClient(L1Fetcher L1Fetcher) *L1BlockRefClient {
	return &L1BlockRefClient{
		L1Fetcher: L1Fetcher,
	}
}

// HeaderHashByNumber implements the espresso.L1Client interface
func (c *L1BlockRefClient) HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error) {
	expectedL1BlockRef, err := c.L1Fetcher.L1BlockRefByNumber(ctx, number.Uint64())
	if err != nil {
		return common.Hash{}, err
	}

	return expectedL1BlockRef.Hash, nil
}
