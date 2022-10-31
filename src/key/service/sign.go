package key_service

import (
	"github.com/hashicorp/errwrap"
	"github.com/hnamzian/hedera-vault-plugin/src/core/key"
	key_entity "github.com/hnamzian/hedera-vault-plugin/src/key/entity"
)

func (k_svc *KeyService) Sign(id, message string) (*key_entity.Key, []byte, error) {
	key_vault, err := k_svc.storage.Read(id)
	if err != nil {
		return nil, nil, errwrap.Wrapf("read key by ID failed: {{err}}", err)
	}

	signature, err := key.Sign(key_vault.PrivateKey, []byte(message))
	if err != nil {
		return nil, nil, errwrap.Wrapf("signing message failed: {{err}}", err)
	}

	return key_vault, signature, nil
}