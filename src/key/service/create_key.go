package key_service

import (
	"github.com/hashicorp/errwrap"
	key_entity "github.com/hnamzian/hedera-vault-plugin/src/key/entity"
	"github.com/hnamzian/hedera-vault-plugin/src/core/key"
)

func (k_svc *KeyService) Create(id, algo, curve string) (*key_entity.Key, error) {
	keypair, err := key.CreateKey(algo, curve)
	if err != nil {
		return nil, errwrap.Wrapf("generate key pair failed: {{err}}", err)
	}

	key_vault := key_entity.FromKeyPair(id, keypair)

	if err := k_svc.storage.Write(id, key_vault); err != nil {
		return nil, err
	}

	return key_vault, nil
}