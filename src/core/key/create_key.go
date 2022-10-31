package key

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

const (
	ALGORITHM_ED25519     = "ED25519"
	ALGORITHM_ECDSA       = "ECDSA"
	CURVE_ECDSA_SECP256K1 = "secp256k1"
)

type KeyPair struct {
	PublicKey  hedera.PublicKey
	PrivateKey hedera.PrivateKey
	Algorithm  string
	Curve      string
}

func NewKeyPair(pub hedera.PublicKey, priv hedera.PrivateKey, algo, curve string) *KeyPair {
	return &KeyPair{
		PublicKey:  pub,
		PrivateKey: priv,
		Algorithm:  algo,
		Curve:      curve,
	}
}

func CreateKey(algo, curve string) (*KeyPair, error) {
	var prv hedera.PrivateKey
	var err error

	if algo == ALGORITHM_ED25519 {
		prv, err = hedera.PrivateKeyGenerateEd25519()
	} else if algo == ALGORITHM_ECDSA && curve == CURVE_ECDSA_SECP256K1 {
		prv, err = hedera.PrivateKeyGenerateEcdsa()
	} else {
		return &KeyPair{}, fmt.Errorf("Invalid Algorithm or Curve")
	}

	if err != nil {
		return &KeyPair{}, err
	}

	pub := prv.PublicKey()
	return NewKeyPair(pub, prv, algo, curve), nil
}

func FromPrivateKey(s string, algo string) (*KeyPair, error) {
	var curve string
	var priv hedera.PrivateKey
	var err error

	if algo == ALGORITHM_ECDSA {
		curve = CURVE_ECDSA_SECP256K1
		priv, err = hedera.PrivateKeyFromStringECSDA(s)
	} else {
		algo = ALGORITHM_ED25519
		priv, err = hedera.PrivateKeyFromStringEd25519(s)
	}

	pub := priv.PublicKey()

	return NewKeyPair(pub, priv, algo, curve), err
}
