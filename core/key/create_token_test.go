package key

import (
	"fmt"
	"testing"
)

func TestCreateKeyED25519(t *testing.T) {
	key, err := CreateKey(ALGORITHM_ED25519, "")
	if err != nil {
		t.Fatalf("Unable to generate key pair %s", err)
	}

	fmt.Printf("Algorithm %s\n", key.Algorithm)
	fmt.Printf("Private Key %s\n", key.PrivateKey)
	fmt.Printf("Public Key %s\n", key.PublicKey)
}

func TestCreateKeyECDSA(t *testing.T) {
	key, err := CreateKey(ALGORITHM_ECDSA, CURVE_ECDSA_SECP256K1)
	if err != nil {
		t.Fatalf("Unable to generate key pair %s", err)
	}

	fmt.Printf("Algorithm %s\n", key.Algorithm)
	fmt.Printf("Private Key %s\n", key.PrivateKey)
	fmt.Printf("Public Key %s\n", key.PublicKey)
}