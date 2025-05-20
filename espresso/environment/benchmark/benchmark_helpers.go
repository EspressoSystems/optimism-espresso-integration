package benchmark

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"
	"runtime"
	"slices"
	"sync"
	"time"

	env "github.com/ethereum-optimism/optimism/espresso/environment"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SingleL2TransactionMetric is a struct that holds metric information for a
// single transaction submitted to the L2 Sequencer.
// It is meant to hold information about the transaction's lifecycle, and
// progression through the system and it's various milestones.
//
// This information can then be aggregated and used to calculate and compare
// the performance of the system under various confirmations and load
// conditions.
type SingleL2TransactionMetric struct {
	SignedTransaction *geth_types.Transaction
	Receipt           *geth_types.Receipt

	LocalCreated   time.Time
	LocalSubmitted time.Time
	SeqReceipt     time.Time
	CaffReceipt    time.Time
	VerifyReceipt  time.Time
}

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
			// TODO: record transaction submission failure for
			// later analysis
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
	CaffReceipts   map[common.Hash]TimestampedValue[common.Hash]
	VerifyReceipts map[common.Hash]TimestampedValue[common.Hash]
}

// Convert to Tracked Transactions
func (s *BenchmarkStats) IndividualTransactionMetrics() []SingleL2TransactionMetric {
	transactions := make([]SingleL2TransactionMetric, 0, len(s.Created))
	for txHash, tx := range s.Created {
		value := SingleL2TransactionMetric{
			SignedTransaction: tx.Value,
			LocalCreated:      tx.Timestamp,
		}

		if submitted, ok := s.Submitted[txHash]; ok {
			value.LocalSubmitted = submitted.Timestamp
		}

		if receipt, ok := s.Receipts[txHash]; ok {
			value.Receipt = receipt.Value
			value.SeqReceipt = receipt.Timestamp
		}

		receipt := value.Receipt
		if receipt == nil {
			transactions = append(transactions, value)
			continue
		}

		if caffReceipt, ok := s.CaffReceipts[receipt.BlockHash]; ok {
			value.CaffReceipt = caffReceipt.Timestamp
		}

		if verifyReceipt, ok := s.VerifyReceipts[receipt.BlockHash]; ok {
			value.VerifyReceipt = verifyReceipt.Timestamp
		}

		transactions = append(transactions, value)
	}

	return transactions
}

// SplitIndividualTransactionMetrics is a function that takes a slice of
// SingleL2TransactionMetric and splits it into two slices: one for
// complete transactions and one for incomplete transactions. A
// transaction is considered complete if it has a receipt and a
// caff receipt. Otherwise, it is considered incomplete.
func SplitIndividualTransactionMetrics(
	transactions []SingleL2TransactionMetric,
) (complete, incomplete []SingleL2TransactionMetric) {
	complete = make([]SingleL2TransactionMetric, 0, len(transactions))
	incomplete = make([]SingleL2TransactionMetric, 0, len(transactions))
	var zeroTime time.Time

	for _, tx := range transactions {
		if tx.Receipt != nil && tx.CaffReceipt != zeroTime && tx.VerifyReceipt != zeroTime {
			complete = append(complete, tx)
		} else {
			incomplete = append(incomplete, tx)
		}
	}

	return complete, incomplete
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type SampleSummary[T comparable] struct {
	Count  int
	Min    T
	Max    T
	Mean   T
	Median T
	StdDev T
}

func SummarizeSamples[T integer](samples []T) SampleSummary[T] {
	if len(samples) <= 0 {
		return SampleSummary[T]{}
	}

	// Sort the samples
	slices.Sort(samples)

	mid := len(samples) / 2

	metric := SampleSummary[T]{
		Count:  len(samples),
		Min:    samples[0],
		Max:    samples[len(samples)-1],
		Median: samples[mid],
	}

	total := new(big.Int)
	for _, sample := range samples {
		total.Add(total, big.NewInt(int64(sample)))
	}
	mean := new(big.Int).Div(total, big.NewInt(int64(len(samples))))
	metric.Mean = T(mean.Int64())

	// Calculate the standard deviation
	variance := new(big.Int)
	for _, duration := range samples {
		v := duration - metric.Mean
		variance.Add(variance, big.NewInt(int64(v*v)))
	}
	variance = variance.Div(variance, big.NewInt(int64(len(samples))))
	metric.StdDev = T(math.Sqrt(float64(variance.Int64())))

	return metric
}

type TimingMetrics struct {
	CreatedToSubmitted SampleSummary[time.Duration]
	SubmittedToReceipt SampleSummary[time.Duration]
	ReceiptToCaff      SampleSummary[time.Duration]
	ReceiptToVerify    SampleSummary[time.Duration]
}

func ComputeCompletedTransactionStatistics(completed []SingleL2TransactionMetric) TimingMetrics {
	var zeroTime time.Time
	var createdToSubmittedSamples []time.Duration
	var submittedToReceiptSamples []time.Duration
	var receiptToCaffSamples []time.Duration
	var receiptToVerifySamples []time.Duration

	for _, tx := range completed {
		if tx.LocalCreated == zeroTime {
			continue
		}

		if tx.LocalSubmitted == zeroTime {
			continue
		}

		createdToSubmittedSample := tx.SeqReceipt.Sub(tx.LocalSubmitted)
		createdToSubmittedSamples = append(createdToSubmittedSamples, createdToSubmittedSample)

		if tx.SeqReceipt == zeroTime {
			continue
		}

		submittedToReceiptSample := tx.SeqReceipt.Sub(tx.LocalCreated)
		submittedToReceiptSamples = append(submittedToReceiptSamples, submittedToReceiptSample)

		if tx.CaffReceipt != zeroTime {
			receiptToCaffSample := tx.CaffReceipt.Sub(tx.SeqReceipt)
			receiptToCaffSamples = append(receiptToCaffSamples, receiptToCaffSample)
		}

		if tx.VerifyReceipt != zeroTime {
			receiptToVerifySample := tx.VerifyReceipt.Sub(tx.SeqReceipt)
			receiptToVerifySamples = append(receiptToVerifySamples, receiptToVerifySample)
		}
	}

	return TimingMetrics{
		CreatedToSubmitted: SummarizeSamples(createdToSubmittedSamples),
		SubmittedToReceipt: SummarizeSamples(submittedToReceiptSamples),
		ReceiptToCaff:      SummarizeSamples(receiptToCaffSamples),
		ReceiptToVerify:    SummarizeSamples(receiptToVerifySamples),
	}
}

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
			case e2esys.RoleVerif:
				stats.VerifyReceipts[annotatedBlock.Value.BlockHash] = TimestampedValue[common.Hash]{Value: annotatedBlock.Value.BlockHash, Timestamp: annotatedBlock.Timestamp}
			}
		}
	}
}

// TeeChan is a helper function that takes a channel that is expected to be
// send to, (the source channel) and returns two channels that should be
// submitted to.
func TeeChan[T any](src <-chan T) (<-chan T, <-chan T) {
	dst1 := make(chan T, cap(src))
	dst2 := make(chan T, cap(src))

	go func() {
		for v := range src {
			dst1 <- v
			dst2 <- v
		}

		close(dst1)
		close(dst2)
	}()

	return dst1, dst2
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

func WithSeqClient(client *ethclient.Client) BenchmarkOption {
	return func(config *BenchmarkConfig) {
		config.SeqClient = client
	}
}

func WithCaffClient(client *ethclient.Client) BenchmarkOption {
	return func(config *BenchmarkConfig) {
		config.CaffClient = client
	}
}

func WithVerifyClient(client *ethclient.Client) BenchmarkOption {
	return func(config *BenchmarkConfig) {
		config.VerifyClient = client
	}
}

type Benchmarker interface {
	RunWithContext(ctx context.Context) (BenchmarkStats, error)
}

type waitGroupContext struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// Deadline implements context.Context.
func (w *waitGroupContext) Deadline() (deadline time.Time, ok bool) {
	return w.ctx.Deadline()
}

// Done implements context.Context.
func (w *waitGroupContext) Done() <-chan struct{} {
	return w.ctx.Done()
}

// Err implements context.Context.
func (w *waitGroupContext) Err() error {
	return w.ctx.Err()
}

// Value implements context.Context.
func (w *waitGroupContext) Value(key any) any {
	return w.ctx.Value(key)
}

var _ context.Context = (*waitGroupContext)(nil)

func NewCancelContext(ctx context.Context) waitGroupContext {
	ctx, cancel := context.WithCancel(ctx)
	return waitGroupContext{
		wg:     sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
	}
}

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
	caffSubscription   geth.Subscription
	verifySubscription geth.Subscription

	// The collected statistics for the benchmark
	stats *BenchmarkStats
}

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
	b.annotatedBlockChanSrc = make(chan TimestampedValue[AnnotatedBlockHash], 10)

	b.stats = &BenchmarkStats{
		Created:        make(map[common.Hash]TimestampedValue[*geth_types.Transaction], 1024),
		Submitted:      make(map[common.Hash]TimestampedValue[common.Hash], 1024),
		Receipts:       make(map[common.Hash]TimestampedValue[*geth_types.Receipt], 1024),
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

	b.subscriptions.wg.Add(1)
	go WorkerConsumeBlockHeaders(b.subscriptions.ctx, &b.subscriptions.wg, env.RoleCaffNode, caffBlockChan, b.annotatedBlockChanSrc)
	b.subscriptions.wg.Add(1)
	go WorkerConsumeBlockHeaders(b.subscriptions.ctx, &b.subscriptions.wg, e2esys.RoleVerif, verifyBlockChan, b.annotatedBlockChanSrc)

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

	verifSub, err := l2Verif.SubscribeNewHead(ctx, verifyBlockChan)
	if err != nil {
		panic(err)
	}
	b.verifySubscription = verifSub
	cafSub, err := l2Caff.SubscribeNewHead(ctx, caffBlockChan)
	b.caffSubscription = cafSub
	if err != nil {
		panic(err)
	}
}

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

	// Stop the Signers
	b.signers.cancel()
	b.signers.wg.Wait()
	close(b.txSignedChanSrc)

	// Now the submitters can be stopped.
	// They should automatically stop once they see that the channel they
	// are reading from is closed.
	b.submitter.wg.Wait()
	b.submitter.cancel()
	close(b.txSubmitterChanSrc)

	// Now the receipt workers can be stopped.
	// They should automatically stop once they see that the channel they
	// are reading from is closed.
	b.receipt.wg.Wait()
	b.receipt.cancel()

	// Stop the subscriptions
	b.verifySubscription.Unsubscribe()
	b.caffSubscription.Unsubscribe()
	b.subscriptions.wg.Wait()
	b.subscriptions.cancel()
	close(b.annotatedBlockChanSrc)

	// Stop the metric recorder
	b.metricRecorder.cancel()
	b.metricRecorder.wg.Wait()

	// At this point all goroutines and workers should be stopped and we
	// should be synchronized.
	b.running = false
}

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
