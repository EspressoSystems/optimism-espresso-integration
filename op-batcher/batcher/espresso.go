package batcher

import (
	// #cgo darwin,arm64 LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
	"C"

	"context"
	"encoding/json"
	"fmt"
	"time"

	espressoCommon "github.com/EspressoSystems/espresso-network-go/types"
	espressoVerification "github.com/EspressoSystems/espresso-network-go/verification"
	"github.com/ethereum-optimism/optimism/op-batcher/bindings"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

const exampleNamespace = 42

// TODO: Pull out to be re-used in op-node for derivation from Espresso
type Transaction struct {
	// Namespace of transaction to be published
	Namespace uint64
	// TODO: placeholder for sequencer's signature
	BatcherSignature []byte
	// Frames serialized as they would be for posting to L1 as calldata
	CallData []byte
}

// Parameters for transaction fetching loop, which waits for transactions
// to be sequenced on Espresso
const (
	transactionFetchTimeout  = 2 * time.Minute
	transactionFetchInterval = 100 * time.Millisecond
)

// Parameters for finality checking loop, which waits for merkle proof for
// Espresso transaction to be available from Light Client contract
const (
	finalityTimeout       = 2 * time.Minute
	finalityCheckInterval = 100 * time.Millisecond
)

func (l *BatchSubmitter) registerBatcher(ctx context.Context) error {
	if l.Attestation == nil {
		l.Log.Warn("Attestation is nil, skipping registration")
		return nil
	}

	batchInbox, err := bindings.NewBatchInbox(l.RollupConfig.BatchInboxAddress, l.L1Client)
	if err != nil {
		return fmt.Errorf("failed to create batch inbox contract bindings: %w", err)
	}

	// Decode the attestation off-chain to conserve gas
	decoded, err := batchInbox.DecodeAttestationTbs(&bind.CallOpts{}, l.Attestation)
	if err != nil {
		return fmt.Errorf("failed to decode attestation: %w", err)
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(l.Config.EphemeralPrivateKey, l.RollupConfig.L1ChainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	// Submit decoded attestation to batch inbox contract
	tx, err := batchInbox.RegisterSigner(txOpts, decoded.AttestationTbs, decoded.Signature)
	if err != nil {
		return fmt.Errorf("failed to create RegisterSigner transaction: %w", err)
	}

	candidate := txmgr.TxCandidate{
		TxData: tx.Data(),
		To:     tx.To(),
	}

	_, err = l.Txmgr.Send(ctx, candidate)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	return nil
}

func (t Transaction) toEspresso() espressoCommon.Transaction {
	payload := append(t.BatcherSignature, t.CallData...)
	return espressoCommon.Transaction{
		Namespace: t.Namespace,
		Payload:   payload,
	}
}

func (l *BatchSubmitter) waitForFinality(height uint64, rawHeader json.RawMessage, header *espressoCommon.HeaderImpl) error {
	timer := time.NewTimer(finalityTimeout)
	defer timer.Stop()

	ticker := time.NewTicker(finalityCheckInterval)
	defer ticker.Stop()

	var snapshot espressoCommon.BlockMerkleSnapshot

Loop:
	for {
		select {
		case <-ticker.C:
			res, err := l.EspressoLightClient.FetchMerkleRoot(height, nil)
			if err == nil {
				snapshot = res
				break Loop
			}
		case <-timer.C:
			return fmt.Errorf("failed to fetch merkle root")
		}
	}

	if snapshot.Height <= height {
		return fmt.Errorf("snapshot height is less than or equal to the requested height")
	}

	nextHeader, err := l.Espresso.FetchHeaderByHeight(l.shutdownCtx, snapshot.Height)
	if err != nil {
		return fmt.Errorf("error fetching the snapshot header (height: %d): %w", snapshot.Height, err)
	}

	proof, err := l.Espresso.FetchBlockMerkleProof(l.shutdownCtx, snapshot.Height, height)
	if err != nil {
		return fmt.Errorf("error fetching merkle proof")
	}

	blockMerkleTreeRoot := nextHeader.Header.GetBlockMerkleTreeRoot()

	log.Info("Verifying merkle proof", "height", height)
	ok := espressoVerification.VerifyMerkleProof(proof.Proof, rawHeader, *blockMerkleTreeRoot, snapshot.Root)
	if !ok {
		return fmt.Errorf("error validating merkle proof (height: %d, snapshot height: %d)", height, snapshot.Height)
	}

	// Verify the namespace proof
	log.Info("Verifying namespace proof", "height", height)
	resp, err := l.Espresso.FetchTransactionsInBlock(l.shutdownCtx, height, 42)
	if err != nil {
		return fmt.Errorf("failed to fetch the transactions in block")
	}

	namespaceOk := espressoVerification.VerifyNamespace(
		exampleNamespace,
		resp.Proof,
		*header.Header.GetPayloadCommitment(),
		*header.Header.GetNsTable(),
		resp.Transactions,
		resp.VidCommon,
	)

	if !namespaceOk {
		return fmt.Errorf("error validating namespace proof (height: %d)", height)
	}

	return nil
}

func (l *BatchSubmitter) submitToEspresso(txdata txData, sig, batcherSignature []byte) (*EspressoCommitment, error) {
	transaction := Transaction{
		Namespace:        exampleNamespace,
		BatcherSignature: batcherSignature,
		CallData:         txdata.CallData(),
	}.toEspresso()
	txHash, err := l.Espresso.SubmitTransaction(l.shutdownCtx, transaction)
	if err != nil {
		l.Log.Error("Failed to submit transaction", "transaction", transaction, "error", err)
		l.recordFailedDARequest(txdata.ID(), err)
		return nil, fmt.Errorf("failed to submit transaction: %w", err)
	}

	timer := time.NewTimer(transactionFetchTimeout)
	defer timer.Stop()

	ticker := time.NewTicker(transactionFetchInterval)
	defer ticker.Stop()

	var txQueryData espressoCommon.TransactionQueryData
Loop:
	for {
		select {
		case <-ticker.C:
			txQueryData, err = l.Espresso.FetchTransactionByHash(l.shutdownCtx, txHash)
			if err == nil {
				break Loop
			}
			l.Log.Warn("Retry fetching transaction by hash", "txHash", txHash, "error", err)
		case <-timer.C:
			l.Log.Error("Failed to fetch transaction by hash after multiple attempts", "txHash", txHash)
			l.recordFailedDARequest(txdata.ID(), err)
			return nil, fmt.Errorf("failed to fetch transaction by hash: %w", err)
		}
	}

	rawHeader, err := l.Espresso.FetchRawHeaderByHeight(l.shutdownCtx, txQueryData.BlockHeight)
	if err != nil {
		return nil, err
	}

	var header espressoCommon.HeaderImpl
	err = json.Unmarshal(rawHeader, &header)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal header from bytes")
	}

	height := header.Header.GetBlockHeight()

	err = l.waitForFinality(height, rawHeader, &header)
	if err != nil {
		return nil, err
	}

	espComm := EspressoCommitment{
		Signature: sig,
		TxHash:    txQueryData.Hash.Value(),
	}

	return &espComm, nil
}

// sendEspressoTx uses the txmgr queue to send the given transaction candidate after setting its
// gaslimit. It will block if the txmgr queue has reached its MaxPendingTransactions limit.
func (l *BatchSubmitter) sendEspressoTx(commitment *EspressoCommitment, txId txID, queue *txmgr.Queue[txRef], receiptsCh chan txmgr.TxReceipt[txRef]) error {
	encodedCommitment := commitment.toGeneric().TxData()

	signature, err := crypto.Sign(crypto.Keccak256(encodedCommitment), l.Config.EphemeralPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction data: %w", err)
	}

	batchInboxAbi, err := bindings.BatchInboxMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get batch inbox ABI: %w", err)
	}

	data, err := batchInboxAbi.Pack("submitBatch", encodedCommitment, signature)

	if err != nil {
		return fmt.Errorf("failed to pack transaction data: %w", err)
	}

	candidate := &txmgr.TxCandidate{
		TxData:   data,
		To:       &l.RollupConfig.BatchInboxAddress,
		GasLimit: 210_000,
	}

	queue.Send(txRef{id: txId, isCancel: false, isBlob: false}, *candidate, receiptsCh)

	return nil
}
