package derive

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// CalldataSource is a fault tolerant approach to fetching data.
// The constructor will never fail & it will instead re-attempt the fetcher
// at a later point.
type CalldataSource struct {
	// Internal state + data
	open bool
	data []eth.Data
	// Required to re-attempt fetching
	ref     eth.L1BlockRef
	dsCfg   DataSourceConfig
	fetcher L1TransactionFetcher
	log     log.Logger

	batcherAddr common.Address
}

// NewCalldataSource creates a new calldata source. It suppresses errors in fetching the L1 block if they occur.
// If there is an error, it will attempt to fetch the result on the next call to `Next`.
func NewCalldataSource(ctx context.Context, log log.Logger, dsCfg DataSourceConfig, fetcher L1TransactionFetcher, ref eth.L1BlockRef, batcherAddr common.Address) DataIter {
	closedSource := &CalldataSource{
		open:        false,
		ref:         ref,
		dsCfg:       dsCfg,
		fetcher:     fetcher,
		log:         log,
		batcherAddr: batcherAddr,
	}

	_, txs, err := fetcher.InfoAndTxsByHash(ctx, ref.Hash)
	if err != nil {
		return closedSource
	}
	_, receipts, err := fetcher.FetchReceipts(ctx, ref.Hash)
	if err != nil {
		return closedSource
	}
	if len(txs) != len(receipts) {
		return closedSource
	}
	return &CalldataSource{
		open: true,
		data: DataFromEVMTransactions(dsCfg, batcherAddr, txs, receipts, log.New("origin", ref)),
	}
}

// Next returns the next piece of data if it has it. If the constructor failed, this
// will attempt to reinitialize itself. If it cannot find the block it returns a ResetError
// otherwise it returns a temporary error if fetching the block returns an error.
func (ds *CalldataSource) Next(ctx context.Context) (eth.Data, error) {
	if !ds.open {
		_, txs, err := ds.fetcher.InfoAndTxsByHash(ctx, ds.ref.Hash)
		if errors.Is(err, ethereum.NotFound) {
			return nil, NewResetError(fmt.Errorf("failed to open calldata source: %w", err))
		} else if err != nil {
			return nil, NewTemporaryError(fmt.Errorf("failed to open calldata source: %w", err))
		}
		_, receipts, err := ds.fetcher.FetchReceipts(ctx, ds.ref.Hash)
		if errors.Is(err, ethereum.NotFound) {
			return nil, NewResetError(fmt.Errorf("failed to open calldata source: %w", err))
		} else if err != nil {
			return nil, NewTemporaryError(fmt.Errorf("failed to open calldata source: %w", err))
		}
		if len(txs) != len(receipts) {
			return nil, NewTemporaryError(fmt.Errorf("failed to open calldata source: L1 fetcher provided inconsistent number of transactions and receipts"))
		}
		ds.open = true
		ds.data = DataFromEVMTransactions(ds.dsCfg, ds.batcherAddr, txs, receipts, ds.log)
	}
	if len(ds.data) == 0 {
		return nil, io.EOF
	} else {
		data := ds.data[0]
		ds.data = ds.data[1:]
		return data, nil
	}
}

// DataFromEVMTransactions filters all of the transactions and returns the calldata from transactions
// that are sent to the batch inbox address from the batch sender address.
// This will return an empty array if no valid transactions are found.
func DataFromEVMTransactions(dsCfg DataSourceConfig, batcherAddr common.Address, txs types.Transactions, receipts types.Receipts, log log.Logger) []eth.Data {
	out := []eth.Data{}
	for i, tx := range txs {
		if isValidBatchTx(tx, receipts[i], dsCfg.l1Signer, dsCfg.batchInboxAddress, batcherAddr, log) {
			out = append(out, tx.Data())
		}
	}
	return out
}
