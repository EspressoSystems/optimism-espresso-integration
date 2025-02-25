package batcher

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestBatchSubmitter_SignatureGeneration(t *testing.T) {
	bs, _ := setup(t)

	txdata := emptyTxData

	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate key pair for batcher: %v", err)
	}
	bs.Config.BatcherPrivateKey = key
	bs.Config.BatcherPublicKey = &key.PublicKey

	// add batcher's signature on txdata sent to L1
	sig, err := txdata.signTx(bs.Config.BatcherPrivateKey)
	require.NoError(t, err)

	// test that the valid signature can be verified
	pubKeyBytes := crypto.FromECDSAPub(bs.Config.BatcherPublicKey)
	require.True(t, crypto.VerifySignature(pubKeyBytes, crypto.Keccak256(txdata.CallData()), sig[:len(sig)-1]))

	// test that the invalid signature cannot be verified
	badSig := []byte{1, 2, 3, 4}
	require.False(t, crypto.VerifySignature(pubKeyBytes, crypto.Keccak256(txdata.CallData()), badSig[:len(badSig)-1]))
}
