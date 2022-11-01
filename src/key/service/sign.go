package key_service

import (
	"fmt"

	"github.com/hnamzian/hedera-vault-plugin/src/core/key"
	key_entity "github.com/hnamzian/hedera-vault-plugin/src/key/entity"
)

func (k_svc *KeyService) Sign(id, message string) (*key_entity.Key, []byte, error) {
	key_vault, err := k_svc.storage.Read(id)
	if err != nil {
		return nil, nil, fmt.Errorf("read key by ID failed: %s", err)
	}

	signature, err := key.Sign(key_vault.PrivateKey, []byte(message))
	if err != nil {
		return nil, nil, fmt.Errorf("signing message failed: %s", err)
	}

	return key_vault, signature, nil
}
