package batcher

import (
	"fmt"
	"time"

	espressoCommon "github.com/EspressoSystems/espresso-sequencer-go/types"
	espressoVerification "github.com/EspressoSystems/espresso-sequencer-go/verification"
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
		l.recordFailedDARequest(txdata.ID(), err)
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
			return nil, fmt.Errorf("failed to fetch transaction by hash: %w", err)
		}
	}

	// TODO: Fetch and verify proofs here
	header, err := l.Espresso.FetchHeaderByHeight(l.shutdownCtx, txQueryData.BlockHeight)
	if err != nil {
		return nil, err
	}
	_ = header

	// TODO: Generate a real attestation
	teeAttestation := []byte{1, 2, 3, 4}

	espressoVerification.VerifyMerkleProof(
		txQueryData.Proof,
	)

	espComm := EspressoCommitment{
		TeeAttestation: teeAttestation,
		TxHash:         txQueryData.Hash.Value(),
	}

	return &espComm, nil
}
