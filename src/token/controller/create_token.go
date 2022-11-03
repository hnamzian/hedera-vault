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
	kycID := data.Get("kycId").(string)
	freezeID := data.Get("freezeId").(string)
	wipeID := data.Get("wipeId").(string)
	supplyID := data.Get("supplyId").(string)
	feeScheduleID := data.Get("feeScheduleId").(string)
	pauseID := data.Get("pauseId").(string)

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
		CustomFees:       nil,
		MaxSupply:        maxSupply,
		SupplyType:       supplyType,
		FreezeDefault:    freezeDefault,
		ExpirationTime:   expirationTime,
		AutoRenewAccount: autoRenewAccount,
		Memo:             memo,
	}

	t_svc := service.New(ctx, req.Storage, req.ClientToken)

	tokenID, err := t_svc.CreateToken(
		tokenCreation,
		operatorID,
		adminID,
		treasuryID,
		pauseID,
		freezeID,
		kycID,
		feeScheduleID,
		supplyID,
		wipeID)
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
			"kycKey":           kycID,
			"freezeKey":        freezeID,
			"wipeKey":          wipeID,
			"supplyID":         supplyID,
			"feeScheduleKey":   feeScheduleID,
			"pauseKey":         pauseID,
			"supplyType":       supplyType,
			"freezeDefault":    freezeDefault,
			"autoRenewAccount": autoRenewAccountString,
			"expirationTime":   expirationTimeString,
			"memo":             memo,
			"tokenID":          tokenID.String(),
		},
	}, nil
}
