package batcher_test

import (
	"context"
	"encoding/json"
	"sync"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	"github.com/EspressoSystems/espresso-network/sdks/go/types"
	common "github.com/EspressoSystems/espresso-network/sdks/go/types/common"
)

// EspressoClientSwappableImplementation is as implementation of EspressoClient
// that is just a proxy.
//
// This allows it to be created and swapped easily as needed for testing.
type EspressoClientSwappableImplementation struct {
	sync.RWMutex
	espClient espressoClient.EspressoClient
}

// Compile time check to ensure adherence to EspressoClient interface
var _ espressoClient.EspressoClient = (*EspressoClientSwappableImplementation)(nil)

func (c *EspressoClientSwappableImplementation) SetEspressoClient(client espressoClient.EspressoClient) {
	c.Lock()
	defer c.Unlock()

	c.espClient = client
}

func (c *EspressoClientSwappableImplementation) FetchLatestBlockHeight(ctx context.Context) (uint64, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchLatestBlockHeight(ctx)
}

func (c *EspressoClientSwappableImplementation) FetchHeaderByHeight(ctx context.Context, height uint64) (types.HeaderImpl, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchHeaderByHeight(ctx, height)
}

func (c *EspressoClientSwappableImplementation) FetchRawHeaderByHeight(ctx context.Context, height uint64) (json.RawMessage, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchRawHeaderByHeight(ctx, height)
}

func (c *EspressoClientSwappableImplementation) FetchHeadersByRange(ctx context.Context, from uint64, until uint64) ([]types.HeaderImpl, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchHeadersByRange(ctx, from, until)
}

func (c *EspressoClientSwappableImplementation) FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (espressoClient.TransactionsInBlock, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchTransactionsInBlock(ctx, blockHeight, namespace)
}

func (c *EspressoClientSwappableImplementation) FetchTransactionByHash(ctx context.Context, hash *types.TaggedBase64) (types.TransactionQueryData, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchTransactionByHash(ctx, hash)
}

func (c *EspressoClientSwappableImplementation) FetchVidCommonByHeight(ctx context.Context, blockHeight uint64) (types.VidCommon, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchVidCommonByHeight(ctx, blockHeight)
}

func (c *EspressoClientSwappableImplementation) FetchExplorerTransactionByHash(ctx context.Context, hash *types.TaggedBase64) (types.ExplorerTransactionQueryData, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchExplorerTransactionByHash(ctx, hash)
}

func (c *EspressoClientSwappableImplementation) StreamPayloads(ctx context.Context, height uint64) (espressoClient.Stream[types.PayloadQueryData], error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.StreamPayloads(ctx, height)
}

func (c *EspressoClientSwappableImplementation) StreamTransactions(ctx context.Context, height uint64) (espressoClient.Stream[types.TransactionQueryData], error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.StreamTransactions(ctx, height)
}

func (c *EspressoClientSwappableImplementation) StreamTransactionsInNamespace(ctx context.Context, height uint64, namespace uint64) (espressoClient.Stream[types.TransactionQueryData], error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.StreamTransactionsInNamespace(ctx, height, namespace)
}

func (c *EspressoClientSwappableImplementation) FetchNamespaceTransactionsInRange(ctx context.Context, fromHeight uint64, toHeight uint64, namespace uint64) ([]types.NamespaceTransactionsRangeData, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.FetchNamespaceTransactionsInRange(ctx, fromHeight, toHeight, namespace)
}

func (c *EspressoClientSwappableImplementation) SubmitTransaction(ctx context.Context, tx common.Transaction) (*common.TaggedBase64, error) {
	c.RLock()
	defer c.RUnlock()

	return c.espClient.SubmitTransaction(ctx, tx)
}
