package controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hnamzian/hedera-vault-plugin/src/account/service"
	"github.com/hnamzian/hedera-vault-plugin/src/core/formatters"
)

func GetAccount(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	id := data.Get("id").(string)

	a_svc := service.New(ctx, req.Storage, req.ClientToken)

	account, err := a_svc.GetAccount(id)
	if err != nil {
		return nil, fmt.Errorf("read Account form storage failed: %s", err)
	}

	return &logical.Response{
		Data: formatters.FormatResponse(account),
	}, nil
}
