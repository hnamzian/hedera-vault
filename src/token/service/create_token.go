package service

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	hc "github.com/hnamzian/hedera-vault-plugin/src/core/hedera"
)

func (t_svc *TokenService) CreateToken(tokenCreation *hc.FTokenCreation, operatorID, adminID, treasuryID string) (*hedera.TokenID, error) {
	operator_account, err := t_svc.a_svc.GetAccount(operatorID)
	if err != nil {
		return nil, fmt.Errorf("retreive operator account from vault failed: %s", err)
	}
	operator_key, _ := t_svc.k_svc.GetKey(operator_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive operator key from vault failed: %s", err)
	}

	admin_account, _ := t_svc.a_svc.GetAccount(adminID)
	if err != nil {
		return nil, fmt.Errorf("retreive admin account from vault failed: %s", err)
	}
	admin_key, _ := t_svc.k_svc.GetKey(admin_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive admin key from vault failed: %s", err)
	}

	treasury_account, _ := t_svc.a_svc.GetAccount(treasuryID)
	if err != nil {
		return nil, fmt.Errorf("retreive treasury account from vault failed: %s", err)
	}
	treasury_key, _ := t_svc.k_svc.GetKey(treasury_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive treasury key from vault failed: %s", err)
	}

	operatorAccountID, _ := hedera.AccountIDFromString(operator_account.AccountID)
	if err != nil {
		return nil, fmt.Errorf("parse operator account id failed: %s", err)
	}
	operatorPrivateKey, _ := hedera.PrivateKeyFromString(operator_key.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("parse operator key failed: %s", err)
	}

	adminPrivateKey, _ := hedera.PrivateKeyFromString(admin_key.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("parse admin private key failed: %s", err)
	}
	adminPublicKey, _ := hedera.PublicKeyFromString(admin_key.Publickey)
	if err != nil {
		return nil, fmt.Errorf("parse admin public key failed: %s", err)
	}

	treasuryAccountID, _ := hedera.AccountIDFromString(treasury_account.AccountID)
	if err != nil {
		return nil, fmt.Errorf("parse treasury account id failed: %s", err)
	}
	treasuryPrivateKey, _ := hedera.PrivateKeyFromString(treasury_key.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("parse treasury private key failed: %s", err)
	}
	treasuryPublicKey, _ := hedera.PublicKeyFromString(treasury_key.Publickey)
	if err != nil {
		return nil, fmt.Errorf("parse treasury public key failed: %s", err)
	}

	tokenCreation.AdminPrivateKey = adminPrivateKey
	tokenCreation.AdminPublicKey = adminPublicKey

	tokenCreation.TreasuryAccountID = treasuryAccountID
	tokenCreation.TreasuryPrivateKey = treasuryPrivateKey
	tokenCreation.TreasuryPublicKey = treasuryPublicKey

	client := hc.NewClient(hc.LocalTestNetClientConfig())

	return client.
		WithOperator(operatorAccountID, operatorPrivateKey).
		CreateFT(tokenCreation)
}
