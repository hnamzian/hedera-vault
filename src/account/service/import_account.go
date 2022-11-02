package service

import (
	"fmt"

	"github.com/hnamzian/hedera-vault-plugin/src/account/entity"
)

func (a_svc *AccountService) ImportAccount(id, keyID, accountID string) (*entity.Account, error) {
	account := entity.New(id, accountID, keyID)

	if err := a_svc.storage.Write(id, account); err != nil {
		return nil, fmt.Errorf("import account failed: %s", err)
	}

	return account, nil
}
