package controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/token/service"
)

func GrantKyc(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	tokenID := data.Get("tokenId").(string)
	userID := data.Get("userId").(string)
	kycID := data.Get("kycId").(string)
	operatorID := data.Get("operatorId").(string)

	t_svc := service.New(ctx, req.Storage, req.ClientToken)
	minted, err := t_svc.GrantKyc(tokenID, userID, kycID, operatorID)
	if err != nil {
		return nil, fmt.Errorf("minting token failed: %s", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"kyc_granted": minted.String(),
		},
	}, nil
}
