package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/token/service"
)

func MintToken(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	logger := hclog.Default()


	tokenID := data.Get("tokenId").(string)
	amountString := data.Get("amount").(string)
	amount, _ := strconv.Atoi(amountString)
	supplyAccountID := data.Get("supplyAccountId").(string)
	operatorAccountID := data.Get("operatorAccountId").(string)

	logger.Info("MINT TOKEN XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "tokenId", tokenID, amount, supplyAccountID, operatorAccountID)


	logger.Debug("MINT TOKEN", "tokenId", tokenID)
	t_svc := service.New(ctx, req.Storage, req.ClientToken)
	minted, err := t_svc.MintToken(tokenID, uint64(amount), supplyAccountID, operatorAccountID)
	if err != nil {
		return nil, fmt.Errorf("minting token failed: %s", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"minted": minted.String(),
		},
	}, nil
}