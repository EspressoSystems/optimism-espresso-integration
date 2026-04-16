package batcher

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/espresso/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
)

// computeCommitment computes the batch commitment hash from a transaction candidate.
// For calldata transactions, it returns keccak256(calldata).
// For blob transactions, it returns keccak256(concat(blobVersionedHashes)).
func computeCommitment(candidate *txmgr.TxCandidate) ([32]byte, error) {
	if len(candidate.Blobs) == 0 {
		return crypto.Keccak256Hash(candidate.TxData), nil
	}

	concatenatedBlobHashes := make([]byte, 0)
	for _, blob := range candidate.Blobs {
		blobCommitment, err := blob.ComputeKZGCommitment()
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to compute KZG commitment for blob: %w", err)
		}
		blobHash := eth.KZGToVersionedHash(blobCommitment)
		concatenatedBlobHashes = append(concatenatedBlobHashes, blobHash.Bytes()...)
	}
	return crypto.Keccak256Hash(concatenatedBlobHashes), nil
}

// sendTxWithFallbackAuth authenticates a batch transaction via the BatchAuthenticator contract
// using the fallback batcher's signing key (raw ECDSA, no TEE attestation), then sends the
// batch data to the BatchInbox address.
//
// This is the non-Espresso counterpart to sendTxWithEspresso: it uses the ChainSigner
// (the TxManager's key, which corresponds to the SystemConfig batcher address) to sign
// the commitment, calls authenticateBatchInfo on L1, and then submits the batch data.
func (l *BatchSubmitter) sendTxWithFallbackAuth(txdata txData, isCancel bool, candidate *txmgr.TxCandidate, queue TxSender[txRef], receiptsCh chan txmgr.TxReceipt[txRef]) {
	transactionReference := txRef{id: txdata.ID(), isCancel: isCancel, isBlob: txdata.daType == DaTypeBlob, daType: txdata.daType, size: txdata.Len()}
	l.Log.Debug("Sending fallback-authenticated L1 transaction", "txRef", transactionReference)

	commitment, err := computeCommitment(candidate)
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to compute commitment: %w", err),
		}
		return
	}
	l.Log.Debug("Computed fallback batch commitment", "txRef", transactionReference, "commitment", hexutil.Encode(commitment[:]))

	// Sign the raw commitment hash with the ChainSigner (TxManager key = SystemConfig batcher address).
	// The contract verifies this with ECDSA.recover(commitment, signature) against systemConfig.batcherHash().
	signature, err := l.ChainSigner.Sign(l.killCtx, commitment[:])
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to sign commitment for fallback auth: %w", err),
		}
		return
	}

	// Normalize the recovery ID (v) from 0/1 to 27/28 for Solidity's ECDSA.recover
	if signature[64] < 27 {
		signature[64] += 27
	}

	l.Log.Debug("Signed fallback batch commitment", "txRef", transactionReference, "commitment", hexutil.Encode(commitment[:]), "sig", hexutil.Encode(signature))

	batchAuthenticatorAbi, err := bindings.BatchAuthenticatorMetaData.GetAbi()
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to get batch authenticator ABI: %w", err),
		}
		return
	}

	authenticateBatchCalldata, err := batchAuthenticatorAbi.Pack("authenticateBatchInfo", commitment, signature)
	if err != nil {
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to pack authenticateBatchInfo calldata: %w", err),
		}
		return
	}

	verifyCandidate := txmgr.TxCandidate{
		TxData: authenticateBatchCalldata,
		To:     &l.RollupConfig.BatchAuthenticatorAddress,
	}

	l.Log.Debug(
		"Sending fallback authenticateBatchInfo transaction",
		"txRef", transactionReference,
		"commitment", hexutil.Encode(commitment[:]),
		"address", l.RollupConfig.BatchAuthenticatorAddress.String(),
	)
	verificationReceipt, err := l.Txmgr.Send(l.killCtx, verifyCandidate)
	if err != nil {
		l.Log.Error("Failed to send fallback authenticateBatchInfo transaction", "txRef", transactionReference, "err", err)
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to send fallback authenticateBatchInfo transaction: %w", err),
		}
		return
	}

	receipt, err := l.Txmgr.Send(l.killCtx, *candidate)
	if err != nil {
		l.Log.Error("Failed to send batch inbox transaction", "txRef", transactionReference, "err", err)
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("failed to send batch inbox transaction: %w", err),
		}
		return
	}

	distance := new(big.Int).Sub(receipt.BlockNumber, verificationReceipt.BlockNumber)
	lookbackWindow := new(big.Int).SetUint64(uint64(derive.BatchAuthLookbackWindow))
	if distance.Sign() < 0 || distance.Cmp(lookbackWindow) >= 0 {
		l.Log.Error("authenticateBatchInfo transaction too far from batch inbox transaction", "txRef", transactionReference, "distance", distance)
		receiptsCh <- txmgr.TxReceipt[txRef]{
			ID:  transactionReference,
			Err: fmt.Errorf("authenticateBatchInfo transaction too far from batch inbox transaction: %s", distance),
		}
		return
	}

	receiptsCh <- txmgr.TxReceipt[txRef]{
		ID:      transactionReference,
		Receipt: receipt,
		Err:     nil,
	}
}
