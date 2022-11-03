package service

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	hedera_client "github.com/hnamzian/hedera-vault-plugin/src/core/hedera"
	hedera_token "github.com/hnamzian/hedera-vault-plugin/src/core/hedera/token"
)

func (t_svc *TokenService) WipeToken(tokenID string, account string, amount uint64, wipeID, operatorID string) (*hedera.Status, error) {
	operator_account, err := t_svc.a_svc.GetAccount(operatorID)
	if err != nil {
		return nil, fmt.Errorf("retreive operator account from vault failed: %s", err)
	}
	operator_key, _ := t_svc.k_svc.GetKey(operator_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive operator key from vault failed: %s", err)
	}
	operatorAccountID, _ := hedera.AccountIDFromString(operator_account.AccountID)
	if err != nil {
		return nil, fmt.Errorf("parse operator account id failed: %s", err)
	}
	operatorPrivateKey, _ := hedera.PrivateKeyFromString(operator_key.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("parse operator key failed: %s", err)
	}

	wipe_account, err := t_svc.a_svc.GetAccount(wipeID)
	if err != nil {
		return nil, fmt.Errorf("retreive wipe account from vault failed: %s", err)
	}
	wipe_key, _ := t_svc.k_svc.GetKey(wipe_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive wipe key from vault failed: %s", err)
	}

	user_account, err := t_svc.a_svc.GetAccount(account)
	if err != nil {
		return nil, fmt.Errorf("retreive user account from vault failed: %s", err)
	}

	client := hedera_client.
		NewClient(hedera_client.LocalTestNetClientConfig()).
		WithOperator(operatorAccountID, operatorPrivateKey).
		GetClient()
	ht := hedera_token.New(client)

	return ht.WipeToken(tokenID, user_account.AccountID, amount, wipe_key.PrivateKey)
}
