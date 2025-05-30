package benchmark

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// WorkerSignTransaction is a function that is meant to be run as a goroutine.
// It will continually sign transactions with increasing nonces with the given
// details, until the given context is done.
func WorkerSignTransaction(
	ctx context.Context,
	wg *sync.WaitGroup,
	interval time.Duration,
	txChan chan<- TimestampedValue[*geth_types.Transaction],
	key *ecdsa.PrivateKey,
	signer geth_types.Signer,
	chainID *big.Int,
	to *common.Address,
	value *big.Int,
) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer wg.Done()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for nonce := uint64(0); true; nonce++ {
		// Check to see if we should stop
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			// Time to submit a new transaction.
		}

		tx := geth_types.MustSignNewTx(key, signer, &geth_types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			To:        to,
			Value:     big.NewInt(1),
			GasTipCap: big.NewInt(10),
			GasFeeCap: big.NewInt(200),
			Gas:       21_000,
		})

		// Submit the Signed Transaction to the given channel
		select {
		case txChan <- WithTimestamp(tx):
		case <-ctx.Done():
			return
		}
	}
}

// WorkerSubmitSignedTransaction is a function that is meant to be run as a
// goroutine.  It will continually request new signed transactions to submit,
// and will record the submission time of the transaction.
func WorkerSubmitSignedTransaction(
	ctx context.Context,
	wg *sync.WaitGroup,
	signedTxChannel <-chan TimestampedValue[*geth_types.Transaction],
	txSubmittedChan chan<- TimestampedValue[common.Hash],
	client *ethclient.Client,
) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer wg.Done()

	for {
		var tx TimestampedValue[*geth_types.Transaction]
		var ok bool

		// Wait for work
		select {
		case <-ctx.Done():
			return

		case tx, ok = <-signedTxChannel:
			if !ok {
				// The channel was closed, so we have nothing more
				// to process.
				return
			}
		}

		// Submit the Signed Transaction to the given channel
		err := client.SendTransaction(ctx, tx.Value)
		if err != nil {
			continue
		}

		select {
		case txSubmittedChan <- WithTimestamp(tx.Value.Hash()):
		case <-ctx.Done():
			return
		}
	}
}

// WorkerProcessL2Receipt is a function that is meant to be run as a
// goroutine. When a submitted transaction is received, it will wait for the
// receipt of the transaction to be available, and then send the receipt to
// the given channel. It will also record the time at which the receipt was
// received.
func WorkerProcessL2Receipt(
	ctx context.Context,
	wg *sync.WaitGroup,
	txSubmittedChan <-chan TimestampedValue[common.Hash],
	receiptChan chan<- TimestampedValue[*geth_types.Receipt],
	client *ethclient.Client,
) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer wg.Done()

	for {
		var txHash TimestampedValue[common.Hash]
		var ok bool

		// Wait for work
		select {
		case <-ctx.Done():
			return

		case txHash, ok = <-txSubmittedChan:
			if !ok {
				// The channel was closed, so we have nothing more
				// to process.
				return
			}
		}

		receipt, err := wait.ForReceiptOK(ctx, client, txHash.Value)
		if err != nil {
			log.Error("Failed to get receipt", "err", err)
			continue
		}

		select {
		case receiptChan <- WithTimestamp(receipt):
		case <-ctx.Done():
			return
		}
	}
}

// AnnotatedBlockHash is a struct that holds a block hash and a label.
// It is used to annotate block hashes with a label for easier
// identification in logs and metrics.
type AnnotatedBlockHash struct {
	Label     string
	BlockHash common.Hash
}

// WorkerConsumeBlockHeaders is a function that is meant to be run as a
// goroutine. It will continually receive block headers from the given
// channel, and will send the block hash to the given channel with a
// timestamp.
func WorkerConsumeBlockHeaders(
	ctx context.Context,
	wg *sync.WaitGroup,
	label string,
	headerChan <-chan *geth_types.Header,
	receivedChain chan<- TimestampedValue[AnnotatedBlockHash],
) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer wg.Done()
	for {
		var header *geth_types.Header
		var ok bool
		select {
		case <-ctx.Done():
			return
		case header, ok = <-headerChan:
			if !ok {
				// The channel was closed, so we have nothing more
				// to process.
				return
			}
		}

		select {
		case receivedChain <- WithTimestamp(AnnotatedBlockHash{
			BlockHash: header.Hash(),
			Label:     label,
		}):
		case <-ctx.Done():
			return
		}
	}
}

// WorkerRecordTimestampedEvents is a function that is meant to be run as a
// goroutine. It will continually receive events from the given channels
// and will record the events in the given stats struct.
func WorkerRecordTimestampedEvents(
	ctx context.Context,
	wg *sync.WaitGroup,
	txCreated <-chan TimestampedValue[*geth_types.Transaction],
	seqSubmissions <-chan TimestampedValue[common.Hash],
	seqReceipts <-chan TimestampedValue[*geth_types.Receipt],
	annotatedBlockHeaders <-chan TimestampedValue[AnnotatedBlockHash],
	stats *BenchmarkStats,
) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case tx, ok := <-txCreated:
			if !ok {
				continue
			}
			stats.Created[tx.Value.Hash()] = tx
		case tx, ok := <-seqSubmissions:
			if !ok {
				continue
			}
			stats.Submitted[tx.Value] = tx
		case receipt, ok := <-seqReceipts:
			if !ok {
				continue
			}
			stats.Receipts[receipt.Value.TxHash] = receipt
		case annotatedBlock, ok := <-annotatedBlockHeaders:
			if !ok {
				continue
			}

			switch annotatedBlock.Value.Label {
			default:
				// Not sure what to do for this case
			case env.RoleCaffNode:
				stats.CaffReceipts[annotatedBlock.Value.BlockHash] = TimestampedValue[common.Hash]{Value: annotatedBlock.Value.BlockHash, Timestamp: annotatedBlock.Timestamp}
			case e2esys.RoleSeq:
				stats.SeqReceipts[annotatedBlock.Value.BlockHash] = TimestampedValue[common.Hash]{Value: annotatedBlock.Value.BlockHash, Timestamp: annotatedBlock.Timestamp}
			case e2esys.RoleVerif:
				stats.VerifyReceipts[annotatedBlock.Value.BlockHash] = TimestampedValue[common.Hash]{Value: annotatedBlock.Value.BlockHash, Timestamp: annotatedBlock.Timestamp}
			}
		}
	}
}
