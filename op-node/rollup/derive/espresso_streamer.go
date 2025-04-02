package derive

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"sync"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network-go/client"
	espressoTypes "github.com/EspressoSystems/espresso-network-go/types"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

type EspressoClientInterface interface {
	FetchLatestBlockHeight(ctx context.Context) (uint64, error)
	FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (espressoClient.TransactionsInBlock, error)
}

type MessageWithHeight struct {
	batch         *SingularBatch
	HotShotHeight uint64
}

type EspressoStreamer struct {
	espressoClient                EspressoClientInterface
	nextHotShotBlockNum           uint64
	currentMessagePos             uint64
	namespace                     uint64
	pollingHotShotPollingInterval time.Duration
	messagesWithHeights           []*MessageWithHeight
	log                           log.Logger
	batchInboxAddr                common.Address
	rollupConfig                  *rollup.Config
	messageMutex                  sync.Mutex
}

func NewEspressoStreamer(namespace uint64,
	nextHotShotBlockNum uint64,
	pollingHotShotPollingInterval time.Duration,
	espressoClientInterface EspressoClientInterface,
	log log.Logger,
	batchInboxAddr common.Address,
	rollupConfig *rollup.Config,
) *EspressoStreamer {

	return &EspressoStreamer{
		espressoClient:                espressoClientInterface,
		nextHotShotBlockNum:           nextHotShotBlockNum,
		pollingHotShotPollingInterval: pollingHotShotPollingInterval,
		namespace:                     namespace,
		log:                           log,
		batchInboxAddr:                batchInboxAddr,
		rollupConfig:                  rollupConfig,
	}
}

func (s *EspressoStreamer) Reset(currentMessagePos uint64, currentHostshotBlock uint64) {
	s.messageMutex.Lock()
	defer s.messageMutex.Unlock()
	s.currentMessagePos = currentMessagePos
	s.nextHotShotBlockNum = currentHostshotBlock
	s.messagesWithHeights = []*MessageWithHeight{}
}

func CheckBatchEspresso(ctx context.Context, cfg *rollup.Config, log log.Logger, l2SafeHead eth.L2BlockRef, batch *SingularBatch) BatchValidity {
	// add details to the log
	log = batch.LogContext(log)

	// Sishan TODO: check the L1 origin is already finalized

	// Sishan TODO: these checks are copy-pasted from OP's checkSingularBatch(), we should check whether these apply to caff node
	nextTimestamp := l2SafeHead.Time + cfg.BlockTime
	if batch.Timestamp > nextTimestamp {
		log.Trace("received out-of-order batch for future processing after next batch", "next_timestamp", nextTimestamp)
		return BatchFuture
	}
	if batch.Timestamp < nextTimestamp {
		log.Warn("dropping past batch with old timestamp", "min_timestamp", nextTimestamp)
		return BatchDrop
	}

	// dependent on above timestamp check. If the timestamp is correct, then it must build on top of the safe head.
	if batch.ParentHash != l2SafeHead.Hash {
		log.Warn("ignoring batch with mismatching parent hash", "current_safe_head", l2SafeHead.Hash)
		return BatchDrop
	}

	// We can do this check earlier, but it's a more intensive one, so we do this last.
	for i, txBytes := range batch.Transactions {
		if len(txBytes) == 0 {
			log.Warn("transaction data must not be empty, but found empty tx", "tx_index", i)
			return BatchDrop
		}
		if txBytes[0] == types.DepositTxType {
			log.Warn("sequencers may not embed any deposits into batch data, but found tx that has one", "tx_index", i)
			return BatchDrop
		}
	}

	return BatchAccept
}

func (s *EspressoStreamer) NextBatch(ctx context.Context, parent eth.L2BlockRef) (*SingularBatch, bool, error) {
	s.messageMutex.Lock()
	defer s.messageMutex.Unlock()

	// Sishan TODO: Find the batch that match the parent block, concluding is assignedto false for now
	var returnBatch *SingularBatch
	var remaining []*MessageWithHeight
batchLoop:
	for i, message := range s.messagesWithHeights {
		validity := CheckBatchEspresso(ctx, s.rollupConfig, s.log.New("batch_index", i), parent, message.batch)
		// sort out the next batch and drop batch in existing batches
		switch validity {
		case BatchFuture:
			remaining = append(remaining, message)
			continue
		case BatchDrop:
			message.batch.LogContext(s.log).Warn("Dropping batch",
				"parent", parent.ID(),
				"parent_time", parent.Time,
			)
			continue
		case BatchAccept:
			returnBatch = message.batch
			// don't keep the current batch in the remaining items since we are processing it now,
			// but retain every batch we didn't get to yet.
			remaining = append(remaining, s.messagesWithHeights[i+1:]...)
			break batchLoop
		case BatchUndecided: // Sishan TODO: remove if this is not needed
			remaining = append(remaining, s.messagesWithHeights[i:]...)
			s.messagesWithHeights = remaining
			return nil, false, io.EOF
		default:
			return nil, false, NewCriticalError(fmt.Errorf("unknown batch validity type: %d", validity))
		}
	}
	s.messagesWithHeights = remaining
	s.log.Info("NextBatch", "returnBatch", returnBatch)
	return returnBatch, false, nil
}

func ParseHotShotPayload(log log.Logger, payload []byte) (batcherSignature []byte, batchByte []byte, err error) {

	log.Info("Parsing hotshot payload", "payload", hex.EncodeToString(payload))

	batcherSignature, batchByte = payload[:ethCrypto.SignatureLength], payload[ethCrypto.SignatureLength:]

	log.Info("Parsed Batcher signature", "signature", hex.EncodeToString(batcherSignature), "espressoBatch byte", hex.EncodeToString(batchByte))
	return batcherSignature, batchByte, nil
}

type EspressoBatch struct {
	Header types.Header
	Batch  SingularBatch
}

func (s *EspressoStreamer) parseEspressoTransaction(tx espressoTypes.Bytes) ([]*MessageWithHeight, error) {
	s.log.Info("Parsing espresso transaction", "tx", hex.EncodeToString(tx))
	batcherSignature, batchByte, err := ParseHotShotPayload(s.log, tx)
	if err != nil {
		s.log.Warn("failed to parse hotshot payload", "err", err)
		return nil, err
	}
	// if batcher'ssignature verification fails, we should skip this message

	batchHash := ethCrypto.Keccak256(batchByte)
	err = crypto.Verify(batchHash, batcherSignature, s.batchInboxAddr)
	if err != nil {
		s.log.Warn("failed to verify signature", "err", err)
		// Sishan TODO: return an error instead of continuing
	}

	var batch EspressoBatch
	if err := rlp.DecodeBytes(batchByte, &batch); err != nil {
		return nil, err
	}

	s.log.Info("Parsed espresso batch", "batch", batch)
	result := &MessageWithHeight{
		batch:         &batch.Batch,
		HotShotHeight: s.nextHotShotBlockNum,
	}

	return []*MessageWithHeight{result}, nil
}

/*
*
* Create a queue of messages from the hotshot to be processed by the node
* It will sort the messages by the message index
* and store the messages in `messagesWithMetadata` queue
*
* Expose the *parseHotShotPayloadFn* to the caller for testing purposes
 */
func (s *EspressoStreamer) QueueMessagesFromHotShot(
	ctx context.Context,
	parseHotShotPayloadFn func(tx espressoTypes.Bytes) ([]*MessageWithHeight, error),
) error {
	// Note: Adding the lock on top level
	// because s.nextHotShotBlockNum is updated if n.nextHotShotBlockNum == 0
	s.messageMutex.Lock()
	defer s.messageMutex.Unlock()

	if s.nextHotShotBlockNum == 0 {
		// We dont need to check majority here  because when we eventually go
		// to fetch a block at a certain height,
		// we will check that a quorum of nodes agree on the block at that height,
		// which wouldn't be possible if we were somehow are given a height
		// that wasn't finalized at all
		latestBlock, err := s.espressoClient.FetchLatestBlockHeight(ctx)
		if err != nil {
			s.log.Warn("unable to fetch latest hotshot block", "err", err)
			return err
		}
		s.log.Info("Started node at the latest hotshot block", "block number", latestBlock)
		s.nextHotShotBlockNum = latestBlock
	}

	txns, err := s.espressoClient.FetchTransactionsInBlock(ctx, s.nextHotShotBlockNum, s.namespace)
	if err != nil {
		s.log.Warn("failed to fetch the transactions", "err", err)
		return err
	}

	if len(txns.Transactions) == 0 {
		s.log.Info("No transactions found in the hotshot block", "block number", s.nextHotShotBlockNum)
		s.nextHotShotBlockNum += 1
		return nil
	}

	for _, tx := range txns.Transactions {
		s.log.Info("Parsing espresso transaction", "tx", hex.EncodeToString(tx))
		messages, err := parseHotShotPayloadFn(tx)
		if err != nil {
			s.log.Warn("failed to verify espresso transaction", "err", err)
			continue
		}
		// Sishan TODO: Filter out the messages have already been seen
		s.messagesWithHeights = append(s.messagesWithHeights, messages...)
	}

	s.nextHotShotBlockNum += 1

	return nil
}

func (s *EspressoStreamer) Start(ctx context.Context) error {

	s.log.Info("In the function, Starting espresso streamer")
	bigTimeout := 2 * time.Minute
	timer := time.NewTimer(bigTimeout)
	defer timer.Stop()

	// Sishan TODO: maybe use better handler with dynamic interval in the future
	ticker := time.NewTicker(s.pollingHotShotPollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := s.QueueMessagesFromHotShot(ctx, s.parseEspressoTransaction)
			if err != nil {
				s.log.Error("error while queueing messages", "err", err)
			} else {
				s.log.Info("Processing block", "block number", s.nextHotShotBlockNum)
				// Successful execution: reset the timer to start the timeout period over.
				// Stop the timer and drain if needed.
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(bigTimeout)
			}
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			return fmt.Errorf("timeout while queueing messages from hotshot")
		}
	}

}
