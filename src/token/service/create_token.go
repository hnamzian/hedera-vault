package service

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	hedera_client "github.com/hnamzian/hedera-vault-plugin/src/core/hedera"
	hedera_token "github.com/hnamzian/hedera-vault-plugin/src/core/hedera/token"
)

func (t_svc *TokenService) CreateToken(
	tokenCreation *hedera_token.FTokenCreation,
	operatorID,
	adminID,
	treasuryID,
	pauseID,
	freezeID,
	kycID,
	feeScheduleID,
	supplyID,
	wipeID string) (*hedera.TokenID, error) {
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

	admin_account, _ := t_svc.a_svc.GetAccount(adminID)
	if err != nil {
		return nil, fmt.Errorf("retreive admin account from vault failed: %s", err)
	}
	admin_key, _ := t_svc.k_svc.GetKey(admin_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive admin key from vault failed: %s", err)
	}
	adminPrivateKey, _ := hedera.PrivateKeyFromString(admin_key.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("parse admin private key failed: %s", err)
	}
	adminPublicKey, _ := hedera.PublicKeyFromString(admin_key.Publickey)
	if err != nil {
		return nil, fmt.Errorf("parse admin public key failed: %s", err)
	}

	treasury_account, _ := t_svc.a_svc.GetAccount(treasuryID)
	if err != nil {
		return nil, fmt.Errorf("retreive treasury account from vault failed: %s", err)
	}
	treasury_key, _ := t_svc.k_svc.GetKey(treasury_account.KeyID)
	if err != nil {
		return nil, fmt.Errorf("retreive treasury key from vault failed: %s", err)
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

	if pauseID != "" {
		pause_account, _ := t_svc.a_svc.GetAccount(pauseID)
		if err != nil {
			return nil, fmt.Errorf("retreive pause account from vault failed: %s", err)
		}
		pause_key, _ := t_svc.k_svc.GetKey(pause_account.KeyID)
		if err != nil {
			return nil, fmt.Errorf("retreive pause key from vault failed: %s", err)
		}
		pausePublicKey, _ := hedera.PublicKeyFromString(pause_key.Publickey)
		if err != nil {
			return nil, fmt.Errorf("parse pause public key failed: %s", err)
		}
		tokenCreation.PauseKey = pausePublicKey
	}

	if freezeID != "" {
		freeze_account, _ := t_svc.a_svc.GetAccount(freezeID)
		if err != nil {
			return nil, fmt.Errorf("retreive freeze account from vault failed: %s", err)
		}
		freeze_key, _ := t_svc.k_svc.GetKey(freeze_account.KeyID)
		if err != nil {
			return nil, fmt.Errorf("retreive freeze key from vault failed: %s", err)
		}
		freezePublicKey, _ := hedera.PublicKeyFromString(freeze_key.Publickey)
		if err != nil {
			return nil, fmt.Errorf("parse freeze public key failed: %s", err)
		}
		tokenCreation.FreezeKey = freezePublicKey

	}

	if kycID != "" {
		kyc_account, _ := t_svc.a_svc.GetAccount(kycID)
		if err != nil {
			return nil, fmt.Errorf("retreive kyc account from vault failed: %s", err)
		}
		kyc_key, _ := t_svc.k_svc.GetKey(kyc_account.KeyID)
		if err != nil {
			return nil, fmt.Errorf("retreive kyc key from vault failed: %s", err)
		}
		kycPublicKey, _ := hedera.PublicKeyFromString(kyc_key.Publickey)
		if err != nil {
			return nil, fmt.Errorf("parse kyc public key failed: %s", err)
		}
		tokenCreation.KycKey = kycPublicKey
	}

	if feeScheduleID != "" {
		feeSchedule_account, _ := t_svc.a_svc.GetAccount(feeScheduleID)
		if err != nil {
			return nil, fmt.Errorf("retreive feeSchedule account from vault failed: %s", err)
		}
		feeSchedule_key, _ := t_svc.k_svc.GetKey(feeSchedule_account.KeyID)
		if err != nil {
			return nil, fmt.Errorf("retreive feeSchedule key from vault failed: %s", err)
		}
		feeSchedulePublicKey, _ := hedera.PublicKeyFromString(feeSchedule_key.Publickey)
		if err != nil {
			return nil, fmt.Errorf("parse feeSchedule public key failed: %s", err)
		}
		tokenCreation.FeeScheduleKey = feeSchedulePublicKey
	}

	if supplyID != "" {
		supply_account, _ := t_svc.a_svc.GetAccount(supplyID)
		if err != nil {
			return nil, fmt.Errorf("retreive supply account from vault failed: %s", err)
		}
		supply_key, _ := t_svc.k_svc.GetKey(supply_account.KeyID)
		if err != nil {
			return nil, fmt.Errorf("retreive supply key from vault failed: %s", err)
		}
		supplyPublicKey, _ := hedera.PublicKeyFromString(supply_key.Publickey)
		if err != nil {
			return nil, fmt.Errorf("parse supply public key failed: %s", err)
		}
		tokenCreation.SupplyKey = supplyPublicKey
	}

	if wipeID != "" {
		wipe_account, _ := t_svc.a_svc.GetAccount(wipeID)
		if err != nil {
			return nil, fmt.Errorf("retreive wipe account from vault failed: %s", err)
		}
		wipe_key, _ := t_svc.k_svc.GetKey(wipe_account.KeyID)
		if err != nil {
			return nil, fmt.Errorf("retreive wipe key from vault failed: %s", err)
		}
		wipePublicKey, _ := hedera.PublicKeyFromString(wipe_key.Publickey)
		if err != nil {
			return nil, fmt.Errorf("parse wipe public key failed: %s", err)
		}
		tokenCreation.WipeKey = wipePublicKey
	}

	tokenCreation.AdminPrivateKey = adminPrivateKey
	tokenCreation.AdminPublicKey = adminPublicKey
	tokenCreation.TreasuryAccountID = treasuryAccountID
	tokenCreation.TreasuryPrivateKey = treasuryPrivateKey
	tokenCreation.TreasuryPublicKey = treasuryPublicKey

	client := hedera_client.
		NewClient(hedera_client.LocalTestNetClientConfig()).
		WithOperator(operatorAccountID, operatorPrivateKey).
		GetClient()
	ht := hedera_token.New(client)

	return ht.CreateFT(tokenCreation)
}
