package espressostreamer

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-sequencer-go/client"
	espressoTypes "github.com/EspressoSystems/espresso-sequencer-go/types"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/log"
)

type Transaction struct {
	// Namespace of transaction to be published
	Namespace uint64
	// TODO: placeholder for sequencer's signature
	SequencerSignature []byte
	// Frames serialized as they would be for posting to L1 as calldata
	CallData []byte
}

type EspressoClientInterface interface {
	FetchLatestBlockHeight(ctx context.Context) (uint64, error)
	FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (espressoClient.TransactionsInBlock, error)
}

type MessageWithHeight struct {
	SequencerBatches *derive.SingularBatch
	HotshotHeight    uint64
}

type EspressoStreamer struct {
	espressoClient                EspressoClientInterface
	nextHotshotBlockNum           uint64
	currentMessagePos             uint64
	namespace                     uint64
	pollingHotshotPollingInterval time.Duration
	messageWithHeight             []*MessageWithHeight
	log                           log.Logger

	messageMutex sync.Mutex
}

func NewEspressoStreamer(namespace uint64,
	nextHotshotBlockNum uint64,
	pollingHotshotPollingInterval time.Duration,
	espressoClientInterface EspressoClientInterface,
	log log.Logger,
) *EspressoStreamer {

	return &EspressoStreamer{
		espressoClient:                espressoClientInterface,
		nextHotshotBlockNum:           nextHotshotBlockNum,
		pollingHotshotPollingInterval: pollingHotshotPollingInterval,
		namespace:                     namespace,
		log:                           log,
	}
}

func (s *EspressoStreamer) Reset(currentMessagePos uint64, currentHostshotBlock uint64) {
	s.messageMutex.Lock()
	defer s.messageMutex.Unlock()
	s.currentMessagePos = currentMessagePos
	s.nextHotshotBlockNum = currentHostshotBlock
	s.messageWithHeight = []*MessageWithHeight{}
}

func (s *EspressoStreamer) NextBatch(parent eth.L2BlockRef) (MessageWithHeight, error) {
	s.messageMutex.Lock()
	defer s.messageMutex.Unlock()

	// Sishan TODO: Find the batch that match the parent block
	message := s.messageWithHeight[s.currentMessagePos]
	s.currentMessagePos += 1
	return *message, nil
}

// Sishan TODO: not sure whether we'll also parse namespace here, maybe there is no namespace in the input payload
func ParseHotShotPayload(payload []byte) (batcherSignature []byte, sequencerBatchesByte []byte, err error) {

	// Sishan TODO: do real parse, blocked by batcher submitter changes.
	// now the payload is append(batcherSignature, txdata.CallData()...),
	// what we need will be append(batcherSignature,sequencerBatches...)

	// placeholder
	batcherSignature = []byte{1, 2, 3, 4}
	sequencerBatchesByte = []byte{5, 6, 7, 8}

	return batcherSignature, sequencerBatchesByte, nil
}

func (s *EspressoStreamer) parseEspressoTransaction(tx espressoTypes.Bytes) ([]*MessageWithHeight, error) {
	s.log.Info("Parsing espresso transaction", "tx", hex.EncodeToString(tx))
	// placeholder for batcherSignature and sequencerBatchesByte
	_, _, err := ParseHotShotPayload(tx)
	if err != nil {
		s.log.Warn("failed to parse hotshot payload", "err", err)
		return nil, err
	}
	// if signature verification fails, we should skip this message
	// Sishan TODO: verify the signature instead
	// verifySignature(batcherSignature, sequencerBatchesByte, )

	// placeholder for sequencer batches, it should be derived from sequencerBatchesByte
	rng := rand.New(rand.NewSource(0x543331))
	chainID := big.NewInt(rng.Int63n(1000))
	txCount := 1 + rng.Intn(8)
	sequencerBatches := derive.RandomSingularBatch(rng, txCount, chainID) // It should be *derive.SingularBatch

	result := &MessageWithHeight{
		SequencerBatches: sequencerBatches,
		HotshotHeight:    s.nextHotshotBlockNum,
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
func (s *EspressoStreamer) QueueMessagesFromHotshot(
	ctx context.Context,
	parseHotShotPayloadFn func(tx espressoTypes.Bytes) ([]*MessageWithHeight, error),
) error {
	// Note: Adding the lock on top level
	// because s.nextHotshotBlockNum is updated if n.nextHotshotBlockNum == 0
	s.messageMutex.Lock()
	defer s.messageMutex.Unlock()

	if s.nextHotshotBlockNum == 0 {
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
		s.nextHotshotBlockNum = latestBlock
	}

	txns, err := s.espressoClient.FetchTransactionsInBlock(ctx, s.nextHotshotBlockNum, s.namespace)
	if err != nil {
		s.log.Warn("failed to fetch the transactions", "err", err)
		return err
	}

	if len(txns.Transactions) == 0 {
		s.log.Info("No transactions found in the hotshot block", "block number", s.nextHotshotBlockNum)
		s.nextHotshotBlockNum += 1
		return nil
	}

	for _, tx := range txns.Transactions {
		s.log.Info("Parsing espresso transaction", "tx", hex.EncodeToString(tx))
		messages, err := parseHotShotPayloadFn(tx)
		if err != nil {
			s.log.Warn("failed to verify espresso transaction", "err", err)
			continue
		}
		// Sishan TODO: Sort it, and filter out the messages have already been seen
		s.messageWithHeight = append(s.messageWithHeight, messages...)
	}

	s.nextHotshotBlockNum += 1

	return nil
}

func (s *EspressoStreamer) Start(ctx context.Context) error {

	s.log.Info("In the function, Starting espresso streamer")
	bigTimeout := 2 * time.Minute
	timer := time.NewTimer(bigTimeout)
	defer timer.Stop()

	// Sishan TODO: maybe use better handler with dynamic interval in the future
	ticker := time.NewTicker(s.pollingHotshotPollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := s.QueueMessagesFromHotshot(ctx, s.parseEspressoTransaction)
			if err != nil {
				s.log.Error("error while queueing messages", "err", err)
			} else {
				s.log.Info("Processing block", "block number", s.nextHotshotBlockNum)
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
