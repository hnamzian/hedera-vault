package key_service

import (
	key_entity "github.com/hnamzian/hedera-vault-plugin/src/key/entity"
)

func (k_svc *KeyService) Get(id string) (*key_entity.Key, error) {
	key, err := k_svc.storage.Read(id)
	if err != nil {
		return nil, nil
	}

	return key, nil
}
