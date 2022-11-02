package service

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/hashicorp/go-hclog"
	hedera_client "github.com/hnamzian/hedera-vault-plugin/src/core/hedera"
	hedera_token "github.com/hnamzian/hedera-vault-plugin/src/core/hedera/token"
)

func (t_svc *TokenService) CreateToken(tokenCreation *hedera_token.FTokenCreation, operatorID, adminID, treasuryID, supplyID string) (*hedera.TokenID, error) {
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

	supply_account, _ := t_svc.a_svc.GetAccount(supplyID)
	if err != nil {
		return nil, fmt.Errorf("retreive supply account from vault failed: %s", err)
	}
	supply_key, _ := t_svc.k_svc.GetKey(supply_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive supply key from vault failed: %s", err)
	}
	hclog.Default().Info("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", supplyID, supply_key.Publickey)

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

	supplyPublicKey, _ := hedera.PublicKeyFromString(supply_key.Publickey)
	if err != nil {
		return nil, fmt.Errorf("parse supply public key failed: %s", err)
	}


	tokenCreation.AdminPrivateKey = adminPrivateKey
	tokenCreation.AdminPublicKey = adminPublicKey

	tokenCreation.TreasuryAccountID = treasuryAccountID
	tokenCreation.TreasuryPrivateKey = treasuryPrivateKey
	tokenCreation.TreasuryPublicKey = treasuryPublicKey
	tokenCreation.SupplyKey = supplyPublicKey

	client := hedera_client.
		NewClient(hedera_client.LocalTestNetClientConfig()).
		WithOperator(operatorAccountID, operatorPrivateKey).
		GetClient()
	ht := hedera_token.New(client)

	return ht.CreateFT(tokenCreation)
}
