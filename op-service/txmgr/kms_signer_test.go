package txmgr

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// ecdsaDERSignature is used in tests to construct DER-encoded signatures that mimic
// what AWS KMS returns, so that derToEthereumSig can be tested without a real KMS key.
type ecdsaDERSignature struct {
	R, S *big.Int
}

// ---------------------------------------------------------------------------
// derToEthereumSig unit tests (no AWS dependency)
// ---------------------------------------------------------------------------

// TestDerToEthereumSig_RoundTrip verifies the DER → Ethereum 65-byte conversion
// using a real secp256k1 key pair. This exercises the core conversion that the
// KMS signer performs on every signature returned by AWS.
func TestDerToEthereumSig_RoundTrip(t *testing.T) {
	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	hash := crypto.Keccak256([]byte("test message"))
	from := crypto.PubkeyToAddress(privKey.PublicKey)

	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash)
	require.NoError(t, err)

	der, err := asn1.Marshal(ecdsaDERSignature{R: r, S: s})
	require.NoError(t, err)

	sig, err := derToEthereumSig(der, hash, from)
	require.NoError(t, err)
	require.Len(t, sig, 65)

	// Recovery bit must be 0 or 1.
	require.LessOrEqual(t, sig[64], byte(1))

	// The signature must recover to the correct address.
	recovered, err := crypto.SigToPub(hash, sig)
	require.NoError(t, err)
	require.Equal(t, from, crypto.PubkeyToAddress(*recovered))
}

// TestDerToEthereumSig_HighSNormalization verifies EIP-2 low-S normalization.
// KMS may return a high-S signature; the signer must normalise it before returning.
func TestDerToEthereumSig_HighSNormalization(t *testing.T) {
	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	hash := crypto.Keccak256([]byte("high-s test"))
	from := crypto.PubkeyToAddress(privKey.PublicKey)

	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash)
	require.NoError(t, err)

	// Force s into the high-S range to simulate what KMS may return.
	curveOrder := crypto.S256().Params().N
	halfOrder := new(big.Int).Rsh(curveOrder, 1)
	if s.Cmp(halfOrder) <= 0 {
		s = new(big.Int).Sub(curveOrder, s)
	}

	der, err := asn1.Marshal(ecdsaDERSignature{R: r, S: s})
	require.NoError(t, err)

	sig, err := derToEthereumSig(der, hash, from)
	require.NoError(t, err)

	// The output S must be in the low-S range.
	outS := new(big.Int).SetBytes(sig[32:64])
	require.LessOrEqual(t, outS.Cmp(halfOrder), 0, "S must be normalised to low-S")

	// Address recovery must still succeed after normalisation.
	recovered, err := crypto.SigToPub(hash, sig)
	require.NoError(t, err)
	require.Equal(t, from, crypto.PubkeyToAddress(*recovered))
}

// TestDerToEthereumSig_BadDER ensures malformed DER input is rejected cleanly.
func TestDerToEthereumSig_BadDER(t *testing.T) {
	_, err := derToEthereumSig([]byte("not valid der"), []byte("hash"), common.Address{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to parse DER signature")
}

// TestDerToEthereumSig_WrongAddress ensures that if the expected address does not match
// either recovery candidate, an error is returned rather than a silently wrong signature.
func TestDerToEthereumSig_WrongAddress(t *testing.T) {
	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	hash := crypto.Keccak256([]byte("wrong address test"))

	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash)
	require.NoError(t, err)

	der, err := asn1.Marshal(ecdsaDERSignature{R: r, S: s})
	require.NoError(t, err)

	// Wrong address — belongs to a different key.
	wrongKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	wrongAddr := crypto.PubkeyToAddress(wrongKey.PublicKey)

	_, err = derToEthereumSig(der, hash, wrongAddr)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to determine recovery bit")
}

// ---------------------------------------------------------------------------
// Mock KMS HTTP server helpers
// ---------------------------------------------------------------------------

// marshalSPKI encodes a secp256k1 public key as a DER SubjectPublicKeyInfo, matching
// the format that AWS KMS GetPublicKey returns. We use the same ecSubjectPublicKeyInfo
// struct that the production parser uses, so the round-trip is structurally identical.
func marshalSPKI(t *testing.T, pub *ecdsa.PublicKey) []byte {
	t.Helper()

	// Uncompressed EC point: 04 || X (32) || Y (32).
	raw := make([]byte, 65)
	raw[0] = 0x04
	xb := pub.X.Bytes()
	yb := pub.Y.Bytes()
	copy(raw[1+32-len(xb):33], xb)
	copy(raw[33+32-len(yb):65], yb)

	der, err := asn1.Marshal(ecSubjectPublicKeyInfo{
		Algorithm: ecPublicKeyAlgorithm{
			Algorithm:  oidECPublicKey,
			Parameters: oidSecp256k1,
		},
		PublicKey: asn1.BitString{
			Bytes:     raw,
			BitLength: 8 * len(raw),
		},
	})
	require.NoError(t, err)
	return der
}

// mockKMSServer starts a local HTTP server that mimics the AWS KMS JSON API.
// privKey is the "KMS key" used to respond to GetPublicKey and Sign requests.
func mockKMSServer(t *testing.T, privKey *ecdsa.PrivateKey) *httptest.Server {
	t.Helper()

	spkiDER := marshalSPKI(t, &privKey.PublicKey)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := r.Header.Get("X-Amz-Target")
		w.Header().Set("Content-Type", "application/x-amz-json-1.1")

		switch target {
		case "TrentService.GetPublicKey":
			json.NewEncoder(w).Encode(map[string]interface{}{
				"KeyId":     "test-key-id",
				"KeySpec":   "ECC_SECG_P256K1",
				"PublicKey": base64.StdEncoding.EncodeToString(spkiDER),
			})

		case "TrentService.Sign":
			var req struct {
				Message string `json:"Message"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "bad request body", http.StatusBadRequest)
				return
			}
			hash, err := base64.StdEncoding.DecodeString(req.Message)
			if err != nil {
				http.Error(w, "bad Message base64", http.StatusBadRequest)
				return
			}
			sigR, sigS, err := ecdsa.Sign(rand.Reader, privKey, hash)
			if err != nil {
				http.Error(w, "sign failed", http.StatusInternalServerError)
				return
			}
			der, err := asn1.Marshal(ecdsaDERSignature{R: sigR, S: sigS})
			if err != nil {
				http.Error(w, "marshal failed", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]interface{}{
				"KeyId":            "test-key-id",
				"Signature":        base64.StdEncoding.EncodeToString(der),
				"SigningAlgorithm": "ECDSA_SHA_256",
			})

		default:
			http.Error(w, "unexpected target: "+target, http.StatusBadRequest)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

// setFakeAWSEnv sets minimal environment variables so that the AWS SDK can load
// credentials and a region without hitting EC2 metadata or real config files.
func setFakeAWSEnv(t *testing.T) {
	t.Helper()
	t.Setenv("AWS_ACCESS_KEY_ID", "test-access-key")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "test-secret-key")
	t.Setenv("AWS_REGION", "us-east-1")
	// Prevent the SDK from probing 169.254.169.254 during credential resolution.
	t.Setenv("AWS_EC2_METADATA_DISABLED", "true")
}

// ---------------------------------------------------------------------------
// Integration tests using the mock KMS server
// ---------------------------------------------------------------------------

// TestNewKMSChainSigner_AddressDerivation verifies that NewKMSChainSigner derives
// the correct Ethereum address from a KMS-returned DER public key.
func TestNewKMSChainSigner_AddressDerivation(t *testing.T) {
	setFakeAWSEnv(t)

	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	srv := mockKMSServer(t, privKey)
	expectedAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	signer, addr, err := NewKMSChainSigner(t.Context(), "test-key-id", srv.URL, big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, expectedAddr, addr)
	require.Equal(t, expectedAddr, signer.from)
}

// TestKMSChainSigner_SignTransaction verifies the full sign path: a transaction hash is
// sent to the mock KMS server, the DER response is converted to an Ethereum signature,
// and the signed transaction recovers to the expected sender address.
func TestKMSChainSigner_SignTransaction(t *testing.T) {
	setFakeAWSEnv(t)

	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	chainID := big.NewInt(1337)
	srv := mockKMSServer(t, privKey)
	expectedAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	signer, _, err := NewKMSChainSigner(t.Context(), "test-key-id", srv.URL, chainID)
	require.NoError(t, err)

	tx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     1,
		GasTipCap: big.NewInt(1e9),
		GasFeeCap: big.NewInt(2e9),
		Gas:       21000,
		To:        &expectedAddr,
		Value:     big.NewInt(0),
	})

	signed, err := signer.SignTransaction(t.Context(), expectedAddr, tx)
	require.NoError(t, err)

	// The signed transaction must recover to the expected sender.
	txSigner := ethtypes.LatestSignerForChainID(chainID)
	sender, err := txSigner.Sender(signed)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, sender)
}

// TestKMSChainSigner_Sign verifies the Sign method (used for arbitrary hash signing,
// e.g. Espresso commitment signing) against the mock KMS server.
func TestKMSChainSigner_Sign(t *testing.T) {
	setFakeAWSEnv(t)

	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	srv := mockKMSServer(t, privKey)
	expectedAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	signer, _, err := NewKMSChainSigner(t.Context(), "test-key-id", srv.URL, big.NewInt(1))
	require.NoError(t, err)

	hash := crypto.Keccak256([]byte("arbitrary payload"))
	sig, err := signer.Sign(t.Context(), hash)
	require.NoError(t, err)
	require.Len(t, sig, 65)

	recovered, err := crypto.SigToPub(hash, sig)
	require.NoError(t, err)
	require.Equal(t, expectedAddr, crypto.PubkeyToAddress(*recovered))
}

// TestKMSChainSigner_SignTransaction_AddressMismatch verifies that SignTransaction
// returns an error when the requested signer address differs from the loaded key.
func TestKMSChainSigner_SignTransaction_AddressMismatch(t *testing.T) {
	setFakeAWSEnv(t)

	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	chainID := big.NewInt(1)
	srv := mockKMSServer(t, privKey)

	signer, _, err := NewKMSChainSigner(t.Context(), "test-key-id", srv.URL, chainID)
	require.NoError(t, err)

	wrongKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	wrongAddr := crypto.PubkeyToAddress(wrongKey.PublicKey)

	tx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{ChainID: chainID, Gas: 21000})
	_, err = signer.SignTransaction(t.Context(), wrongAddr, tx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "address mismatch")
}

// ---------------------------------------------------------------------------
// CLI config validation
// ---------------------------------------------------------------------------

// TestCLIConfig_KMSKeyIDAloneIsValid verifies that --kms-key-id by itself passes Check().
func TestCLIConfig_KMSKeyIDAloneIsValid(t *testing.T) {
	cfg := NewCLIConfig(l1EthRpcValue, DefaultBatcherFlagValues)
	cfg.KMSKeyID = "arn:aws:kms:us-east-1:123456789012:key/abc"
	require.NoError(t, cfg.Check())
}

// TestCLIConfig_KMSKeyIDWithPrivateKeyIsInvalid verifies that setting both
// --kms-key-id and --private-key is rejected.
func TestCLIConfig_KMSKeyIDWithPrivateKeyIsInvalid(t *testing.T) {
	cfg := NewCLIConfig(l1EthRpcValue, DefaultBatcherFlagValues)
	cfg.KMSKeyID = "arn:aws:kms:us-east-1:123456789012:key/abc"
	cfg.PrivateKey = "0xdeadbeef"
	require.Error(t, cfg.Check())
}

// TestCLIConfig_KMSEndpointURLWithoutKeyIDIsValid verifies that --kms-endpoint-url
// without --kms-key-id is not an error (the endpoint flag is simply unused).
func TestCLIConfig_KMSEndpointURLWithoutKeyIDIsValid(t *testing.T) {
	cfg := NewCLIConfig(l1EthRpcValue, DefaultBatcherFlagValues)
	cfg.KMSEndpointURL = "http://localhost:8339"
	require.NoError(t, cfg.Check())
}
