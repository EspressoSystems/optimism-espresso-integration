package batcher_test

import (
	"context"
	"encoding/json"
	"errors"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	"github.com/EspressoSystems/espresso-network/sdks/go/types"
	common "github.com/EspressoSystems/espresso-network/sdks/go/types/common"
)

// ErrNotImplemented is a sentinel error used to indicate that a method in the
// was not implement.
var ErrNotImplemented = errors.New("not implemented")

// AlwaysFailingEspressoClient is a mock implementation of the EspressoClient
// interface that always returns an error for every method call. This is
// useful for testing error handling in the batcher without relying on a
// real Espresso client.
type AlwaysFailingEspressoClient struct{}

// Compile time check to ensure adherence to EspressoClient interface
var _ espressoClient.EspressoClient = (*AlwaysFailingEspressoClient)(nil)

func (*AlwaysFailingEspressoClient) FetchLatestBlockHeight(ctx context.Context) (uint64, error) {
	return 0, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchHeaderByHeight(ctx context.Context, height uint64) (types.HeaderImpl, error) {
	return types.HeaderImpl{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchRawHeaderByHeight(ctx context.Context, height uint64) (json.RawMessage, error) {
	return json.RawMessage{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchHeadersByRange(ctx context.Context, from uint64, until uint64) ([]types.HeaderImpl, error) {
	return []types.HeaderImpl{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (espressoClient.TransactionsInBlock, error) {
	return espressoClient.TransactionsInBlock{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchTransactionByHash(ctx context.Context, hash *types.TaggedBase64) (types.TransactionQueryData, error) {
	return types.TransactionQueryData{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchVidCommonByHeight(ctx context.Context, blockHeight uint64) (types.VidCommon, error) {
	return types.VidCommon{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchExplorerTransactionByHash(ctx context.Context, hash *types.TaggedBase64) (types.ExplorerTransactionQueryData, error) {
	return types.ExplorerTransactionQueryData{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) StreamPayloads(ctx context.Context, height uint64) (espressoClient.Stream[types.PayloadQueryData], error) {
	return nil, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) StreamTransactions(ctx context.Context, height uint64) (espressoClient.Stream[types.TransactionQueryData], error) {
	return nil, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) StreamTransactionsInNamespace(ctx context.Context, height uint64, namespace uint64) (espressoClient.Stream[types.TransactionQueryData], error) {
	return nil, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) FetchNamespaceTransactionsInRange(ctx context.Context, fromHeight uint64, toHeight uint64, namespace uint64) ([]types.NamespaceTransactionsRangeData, error) {
	return []types.NamespaceTransactionsRangeData{}, ErrNotImplemented
}

func (*AlwaysFailingEspressoClient) SubmitTransaction(ctx context.Context, tx common.Transaction) (*common.TaggedBase64, error) {
	return nil, ErrNotImplemented
}
