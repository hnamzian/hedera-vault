package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/token/service"
)

func WipeToken(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	tokenID := data.Get("tokenId").(string)
	userID := data.Get("userId").(string)
	amountString := data.Get("amount").(string)
	amount, _ := strconv.Atoi(amountString)
	wipeID := data.Get("wipeId").(string)
	operatorID := data.Get("operatorId").(string)

	t_svc := service.New(ctx, req.Storage, req.ClientToken)
	minted, err := t_svc.WipeToken(tokenID, userID, uint64(amount), wipeID, operatorID)
	if err != nil {
		return nil, fmt.Errorf("minting token failed: %s", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"wiped": minted.String(),
		},
	}, nil
}
