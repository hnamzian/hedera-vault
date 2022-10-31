package service

import "github.com/hnamzian/hedera-vault-plugin/src/account/entity"

func (a_svc *AccountService) GetAccount(id string) (*entity.Account, error) {
	account, err := a_svc.storage.Read(id)
	if err != nil {
		return nil, nil
	}

	return account, nil
}