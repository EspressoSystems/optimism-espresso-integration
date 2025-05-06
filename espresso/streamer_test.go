package espresso_test

import (
	"context"
	"encoding/binary"
	"errors"
	"log/slog"
	"math/big"
	"math/rand"
	"testing"
	"time"

	esp_client "github.com/EspressoSystems/espresso-network-go/client"
	esp_common "github.com/EspressoSystems/espresso-network-go/types"
	"github.com/ethereum-optimism/optimism/espresso"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	opsigner "github.com/ethereum-optimism/optimism/op-service/signer"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// TestNewEspressoStreamer tests that we can create a new EspressoStreamer
// without any panic being thrown.

func TestNewEspressoStreamer(t *testing.T) {
	_ = espresso.NewEspressoStreamer(
		1,
		nil,
		nil, nil, nil, derive.CreateEspressoBatchUnmarshaler(common.Address{}),
		50*time.Millisecond,
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

	EspTransactionData map[EspBlockAndNamespace]esp_client.TransactionsInBlock
	LatestEspHeight    uint64
}

// AdvanceFinalizedL1ByNBlocks advances the FinalizedL1 block reference by n blocks.
func (m *MockStreamerSource) AdvanceFinalizedL1ByNBlocks(n uint) {
	m.FinalizedL1 = createL1BlockRef(m.FinalizedL1.Number + uint64(n))
}

// AdvanceFinalizedL1 advances the FinalizedL1 block reference by one block.
func (m *MockStreamerSource) AdvanceFinalizedL1() {
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

func (m *MockStreamerSource) AddEspressoTransactionData(height, namespace uint64, txData esp_client.TransactionsInBlock) {
	if m.EspTransactionData == nil {
		m.EspTransactionData = make(map[EspBlockAndNamespace]esp_client.TransactionsInBlock)
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

func (m *MockStreamerSource) FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (esp_client.TransactionsInBlock, error) {
	if m.LatestEspHeight < blockHeight {
		return esp_client.TransactionsInBlock{}, ErrNotFound
	}

	// NOTE: if this combination is not found, we will end up returning an
	//       empty TransactionsInBlock, which is intentional.  It will allow
	//       the consumer to know that this block exists, but no transactions
	//       for the requested namespace exist.
	return m.EspTransactionData[BlockAndNamespace(blockHeight, namespace)], nil
}

// Espresso Light Client implementation
var _ espresso.LightClientReaderInterface = (*MockStreamerSource)(nil)

// NoOpLogger is a no-op implementation of the log.Logger interface.
// It is used to pass a non-nil logger to the EspressoStreamer without
// producing any output.
type NoOpLogger struct{}

var _ log.Logger = (*NoOpLogger)(nil)

func (l *NoOpLogger) With(ctx ...interface{}) log.Logger                   { return l }
func (l *NoOpLogger) New(ctx ...interface{}) log.Logger                    { return l }
func (l *NoOpLogger) Log(level slog.Level, msg string, ctx ...interface{}) {}
func (l *NoOpLogger) Trace(msg string, ctx ...interface{})                 {}
func (l *NoOpLogger) Debug(msg string, ctx ...interface{})                 {}
func (l *NoOpLogger) Info(msg string, ctx ...interface{})                  {}
func (l *NoOpLogger) Warn(msg string, ctx ...interface{})                  {}
func (l *NoOpLogger) Error(msg string, ctx ...interface{})                 {}
func (l *NoOpLogger) Crit(msg string, ctx ...interface{})                  { panic("critical error") }
func (l *NoOpLogger) Write(level slog.Level, msg string, attrs ...any)     {}
func (l *NoOpLogger) Enabled(ctx context.Context, level slog.Level) bool   { return true }
func (l *NoOpLogger) Handler() slog.Handler                                { return nil }

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
func setupStreamerTesting(namespace uint64, batcherAddress common.Address) (*MockStreamerSource, espresso.EspressoStreamer[derive.EspressoBatch]) {
	state := new(MockStreamerSource)
	state.AdvanceFinalizedL1()

	logger := new(NoOpLogger)
	streamer := espresso.NewEspressoStreamer(
		namespace,
		state,
		state,
		state,
		logger,
		derive.CreateEspressoBatchUnmarshaler(batcherAddress),
		50*time.Millisecond,
	)

	return state, streamer
}

// createSingularBatch creates a mock SingularBatch for testing purposes
// containing the provided number of transactions.
// It is generated using a random number generator to create the transactions
// contained within.  Everything else is derived from the provided parameters
// for repeatability.
func (m *MockStreamerSource) createSingularBatch(rng *rand.Rand, txCount int, chainID *big.Int, l2Height uint64) *derive.SingularBatch {
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
		EpochNum:     rollup.Epoch(m.FinalizedL1.Number),
		EpochHash:    m.FinalizedL1.Hash,
		Timestamp:    l2Height,
		Transactions: txsEncoded,
	}
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
func createEspressoTransaction(ctx context.Context, batch *derive.EspressoBatch, namespace uint64, chainSigner crypto.ChainSigner) *esp_common.Transaction {
	tx, err := batch.ToEspressoTransaction(ctx, namespace, chainSigner)
	if have, want := err, error(nil); have != want {
		panic(err)
	}

	return tx
}

// createTransactionsInBlock creates a mock TransactionsInBlock for testing purposes
// containing the provided Espresso transaction.
func createTransactionsInBlock(tx *esp_common.Transaction) esp_client.TransactionsInBlock {
	return esp_client.TransactionsInBlock{
		Transactions: []esp_common.Bytes{tx.Payload},
	}
}

// CreateEspressoTxnData creates a mock Espresso transaction data set
// for testing purposes. It generates a test SingularBatch, and takes it
// through the steps of getting all the way to an Espresso transaction in block.
// Every intermediate step is returned for inspection / utilization in tests.
func (m *MockStreamerSource) CreateEspressoTxnData(
	ctx context.Context,
	namespace uint64,
	rng *rand.Rand,
	chainID *big.Int,
	l2Height uint64,
	chainSigner crypto.ChainSigner,
) (*derive.SingularBatch, *derive.EspressoBatch, *esp_common.Transaction, esp_client.TransactionsInBlock) {
	txCount := rng.Intn(10)
	batch := m.createSingularBatch(rng, txCount, chainID, l2Height)
	espBatch := createEspressoBatch(batch)
	espTxn := createEspressoTransaction(ctx, espBatch, namespace, chainSigner)
	espTxnInBlock := createTransactionsInBlock(espTxn)

	return batch, espBatch, espTxn, espTxnInBlock
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
	updated, err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number)
	if have, want := updated, false; have != want {
		t.Fatalf("failed to refresh streamer state:\nhave:\n\t%v\nwant:\n\t%v\n", updated, want)
	}

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
	const N = 10

	for i := 0; i < N; i++ {
		// update the state of our streamer
		syncStatus := state.SyncStatus()
		_, err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number)

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

		if have, want := batchFromEsp, (*derive.EspressoBatch)(nil); have == want {
			t.Fatalf("unexpectedly did not received batch from streamer:\nhave:\n\t%v\nwant:\n\t%v\n", have, want)
		}

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
	const N = 10

	var batches []*derive.SingularBatch

	for i := 0; i < N; i++ {
		// update the state of our streamer
		syncStatus := state.SyncStatus()
		_, err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number)

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
		batches = append(batches, batch)
	}

	// Update the state of our streamer
	if have, want := streamer.Update(ctx), error(nil); !errors.Is(have, want) {
		t.Fatalf("failed to update streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	for i := 0; i < N; i++ {
		batch := batches[i]

		batchFromEsp := streamer.Next(ctx)

		if have, want := batchFromEsp, (*derive.EspressoBatch)(nil); have == want {
			t.Fatalf("unexpectedly did not received batch from streamer:\nhave:\n\t%v\nwant:\n\t%v\n", have, want)
		}

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
	_, err := streamer.Refresh(ctx, syncStatus.FinalizedL1, syncStatus.SafeL2.Number)

	if have, want := err, error(nil); have != want {
		t.Fatalf("failed to refresh streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	const N = 10
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

	// Update the state of our streamer
	if have, want := streamer.Update(ctx), error(nil); !errors.Is(have, want) {
		t.Fatalf("failed to update streamer state encountered error:\nhave:\n\t\"%v\"\nwant:\n\t\"%v\"\n", have, want)
	}

	{

		for i := 0; i < N; i++ {
			batch := batches[i]
			batchFromEsp := streamer.Next(ctx)
			if have, want := batchFromEsp, (*derive.EspressoBatch)(nil); have == want {
				t.Fatalf("unexpectedly did not received batch from streamer:\nhave:\n\t%v\ndo not want:\n\t%v\n", have, want)
			}

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
