package key_service

import (
	"github.com/hnamzian/hedera-vault-plugin/src/core/key"
	"github.com/hnamzian/hedera-vault-plugin/src/key/entity"
)

func (k_svc *KeyService) ImportKey(id, priv, algo, curve string) (*key_entity.Key, error) {
	keypair, err := key.FromPrivateKey(priv, algo)
	if err != nil {
		return nil, nil
	}

	keybuf, err := key_entity.FromKeyPair(id, keypair).ToBytes()
	if err != nil {
		return nil, nil
	}

	key, err := key_entity.FromBytes(keybuf)
	if err != nil {
		return nil, err
	}

	if err := k_svc.storage.Write(id, key); err != nil {
		return nil, err
	}

	return key, err
}