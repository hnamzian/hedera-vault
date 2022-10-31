package service

import (
	"github.com/hnamzian/hedera-vault-plugin/src/account/entity"
)

func (a_svc *AccountService) ImportAccount(id, keyID, accountID string) (*entity.Account, error) {
	account := entity.New(id, accountID, keyID)

	if err := a_svc.storage.Write(id, account); err != nil {
		return nil, err
	}

	return account, nil
}
