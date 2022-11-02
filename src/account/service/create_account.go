package service

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/hnamzian/hedera-vault-plugin/src/account/entity"
	hc "github.com/hnamzian/hedera-vault-plugin/src/core/hedera"
)

func (a_svc *AccountService) CreateAccount(id, newID string) (*entity.Account, error) {
	operator_account, err := a_svc.storage.Read(id)
	if err != nil {
		return nil, err
	}

	operator_key, err := a_svc.k_svc.GetKey(operator_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("decode account failed: %s", err)
	}

	client := hc.NewClient(hc.LocalTestNetClientConfig())
	operator_account_id, err := hedera.AccountIDFromString(operator_account.AccountID)
	if err != nil {
		return nil, fmt.Errorf("invalid account ID: %s", err)
	}
	operator_private_key, err := hedera.PrivateKeyFromString(operator_key.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %s", err)
	}
	new_account_id, err := client.WithOperator(operator_account_id, operator_private_key).NewAccount(operator_private_key)
	if err != nil {
		return nil, fmt.Errorf("create new account failed: %s", err)
	}

	new_account := entity.New(newID, new_account_id.String(), operator_account.KeyID)
	if err := a_svc.storage.Write(newID, new_account); err != nil {
		return nil, err
	}

	return new_account, nil
}
