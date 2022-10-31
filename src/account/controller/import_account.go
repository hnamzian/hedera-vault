package controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/account/service"
	"github.com/hnamzian/hedera-vault-plugin/src/core/formatters"
)

func ImportAccount(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	id := data.Get("id").(string)
	keyID := data.Get("keyId").(string)
	accountID := data.Get("accountId").(string)
	

	a_svc := service.New(ctx, req.Storage, req.ClientToken)
	account, err := a_svc.ImportAccount(id, keyID, accountID)
	if err != nil {
		return nil, err
	}

	return &logical.Response{
		Data: formatters.FormatResponse(account),
	}, nil
}