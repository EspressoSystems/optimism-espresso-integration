package enclave

import (
	"errors"

	"github.com/hf/nsm"
	"github.com/hf/nsm/request"
)

func attest(nonce, userData, publicKey []byte) ([]byte, error) {
	sess, err := nsm.OpenDefaultSession()
	defer sess.Close()

	if nil != err {
		return nil, err
	}

	res, err := sess.Send(&request.Attestation{
		Nonce:     nonce,
		UserData:  userData,
		PublicKey: publicKey,
	})
	if nil != err {
		return nil, err
	}

	if "" != res.Error {
		return nil, errors.New(string(res.Error))
	}

	if nil == res.Attestation || nil == res.Attestation.Document {
		return nil, errors.New("NSM device did not return an attestation")
	}

	return res.Attestation.Document, nil
}

func GetAttestationWithTxData(txData []byte) ([]byte, error) {
	// Use empty slices for nonce and publicKey when they're not needed
	nonce := make([]byte, 0)
	publicKey := make([]byte, 0)

	// Call the existing attest function with txData as userData
	return attest(nonce, txData, publicKey)
}
