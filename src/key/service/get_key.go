package key_service

import (
	"fmt"

	key_entity "github.com/hnamzian/hedera-vault-plugin/src/key/entity"
)

func (k_svc *KeyService) GetKey(id string) (*key_entity.Key, error) {
	
	key, err := k_svc.storage.Read(id)
	if err != nil {
		return nil, fmt.Errorf("read key from storage failed: %s", err)
	}

	return key, nil
}
