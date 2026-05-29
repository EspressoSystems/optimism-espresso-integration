// Replay-DoS micro-benchmark for the derivation pipeline.
//
// Background
// ----------
// celo-org/optimism#445 introduces event-based batch authentication for
// Espresso. A reviewer (piersy) raised the concern that, post-fork, any
// account can re-submit a previously-authenticated batch (the
// authentication check is just `authenticatedHashes[batchHash]`), which
// trivially passes for any replay of real Espresso-batcher data. For altDA
// the per-replay L1 cost is ~$0.10. The question is: how much derivation
// work does this impose on every honest node per replayed batch?
//
// What this benchmark does
// ------------------------
// We build a single "representative" OP-Mainnet-style compressed channel:
//   - a span batch containing ~5 L2 blocks
//   - ~30 transactions per block (close to OP Mainnet average)
//   - zlib compression
// We wrap the resulting channel bytes into one L1 batcher transaction whose
// calldata is the *exact* payload a real altDA replay would surface to the
// derivation pipeline. (We use the calldata path rather than the altDA path
// because the post-fetch work is identical and this lets us avoid stubbing
// the DA layer.)
//
// We then instantiate the *full* derivation pipeline via
// derive.NewDerivationPipeline (the same constructor op-node uses in
// production) wired against in-memory stubs for L1 / L2. Each benchmark
// iteration advances the L1 head by one block — each L1 block contains the
// same replayed batcher tx — and drives the pipeline until it has processed
// the block. The pipeline drops the batch on the parent-hash check inside
// checkSpanBatchPrefix (or checkSingularBatch), but only *after* having
// done the full per-replay derivation work: frame parsing, channel
// reassembly, decompression, RLP decode of the RawSpanBatch, recoverV and
// fullTxs reconstruction of every transaction.
//
// Why this is conservative (overestimates attacker cost)
// ------------------------------------------------------
// We use the full pipeline construction, not a hand-rolled chain of
// stages. We do not bypass any cache. The pipeline pays the cost of
// resolving each new L1 origin, advancing the L1Traversal, opening a
// CalldataSource, parsing frames, reading the channel, decoding the batch,
// and propagating EOF back up. This includes a non-trivial amount of
// scaffolding work that a real attacker scenario only pays once. The pure
// "wasted derivation work per replay" is some *smaller* number than what
// this benchmark reports.
//
// How to run
// ----------
//   go test -run x -bench BenchmarkReplayDoS -benchtime=1000x \
//       ./op-node/benchmarks/...
//
// `-benchtime=Nx` controls how many replay iterations are processed. For a
// million iterations use `-benchtime=1000000x`. Go's `b.ReportMetric` is
// used to surface per-replay timing in human-friendly units. The benchmark
// also prints derived attacker-economics numbers.

package benchmarks

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	gethparams "github.com/ethereum/go-ethereum/params"

	altda "github.com/ethereum-optimism/optimism/op-alt-da"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	derparams "github.com/ethereum-optimism/optimism/op-node/rollup/derive/params"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
)

// ---------------------------------------------------------------------------
// Synthetic-batch construction
// ---------------------------------------------------------------------------

// buildReplayCalldata produces a byte slice that is a valid batcher-tx
// calldata (DerivationVersion0 ++ frames) containing one span batch with
// `blocksPerBatch` L2 blocks and `txsPerBlock` transactions per block,
// compressed with zlib via SpanChannelOut. The returned bytes are what a
// real OP Mainnet batcher would post to L1 calldata (modulo content), or
// what an altDA layer would surface to op-node after GetInput.
//
// Sizing knobs reflect OP-Mainnet-ish realities:
//
//   - blockTime=2s, 5 blocks → 10s of span
//   - txsPerBlock=30 ≈ OP Mainnet average over the last year
//   - zlib (the default OP batcher compressor)
//   - frame size cap 120_000 bytes (op-batcher default for calldata mode)
func buildReplayCalldata(t testing.TB, rollupCfg *rollup.Config, blocksPerBatch, txsPerBlock int) []byte {
	t.Helper()
	chainSpec := rollup.NewChainSpec(rollupCfg)
	// Target output size of "huge" so the compressor never closes early; we
	// fill it with exactly the data we generated.
	out, err := derive.NewSpanChannelOut(targetOutput_huge, derive.Zlib, chainSpec)
	if err != nil {
		t.Fatalf("NewSpanChannelOut: %v", err)
	}

	rng := rand.New(rand.NewSource(0xDEADBEEF))
	baseTime := uint64(1_700_000_000)
	for i := 0; i < blocksPerBatch; i++ {
		ts := baseTime + uint64(i)*rollupCfg.BlockTime
		blk, berr := buildSyntheticL2Block(rollupCfg, rng, txsPerBlock, ts, uint64(i+1))
		if berr != nil {
			t.Fatalf("buildSyntheticL2Block: %v", berr)
		}
		if _, err := out.AddBlock(rollupCfg, blk); err != nil {
			t.Fatalf("AddBlock: %v", err)
		}
	}
	if err := out.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	var buf bytes.Buffer
	// derivation version prefix
	buf.WriteByte(derparams.DerivationVersion0)
	// emit all frames into the buffer; OutputFrame returns io.EOF once the
	// last frame has been written. Any other error (e.g.
	// ErrMaxFrameSizeTooSmall, a compressor error) would mean we silently
	// generated invalid calldata, which would make the pipeline reject
	// every frame and falsely report ~zero per-replay work.
	const maxFrameSize = 120_000
	for {
		_, err := out.OutputFrame(&buf, maxFrameSize)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("OutputFrame: %v", err)
		}
	}
	return buf.Bytes()
}

// buildSyntheticL2Block returns a *types.Block representing a single L2
// block with `txsPerBlock` random non-deposit txs plus a synthetic
// L1InfoDeposit prepended (so that channel_out.BlockToSingularBatch can
// parse it). We use a non-nil BaseFee on the synthetic L1 info to avoid
// the panic in solabi.WriteUint256(nil).
func buildSyntheticL2Block(rollupCfg *rollup.Config, rng *rand.Rand, txsPerBlock int, timestamp, l1Num uint64) (*types.Block, error) {
	batch := derive.RandomSingularBatch(rng, txsPerBlock, rollupCfg.L2ChainID)
	batch.Timestamp = timestamp

	mockL1 := &testutils.MockBlockInfo{
		InfoNum:         l1Num,
		InfoHash:        batch.EpochHash,
		InfoTime:        timestamp,
		InfoBaseFee:     big.NewInt(1_000_000_000),
		InfoBlobBaseFee: big.NewInt(1),
		InfoReceiptRoot: types.EmptyRootHash,
		InfoRoot:        types.EmptyRootHash,
		InfoGasLimit:    30_000_000,
	}
	l1InfoTx, err := derive.L1InfoDeposit(rollupCfg, gethparams.MainnetChainConfig, eth.SystemConfig{}, 0, mockL1, timestamp)
	if err != nil {
		return nil, fmt.Errorf("L1InfoDeposit: %w", err)
	}

	txs := make([]*types.Transaction, 0, 1+len(batch.Transactions))
	txs = append(txs, types.NewTx(l1InfoTx))
	for i, opaque := range batch.Transactions {
		var tx types.Transaction
		if err := tx.UnmarshalBinary(opaque); err != nil {
			return nil, fmt.Errorf("decode tx %d: %w", i, err)
		}
		txs = append(txs, &tx)
	}
	return types.NewBlockWithHeader(&types.Header{
		ParentHash: batch.ParentHash,
		Time:       timestamp,
	}).WithBody(types.Body{Transactions: txs}), nil
}

// ---------------------------------------------------------------------------
// Stub L1 / L2 sources
// ---------------------------------------------------------------------------

// staticL1 implements derive.L1Fetcher (L1BlockRefByLabel,
// L1BlockRefByNumber, L1BlockRefByHash, FetchReceipts,
// InfoAndTxsByHash, InfoByHash) over an in-memory parent-linked chain.
// All blocks have the same single batcher transaction in them. Receipts
// are empty (no system-config updates, no auth events; we configure the
// rollup pre-Espresso so the data source uses sender-based auth).
type staticL1 struct {
	mu       sync.Mutex
	byHash   map[common.Hash]eth.L1BlockRef
	byNumber []eth.L1BlockRef // index 0 is genesis
	headIdx  int              // current head; pipeline advances by 1 each iter

	// All blocks contain this same batcher tx.
	batcherTx *types.Transaction

	chainID *big.Int

	// Cached BlockInfo per hash, reused per call.
	infoCache map[common.Hash]eth.BlockInfo

	// Counters used for sanity-checking what the pipeline actually does.
	fetchReceiptsCalls    int
	infoAndTxsByHashCalls int
}

func newStaticL1(chainID *big.Int, batcherTx *types.Transaction, initialBlocks int) *staticL1 {
	l1 := &staticL1{
		byHash:    make(map[common.Hash]eth.L1BlockRef),
		batcherTx: batcherTx,
		chainID:   chainID,
		infoCache: make(map[common.Hash]eth.BlockInfo),
	}
	// Seed a deterministic genesis ref and grow the chain.
	for i := 0; i < initialBlocks; i++ {
		l1.appendBlock()
	}
	l1.headIdx = 0 // start at genesis
	return l1
}

func (l *staticL1) appendBlock() {
	idx := len(l.byNumber)
	var parent common.Hash
	if idx > 0 {
		parent = l.byNumber[idx-1].Hash
	}
	// Hash derived from index to be deterministic & unique.
	var h common.Hash
	h[0] = byte(idx)
	h[1] = byte(idx >> 8)
	h[2] = byte(idx >> 16)
	h[3] = byte(idx >> 24)
	// Sprinkle bytes so map distribution stays sane.
	h[31] = 0xAB
	ref := eth.L1BlockRef{
		Hash:       h,
		Number:     uint64(idx),
		ParentHash: parent,
		// Times pre-Espresso so sender-based auth path is exercised (matches
		// upstream Optimism behavior; the replay-DoS surface is identical
		// post-Espresso once the batchHash is in the cached set).
		Time: 1_700_000_000 + uint64(idx)*12,
	}
	l.byHash[h] = ref
	l.byNumber = append(l.byNumber, ref)
}

// AdvanceHead makes the pipeline see one new L1 block on the next Step()
// loop. It grows the chain by one block (with the same batcher tx).
func (l *staticL1) AdvanceHead() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.headIdx+1 >= len(l.byNumber) {
		l.appendBlock()
	}
	l.headIdx++
}

func (l *staticL1) head() eth.L1BlockRef {
	return l.byNumber[l.headIdx]
}

func (l *staticL1) L1BlockRefByLabel(_ context.Context, _ eth.BlockLabel) (eth.L1BlockRef, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.head(), nil
}

func (l *staticL1) L1BlockRefByNumber(_ context.Context, num uint64) (eth.L1BlockRef, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if num >= uint64(len(l.byNumber)) {
		// Beyond the current chain; pretend it doesn't exist yet.
		return eth.L1BlockRef{}, ethereumNotFound
	}
	if num > uint64(l.headIdx) {
		return eth.L1BlockRef{}, ethereumNotFound
	}
	return l.byNumber[num], nil
}

func (l *staticL1) L1BlockRefByHash(_ context.Context, h common.Hash) (eth.L1BlockRef, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	r, ok := l.byHash[h]
	if !ok {
		return eth.L1BlockRef{}, ethereumNotFound
	}
	return r, nil
}

func (l *staticL1) FetchReceipts(_ context.Context, h common.Hash) (eth.BlockInfo, types.Receipts, error) {
	info := l.infoFor(h)
	l.mu.Lock()
	l.fetchReceiptsCalls++
	l.mu.Unlock()
	return info, types.Receipts{}, nil
}

func (l *staticL1) InfoAndTxsByHash(_ context.Context, h common.Hash) (eth.BlockInfo, types.Transactions, error) {
	info := l.infoFor(h)
	l.mu.Lock()
	l.infoAndTxsByHashCalls++
	l.mu.Unlock()
	return info, types.Transactions{l.batcherTx}, nil
}

func (l *staticL1) InfoByHash(_ context.Context, h common.Hash) (eth.BlockInfo, error) {
	return l.infoFor(h), nil
}

func (l *staticL1) infoFor(h common.Hash) eth.BlockInfo {
	l.mu.Lock()
	defer l.mu.Unlock()
	if info, ok := l.infoCache[h]; ok {
		return info
	}
	ref, ok := l.byHash[h]
	if !ok {
		// Best-effort: return a zero info; the pipeline should not request
		// info for blocks it didn't first hear about via traversal.
		return &testutils.MockBlockInfo{}
	}
	info := &testutils.MockBlockInfo{
		InfoHash:        ref.Hash,
		InfoParentHash:  ref.ParentHash,
		InfoNum:         ref.Number,
		InfoTime:        ref.Time,
		InfoReceiptRoot: types.EmptyRootHash,
		InfoRoot:        types.EmptyRootHash,
		InfoBaseFee:     big.NewInt(1_000_000_000),
		InfoGasLimit:    30_000_000,
	}
	l.infoCache[h] = info
	return info
}

// staticL2 implements derive.L2Source / SystemConfigL2Fetcher with a fixed
// safe head whose parent hash is guaranteed *not* to match the synthetic
// batch's ParentHash. That makes every replayed batch fail the
// checkSingularBatch / checkSpanBatchPrefix parent-hash check after full
// decoding work has been done.
type staticL2 struct {
	safe        eth.L2BlockRef
	sysCfg      eth.SystemConfig
	chainID     *big.Int
	emptyOutput *eth.OutputV0
}

func newStaticL2(chainID *big.Int, batcherAddr common.Address, l1Genesis eth.L1BlockRef) *staticL2 {
	var safeHash common.Hash
	safeHash[0] = 0xC0
	safeHash[1] = 0xFF
	safeHash[2] = 0xEE
	var parentHash common.Hash
	parentHash[0] = 0xDE
	parentHash[1] = 0xAD
	safe := eth.L2BlockRef{
		Hash:           safeHash,
		Number:         1,
		ParentHash:     parentHash,
		Time:           l1Genesis.Time + 2,
		L1Origin:       l1Genesis.ID(),
		SequenceNumber: 0,
	}
	return &staticL2{
		safe:    safe,
		chainID: chainID,
		sysCfg: eth.SystemConfig{
			BatcherAddr: batcherAddr,
			Overhead:    eth.Bytes32{},
			Scalar:      eth.Bytes32{},
			GasLimit:    30_000_000,
		},
	}
}

func (l *staticL2) L2BlockRefByLabel(_ context.Context, _ eth.BlockLabel) (eth.L2BlockRef, error) {
	return l.safe, nil
}
func (l *staticL2) L2BlockRefByHash(_ context.Context, _ common.Hash) (eth.L2BlockRef, error) {
	return l.safe, nil
}
func (l *staticL2) L2BlockRefByNumber(_ context.Context, _ uint64) (eth.L2BlockRef, error) {
	return l.safe, nil
}
func (l *staticL2) SystemConfigByL2Hash(_ context.Context, _ common.Hash) (eth.SystemConfig, error) {
	return l.sysCfg, nil
}
func (l *staticL2) PayloadByHash(_ context.Context, _ common.Hash) (*eth.ExecutionPayloadEnvelope, error) {
	return nil, ethereumNotFound
}
func (l *staticL2) PayloadByNumber(_ context.Context, _ uint64) (*eth.ExecutionPayloadEnvelope, error) {
	return nil, ethereumNotFound
}

// staticDepSet returns an empty chain set; the pipeline only consults
// DependencySet when interop is active, which we leave disabled.
type staticDepSet struct{}

func (staticDepSet) Chains() []eth.ChainID { return nil }

// nopBlobs returns an error for any blob request; with EcotoneTime unset
// the pipeline never asks for blobs.
type nopBlobs struct{}

func (nopBlobs) GetBlobsByHash(_ context.Context, _ uint64, _ []common.Hash) ([]*eth.Blob, error) {
	return nil, fmt.Errorf("nopBlobs: blob fetch not expected")
}

// nopMetrics satisfies the pipeline's Metrics interface with empty methods.
type nopMetrics struct{}

func (nopMetrics) RecordL1Ref(string, eth.L1BlockRef) {}
func (nopMetrics) RecordL2Ref(string, eth.L2BlockRef) {}
func (nopMetrics) RecordChannelInputBytes(int)        {}
func (nopMetrics) RecordHeadChannelOpened()           {}
func (nopMetrics) RecordChannelTimedOut()             {}
func (nopMetrics) RecordFrame()                       {}
func (nopMetrics) RecordDerivedBatches(string)        {}
func (nopMetrics) SetDerivationIdle(bool)             {}
func (nopMetrics) RecordPipelineReset()               {}

// ethereumNotFound is the standard not-found error the pipeline expects.
// l1_traversal.go uses errors.Is(err, ethereum.NotFound) to detect it.
var ethereumNotFound = ethereum.NotFound

// ---------------------------------------------------------------------------
// Pipeline driver
// ---------------------------------------------------------------------------

// signBatcherTx signs a DynamicFeeTx that posts the given calldata to the
// batch inbox address from the batcher's address. This is exactly the
// shape of a real OP Mainnet batcher submission (calldata mode), and is
// also what an altDA replay surfaces to the data-source layer post-fetch.
func signBatcherTx(t testing.TB, chainID *big.Int, batcherKey *ecdsa.PrivateKey, inbox common.Address, calldata []byte, nonce uint64) *types.Transaction {
	signer := types.NewCancunSigner(chainID)
	tx, err := types.SignNewTx(batcherKey, signer, &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: big.NewInt(2 * gethparams.GWei),
		GasFeeCap: big.NewInt(30 * gethparams.GWei),
		Gas:       10_000_000,
		To:        &inbox,
		Value:     big.NewInt(0),
		Data:      calldata,
	})
	if err != nil {
		t.Fatalf("sign tx: %v", err)
	}
	return tx
}

// BenchmarkReplayDoS measures derivation work per replayed batch.
//
// Methodology:
//  1. Build one realistic-size compressed channel (a span batch of N L2
//     blocks * M txs/block) and embed it into one signed batcher tx.
//  2. Configure a rollup that puts that tx into every L1 block.
//  3. Stand up the full derivation pipeline (NewDerivationPipeline) with
//     in-memory L1/L2 stubs and reset it once.
//  4. Each benchmark iteration: advance the L1 head by one block and drive
//     pipeline.Step() in a loop until it reports io.EOF (i.e. it has
//     finished the L1 block). The pipeline parses the frame, reassembles
//     and decompresses the channel, RLP-decodes the span batch, expands
//     it via recoverV/fullTxs, then drops it on the parent-hash check.
//
// Reported metrics:
//   - ns/op           : nanoseconds per replayed L1 block (= per replay)
//   - kB-replayed/op  : decompressed batch payload bytes per replay
//   - replays/$       : derived from $0.10 per altDA replay
//   - CPU-pct/replay@12s : sustained CPU pct per L1 block on one core
//     assuming the attacker submits one replay per L1 block (12s)
//
// BenchmarkReplayDoS runs a 2×2 matrix:
//
//	{Celo-load, OP-load} × {altDA, no-altDA}
//
// where:
//
//   - Celo-load = ~384 L2 blocks × ~4 txs/block ≈ 1500 L2 txs per channel,
//     based on Celo Mainnet L2 traffic surveyed 2026-05-29.
//
//   - OP-load   = ~150 L2 blocks × ~32 txs/block ≈ 4800 L2 txs per channel,
//     based on OP Mainnet L2 traffic surveyed 2026-05-29.
//
//   - altDA    : attacker rebroadcasts a ~2 KB commitment-only L1 calldata
//     tx; op-node fetches the payload from the DA layer (cost
//     amortised across honest nodes). Real Celo altDA replay
//     tx cost on L1: ~$0.09 at today's ~0.5 gwei gas.
//
//   - no-altDA : attacker rebroadcasts a 5-blob L1 tx carrying the full
//     ~637 KB compressed payload directly. Real OP Mainnet
//     5-blob batch tx cost on L1: ~$0.13 at today's blob fees.
//
// Per-replay *derivation work* depends only on the payload (post-fetch),
// not on whether it came from blobs, calldata, or altDA. The DA mode only
// changes the attacker's L1 cost, which we plug into the reporting at the
// end (ms-CPU per attacker dollar).
//
// Conservatism caveat: RandomSingularBatch produces incompressible random
// tx data, so the synthetic *compressed* sizes shown by the benchmark are
// much larger than what real batchers post for the same tx count
// (Celo-load: ~970 KB synthetic vs ~2.6 KB real; OP-load: ~3 MB synthetic
// vs ~637 KB real). The work measured (decompression + RLP decode +
// recoverV + fullTxs) scales with tx count and uncompressed RLP size,
// which the benchmark *faithfully* reflects, so the ns/op numbers are an
// upper bound on per-replay derivation cost for the configured tx count.
func BenchmarkReplayDoS(b *testing.B) {
	const (
		// Cost in USD that an attacker pays on L1 to post one replay
		// transaction, at today's L1 gas (~0.5 gwei, ETH @ ~$2038).
		//
		// altDAReplayCost is the realised cost of a Celo-style altDA
		// commitment tx (gasUsed≈99,480 at 0.45 gwei = ~$0.09).
		altDAReplayCost = 0.09
		// noAltDAReplayCost is the realised cost of an OP-Mainnet-style
		// 5-blob batcher tx (blobGasUsed=655_360 + plain gas, paid at
		// today's blob basefee + 0.5 gwei = ~$0.13).
		noAltDAReplayCost = 0.13
	)

	cases := []struct {
		name             string
		blocksPerBatch   int
		txsPerBlock      int
		dollarsPerReplay float64
	}{
		// Celo-load × altDA — the PR scenario.
		{"celo_load_altDA", 384, 4, altDAReplayCost},
		// Celo-load × no-altDA — same payload, posted as 5-blob blob tx.
		{"celo_load_noAltDA", 384, 4, noAltDAReplayCost},
		// OP-load × altDA.
		{"op_load_altDA", 150, 32, altDAReplayCost},
		// OP-load × no-altDA.
		{"op_load_noAltDA", 150, 32, noAltDAReplayCost},
	}
	for _, tc := range cases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			runReplayDoS(b, tc.blocksPerBatch, tc.txsPerBlock, tc.dollarsPerReplay)
		})
	}
}

func runReplayDoS(b *testing.B, blocksPerBatch, txsPerBlock int, dollarsPerReplay float64) {
	chainID := big.NewInt(10) // OP Mainnet
	zeroTime := uint64(0)
	rollupCfg := &rollup.Config{
		L2ChainID: chainID,
		L1ChainID: big.NewInt(1),
		BlockTime: 2,
		Genesis: rollup.Genesis{
			L1:     eth.BlockID{Number: 0},
			L2:     eth.BlockID{Number: 0},
			L2Time: 1_700_000_000,
			SystemConfig: eth.SystemConfig{
				GasLimit: 30_000_000,
			},
		},
		BatchInboxAddress: common.HexToAddress("0xbeef000000000000000000000000000000000001"),
		// Activate Delta (so span batches are accepted by the channel-in
		// reader). We leave Ecotone, Espresso, and Holocene inactive,
		// which keeps the data source on the calldata path and selects
		// the legacy BatchQueue. Note: both BatchQueue and BatchStage
		// short-circuit on the parent-hash mismatch inside
		// checkSpanBatchPrefix before any PayloadByNumber probe, so the
		// choice of which path is selected does not measurably affect
		// per-replay cost — both pay the full decompression + RLP decode
		// + recoverV/fullTxs work upstream of the validity check.
		DeltaTime:         &zeroTime,
		MaxSequencerDrift: 600,
		// Very large SeqWindowSize so the benchmark never trips the empty-
		// batch fallback path inside the BatchQueue (which is unrelated to
		// the per-replay work we want to measure).
		SeqWindowSize:         1_000_000_000,
		ChannelTimeoutBedrock: 300,
	}

	// Batcher key & address.
	batcherKey, err := crypto.GenerateKey()
	if err != nil {
		b.Fatalf("generate key: %v", err)
	}
	batcherAddr := crypto.PubkeyToAddress(batcherKey.PublicKey)

	// Build the synthetic representative batch calldata once.
	calldata := buildReplayCalldata(b, rollupCfg, blocksPerBatch, txsPerBlock)
	b.Logf("synthetic batch: blocks=%d txs/block=%d calldata=%d bytes (zlib compressed)",
		blocksPerBatch, txsPerBlock, len(calldata))

	// Sign one batcher tx; we replay this identical tx in every L1 block.
	batcherTx := signBatcherTx(b, big.NewInt(1) /* L1 chainID */, batcherKey,
		rollupCfg.BatchInboxAddress, calldata, 0)

	// --- L1 / L2 stubs ---
	l1 := newStaticL1(big.NewInt(1), batcherTx, 4) // seed a few blocks
	l2 := newStaticL2(chainID, batcherAddr, l1.byNumber[0])

	// Update genesis L1 hash so the rollup cfg's Genesis.L1 matches l1[0].
	rollupCfg.Genesis.L1 = l1.byNumber[0].ID()
	rollupCfg.Genesis.SystemConfig.BatcherAddr = batcherAddr

	// --- Pipeline ---
	logger := log.NewLogger(log.DiscardHandler())

	metrics := nopMetrics{}
	pipeline := derive.NewDerivationPipeline(
		logger,
		rollupCfg,
		staticDepSet{},
		l1,             // L1Fetcher
		nopBlobs{},     // L1BlobsFetcher (unused; no Ecotone)
		altda.Disabled, // altDA disabled; we use calldata path
		l2,             // L2Source
		metrics,
		false, // not managed by supervisor
		gethparams.MainnetChainConfig,
	)

	// Drive the initial reset to completion. The pipeline expects the
	// engine to be "reset" before it starts deriving.
	ctx := context.Background()
	pipeline.ConfirmEngineReset()

	// Drain reset steps until DerivationReady. Each Step() call during
	// reset returns nil error and advances dp.resetting by one stage when
	// the stage's Reset returns io.EOF.
	for i := 0; i < 64 && !pipeline.DerivationReady(); i++ {
		_, _ = pipeline.Step(ctx, l2.safe)
	}
	if !pipeline.DerivationReady() {
		b.Fatalf("pipeline failed to become ready after reset")
	}

	// --- Benchmark loop ---
	//
	// One iteration = one replayed batch on one new L1 block. We grow the
	// L1 chain by one block (containing the same replayed batcher tx) and
	// drive Step() until the pipeline's origin has caught up to the new
	// head. At that point all the per-replay work (frame parsing, channel
	// reassembly, decompression, RLP decode, span-batch derivation,
	// validity check) has been performed for this iteration's block.
	// attribsEmitted counts AttributesWithParent values returned by
	// pipeline.Step. For a faithful "replayed batch is dropped" benchmark,
	// this MUST remain zero: a non-zero count would mean the synthetic
	// batch is accidentally passing validity and producing attributes,
	// which would skew the measurement (we'd be paying attribute-build
	// cost on top of derive-then-drop cost).
	var attribsEmitted int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l1.AdvanceHead()
		targetHead := l1.head()

		stepLoop := 0
		for {
			// If the pipeline's stage origin has reached the new head, we've
			// fully processed this iteration's L1 block.
			if pipeline.Origin().Number >= targetHead.Number {
				break
			}
			attrib, err := pipeline.Step(ctx, l2.safe)
			stepLoop++
			if attrib != nil {
				attribsEmitted++
			}
			// io.EOF here means AdvanceL1Block found nothing new — which
			// shouldn't happen since we just grew the chain. We treat it as
			// a transient and re-loop.
			_ = err
			if stepLoop > 4096 {
				b.Fatalf("Step loop did not terminate; pipeline.Origin=%d targetHead=%d (err=%v)",
					pipeline.Origin().Number, targetHead.Number, err)
			}
		}
	}
	b.StopTimer()

	// --- Self-validation assertions ---
	//
	// These guard against silent regressions where the bench stops
	// measuring what we think it measures. If any of these fire, the
	// reported ns/op number is not "per-replay derivation cost" and must
	// not be trusted.
	if attribsEmitted != 0 {
		b.Fatalf("expected 0 attributes emitted (every replayed batch should be dropped), got %d", attribsEmitted)
	}
	if b.N > 0 {
		fetchPerOp := float64(l1.fetchReceiptsCalls) / float64(b.N)
		infoPerOp := float64(l1.infoAndTxsByHashCalls) / float64(b.N)
		// Expect ~1 of each per iteration once warmed up. The first
		// iteration processes 2 L1 blocks (initial origin 0 + advance to
		// 1), so for very small b.N we tolerate up to 2.0/op; for larger
		// b.N this should converge to ~1.0.
		if fetchPerOp > 2.0 || fetchPerOp < 0.5 {
			b.Fatalf("expected ~1 FetchReceipts per replay; got %.3f (possible pipeline misuse)", fetchPerOp)
		}
		if infoPerOp > 2.0 || infoPerOp < 0.5 {
			b.Fatalf("expected ~1 InfoAndTxsByHash per replay; got %.3f (possible pipeline misuse)", infoPerOp)
		}
	}

	// --- Reporting ---
	if b.N > 0 {
		nsPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N)
		b.ReportMetric(float64(len(calldata)), "calldata-bytes/op")
		b.ReportMetric(float64(l1.fetchReceiptsCalls)/float64(b.N), "FetchReceipts/op")
		b.ReportMetric(float64(l1.infoAndTxsByHashCalls)/float64(b.N), "InfoAndTxsByHash/op")

		// Attacker economics.
		const l1BlockSeconds = 12.0
		secPerOp := nsPerOp / 1e9
		cpuPctOneReplayPerL1Block := 100.0 * secPerOp / l1BlockSeconds
		replaysPerDollar := 1.0 / dollarsPerReplay
		b.ReportMetric(cpuPctOneReplayPerL1Block,
			"%CPU/replay-per-L1-block")
		b.ReportMetric(secPerOp*replaysPerDollar*1000.0,
			"ms-CPU/$attacker")
	}
}
