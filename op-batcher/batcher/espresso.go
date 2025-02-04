package batcher

import (
	"encoding/json"
	"fmt"
	"time"

	espressoVerification "github.com/EspressoSystems/espresso-sequencer-go/verification"

	espressoCommon "github.com/EspressoSystems/espresso-sequencer-go/types"
	"github.com/ethereum/go-ethereum/log"
)

// TODO: Pull out to be re-used in op-node for derivation from Espresso
type Transaction struct {
	// Namespace of transaction to be published
	Namespace uint64
	// TEE attestation to be verified by op-node
	TeeAttn []byte
	// Frames serialized as they would be for posting to L1 as calldata
	CallData []byte
}

func (t Transaction) toEspresso() espressoCommon.Transaction {
	payload := append(t.TeeAttn, t.CallData...)
	return espressoCommon.Transaction{
		Namespace: t.Namespace,
		Payload:   payload,
	}
}

func (l *BatchSubmitter) submitToEspresso(txdata txData) (*EspressoCommitment, error) {
	transaction := Transaction{
		Namespace: 42,
		TeeAttn:   []byte{1, 2, 3, 4},
		CallData:  txdata.CallData(),
	}.toEspresso()
	txHash, err := l.Espresso.SubmitTransaction(l.shutdownCtx, transaction)
	if err != nil {
		l.Log.Error("Failed to submit transaction", "transaction", transaction, "error", err)
		return nil, fmt.Errorf("failed to submit transaction: %w", err)
	}

	timer := time.NewTimer(2 * time.Minute)
	defer timer.Stop()

	ticker := time.NewTicker(100 * time.Millisecond)
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

	log.Info("Fetching Merkle Root at hotshot", "height", height)
	// Verify the merkle proof
	snapshot, err := l.EspressoLightClient.FetchMerkleRoot(height, nil)
	if err != nil {
		return nil, fmt.Errorf("Error fetching merkle root: %w", err)
	}

	if snapshot.Height <= height {
		return nil, fmt.Errorf("snapshot height is less than or equal to the requested height")
	}

	nextHeader, err := l.Espresso.FetchHeaderByHeight(l.shutdownCtx, snapshot.Height)
	if err != nil {
		return nil, fmt.Errorf("error fetching the snapshot header (height: %d): %w", snapshot.Height, err)
	}

	proof, err := l.Espresso.FetchBlockMerkleProof(l.shutdownCtx, snapshot.Height, height)
	if err != nil {
		return nil, fmt.Errorf("error fetching merkle proof")
	}

	blockMerkleTreeRoot := nextHeader.Header.GetBlockMerkleTreeRoot()

	log.Info("Verifying merkle proof", "height", height)
	ok := espressoVerification.VerifyMerkleProof(proof.Proof, rawHeader, *blockMerkleTreeRoot, snapshot.Root)
	if !ok {
		return nil, fmt.Errorf("error validating merkle proof (height: %d, snapshot height: %d)", height, snapshot.Height)
	}

	// Verify the namespace proof
	log.Info("Verifying namespace proof", "height", height)
	resp, err := l.Espresso.FetchTransactionsInBlock(l.shutdownCtx, height, 42)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the transactions in block")
	}

	namespaceOk := espressoVerification.VerifyNamespace(
		42,
		resp.Proof,
		*header.Header.GetPayloadCommitment(),
		*header.Header.GetNsTable(),
		resp.Transactions,
		resp.VidCommon,
	)

	if !namespaceOk {
		return nil, fmt.Errorf("error validating namespace proof (height: %d)", height)
	}

	// TODO: Generate a real attestation
	teeAttestation := []byte{1, 2, 3, 4}

	espComm := EspressoCommitment{
		TeeAttestation: teeAttestation,
		TxHash:         txQueryData.Hash.Value(),
	}

	return &espComm, nil
}
