package espresso_test

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"log/slog"
	"math/big"
	"math/rand"
	"testing"
	"time"

	espressoClient "github.com/EspressoSystems/espresso-network/sdks/go/client"
	espressoCommon "github.com/EspressoSystems/espresso-network/sdks/go/types"
	"github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	opsigner "github.com/ethereum-optimism/optimism/op-service/signer"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// TestNewEspressoStreamer tests that we can create a new EspressoStreamer
// without any panic being thrown.

func TestNewEspressoStreamer(t *testing.T) {
	_ = espresso.NewEspressoStreamer(
		1,
		nil,
		nil,
		nil, nil, nil, derive.CreateEspressoBatchUnmarshaler(common.Address{}),
		50*time.Millisecond,
		0,
		1,
	)
}

// EspBlockAndNamespace is a struct that holds the height and namespace
// of an Espresso block. It is used to uniquely identify a block in the
// EspressoStreamer.
type EspBlockAndNamespace struct {
	Height, Namespace uint64
}

// BlockAndNamespace creates a new EspBlockAndNamespace struct
// with the provided height and namespace.
func BlockAndNamespace(height, namespace uint64) EspBlockAndNamespace {
	return EspBlockAndNamespace{
		Height:    height,
		Namespace: namespace,
	}
}

// MockStreamerSource is a mock implementation of the various interfaces
// required by the EspressoStreamer.  The idea behind this mock is to allow
// for the specific progression of the L1, L2, and Espresso states, so we can
// verify the implementation of our Streamer, in relation to specific scenarios
// and edge cases, without needing to forcibly simulate them via a live test
// environment.
//
// As we progress through the tests, we should be able to update our local mock
// state, and then perform our various `.Update` and `.Next` calls, in order to
// verify that we end up with the expected state.
//
// The current expected use case for the Streamer is for the user to "Refresh"
// the state of the streamer by calling `.Refresh`.
type MockStreamerSource struct {
	// At the moment the Streamer utilizes the SyncStatus in order to update
	// it's local state.  But, in general the Streamer doesn't consume all
	// of the fields provided within the SyncStatus.  At the moment it only
	// cares about SafeL2, and FinalizedL1. So this is what we will track

	FinalizedL1 eth.L1BlockRef
	SafeL2      eth.L2BlockRef

	EspTransactionData     map[EspBlockAndNamespace]espressoClient.TransactionsInBlock
	LatestEspHeight        uint64
	finalizedHeightHistory map[uint64]uint64
}

// FetchNamespaceTransactionsInRange implements espresso.EspressoClient.
func (m *MockStreamerSource) FetchNamespaceTransactionsInRange(ctx context.Context, fromHeight uint64, toHeight uint64, namespace uint64) ([]espressoCommon.NamespaceTransactionsRangeData, error) {
	var result []espressoCommon.NamespaceTransactionsRangeData

	if fromHeight > toHeight {
		return nil, ErrNotFound
	}
	for height := fromHeight; height <= toHeight; height++ {
		transactionsInBlock, ok := m.EspTransactionData[BlockAndNamespace(height, namespace)]
		if !ok {
			// Preserve alignment with the requested range even if the block
			// has no transactions in this namespace.
			result = append(result, espressoCommon.NamespaceTransactionsRangeData{})
			continue
		}

		var txs []espressoCommon.Transaction
		for _, txPayload := range transactionsInBlock.Transactions {
			tx := espressoCommon.Transaction{
				Namespace: namespace,
				Payload:   txPayload,
			}
			txs = append(txs, tx)
		}

		result = append(result, espressoCommon.NamespaceTransactionsRangeData{
			Transactions: txs})
	}
	return result, nil
}

func NewMockStreamerSource() *MockStreamerSource {
	finalizedL1 := createL1BlockRef(1)
	return &MockStreamerSource{
		FinalizedL1:            finalizedL1,
		SafeL2:                 createL2BlockRef(0, finalizedL1),
		EspTransactionData:     make(map[EspBlockAndNamespace]espressoClient.TransactionsInBlock),
		finalizedHeightHistory: make(map[uint64]uint64),
		LatestEspHeight:        0,
	}
}

// AdvanceFinalizedL1ByNBlocks advances the FinalizedL1 block reference by n blocks.
func (m *MockStreamerSource) AdvanceFinalizedL1ByNBlocks(n uint) {
	for range n {
		m.AdvanceFinalizedL1()
	}
}

// AdvanceFinalizedL1 advances the FinalizedL1 block reference by one block.
func (m *MockStreamerSource) AdvanceFinalizedL1() {
	m.finalizedHeightHistory[m.FinalizedL1.Number] = m.LatestEspHeight
	m.FinalizedL1 = createL1BlockRef(m.FinalizedL1.Number + 1)
}

// AdvanceL2ByNBlocks advances the SafeL2 block reference by n blocks.
func (m *MockStreamerSource) AdvanceL2ByNBlocks(n uint) {
	m.SafeL2 = createL2BlockRef(m.SafeL2.Number+uint64(n), m.FinalizedL1)
}

// AdvanceSafeL2 advances the SafeL2 block reference by one block.
func (m *MockStreamerSource) AdvanceSafeL2() {
	m.SafeL2 = createL2BlockRef(m.SafeL2.Number+1, m.FinalizedL1)
}

// AdvanceEspressoHeightByNBlocks advances the LatestEspHeight by n blocks.
func (m *MockStreamerSource) AdvanceEspressoHeightByNBlocks(n int) {
	m.LatestEspHeight += uint64(n)
}

// AdvanceEspressoHeight advances the LatestEspHeight by one block.
func (m *MockStreamerSource) AdvanceEspressoHeight() {
	m.LatestEspHeight++
}

// SyncStatus returns the current sync status of the mock streamer source.
// Only the fields FinalizedL1, FinalizedL1, and SafeL2 are populated, as those
// are the only fields explicitly inspected by the EspressoStreamer.
func (m *MockStreamerSource) SyncStatus() *eth.SyncStatus {
	return &eth.SyncStatus{
		FinalizedL1: m.FinalizedL1,
		SafeL2:      m.SafeL2,
	}
}

func (m *MockStreamerSource) AddEspressoTransactionData(height, namespace uint64, txData espressoClient.TransactionsInBlock) {
	if m.EspTransactionData == nil {
		m.EspTransactionData = make(map[EspBlockAndNamespace]espressoClient.TransactionsInBlock)
	}

	m.EspTransactionData[BlockAndNamespace(height, namespace)] = txData

	if m.LatestEspHeight < height {
		m.LatestEspHeight = height
	}
}

var _ espresso.L1Client = (*MockStreamerSource)(nil)

// L1 Client methods

func (m *MockStreamerSource) HeaderHashByNumber(ctx context.Context, number *big.Int) (common.Hash, error) {
	l1Ref := createL1BlockRef(number.Uint64())
	return l1Ref.Hash, nil
}

// Espresso Client Methods
var _ espresso.EspressoClient = (*MockStreamerSource)(nil)

func (m *MockStreamerSource) FetchLatestBlockHeight(ctx context.Context) (uint64, error) {
	return m.LatestEspHeight, nil
}

// ErrorNotFound is a custom error type used to indicate that a requested
// resource was not found.
type ErrorNotFound struct{}

// Error implements error.
func (ErrorNotFound) Error() string {
	return "not found"
}

// ErrNotFound is an instance of ErrorNotFound that can be used to indicate
// that a requested resource was not found.
var ErrNotFound error = ErrorNotFound{}

type MockTransactionStream struct {
	pos       uint64
	subPos    uint64
	end       uint64
	namespace uint64
	source    *MockStreamerSource
}

func (ms *MockTransactionStream) Next(ctx context.Context) (*espressoCommon.TransactionQueryData, error) {
	raw, err := ms.NextRaw(ctx)
	if err != nil {
		return nil, err
	}
	var transaction espressoCommon.TransactionQueryData
	if err := json.Unmarshal(raw, &transaction); err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (ms *MockTransactionStream) NextRaw(ctx context.Context) (json.RawMessage, error) {
	for {
		// get the latest block number
		latestHeight, err := ms.source.FetchLatestBlockHeight(ctx)
		if err != nil {
			// We will return error on NotFound as well to speed up tests.
			// More faithful imitation of HotShot streaming API would be to hang
			// until we receive new transactions, but that would slow down some
			// tests significantly, because streamer would wait for full timeout
			// threshold here before finishing update.
			return nil, err
		}

		if ms.pos > latestHeight {
			return nil, ErrNotFound
		}

		namespaceTransactions, err := ms.source.FetchNamespaceTransactionsInRange(ctx, ms.pos, latestHeight, ms.namespace)
		if err != nil {
			return nil, err
		}

		// Each element in the returned slice corresponds to a block starting
		// at fromHeight. We only need the current block (index 0) because
		// fromHeight == ms.pos.
		if len(namespaceTransactions) == 0 {
			return nil, ErrNotFound
		}

		currentBlock := namespaceTransactions[0]

		if len(currentBlock.Transactions) > int(ms.subPos) {
			tx := currentBlock.Transactions[int(ms.subPos)]
			transaction := &espressoCommon.TransactionQueryData{
				BlockHeight: ms.pos,
				Index:       ms.subPos,
				Transaction: espressoCommon.Transaction{
					Payload:   tx.Payload,
					Namespace: ms.namespace,
				},
			}
			ms.subPos++
			return json.Marshal(transaction)
		}

		// Move on to the next block.
		ms.subPos = 0
		ms.pos++
	}
}

func (ms *MockTransactionStream) Close() error {
	return nil
}

func (m *MockStreamerSource) StreamTransactionsInNamespace(ctx context.Context, height uint64, namespace uint64) (espressoClient.Stream[espressoCommon.TransactionQueryData], error) {
	if m.LatestEspHeight < height {
		return nil, ErrNotFound
	}

	return &MockTransactionStream{
		pos:       height,
		subPos:    0,
		end:       m.LatestEspHeight,
		namespace: namespace,
		source:    m,
	}, nil
}

// Espresso Light Client implementation
var _ espresso.LightClientCallerInterface = (*MockStreamerSource)(nil)

// LightClientCallerInterface implementation
func (m *MockStreamerSource) FinalizedState(opts *bind.CallOpts) (espresso.FinalizedState, error) {
	height, ok := m.finalizedHeightHistory[opts.BlockNumber.Uint64()]
	if !ok {
		height = m.LatestEspHeight
	}
	return espresso.FinalizedState{
		ViewNum:     height,
		BlockHeight: height,
	}, nil
}

// NoOpLogger is a no-op implementation of the log.Logger interface.
// It is used to pass a non-nil logger to the EspressoStreamer without
// producing any output.
type NoOpLogger struct{}

var _ log.Logger = (*NoOpLogger)(nil)

func (l *NoOpLogger) With(ctx ...interface{}) log.Logger                                   { return l }
func (l *NoOpLogger) New(ctx ...interface{}) log.Logger                                    { return l }
func (l *NoOpLogger) Log(level slog.Level, msg string, ctx ...interface{})                 {}
func (l *NoOpLogger) Trace(msg string, ctx ...interface{})                                 {}
func (l *NoOpLogger) Debug(msg string, ctx ...interface{})                                 {}
func (l *NoOpLogger) Info(msg string, ctx ...interface{})                                  {}
func (l *NoOpLogger) Warn(msg string, ctx ...interface{})                                  {}
func (l *NoOpLogger) Error(msg string, ctx ...interface{})                                 {}
func (l *NoOpLogger) Crit(msg string, ctx ...interface{})                                  { panic("critical error") }
func (l *NoOpLogger) Write(level slog.Level, msg string, attrs ...any)                     {}
func (l *NoOpLogger) Enabled(ctx context.Context, level slog.Level) bool                   { return true }
func (l *NoOpLogger) Handler() slog.Handler                                                { return nil }
func (l *NoOpLogger) TraceContext(ctx context.Context, msg string, ctxArgs ...interface{}) {}
func (l *NoOpLogger) DebugContext(ctx context.Context, msg string, ctxArgs ...interface{}) {}
func (l *NoOpLogger) InfoContext(ctx context.Context, msg string, ctxArgs ...interface{})  {}
func (l *NoOpLogger) WarnContext(ctx context.Context, msg string, ctxArgs ...interface{})  {}
func (l *NoOpLogger) ErrorContext(ctx context.Context, msg string, ctxArgs ...interface{}) {}
func (l *NoOpLogger) CritContext(ctx context.Context, msg string, ctxArgs ...interface{}) {
	panic("critical error")
}
func (l *NoOpLogger) LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
}
func (l *NoOpLogger) SetContext(ctx context.Context)                                          {}
func (l *NoOpLogger) WriteCtx(ctx context.Context, level slog.Level, msg string, args ...any) {}

func createHashFromHeight(height uint64) common.Hash {
	var hash common.Hash
	binary.LittleEndian.PutUint64(hash[(len(hash)-8):], height)
	return hash
}

// createL1BlockRef creates a mock L1BlockRef for testing purposes, with the
// every field being derived from the provided height.  This should be
// sufficient for testing purposes.
func createL1BlockRef(height uint64) eth.L1BlockRef {
	var parentHash common.Hash
	if height > 0 {
		parentHash = createHashFromHeight(height - 1)
	}
	return eth.L1BlockRef{
		Number:     height,
		Hash:       createHashFromHeight(height),
		ParentHash: parentHash,
		Time:       height,
	}
}

// createL2BlockRef creates a mock L2BlockRef for testing purposes, with the
// every field being derived from the provided height and L1BlockRef.  This
// should be sufficient for testing purposes.
func createL2BlockRef(height uint64, l1Ref eth.L1BlockRef) eth.L2BlockRef {
	return eth.L2BlockRef{
		Number:         height,
		Hash:           createHashFromHeight(height),
		ParentHash:     createHashFromHeight(height - 1),
		Time:           height,
		SequenceNumber: 1,
		L1Origin: eth.BlockID{
			Hash:   l1Ref.Hash,
			Number: l1Ref.Number,
		},
	}
}

// setupStreamerTesting initializes a MockStreamerSource and an EspressoStreamer
// for testing purposes. It sets up the initial state of the MockStreamerSource
// and returns both the MockStreamerSource and the EspressoStreamer.
func setupStreamerTesting(namespace uint64, batcherAddress common.Address) (*MockStreamerSource, *espresso.BatchStreamer[derive.EspressoBatch]) {
	state := NewMockStreamerSource()

	logger := new(NoOpLogger)
	streamer := espresso.NewEspressoStreamer(
		namespace,
		state,
		state,
		state,
		state,
		logger,
		derive.CreateEspressoBatchUnmarshaler(batcherAddress),
		50*time.Millisecond,
		0,
		1,
	)

	return state, streamer
}

// createEspressoBatch creates a mock EspressoBatch for testing purposes
// containing the provided SingularBatch.
func createEspressoBatch(batch *derive.SingularBatch) *derive.EspressoBatch {
	return &derive.EspressoBatch{
		BatchHeader: &geth_types.Header{
			ParentHash: batch.ParentHash,
			Number:     big.NewInt(int64(batch.Timestamp)),
		},
		Batch:         *batch,
		L1InfoDeposit: geth_types.NewTx(&geth_types.DepositTx{}),
	}
}

// createEspressoTransaction creates a mock Espresso transaction for testing purposes
// containing the provided Espresso batch.
func createEspressoTransaction(ctx context.Context, batch *derive.EspressoBatch, namespace uint64, chainSigner crypto.ChainSigner) *espressoCommon.Transaction {
	tx, err := batch.ToEspressoTransaction(ctx, namespace, chainSigner)
	if have, want := err, error(nil); have != want {
		panic(err)
	}

	return tx
}

// createTransactionsInBlock creates a mock TransactionsInBlock for testing purposes
// containing the provided Espresso transaction.
func createTransactionsInBlock(tx *espressoCommon.Transaction) espressoClient.TransactionsInBlock {
	return espressoClient.TransactionsInBlock{
		Transactions: []espressoCommon.Bytes{tx.Payload},
	}
}

// CreateEspressoTxnData creates a mock Espresso transaction data set
// for testing purposes. It generates a test SingularBatch, and takes it
// through the steps of getting all the way to an Espresso transaction in block.
// Every intermediate step is returned for inspection / utilization in tests.
// Uses m.FinalizedL1 as the L1 origin.
func (m *MockStreamerSource) CreateEspressoTxnData(
	ctx context.Context,
	namespace uint64,
	rng *rand.Rand,
	chainID *big.Int,
	l2Height uint64,
	chainSigner crypto.ChainSigner,
) (*derive.SingularBatch, *derive.EspressoBatch, *espressoCommon.Transaction, espressoClient.TransactionsInBlock) {
	return m.CreateEspressoTxnDataWithL1Origin(ctx, namespace, rng, chainID, l2Height, chainSigner, m.FinalizedL1.Number, m.FinalizedL1.Hash)
}

// TestStreamerSmoke tests the basic functionality of the EspressoStreamer
// ensuring that it behaves as expected from an empty state with no
// iterations, batches, or blocks.
func TestStreamerSmoke(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	state, streamer := setupStreamerTesting(42, common.Address{})

	// update the state of our streamer
	syncStatus := state.SyncStatus()
	err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)

	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to refresh streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// Update the state of our streamer
	if have, want := streamer.Update(ctx), error(nil); !errors.Is(have, want) {
		t.Fatalf("failed to update streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	// We should not get any batches from the Streamer at this point.
	if have, want := streamer.Next(ctx), (*derive.EspressoBatch)(nil); have != want {
		t.Fatalf("failed to get next batch from streamer:\nhave:\n\t%v\nwant:\n\t%v\n", have, want)
	}
}

// TestEspressoStreamerSimpleIncremental tests the EspressoStreamer by
// incrementally adding batches to the state and verifying that the streamer
// can retrieve them in the correct order.
func TestEspressoStreamerSimpleIncremental(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	state, streamer := setupStreamerTesting(namespace, signerAddress)
	rng := rand.New(rand.NewSource(0))
	// The number of batches to create
	const N = 1000

	for i := 0; i < N; i++ {
		// update the state of our streamer
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)

		if have, want := err, error(nil); have != want {
			t.Fatalf("failed to refresh streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		batch, _, _, espTxnInBlock := state.CreateEspressoTxnData(
			ctx,
			namespace,
			rng,
			chainID,
			uint64(i)+1,
			chainSigner,
		)

		state.AddEspressoTransactionData(uint64(5*i), namespace, espTxnInBlock)

		// Update the state of our streamer
		if have, want := streamer.Update(ctx), error(nil); !errors.Is(have, want) {
			t.Fatalf("failed to update streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
		}

		batchFromEsp := streamer.Next(ctx)
		require.NotNil(t, batchFromEsp, "unexpectedly did not receive a batch from streamer")

		// This batch ** should ** match the one we created above.

		if have, want := batchFromEsp.Batch.GetEpochNum(), batch.GetEpochNum(); have != want {
			t.Fatalf("batch epoch number does not match:\nhave:\n\t%v\ndo not want:\n\t%v\n", have, want)
		}

		state.AdvanceSafeL2()
		state.AdvanceFinalizedL1()
	}

	if have, want := len(state.EspTransactionData), N; have != want {
		t.Fatalf("unexpected number of batches in state:\nhave:\n\t%v\nwant:\n\t%v\n", have, want)
	}
}

// TestEspressoStreamerIncrementalDelayedConsumption tests the EspressoStreamer
// by populating all of the batches in the state before incrementing over them
func TestEspressoStreamerIncrementalDelayedConsumption(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	state, streamer := setupStreamerTesting(namespace, signerAddress)
	rng := rand.New(rand.NewSource(0))

	// The number of batches to create
	const N = 1000

	var batches []*derive.SingularBatch

	// update the state of our streamer
	syncStatus := state.SyncStatus()
	err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)

	for i := 0; i < N; i++ {
		batch, _, _, espTxnInBlock := state.CreateEspressoTxnData(
			ctx,
			namespace,
			rng,
			chainID,
			uint64(i)+1,
			chainSigner,
		)

		state.AddEspressoTransactionData(uint64(5*i), namespace, espTxnInBlock)
		batches = append(batches, batch)
	}

	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to refresh streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	for i := 0; i < N; i++ {
		if !streamer.HasNext(ctx) {
			// Update the state of our streamer
			if have, want := streamer.Update(ctx), error(nil); !errors.Is(have, want) {
				t.Fatalf("failed to update streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}
		}

		batch := batches[i]

		batchFromEsp := streamer.Next(ctx)
		require.NotNil(t, batchFromEsp, "unexpectedly did not receive a batch from streamer")

		// This batch ** should ** match the one we created above.

		if have, want := batchFromEsp.Batch.GetEpochNum(), batch.GetEpochNum(); have != want {
			t.Fatalf("batch epoch number does not match:\nhave:\n\t%v\ndo not want:\n\t%v\n", have, want)
		}

		state.AdvanceSafeL2()
		state.AdvanceFinalizedL1()
	}

	if have, want := len(state.EspTransactionData), N; have != want {
		t.Fatalf("unexpected number of batches in state:\nhave:\n\t%v\nwant:\n\t%v\n", have, want)
	}
}

// TestStreamerEspressoOutOfOrder tests the behavior of the EspressoStreamer
// when the batches coming from Espresso are not in sequential order.
//
// The Streamer is expected to be able to reorder these batches before
// iterating over them.
func TestStreamerEspressoOutOfOrder(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	state, streamer := setupStreamerTesting(namespace, signerAddress)
	rng := rand.New(rand.NewSource(0))

	// update the state of our streamer
	syncStatus := state.SyncStatus()
	err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)

	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to refresh streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	const N = 1000
	var batches []*derive.SingularBatch
	for i := 0; i < N; i++ {
		batch, _, _, block := state.CreateEspressoTxnData(
			ctx,
			namespace,
			rng,
			chainID,
			uint64(i)+1,
			chainSigner,
		)

		rollEspBlockNumber := rng.Intn(N * 5)
		for {
			_, ok := state.EspTransactionData[BlockAndNamespace(uint64(rollEspBlockNumber), namespace)]
			if ok {
				// re-roll, if already populated.
				rollEspBlockNumber = rng.Intn(N * 5)
				continue
			}

			break
		}

		state.AddEspressoTransactionData(uint64(rollEspBlockNumber), namespace, block)
		batches = append(batches, batch)
	}

	{

		for i := 0; i < N; i++ {
			for j := 0; j < int(state.LatestEspHeight/100); j++ {
				// Update the state of our streamer
				if have, want := streamer.Update(ctx), error(nil); !errors.Is(have, want) {
					t.Fatalf("failed to update streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
				}
				if streamer.HasNext(ctx) {
					break
				}
			}

			batch := batches[i]
			batchFromEsp := streamer.Next(ctx)
			require.NotNil(t, batchFromEsp, "unexpectedly did not receive a batch from streamer")

			// This batch ** should ** match the one we created above.

			if have, want := batchFromEsp.Batch.GetEpochNum(), batch.GetEpochNum(); have != want {
				t.Fatalf("batch epoch number does not match:\nhave:\n\t%v\ndo not want:\n\t%v\n", have, want)
			}

			state.AdvanceSafeL2()
		}
	}

	if have, want := len(state.EspTransactionData), N; have != want {
		t.Fatalf("unexpected number of batches in state:\nhave:\n\t%v\nwant:\n\t%v\n", have, want)
	}
}

// TestEspressoStreamerDuplicationHandling tests the behavior of the EspressoStreamer
// when a duplicated batch is received.
//
// The Streamer is expected to skip the duplicated batch and only return once for each batch.
func TestEspressoStreamerDuplicationHandling(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	state, streamer := setupStreamerTesting(namespace, signerAddress)
	rng := rand.New(rand.NewSource(0))

	// update the state of our streamer
	syncStatus := state.SyncStatus()
	err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)

	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to refresh streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	const N = 1000
	for i := 0; i < N; i++ {
		batch, _, _, espTxnInBlock := state.CreateEspressoTxnData(
			ctx,
			namespace,
			rng,
			chainID,
			uint64(i)+1,
			chainSigner,
		)

		// duplicate the batch
		for j := 0; j < 2; j++ {
			// update the state of our streamer
			syncStatus := state.SyncStatus()
			err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)

			require.NoError(t, err)

			// add the batch to the state, and make sure duplicate batches are also added with a different height
			state.AddEspressoTransactionData(uint64(5*i+j), namespace, espTxnInBlock)

			// Update the state of our streamer
			if have, want := streamer.Update(ctx), error(nil); !errors.Is(have, want) {
				t.Fatalf("failed to update streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
			}
		}

		batchFromEsp := streamer.Next(ctx)
		require.NotNil(t, batchFromEsp, "unexpectedly did not receive a batch from streamer")

		// This batch ** should ** match the one we created above.
		// If the duplicate one is NOT skipped, this will FAIL.
		require.Equal(t, batchFromEsp.Batch.GetEpochNum(), batch.GetEpochNum())

		state.AdvanceSafeL2()
		state.AdvanceFinalizedL1()

	}

	// Check that the state has the correct number of duplicated batches
	require.Equal(t, len(state.EspTransactionData), 2*N)
}

// createSingularBatch creates a mock SingularBatch for testing purposes
// with a specific L1 origin (epoch number and hash).
func (m *MockStreamerSource) createSingularBatch(rng *rand.Rand, txCount int, chainID *big.Int, l2Height uint64, epochNum uint64, epochHash common.Hash) *derive.SingularBatch {
	signer := geth_types.NewLondonSigner(chainID)
	baseFee := big.NewInt(rng.Int63n(300_000_000_000))
	txsEncoded := make([]hexutil.Bytes, 0, txCount)
	for i := 0; i < txCount; i++ {
		tx := testutils.RandomTx(rng, baseFee, signer)
		txEncoded, err := tx.MarshalBinary()
		if err != nil {
			panic("tx Marshal binary" + err.Error())
		}
		txsEncoded = append(txsEncoded, txEncoded)
	}

	return &derive.SingularBatch{
		ParentHash:   createHashFromHeight(l2Height),
		EpochNum:     rollup.Epoch(epochNum),
		EpochHash:    epochHash,
		Timestamp:    l2Height,
		Transactions: txsEncoded,
	}
}

// CreateEspressoTxnDataWithL1Origin creates a mock Espresso transaction data set
// for testing purposes with a specific L1 origin.
func (m *MockStreamerSource) CreateEspressoTxnDataWithL1Origin(
	ctx context.Context,
	namespace uint64,
	rng *rand.Rand,
	chainID *big.Int,
	l2Height uint64,
	chainSigner crypto.ChainSigner,
	epochNum uint64,
	epochHash common.Hash,
) (*derive.SingularBatch, *derive.EspressoBatch, *espressoCommon.Transaction, espressoClient.TransactionsInBlock) {
	txCount := rng.Intn(10)
	batch := m.createSingularBatch(rng, txCount, chainID, l2Height, epochNum, epochHash)
	espBatch := createEspressoBatch(batch)
	espTxn := createEspressoTransaction(ctx, espBatch, namespace, chainSigner)
	espTxnInBlock := createTransactionsInBlock(espTxn)

	return batch, espBatch, espTxn, espTxnInBlock
}

// TestStreamerHeadBatchHandling tests the headBatch direct assignment and buffer promotion behavior.
func TestStreamerHeadBatchHandling(t *testing.T) {
	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	t.Run("batch matching BatchPos assigned directly to headBatch when headBatch is nil", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(0))

		// Refresh state - after this, BatchPos becomes 1 (fallbackBatchPos=0, BatchPos=0+1=1)
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// Create batch with number matching BatchPos (which is 1 after Refresh with SafeL2.Number=0)
		_, _, _, espTxnInBlock := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, 1, chainSigner)
		state.AddEspressoTransactionData(0, namespace, espTxnInBlock)

		// Update to fetch the batch
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// The batch should be assigned directly to headBatch
		require.True(t, streamer.HasNext(ctx), "batch should be available")

		// Next should return the batch
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())
	})

	t.Run("HasNext promotes batch from buffer when headBatch is nil", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(1))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// Create batch 2 first (goes to buffer since it's ahead of BatchPos=1)
		_, _, _, espTxnInBlock2 := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, 2, chainSigner)
		state.AddEspressoTransactionData(0, namespace, espTxnInBlock2)

		// Create batch 1 (becomes headBatch)
		_, _, _, espTxnInBlock1 := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, 1, chainSigner)
		state.AddEspressoTransactionData(1, namespace, espTxnInBlock1)

		// Update to fetch both batches
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// Consume batch 1
		require.True(t, streamer.HasNext(ctx))
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())

		// Now BatchPos is 2, and batch 2 should be promoted from buffer
		require.True(t, streamer.HasNext(ctx), "batch 2 should be promoted from buffer")
		batch = streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(2), batch.Number())
	})

	t.Run("invalid headBatch is discarded and next candidate promoted from buffer", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(2))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// Create batch 1 with INVALID L1 origin hash (using a hash that won't match)
		invalidHash := common.HexToHash("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")
		_, _, _, espTxnInBlockInvalid := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, invalidHash,
		)
		state.AddEspressoTransactionData(0, namespace, espTxnInBlockInvalid)

		// Create batch 1 with VALID L1 origin (using the correct hash)
		_, _, _, espTxnInBlockValid := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, state.FinalizedL1.Hash,
		)
		state.AddEspressoTransactionData(1, namespace, espTxnInBlockValid)

		// Update to fetch both batches
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// HasNext should drop the invalid batch and find the valid one
		require.True(t, streamer.HasNext(ctx), "valid batch should be available after invalid is dropped")

		// Next should return the valid batch
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())
	})
}

// TestStreamerMultipleBatchesSameNumber tests handling of multiple batches with
// the same batch number but different validity.
func TestStreamerMultipleBatchesSameNumber(t *testing.T) {
	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	t.Run("invalid batches dropped during HasNext iteration until valid found", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(3))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		invalidHash := common.HexToHash("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")

		// Create 3 batches all with number 1:
		// Batch A: invalid L1 origin hash
		_, _, _, espTxnA := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, invalidHash,
		)
		state.AddEspressoTransactionData(0, namespace, espTxnA)

		// Batch B: invalid L1 origin hash
		_, _, _, espTxnB := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, invalidHash,
		)
		state.AddEspressoTransactionData(1, namespace, espTxnB)

		// Batch C: valid L1 origin hash
		_, _, _, espTxnC := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, state.FinalizedL1.Hash,
		)
		state.AddEspressoTransactionData(2, namespace, espTxnC)

		// Update to fetch all batches
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// HasNext should return true (found valid batch C)
		require.True(t, streamer.HasNext(ctx))

		// Next should return batch C (the valid one)
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())

		// BatchPos should have advanced to 2
		require.Equal(t, uint64(2), streamer.BatchPos)
	})

	t.Run("BatchPos does NOT advance when all candidates for batch number are invalid", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(4))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		invalidHash := common.HexToHash("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")

		// Create 3 batches all with number 1, ALL with invalid L1 origins
		for i := 0; i < 3; i++ {
			_, _, _, espTxn := state.CreateEspressoTxnDataWithL1Origin(
				ctx, namespace, rng, chainID, 1, chainSigner,
				state.FinalizedL1.Number, invalidHash,
			)
			state.AddEspressoTransactionData(uint64(i), namespace, espTxn)
		}

		// Update to fetch all batches
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// All candidates should be dropped (BatchDrop)
		// HasNext should return false (no valid batch available)
		require.False(t, streamer.HasNext(ctx))

		// BatchPos should still be 1 (NOT advanced)
		require.Equal(t, uint64(1), streamer.BatchPos)
	})

	t.Run("first valid batch returned when multiple valid candidates exist", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(5))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// Create 2 valid batches for number 1 with different hashes
		_, espBatch1, _, espTxn1 := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, state.FinalizedL1.Hash,
		)
		state.AddEspressoTransactionData(0, namespace, espTxn1)
		firstBatchHash := espBatch1.Hash()

		_, _, _, espTxn2 := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, state.FinalizedL1.Hash,
		)
		state.AddEspressoTransactionData(1, namespace, espTxn2)

		// Update to fetch both batches
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// HasNext should return true
		require.True(t, streamer.HasNext(ctx))

		// Next should return the first valid batch (insertion order matters)
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())
		require.Equal(t, firstBatchHash, batch.Hash(), "first inserted batch should be returned")

		// Second batch should be skipped as BatchPast
		require.False(t, streamer.HasNext(ctx), "no more batches should be available")
	})
}

// TestStreamerBufferCapacityAndSkipPos tests the skip position mechanism when the buffer fills up.
func TestStreamerBufferCapacityAndSkipPos(t *testing.T) {
	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	t.Run("skipPos not overwritten across multiple fetch ranges", func(t *testing.T) {
		// Regression test: when the Update loop iterates through multiple
		// HotShot block ranges, hitting ErrAtCapacity in a later range must
		// NOT overwrite skipPos set by an earlier range. Otherwise the rewind
		// skips the earlier range's batches permanently.
		//
		// Scenario:
		//   - Enough batches (starting from 2, skipping 1) are placed to fill
		//     the buffer, plus an extra fetch range worth of batches beyond it.
		//   - The extra batches are dropped because the buffer is full.
		//     skipPos should record the earliest range where capacity was hit.
		//   - Batch 1 is injected later, consumed, and triggers a rewind.
		//   - After draining the buffer, the next batch must come from the
		//     re-fetched overflow. If skipPos was overwritten to a later range
		//     start, the rewind won't go far enough and those batches are lost.
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state := NewMockStreamerSource()
		logger := new(NoOpLogger)

		streamer := espresso.NewEspressoStreamer(
			namespace,
			state,
			state,
			state,
			state,
			logger,
			derive.CreateEspressoBatchUnmarshaler(signerAddress),
			50*time.Millisecond,
			0,
			0, // originBatchPos=0, so BatchPos starts at 1
		)

		rng := rand.New(rand.NewSource(99))

		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// Place enough batches to fill the buffer and overflow by one full
		// fetch range. Batch 1 is intentionally missing so HasNext stays
		// false, forcing the Update loop to keep iterating across ranges.
		totalBatches := int(espresso.BatchBufferCapacity) + int(espresso.HOTSHOT_BLOCK_FETCH_LIMIT)
		for i := 0; i < totalBatches; i++ {
			_, _, _, espTxn := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, uint64(i+2), chainSigner)
			state.AddEspressoTransactionData(uint64(i), namespace, espTxn)
		}

		// Update processes all ranges. The buffer fills up partway through,
		// and all subsequent batches are dropped with ErrAtCapacity.
		err = streamer.Update(ctx)
		require.NoError(t, err)
		require.False(t, streamer.HasNext(ctx))

		// Inject batch 1 beyond all existing data.
		batch1Pos := uint64(totalBatches + 10)
		_, _, _, espTxn1 := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, 1, chainSigner)
		state.AddEspressoTransactionData(batch1Pos, namespace, espTxn1)

		// Fetch and consume batch 1 — triggers the rewind via skipPos.
		err = streamer.Update(ctx)
		require.NoError(t, err)
		require.True(t, streamer.HasNext(ctx))
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())

		// Drain the entire buffer of previously-buffered batches.
		firstOverflow := uint64(espresso.BatchBufferCapacity) + 2
		for expectedNum := uint64(2); expectedNum < firstOverflow; expectedNum++ {
			err = streamer.Update(ctx)
			require.NoError(t, err)
			require.True(t, streamer.HasNext(ctx), "expected batch %d to be available", expectedNum)
			batch = streamer.Next(ctx)
			require.NotNil(t, batch)
			require.Equal(t, expectedNum, batch.Number())
		}

		// The first batch that was dropped due to capacity must now be
		// recoverable via the rewind. If skipPos was overwritten to a later
		// range, this batch is permanently lost.
		err = streamer.Update(ctx)
		require.NoError(t, err)
		require.True(t, streamer.HasNext(ctx), "first overflow batch must be available after rewind — skipPos must preserve the earliest range")
		batch = streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, firstOverflow, batch.Number(), "first batch after buffer drain must not be skipped")
	})

	t.Run("skipPos set and rewind after Next drains buffer", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state := NewMockStreamerSource()
		logger := new(NoOpLogger)

		// Create streamer - after Refresh with SafeL2.Number=0, BatchPos becomes 1
		streamer := espresso.NewEspressoStreamer(
			namespace,
			state,
			state,
			state,
			state,
			logger,
			derive.CreateEspressoBatchUnmarshaler(signerAddress),
			50*time.Millisecond,
			0,
			0, // originBatchPos=0, so BatchPos starts at 1
		)

		rng := rand.New(rand.NewSource(6))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// Fill the buffer with out-of-order batches (batches 2, 3, 4, ... exceeding capacity)
		// We need batch 1 to be missing initially to fill the buffer
		for i := 0; i < int(espresso.BatchBufferCapacity)+5; i++ {
			// Create batches starting from 2 (skipping 1)
			_, _, _, espTxn := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, uint64(i+2), chainSigner)
			state.AddEspressoTransactionData(uint64(i), namespace, espTxn)
		}

		// Update - this should fill buffer and set skipPos
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// HasNext should return false (batch 1 is missing)
		require.False(t, streamer.HasNext(ctx))

		// Now add batch 1 at a later hotshot position
		laterHotshotPos := uint64(espresso.BatchBufferCapacity + 10)
		_, _, _, espTxn1 := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, 1, chainSigner)
		state.AddEspressoTransactionData(laterHotshotPos, namespace, espTxn1)

		// Update again to fetch batch 1
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// Now HasNext should return true
		require.True(t, streamer.HasNext(ctx))

		// Consume batch 1
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())

		// After consuming, the streamer should have rewound to re-fetch skipped batches
		// We should be able to get batch 2 now
		err = streamer.Update(ctx)
		require.NoError(t, err)

		require.True(t, streamer.HasNext(ctx))
		batch = streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(2), batch.Number())
	})

	t.Run("new batch for current BatchPos arrives when buffer full", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state := NewMockStreamerSource()
		logger := new(NoOpLogger)

		// Create streamer - after Refresh with SafeL2.Number=0, BatchPos becomes 1
		streamer := espresso.NewEspressoStreamer(
			namespace,
			state,
			state,
			state,
			state,
			logger,
			derive.CreateEspressoBatchUnmarshaler(signerAddress),
			50*time.Millisecond,
			0,
			0, // originBatchPos=0, so BatchPos starts at 1
		)

		rng := rand.New(rand.NewSource(7))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// Fill buffer with future batches (2, 3, 4, ...)
		for i := 0; i < int(espresso.BatchBufferCapacity); i++ {
			_, _, _, espTxn := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, uint64(i+2), chainSigner)
			state.AddEspressoTransactionData(uint64(i), namespace, espTxn)
		}

		// Update to fill buffer
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// HasNext should be false (batch 1 is missing)
		require.False(t, streamer.HasNext(ctx))

		// Now add batch 1 (the one we need)
		laterPos := uint64(espresso.BatchBufferCapacity + 1)
		_, _, _, espTxn1 := state.CreateEspressoTxnData(ctx, namespace, rng, chainID, 1, chainSigner)
		state.AddEspressoTransactionData(laterPos, namespace, espTxn1)

		// Update to get batch 1
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// Batch 1 should be assigned to headBatch directly (not buffered)
		require.True(t, streamer.HasNext(ctx))
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())
	})
}

// TestStreamerBatchOrderingDeterminism tests that the streamer processes batches
// deterministically when multiple batches have the same number - insertion order
// must be respected.
func TestStreamerBatchOrderingDeterminism(t *testing.T) {
	namespace := uint64(42)
	chainID := big.NewInt(int64(namespace))
	privateKeyString := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	chainSignerFactory, signerAddress, _ := crypto.ChainSignerFactoryFromConfig(&NoOpLogger{}, privateKeyString, "", "", opsigner.CLIConfig{})
	chainSigner := chainSignerFactory(chainID, common.Address{})

	t.Run("must wait for first-inserted batch to become decided before processing later ones", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(8))

		// Advance L1 to have two finalized blocks at heights 1 and 2
		// FinalizedL1 starts at 1
		state.AdvanceFinalizedL1() // Now at 2

		// Refresh state with only height 1 finalized (we'll pretend height 2 is not finalized yet)
		// We need to control what the streamer sees as finalized
		// After this refresh, BatchPos becomes 1
		l1Height1 := createL1BlockRef(1)
		err := streamer.Refresh(ctx, l1Height1, state.SafeL2.Number, state.SafeL2.L1Origin)
		require.NoError(t, err)

		// Insert batch a1 (number 1, L1 origin at height 2 - NOT finalized yet)
		l1Height2 := createL1BlockRef(2)
		_, espBatchA1, _, espTxnA1 := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			l1Height2.Number, l1Height2.Hash,
		)
		state.AddEspressoTransactionData(0, namespace, espTxnA1)
		a1Hash := espBatchA1.Hash()

		// Insert batch a2 (number 1, L1 origin at height 1 - IS finalized)
		_, _, _, espTxnA2 := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			l1Height1.Number, l1Height1.Hash,
		)
		state.AddEspressoTransactionData(1, namespace, espTxnA2)

		// Update to fetch both batches
		err = streamer.Update(ctx)
		require.NoError(t, err)

		// HasNext should return false - must wait for a1 (inserted first) to become decided
		// even though a2 is already valid
		require.False(t, streamer.HasNext(ctx), "should wait for first-inserted batch to become decided")

		// Now advance L1 finalized to height 2
		err = streamer.Refresh(ctx, l1Height2, state.SafeL2.Number, state.SafeL2.L1Origin)
		require.NoError(t, err)

		// HasNext should now return true
		require.True(t, streamer.HasNext(ctx))

		// Next should return a1 (the first-inserted batch)
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, uint64(1), batch.Number())
		require.Equal(t, a1Hash, batch.Hash(), "first-inserted batch should be returned")

		// a2 should subsequently be skipped as BatchPast
		require.False(t, streamer.HasNext(ctx), "second batch should be skipped")
	})

	t.Run("insertion order respected across multiple Update calls", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		state, streamer := setupStreamerTesting(namespace, signerAddress)
		rng := rand.New(rand.NewSource(9))

		// Refresh state - after this, BatchPos becomes 1
		syncStatus := state.SyncStatus()
		err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number, syncStatus.SafeL2.L1Origin)
		require.NoError(t, err)

		// First Update: insert batch a1 (number 1)
		_, espBatchA1, _, espTxnA1 := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, state.FinalizedL1.Hash,
		)
		state.AddEspressoTransactionData(0, namespace, espTxnA1)
		a1Hash := espBatchA1.Hash()

		err = streamer.Update(ctx)
		require.NoError(t, err)

		// Second Update: insert batch a2 (number 1, different hash)
		_, _, _, espTxnA2 := state.CreateEspressoTxnDataWithL1Origin(
			ctx, namespace, rng, chainID, 1, chainSigner,
			state.FinalizedL1.Number, state.FinalizedL1.Hash,
		)
		state.AddEspressoTransactionData(1, namespace, espTxnA2)

		err = streamer.Update(ctx)
		require.NoError(t, err)

		// HasNext should return true
		require.True(t, streamer.HasNext(ctx))

		// Next should return a1 (first inserted)
		batch := streamer.Next(ctx)
		require.NotNil(t, batch)
		require.Equal(t, a1Hash, batch.Hash(), "first-inserted batch should be returned")

		// a2 should be skipped as BatchPast
		require.False(t, streamer.HasNext(ctx))
	})
}
