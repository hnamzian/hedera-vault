package controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/token/service"
)

func DeleteToken(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	tokenID := data.Get("tokenId").(string)
	adminID := data.Get("adminId").(string)
	operatorID := data.Get("operatorId").(string)

	t_svc := service.New(ctx, req.Storage, req.ClientToken)
	minted, err := t_svc.DeleteToken(tokenID, adminID, operatorID)
	if err != nil {
		return nil, fmt.Errorf("minting token failed: %s", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"deleted": minted.String(),
		},
	}, nil
}