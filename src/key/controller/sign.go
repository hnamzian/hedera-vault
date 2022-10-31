package key_controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/core/formatters"
)

func Sign(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	id := data.Get("id").(string)
	message := data.Get("message").(string)

	kc := New(ctx, req)
	key_vault, signature, err := kc.service.Sign(id, message)
	if err != nil {
		return nil, err
	}

	resp := formatters.FormatResponse(key_vault)
	resp["signature"] = signature

	return &logical.Response{
		Data: resp,
	}, nil
}
