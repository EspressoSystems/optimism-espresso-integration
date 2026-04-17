package txmgr

import (
	"context"
	"encoding/asn1"
	"fmt"
	"math/big"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	kmstypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	secp256k1ecdsa "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	opcrypto "github.com/ethereum-optimism/optimism/op-service/crypto"
)

// OIDs required to parse the SubjectPublicKeyInfo structure returned by KMS GetPublicKey.
// Go's crypto/x509 does not support secp256k1, so we parse the DER manually — this is the
// established pattern used by all Go+KMS+Ethereum libraries (e.g. matelang/go-ethereum-aws-kms-tx-signer).
var (
	oidECPublicKey = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
	oidSecp256k1   = asn1.ObjectIdentifier{1, 3, 132, 0, 10}
)

// ecPublicKeyAlgorithm is the AlgorithmIdentifier for an EC key in SubjectPublicKeyInfo.
type ecPublicKeyAlgorithm struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.ObjectIdentifier
}

// ecSubjectPublicKeyInfo is the ASN.1 structure of the DER blob returned by KMS GetPublicKey.
type ecSubjectPublicKeyInfo struct {
	Algorithm ecPublicKeyAlgorithm
	PublicKey asn1.BitString
}

// kmsChainSigner implements opcrypto.ChainSigner using AWS KMS as the signing backend.
// The private key never leaves KMS hardware — signing is performed inside KMS and only
// the signature is returned.
type kmsChainSigner struct {
	client  *kms.Client
	keyID   string
	chainID *big.Int
	from    common.Address
}

var _ opcrypto.ChainSigner = (*kmsChainSigner)(nil)

// NewKMSChainSigner creates a kmsChainSigner backed by the given KMS key.
// endpointURL overrides the KMS endpoint — set to http://127.0.0.1:{KMSProxyPort} when
// running inside a Nitro enclave so calls are routed through the enclaver vsock proxy.
// Leave empty to use the default regional KMS endpoint (e.g. for local testing with localstack).
func NewKMSChainSigner(ctx context.Context, keyID, endpointURL string, chainID *big.Int) (*kmsChainSigner, common.Address, error) {
	cfg, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to load AWS config: %w", err)
	}

	var opts []func(*kms.Options)
	if endpointURL != "" {
		opts = append(opts, func(o *kms.Options) {
			o.BaseEndpoint = aws.String(endpointURL)
			// The enclaver vsock proxy speaks plain HTTP, not HTTPS.
			o.HTTPClient = &http.Client{}
		})
	}
	client := kms.NewFromConfig(cfg, opts...)

	// Derive the Ethereum address from the KMS public key once at startup.
	from, err := kmsEthereumAddress(ctx, client, keyID)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to derive Ethereum address from KMS key %q: %w", keyID, err)
	}

	return &kmsChainSigner{
		client:  client,
		keyID:   keyID,
		chainID: chainID,
		from:    from,
	}, from, nil
}

// Sign signs an arbitrary hash using the KMS key. The returned signature is in
// Ethereum's 65-byte [R || S || V] format.
func (k *kmsChainSigner) Sign(ctx context.Context, hash []byte) ([]byte, error) {
	return k.signHash(ctx, hash)
}

// SignTransaction signs an Ethereum transaction using the KMS key.
func (k *kmsChainSigner) SignTransaction(ctx context.Context, addr common.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
	if addr != k.from {
		return nil, fmt.Errorf("address mismatch: requested %s, signer is %s", addr, k.from)
	}
	txSigner := ethtypes.LatestSignerForChainID(k.chainID)
	hash := txSigner.Hash(tx)
	sig, err := k.signHash(ctx, hash.Bytes())
	if err != nil {
		return nil, fmt.Errorf("KMS signing failed: %w", err)
	}
	return tx.WithSignature(txSigner, sig)
}

// signHash calls KMS to sign hash and converts the DER-encoded response to Ethereum's
// 65-byte [R || S || V] format including low-S normalization (EIP-2) and recovery bit.
func (k *kmsChainSigner) signHash(ctx context.Context, hash []byte) ([]byte, error) {
	out, err := k.client.Sign(ctx, &kms.SignInput{
		KeyId:            aws.String(k.keyID),
		Message:          hash,
		MessageType:      kmstypes.MessageTypeDigest,
		SigningAlgorithm: kmstypes.SigningAlgorithmSpecEcdsaSha256,
	})
	if err != nil {
		return nil, fmt.Errorf("KMS Sign call failed: %w", err)
	}

	return derToEthereumSig(out.Signature, hash, k.from)
}

// derToEthereumSig converts a DER-encoded ECDSA signature from KMS to Ethereum's
// 65-byte format [R (32) || S (32) || V (1)], applying low-S normalization (EIP-2)
// and computing the recovery bit by trial.
//
// secp256k1ecdsa.ParseDERSignature is used instead of asn1.Unmarshal because it
// performs proper secp256k1-specific DER validation, and ModNScalar.Bytes() returns a
// fixed [32]byte so no manual zero-padding is needed. IsOverHalfOrder / Negate handle
// the low-S normalisation without any big.Int arithmetic.
func derToEthereumSig(der []byte, hash []byte, from common.Address) ([]byte, error) {
	parsed, err := secp256k1ecdsa.ParseDERSignature(der)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DER signature: %w", err)
	}

	r := parsed.R()
	s := parsed.S()

	// EIP-2: ensure low-S. KMS may return a high-S value.
	if s.IsOverHalfOrder() {
		s.Negate()
	}

	rBytes := r.Bytes() // [32]byte, already zero-padded
	sBytes := s.Bytes() // [32]byte, already zero-padded

	// Determine the recovery bit (V) by trying both 0 and 1 and checking which
	// recovers to the expected address. KMS does not provide V.
	sig65 := make([]byte, 65)
	copy(sig65[0:32], rBytes[:])
	copy(sig65[32:64], sBytes[:])
	for v := byte(0); v <= 1; v++ {
		sig65[64] = v
		pubKey, err := crypto.SigToPub(hash, sig65)
		if err != nil {
			continue
		}
		if crypto.PubkeyToAddress(*pubKey) == from {
			return sig65, nil
		}
	}

	return nil, fmt.Errorf("failed to determine recovery bit: neither V=0 nor V=1 recovers address %s", from)
}

// kmsEthereumAddress retrieves the public key from KMS and derives the Ethereum address.
// Go's crypto/x509 does not support secp256k1, so the SubjectPublicKeyInfo DER is parsed
// manually: the raw uncompressed EC point is extracted and passed to crypto.UnmarshalPubkey.
func kmsEthereumAddress(ctx context.Context, client *kms.Client, keyID string) (common.Address, error) {
	out, err := client.GetPublicKey(ctx, &kms.GetPublicKeyInput{
		KeyId: aws.String(keyID),
	})
	if err != nil {
		return common.Address{}, fmt.Errorf("GetPublicKey failed: %w", err)
	}

	var spki ecSubjectPublicKeyInfo
	if _, err := asn1.Unmarshal(out.PublicKey, &spki); err != nil {
		return common.Address{}, fmt.Errorf("failed to parse KMS public key DER: %w", err)
	}
	if !spki.Algorithm.Algorithm.Equal(oidECPublicKey) {
		return common.Address{}, fmt.Errorf("unexpected algorithm OID %v (expected ecPublicKey)", spki.Algorithm.Algorithm)
	}
	if !spki.Algorithm.Parameters.Equal(oidSecp256k1) {
		return common.Address{}, fmt.Errorf("unexpected curve OID %v (expected secp256k1)", spki.Algorithm.Parameters)
	}

	// crypto.UnmarshalPubkey expects a 65-byte uncompressed point: 04 || X (32) || Y (32).
	pubKey, err := crypto.UnmarshalPubkey(spki.PublicKey.Bytes)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unmarshal secp256k1 public key: %w", err)
	}

	return crypto.PubkeyToAddress(*pubKey), nil
}
