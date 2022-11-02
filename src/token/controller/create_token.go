package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	ht "github.com/hnamzian/hedera-vault-plugin/src/core/hedera/token"
	"github.com/hnamzian/hedera-vault-plugin/src/token/service"
)

func CreateToken(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	operatorID := data.Get("operatorId").(string)
	adminID := data.Get("adminId").(string)
	treasuryID := data.Get("treasuryId").(string)

	name := data.Get("name").(string)
	symbol := data.Get("symbol").(string)
	decimals := uint(data.Get("decimals").(int))
	initSupply := uint(data.Get("initSupply").(int))

	kycKeyString := data.Get("kycKey").(string)
	kycKey := hedera.PublicKey{}
	if len(kycKeyString) != 0 {
		kycKey, _ = hedera.PublicKeyFromString(kycKeyString)
	}

	freezeKeyString := data.Get("freezeKey").(string)
	freezeKey := hedera.PublicKey{}
	if len(freezeKeyString) != 0 {
		freezeKey, _ = hedera.PublicKeyFromString(freezeKeyString)
	}

	wipeKeyString := data.Get("wipeKey").(string)
	wipeKey := hedera.PublicKey{}
	if len(wipeKeyString) != 0 {
		wipeKey, _ = hedera.PublicKeyFromString(wipeKeyString)
	}

	supplyKeyString := data.Get("supplyKey").(string)
	supplyKey := hedera.PublicKey{}
	if len(supplyKeyString) != 0 {
		supplyKey, _ = hedera.PublicKeyFromString(supplyKeyString)
	}

	feeScheduleKeyString := data.Get("feeScheduleKey").(string)
	feeScheduleKey := hedera.PublicKey{}
	if len(feeScheduleKeyString) != 0 {
		feeScheduleKey, _ = hedera.PublicKeyFromString(feeScheduleKeyString)
	}

	pauseKeyString := data.Get("pauseKey").(string)
	pauseKey := hedera.PublicKey{}
	if len(pauseKeyString) != 0 {
		pauseKey, _ = hedera.PublicKeyFromString(pauseKeyString)
	}

	// customFees := data.Get("customFees").(string)
	maxSupply := uint(data.Get("maxSupply").(int))
	supplyType := data.Get("supplyType").(string)
	freezeDefault := data.Get("freezeDefault").(bool)

	expirationTimeString := data.Get("expirationTime").(string)
	expirationTime, _ := time.Parse("2006-01-02", expirationTimeString)

	autoRenewAccountString := data.Get("autoRenewAccount").(string)
	autoRenewAccount, _ := hedera.AccountIDFromString(autoRenewAccountString)

	memo := data.Get("memo").(string)

	tokenCreation := &ht.FTokenCreation{
		Name:             name,
		Symbol:           symbol,
		Decimals:         decimals,
		InitSupply:       initSupply,
		KycKey:           kycKey,
		FreezeKey:        freezeKey,
		WipeKey:          wipeKey,
		SupplyKey:        supplyKey,
		FeeScheduleKey:   feeScheduleKey,
		PauseKey:         pauseKey,
		CustomFees:       nil,
		MaxSupply:        maxSupply,
		SupplyType:       supplyType,
		FreezeDefault:    freezeDefault,
		ExpirationTime:   expirationTime,
		AutoRenewAccount: autoRenewAccount,
		Memo:             memo,
	}

	t_svc := service.New(ctx, req.Storage, req.ClientToken)

	tokenID, err := t_svc.CreateToken(tokenCreation, operatorID, adminID, treasuryID)
	if err != nil {
		return nil, fmt.Errorf("create token failed: %s", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"name":             name,
			"symbol":           symbol,
			"decimals":         decimals,
			"initSupply":       initSupply,
			"treasuryID":       treasuryID,
			"adminID":          adminID,
			"maxSupply":        maxSupply,
			"kycKey":           kycKeyString,
			"freezeKey":        freezeKeyString,
			"wipeKey":          wipeKeyString,
			"supplyKey":        supplyKeyString,
			"feeScheduleKey":   feeScheduleKeyString,
			"pauseKey":         pauseKeyString,
			"supplyType":       supplyType,
			"freezeDefault":    freezeDefault,
			"autoRenewAccount": autoRenewAccountString,
			"expirationTime":   expirationTimeString,
			"memo":             memo,
			"tokenID":          tokenID.String(),
		},
	}, nil
}
