package derive

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	altda "github.com/ethereum-optimism/optimism/op-alt-da"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type DataIter interface {
	Next(ctx context.Context) (eth.Data, error)
}

type L1TransactionFetcher interface {
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	FetchReceipts(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Receipts, error)
}

type L1BlobsFetcher interface {
	// GetBlobsByHash fetches blobs that were confirmed at the given timestamp with the given versioned hashes.
	GetBlobsByHash(ctx context.Context, time uint64, hashes []common.Hash) ([]*eth.Blob, error)
}

type AltDAInputFetcher interface {
	// GetInput fetches the input for the given commitment at the given block number from the DA storage service.
	GetInput(ctx context.Context, l1 altda.L1Fetcher, c altda.CommitmentData, blockId eth.L1BlockRef) (eth.Data, error)
	// AdvanceL1Origin advances the L1 origin to the given block number, syncing the DA challenge events.
	AdvanceL1Origin(ctx context.Context, l1 altda.L1Fetcher, blockId eth.BlockID) error
	// Reset the challenge origin in case of L1 reorg
	Reset(ctx context.Context, base eth.L1BlockRef, baseCfg eth.SystemConfig) error
}

// DataSourceFactory reads raw transactions from a given block & then filters for
// batch submitter transactions.
// This is not a stage in the pipeline, but a wrapper for another stage in the pipeline
type DataSourceFactory struct {
	log          log.Logger
	dsCfg        DataSourceConfig
	fetcher      L1Fetcher
	blobsFetcher L1BlobsFetcher
	altDAFetcher AltDAInputFetcher
	ecotoneTime  *uint64
}

func NewDataSourceFactory(log log.Logger, cfg *rollup.Config, fetcher L1Fetcher, blobsFetcher L1BlobsFetcher, altDAFetcher AltDAInputFetcher) *DataSourceFactory {
	config := DataSourceConfig{
		l1Signer:                  cfg.L1Signer(),
		batchInboxAddress:         cfg.BatchInboxAddress,
		altDAEnabled:              cfg.AltDAEnabled(),
		batchAuthenticatorAddress: cfg.BatchAuthenticatorAddress,
	}
	return &DataSourceFactory{
		log:          log,
		dsCfg:        config,
		fetcher:      fetcher,
		blobsFetcher: blobsFetcher,
		altDAFetcher: altDAFetcher,
		ecotoneTime:  cfg.EcotoneTime,
	}
}

// OpenData returns the appropriate data source for the L1 block `ref`.
func (ds *DataSourceFactory) OpenData(ctx context.Context, ref eth.L1BlockRef, batcherAddr common.Address) (DataIter, error) {
	// Creates a data iterator from blob or calldata source so we can forward it to the altDA source
	// if enabled as it still requires an L1 data source for fetching input commmitments.
	var src DataIter
	if ds.ecotoneTime != nil && ref.Time >= *ds.ecotoneTime {
		if ds.blobsFetcher == nil {
			return nil, fmt.Errorf("ecotone upgrade active but beacon endpoint not configured")
		}
		src = NewBlobDataSource(ctx, ds.log, ds.dsCfg, ds.fetcher, ds.blobsFetcher, ref, batcherAddr)
	} else {
		src = NewCalldataSource(ctx, ds.log, ds.dsCfg, ds.fetcher, ref, batcherAddr)
	}
	if ds.dsCfg.altDAEnabled {
		// altDA([calldata | blobdata](l1Ref)) -> data
		return NewAltDADataSource(ds.log, src, ds.fetcher, ds.altDAFetcher, ref), nil
	}
	return src, nil
}

// DataSourceConfig regroups the mandatory rollup.Config fields needed for DataFromEVMTransactions.
type DataSourceConfig struct {
	l1Signer          types.Signer
	batchInboxAddress common.Address
	altDAEnabled      bool
	// batchAuthenticatorAddress is the L1 address of the BatchAuthenticator contract.
	// When non-zero, event-based batch authentication is used instead of sender verification.
	// When zero, legacy sender-based authentication is used.
	batchAuthenticatorAddress common.Address
}

// BatchAuthEnabled returns true if event-based batch authentication is configured.
func (c DataSourceConfig) BatchAuthEnabled() bool {
	return c.batchAuthenticatorAddress != (common.Address{})
}

// isValidBatchTx checks basic transaction validity for batch submission:
//  1. the transaction type is any of Legacy, ACL, DynamicFee, Blob, or Deposit (for L3s).
//  2. the transaction has a To() address that matches the batch inbox address
//
// It does NOT check authentication (sender or event-based) — that is handled separately
// by the caller based on whether batch authenticator is configured.
func isValidBatchTx(tx *types.Transaction, batchInboxAddr common.Address, logger log.Logger) bool {
	// For now, we want to disallow the SetCodeTx type or any future types.
	if tx.Type() > types.BlobTxType && tx.Type() != types.DepositTxType {
		return false
	}

	to := tx.To()
	if to == nil || *to != batchInboxAddr {
		return false
	}

	return true
}

// isAuthorizedBatchSender checks that the transaction sender matches the expected batcher address.
// Used in legacy mode when batch authenticator is not configured.
func isAuthorizedBatchSender(tx *types.Transaction, l1Signer types.Signer, batcherAddr common.Address, logger log.Logger) bool {
	sender, err := l1Signer.Sender(tx)
	if err != nil {
		logger.Warn("tx in inbox with invalid signature", "hash", tx.Hash(), "err", err)
		return false
	}
	if sender != batcherAddr {
		logger.Warn("tx in inbox with unauthorized submitter", "addr", sender, "hash", tx.Hash())
		return false
	}
	return true
}

// isBatchTxAuthorized checks whether a batch transaction is authorized, using either
// event-based authentication (when authenticatedHashes is non-nil) or legacy sender
// verification. For event-based auth, batchHash must be the precomputed hash of the
// batch content (calldata or blob hashes). The authenticatedHashes set is obtained
// once per L1 block via CollectAuthenticatedBatches.
//
// When batch auth is enabled, there are two authorization paths:
//  1. TEE batcher: must have a matching BatchInfoAuthenticated event (event-based auth)
//  2. Fallback batcher: authorized via sender verification against batcherAddr, which is
//     the standard OP stack batcher address from SystemConfig.batcherHash. This allows
//     the fallback batcher address to be changed dynamically via SystemConfig.setBatcherHash().
//
// This dual-mode approach allows the fallback (non-TEE) batcher to post batches without
// calling authenticateBatchInfo on L1, while still requiring the TEE batcher to authenticate
// its batches via on-chain events.
func isBatchTxAuthorized(
	tx *types.Transaction,
	dsCfg DataSourceConfig,
	batcherAddr common.Address,
	batchHash common.Hash,
	authenticatedHashes map[common.Hash]bool,
	logger log.Logger,
) bool {
	if dsCfg.BatchAuthEnabled() {
		// Event-based authentication: TEE batcher must have an auth event
		if authenticatedHashes[batchHash] {
			return true
		}
		// Fallback batcher: accept via sender verification against the SystemConfig batcher address.
		// This is the same address used by the standard OP stack batcher, allowing it to be
		// changed dynamically via SystemConfig.setBatcherHash().
		if isAuthorizedBatchSender(tx, dsCfg.l1Signer, batcherAddr, logger) {
			return true
		}
		logger.Warn("batch not authenticated via event or fallback sender",
			"txHash", tx.Hash(), "batchHash", batchHash)
		return false
	}
	// Non-espresso mode: verify sender
	return isAuthorizedBatchSender(tx, dsCfg.l1Signer, batcherAddr, logger)
}
