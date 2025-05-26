package benchmark

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"runtime"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TimestampedValue is a generic struct that holds a value of type T and
// a timestamp. It is used to record the time at which a value was assessed
type TimestampedValue[T any] struct {
	Value     T
	Timestamp time.Time
}

// WithTimestamp is a function that takes a value of type T and returns a
// TimestampedValue[T]. It is used to create a new TimestampedValue with the
// current time as the timestamp.
func WithTimestamp[T any](value T) TimestampedValue[T] {
	return TimestampedValue[T]{
		Value:     value,
		Timestamp: time.Now(),
	}
}

// NOTE: While all of these maps are using a Hash value as the key, they
//
//	are not all a hash of the same thing.
//
// Hash Values:
// - Created: Hash of the Transaction
// - Submitted: Hash of the Transaction
// - Receipts: Hash of the Transaction
// - CaffReceipts: Hash of the Block Header
// - VerifyReceipts: Hash of the Block Header
type BenchmarkStats struct {
	Created        map[common.Hash]TimestampedValue[*geth_types.Transaction]
	Submitted      map[common.Hash]TimestampedValue[common.Hash]
	Receipts       map[common.Hash]TimestampedValue[*geth_types.Receipt]
	SeqReceipts    map[common.Hash]TimestampedValue[common.Hash]
	CaffReceipts   map[common.Hash]TimestampedValue[common.Hash]
	VerifyReceipts map[common.Hash]TimestampedValue[common.Hash]
}

// BenchmarkSubmissionsConfig is a configuration for a single account
// submission to the L2 Sequencer. This helps to govern the load that
// we attempt to place on the system via a single account.
type BenchmarkSubmitterConfig struct {
	Interval time.Duration
	To       *common.Address
	Value    *big.Int
	Signer   geth_types.Signer
	ChainID  *big.Int
	Key      *ecdsa.PrivateKey
}

// BenchmarkConfig is a struct that holds the configuration for the
// benchmarking process.
type BenchmarkConfig struct {
	Submitters                  []BenchmarkSubmitterConfig
	NumSubmitTransactionWorkers int
	NumReceiptWorkers           int

	SeqClient    *ethclient.Client
	CaffClient   *ethclient.Client
	VerifyClient *ethclient.Client
}

// BenchmarkOption is a a configuration option for the StartBenchmarking
// function. It is used to configure the benchmark configuration before
// the benchmark process has started.
type BenchmarkOption func(*BenchmarkConfig)

// AddSubmitter is a a BenchmarkOption that adds a submitter to the list of
// submitters.
//
// NOTE: Since the nonce **MUST** be unique for transaction submissions on a
// given wallet address, and is expected to be sequential, you absolutely
// **SHOULD NOT** use the same wallet address / private key for multiple
// submitters.
func AddSubmitter(
	submitter BenchmarkSubmitterConfig,
) BenchmarkOption {
	return func(config *BenchmarkConfig) {
		config.Submitters = append(config.Submitters, submitter)
	}
}

// WithSeqClient is a BenchmarkOption that sets the L2 Sequencer client
// to be used for the benchmark.  This is the client that will be used
// to submit the transactions to the L2 Sequencer.
//
// NOTE: if you want to submit transactions to the L2 Sequencer with the
// AddSubmitter option, you will need to set the L2 Sequencer Client as well.
func WithSeqClient(client *ethclient.Client) BenchmarkOption {
	return func(config *BenchmarkConfig) {
		config.SeqClient = client
	}
}

// WithCaffClient is a BenchmarkOption that sets the Caff client to be
// used for the benchmark.  This is the client that will be used to
// subscribe to the Caff node for block headers.
//
// NOTE: This option is required if you want to track the Caff Node Receipt
// time.
func WithCaffClient(client *ethclient.Client) BenchmarkOption {
	return func(config *BenchmarkConfig) {
		config.CaffClient = client
	}
}

// WithVerifyClient is a BenchmarkOption that sets the L2 Verifier client
// to be used for the benchmark.  This is the client that will be used
// to subscribe to the L2 Verifier node for block headers.
//
// NOTE: This option is required if you want to track the L2 Verifier
// Receipt time.
func WithVerifyClient(client *ethclient.Client) BenchmarkOption {
	return func(config *BenchmarkConfig) {
		config.VerifyClient = client
	}
}

// Benchmarker is an interface that defines the functionality of running
// a benchmark and retrieving their statistics. It is a useful abstraction
// so that the user need not worry about all of the details about how the
// benchmark is being run by itself.
type Benchmarker interface {
	// RunWithContext will run the benchmarker with the given context.  The
	// Context itself determines when the benchmarker will stop.  Once the
	// benchmarker has completed it's run, it will return the statistics for
	// the benchmark.
	RunWithContext(ctx context.Context) (BenchmarkStats, error)
}

// benchmarkState is a struct that holds the state of the benchmark.
// It includes the configuration for the benchmark, the channels for
// communication between the workers, and the statistics for the
// benchmark.
type benchmarkState struct {
	running bool
	cfg     BenchmarkConfig

	// WaitGroups and Contexts to synchronize the cancellation of the goroutines
	signers        waitGroupContext
	submitter      waitGroupContext
	receipt        waitGroupContext
	subscriptions  waitGroupContext
	metricRecorder waitGroupContext

	// Communication Channels for the Workers
	txSignedChanSrc       chan TimestampedValue[*geth_types.Transaction]
	txSubmitterChanSrc    chan TimestampedValue[common.Hash]
	annotatedBlockChanSrc chan TimestampedValue[AnnotatedBlockHash]

	// Client Subscriptions
	seqSubscription    geth.Subscription
	caffSubscription   geth.Subscription
	verifySubscription geth.Subscription

	// The collected statistics for the benchmark
	stats *BenchmarkStats
}

// start will start spawn the workers and spin up the benchmark criteria.
// After starting, the user will need to call stop in order to stop the
// benchmark.
func (b *benchmarkState) start(ctx context.Context) {
	if b.running {
		// We are already running, doing so again would be a problem.
		return
	}
	b.running = true
	config := b.cfg

	l2SeqClient := config.SeqClient
	l2Caff := config.CaffClient
	l2Verif := config.VerifyClient

	// Create the Channels for the Workers
	b.txSignedChanSrc = make(chan TimestampedValue[*geth_types.Transaction], 10)
	txSignedChanDst1, txSignedChanDst2 := TeeChan(b.txSignedChanSrc)
	b.txSubmitterChanSrc = make(chan TimestampedValue[common.Hash], 10)
	txSubmitterChanDst1, txSubmitterChanDst2 := TeeChan(b.txSubmitterChanSrc)

	txReceiptChan := make(chan TimestampedValue[*geth_types.Receipt], 10)
	caffBlockChan := make(chan *geth_types.Header, 1024)
	verifyBlockChan := make(chan *geth_types.Header, 1024)
	seqBlockChain := make(chan *geth_types.Header, 1024)
	b.annotatedBlockChanSrc = make(chan TimestampedValue[AnnotatedBlockHash], 10)

	b.stats = &BenchmarkStats{
		Created:        make(map[common.Hash]TimestampedValue[*geth_types.Transaction], 1024),
		Submitted:      make(map[common.Hash]TimestampedValue[common.Hash], 1024),
		Receipts:       make(map[common.Hash]TimestampedValue[*geth_types.Receipt], 1024),
		SeqReceipts:    make(map[common.Hash]TimestampedValue[common.Hash], 1024),
		CaffReceipts:   make(map[common.Hash]TimestampedValue[common.Hash], 1024),
		VerifyReceipts: make(map[common.Hash]TimestampedValue[common.Hash], 1024),
	}

	b.signers = NewCancelContext(ctx)
	b.submitter = NewCancelContext(ctx)
	b.receipt = NewCancelContext(ctx)
	b.subscriptions = NewCancelContext(ctx)
	b.metricRecorder = NewCancelContext(ctx)

	b.metricRecorder.wg.Add(1)
	go WorkerRecordTimestampedEvents(b.metricRecorder.ctx, &b.metricRecorder.wg, txSignedChanDst2, txSubmitterChanDst2, txReceiptChan, b.annotatedBlockChanSrc, b.stats)

	if l2Caff != nil {
		b.subscriptions.wg.Add(1)
		go WorkerConsumeBlockHeaders(b.subscriptions.ctx, &b.subscriptions.wg, env.RoleCaffNode, caffBlockChan, b.annotatedBlockChanSrc)
		cafSub, err := l2Caff.SubscribeNewHead(ctx, caffBlockChan)
		if err != nil {
			panic(err)
		}
		b.caffSubscription = cafSub
	}

	if l2Verif != nil {
		b.subscriptions.wg.Add(1)
		go WorkerConsumeBlockHeaders(b.subscriptions.ctx, &b.subscriptions.wg, e2esys.RoleVerif, verifyBlockChan, b.annotatedBlockChanSrc)
		verifSub, err := l2Verif.SubscribeNewHead(ctx, verifyBlockChan)
		if err != nil {
			panic(err)
		}
		b.verifySubscription = verifSub
	}

	if l2SeqClient != nil {
		b.subscriptions.wg.Add(1)
		go WorkerConsumeBlockHeaders(b.subscriptions.ctx, &b.subscriptions.wg, e2esys.RoleSeq, seqBlockChain, b.annotatedBlockChanSrc)
		seqSub, err := l2SeqClient.SubscribeNewHead(ctx, seqBlockChain)
		if err != nil {
			panic(err)
		}
		b.seqSubscription = seqSub
		for i := 0; i < config.NumSubmitTransactionWorkers; i++ {
			b.submitter.wg.Add(1)
			go WorkerSubmitSignedTransaction(b.submitter.ctx, &b.submitter.wg, txSignedChanDst1, b.txSubmitterChanSrc, l2SeqClient)
		}

		for i := 0; i < config.NumReceiptWorkers; i++ {
			b.receipt.wg.Add(1)
			go WorkerProcessL2Receipt(b.receipt.ctx, &b.receipt.wg, txSubmitterChanDst1, txReceiptChan, l2SeqClient)
		}

		for _, submitter := range config.Submitters {
			b.signers.wg.Add(1)
			go WorkerSignTransaction(b.signers.ctx, &b.signers.wg, submitter.Interval, b.txSignedChanSrc, submitter.Key, submitter.Signer, submitter.ChainID, submitter.To, submitter.Value)
		}
	}
}

// stop will stop the workers spawned by the benchmarks.  It does this
// by signalizing to the contexts that govern them that they should stop.
// The method will then wait for all of the workers to exit before returning.
// It will also ensure that it closes all of the channels opened to
// facilitate the communication between the workers.
func (b *benchmarkState) stop() {
	if !b.running {
		// We are not running, so we have nothing to stop.
		return
	}

	b.metricRecorder.cancel()
	b.subscriptions.cancel()
	b.submitter.cancel()
	b.signers.cancel()
	b.receipt.cancel()

	// We want to stop everything, but we want to do so in a manner that allows
	// for us to clean up effectively.
	// We also want to do so in a way that minimizes the data we are going to
	// lose.  That means that we still want to process the data that was already
	// generated in-flight.  All the worker queues should be drained effectively
	// before we we return.
	//
	// In order to do this, we need to cancel the contexts for the workers in
	// order.  This should happen in a down stream fashion.

	// Wait for the signer workers to finish.
	b.signers.wg.Wait()
	close(b.txSignedChanSrc)

	// Now the submitters can be stopped.
	// They should automatically stop once they see that the channel they
	// are reading from is closed.
	b.submitter.wg.Wait()
	close(b.txSubmitterChanSrc)

	// Now the receipt workers can be stopped.
	// They should automatically stop once they see that the channel they
	// are reading from is closed.
	b.receipt.wg.Wait()

	// Cancel the subscriptions to the new heads, and
	// wait for the subscription workers to finish.
	b.verifySubscription.Unsubscribe()
	b.caffSubscription.Unsubscribe()
	b.subscriptions.wg.Wait()
	close(b.annotatedBlockChanSrc)

	// Wait for the metric recorder to finish.
	b.metricRecorder.wg.Wait()

	// At this point all goroutines and workers should be stopped and we
	// should be synchronized.
	b.running = false
}

// RunWithContext is a function that runs the benchmark with the given
// context. It will start the benchmark and then wait for the context to
// be done. Once the context is done, it will stop the benchmark and
// return the statistics for the benchmark.
func (b *benchmarkState) RunWithContext(oCtx context.Context) (BenchmarkStats, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.start(ctx)

	// Wait for the context to be done
	<-oCtx.Done()

	// We are done, so we can stop the benchmark.
	b.stop()

	// Return the stats
	return *b.stats, nil
}

// CreateBenchmarker is a function that creates a new benchmarker with the
// given options. It will create a new benchmarker with the default
// configuration and then apply the given options to the benchmarker.
// It will return the benchmarker.
//
// NOTE: The Benchmarker returned is not started by default. You must call
// the RunWithContext function in order to start the benchmarker.
// Additionally, the Benchmarker will run until the context passed to the
// RunContext function is done.
//
// NOTE: This Benchmarker is designed to supply a load to the L2 Sequencer and
// to track and to track those transactions through their receipt on the
// L2 Sequencer.  From there we can track the transaction via the block
// hashes being received from the block headers streamed from the Caff
// node, the L2 Verifier node, and the L2 Sequencer.
//
// NOTE: The benchmarker is flexible and allows for a subset of Client Nodes
// to be used.  If you want to get the full set of metrics, you should supply
// the L2 Sequencer, the Caff Node, and the L2 Verifier Nodes.  Anything less
// and the full picture cannot be obtained, or even reasoned about.
func CreateBenchmarker(
	ctx context.Context,
	options ...BenchmarkOption,
) Benchmarker {
	config := &BenchmarkConfig{
		NumSubmitTransactionWorkers: runtime.NumCPU(),
		NumReceiptWorkers:           runtime.NumCPU(),
	}

	for _, opt := range options {
		opt(config)
	}

	return &benchmarkState{
		cfg: *config,
	}
}
