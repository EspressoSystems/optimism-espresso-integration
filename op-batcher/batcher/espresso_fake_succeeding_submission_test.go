package batcher_test

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"sync"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	tagged_base64 "github.com/EspressoSystems/espresso-network/sdks/go/tagged-base64"
	"github.com/EspressoSystems/espresso-network/sdks/go/types"
	common "github.com/EspressoSystems/espresso-network/sdks/go/types/common"
)

// FakeSubmissionSucceedingEspressoClient is a mock implementation of the
// EspressoClient for the explicit purposes of submitting transactions, and
// seeing their response.
type FakeSubmissionSucceedingEspressoClient struct {
	sync.RWMutex
	espressoClient.EspressoClient
	txns map[string]common.Transaction
}

// Compile time check to ensure adherence to EspressoClient interface
var _ espressoClient.EspressoClient = (*FakeSubmissionSucceedingEspressoClient)(nil)

var (
	ErrNotInitialized      = errors.New("not initialized")
	ErrHashCannotBeNil     = errors.New("hash cannot be nil")
	ErrTransactionNotFound = errors.New("transaction not found")
)

func (c *FakeSubmissionSucceedingEspressoClient) Init() {
	c.Lock()
	defer c.Unlock()
	c.txns = make(map[string]common.Transaction)
}

// FetchTransactionByHash simulates fetching a transaction by its hash. it
// looks up a transaction for the given hash, and returns the transaction
// if it is found.
func (c *FakeSubmissionSucceedingEspressoClient) FetchTransactionByHash(ctx context.Context, hash *types.TaggedBase64) (types.TransactionQueryData, error) {
	c.RLock()
	defer c.RUnlock()
	if c.txns == nil {
		return types.TransactionQueryData{}, ErrNotInitialized
	}

	if hash == nil {
		return types.TransactionQueryData{}, ErrHashCannotBeNil
	}

	txn, found := c.txns[hash.String()]
	if !found {
		return types.TransactionQueryData{}, ErrTransactionNotFound
	}

	// Just to simulate some processing on the transaction
	height := binary.LittleEndian.Uint64(txn.Payload)

	return types.TransactionQueryData{
		Transaction: txn,
		Hash:        hash,
		Index:       0,
		Proof:       json.RawMessage{},
		BlockHash:   nil,
		BlockHeight: height,
	}, nil
}

// SubmitTransaction simulates a successful transaction submission and stores
// it for future retrieval
func (c *FakeSubmissionSucceedingEspressoClient) SubmitTransaction(ctx context.Context, tx common.Transaction) (*common.TaggedBase64, error) {
	c.Lock()
	defer c.Unlock()
	if c.txns == nil {
		return nil, ErrNotInitialized
	}

	hash, err := tagged_base64.New("TX", tx.Payload)
	if err != nil {
		return nil, err
	}

	c.txns[hash.String()] = tx
	return hash, nil
}
