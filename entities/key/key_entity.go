package key_entity

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/vault/sdk/helper/jsonutil"
	"github.com/hnamzian/hedera-vault-plugin/core/key"
)

type Key struct {
	ID         string `json:"id"`
	Algorithm  string `json:"algorithm"`
	Curve      string `json:"curve"`
	PrivateKey string `json:"-"`
	Publickey  string `json:"publicKey"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func FromKeyPair(id string, kp *key.KeyPair) *Key {
	return &Key{
		ID:         id,
		Algorithm:  kp.Algorithm,
		Curve:      kp.Curve,
		PrivateKey: kp.PrivateKey.String(),
		Publickey:  kp.PublicKey.String(),
		CreatedAt:  time.Now().UTC().String(),
		UpdatedAt:  time.Now().UTC().String(),
	}
}

func FromBytes(buf []byte) (*Key, error) {
	var key Key
	if err := jsonutil.DecodeJSON(buf, &key); err != nil {
		return &Key{}, err
	}
	return &key, nil
}

func (k *Key) ToBytes() ([]byte, error) {
	val, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}
	return val, nil
}
