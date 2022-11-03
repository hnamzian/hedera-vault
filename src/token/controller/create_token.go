package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	ht "github.com/hnamzian/hedera-vault-plugin/src/core/hedera/token"
	"github.com/hnamzian/hedera-vault-plugin/src/token/service"
)

func CreateToken(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	tokenType := data.Get("type").(string)
	name := data.Get("name").(string)
	symbol := data.Get("symbol").(string)
	decimals := uint(data.Get("decimals").(int))
	initSupply := uint(data.Get("initSupply").(int))
	operatorID := data.Get("operatorId").(string)
	adminID := data.Get("adminId").(string)
	treasuryID := data.Get("treasuryId").(string)
	kycID := data.Get("kycId").(string)
	freezeID := data.Get("freezeId").(string)
	wipeID := data.Get("wipeId").(string)
	supplyID := data.Get("supplyId").(string)
	feeScheduleID := data.Get("feeScheduleId").(string)
	pauseID := data.Get("pauseId").(string)
	autoRenewID := data.Get("autoRenewId").(string)

	// customFees := data.Get("customFees").(string)
	maxSupply := uint(data.Get("maxSupply").(int))
	supplyType := data.Get("supplyType").(string)
	freezeDefault := data.Get("freezeDefault").(bool)

	expirationTimeString := data.Get("expirationTime").(string)
	expirationTime, _ := time.Parse("2006-01-02", expirationTimeString)

	memo := data.Get("memo").(string)

	tokenCreation := &ht.FTokenCreation{
		Type:           tokenType,
		Name:           name,
		Symbol:         symbol,
		Decimals:       decimals,
		InitSupply:     initSupply,
		CustomFees:     nil,
		MaxSupply:      maxSupply,
		SupplyType:     supplyType,
		FreezeDefault:  freezeDefault,
		ExpirationTime: expirationTime,
		Memo:           memo,
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
		wipeID,
		autoRenewID)
	if err != nil {
		return nil, fmt.Errorf("create token failed: %s", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"type":           tokenType,
			"name":           name,
			"symbol":         symbol,
			"decimals":       decimals,
			"initSupply":     initSupply,
			"treasuryID":     treasuryID,
			"adminID":        adminID,
			"maxSupplyID":    maxSupply,
			"kycID":          kycID,
			"freezeID":       freezeID,
			"wipeID":         wipeID,
			"supplyID":       supplyID,
			"feeScheduleID":  feeScheduleID,
			"pauseID":        pauseID,
			"autoRenewID":    autoRenewID,
			"supplyType":     supplyType,
			"freezeDefault":  freezeDefault,
			"expirationTime": expirationTimeString,
			"memo":           memo,
			"tokenID":        tokenID.String(),
		},
	}, nil
}
