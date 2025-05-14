package derive

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"io"
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
)

type testTx struct {
	to      *common.Address
	dataLen int
	author  *ecdsa.PrivateKey
	good    bool
	value   int
}

func (tx *testTx) Create(t *testing.T, signer types.Signer, rng *rand.Rand) *types.Transaction {
	t.Helper()
	out, err := types.SignNewTx(tx.author, signer, &types.DynamicFeeTx{
		ChainID:   signer.ChainID(),
		Nonce:     0,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: big.NewInt(30 * params.GWei),
		Gas:       100_000,
		To:        tx.to,
		Value:     big.NewInt(int64(tx.value)),
		Data:      testutils.RandomData(rng, tx.dataLen),
	})
	require.NoError(t, err)
	return out
}

type calldataTest struct {
	name string
	txs  []testTx
}

// TestDataFromEVMTransactions creates some transactions from a specified template and asserts
// that DataFromEVMTransactions properly filters and returns the data from the authorized transactions
// inside the transaction set.
func TestDataFromEVMTransactions(t *testing.T) {
	inboxPriv := testutils.RandomKey()
	batcherPriv := testutils.RandomKey()
	cfg := &rollup.Config{
		L1ChainID:         big.NewInt(100),
		BatchInboxAddress: crypto.PubkeyToAddress(inboxPriv.PublicKey),
	}
	batcherAddr := crypto.PubkeyToAddress(batcherPriv.PublicKey)

	altInbox := testutils.RandomAddress(rand.New(rand.NewSource(1234)))
	altAuthor := testutils.RandomKey()

	testCases := []calldataTest{
		{
			name: "simple",
			txs:  []testTx{{to: &cfg.BatchInboxAddress, dataLen: 1234, author: batcherPriv, good: true}},
		},
		{
			name: "other inbox",
			txs:  []testTx{{to: &altInbox, dataLen: 1234, author: batcherPriv, good: false}}},
		{
			name: "other author",
			txs:  []testTx{{to: &cfg.BatchInboxAddress, dataLen: 1234, author: altAuthor, good: false}}},
		{
			name: "inbox is author",
			txs:  []testTx{{to: &cfg.BatchInboxAddress, dataLen: 1234, author: inboxPriv, good: false}}},
		{
			name: "author is inbox",
			txs:  []testTx{{to: &batcherAddr, dataLen: 1234, author: batcherPriv, good: false}}},
		{
			name: "unrelated",
			txs:  []testTx{{to: &altInbox, dataLen: 1234, author: altAuthor, good: false}}},
		{
			name: "contract creation",
			txs:  []testTx{{to: nil, dataLen: 1234, author: batcherPriv, good: false}}},
		{
			name: "empty tx",
			txs:  []testTx{{to: &cfg.BatchInboxAddress, dataLen: 0, author: batcherPriv, good: true}}},
		{
			name: "value tx",
			txs:  []testTx{{to: &cfg.BatchInboxAddress, dataLen: 1234, value: 42, author: batcherPriv, good: true}}},
		{
			name: "empty block", txs: []testTx{},
		},
		{
			name: "mixed txs",
			txs: []testTx{
				{to: &cfg.BatchInboxAddress, dataLen: 1234, value: 42, author: batcherPriv, good: true},
				{to: &cfg.BatchInboxAddress, dataLen: 3333, value: 32, author: altAuthor, good: false},
				{to: &cfg.BatchInboxAddress, dataLen: 2000, value: 22, author: batcherPriv, good: true},
				{to: &altInbox, dataLen: 2020, value: 12, author: batcherPriv, good: false},
			},
		},
		// TODO: test with different batcher key, i.e. when it's changed from initial config value by L1 contract
	}

	for i, tc := range testCases {
		rng := rand.New(rand.NewSource(int64(i)))
		signer := cfg.L1Signer()

		var expectedData []eth.Data
		var txs []*types.Transaction
		var receipts []*types.Receipt
		for i, tx := range tc.txs {
			transaction := tx.Create(t, signer, rng)
			txs = append(txs, transaction)
			receipts = append(receipts, &types.Receipt{
				Status: types.ReceiptStatusSuccessful,
				TxHash: transaction.Hash(),
			})

			if tx.good {
				expectedData = append(expectedData, txs[i].Data())
			}
		}

		out := DataFromEVMTransactions(DataSourceConfig{cfg.L1Signer(), cfg.BatchInboxAddress, false}, batcherAddr, txs, receipts, testlog.Logger(t, log.LevelCrit))
		require.ElementsMatch(t, expectedData, out)
	}

}

// TestAltDADataSourceL1FetcherErrors tests that the pipeline handles intermittent errors in
// L1Source correctly.
func TestCallDataSourceL1FetcherErrors(t *testing.T) {
	logger := testlog.Logger(t, log.LevelDebug)
	ctx := context.Background()

	rng := rand.New(rand.NewSource(1234))

	l1F := &testutils.MockL1Source{}

	// Create rollup genesis and config
	l1Time := uint64(2)
	refA := testutils.RandomBlockRef(rng)
	refA.Number = 1
	l1Refs := []eth.L1BlockRef{refA}
	refA0 := eth.L2BlockRef{
		Hash:           testutils.RandomHash(rng),
		Number:         0,
		ParentHash:     common.Hash{},
		Time:           refA.Time,
		L1Origin:       refA.ID(),
		SequenceNumber: 0,
	}
	batcherPriv := testutils.RandomKey()
	batcherAddr := crypto.PubkeyToAddress(batcherPriv.PublicKey)
	batcherInbox := common.Address{42}
	cfg := &rollup.Config{
		Genesis: rollup.Genesis{
			L1:     refA.ID(),
			L2:     refA0.ID(),
			L2Time: refA0.Time,
		},
		BlockTime:         1,
		SeqWindowSize:     20,
		BatchInboxAddress: batcherInbox,
	}

	signer := cfg.L1Signer()

	factory := NewDataSourceFactory(logger, cfg, l1F, nil, nil)

	parent := l1Refs[0]
	// create a new mock l1 ref
	ref := eth.L1BlockRef{
		Hash:       testutils.RandomHash(rng),
		Number:     parent.Number + 1,
		ParentHash: parent.Hash,
		Time:       parent.Time + l1Time,
	}

	input := testutils.RandomData(rng, 200)
	tx, err := types.SignNewTx(batcherPriv, signer, &types.DynamicFeeTx{
		ChainID:   signer.ChainID(),
		Nonce:     0,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: big.NewInt(30 * params.GWei),
		Gas:       100_000,
		To:        &batcherInbox,
		Value:     big.NewInt(int64(0)),
		Data:      input,
	})
	require.NoError(t, err)

	txs := []*types.Transaction{tx}
	receipts := types.Receipts{&types.Receipt{TxHash: tx.Hash(), Status: types.ReceiptStatusSuccessful}}

	l1F.ExpectFetchReceipts(ref.Hash, nil, nil, errors.New("Intermittent error"))
	l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), txs, nil)

	src, err := factory.OpenData(ctx, ref, batcherAddr)
	require.IsType(t, &CalldataSource{}, src, src)
	// Data source should still be opened correctly and attempt to fetch receipts
	require.NoError(t, err)

	l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), txs, nil)
	l1F.ExpectFetchReceipts(ref.Hash, nil, nil, errors.New("Intermittent error"))

	// Should fail because receipts are still not delivered
	_, err = src.Next(ctx)
	require.Error(t, err)

	l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), txs, nil)
	l1F.ExpectFetchReceipts(ref.Hash, nil, types.Receipts{}, nil)

	// Should fail because receipts do not match the transactions
	_, err = src.Next(ctx)
	require.Error(t, err)

	l1F.ExpectInfoAndTxsByHash(ref.Hash, testutils.RandomBlockInfo(rng), txs, nil)
	l1F.SetFetchReceipts(ref.Hash, nil, receipts, nil)

	// regular input is passed through
	data, err := src.Next(ctx)
	require.NoError(t, err)
	require.Equal(t, hexutil.Bytes(input), data)

	_, err = src.Next(ctx)
	require.ErrorIs(t, err, io.EOF)

	l1F.AssertExpectations(t)
}
