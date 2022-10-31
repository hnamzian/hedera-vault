package key_controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/core/formatters"
)

func Import(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	id := data.Get("id").(string)
	priv := data.Get("privateKey").(string)
	algo := data.Get("algo").(string)
	curve := data.Get("curve").(string)

	kc := New(ctx, req)
	key_vault, err := kc.service.ImportKey(id, priv, algo, curve)
	if err != nil {
		return nil, err
	}

	return &logical.Response{
		Data: formatters.FormatResponse(key_vault),
	}, nil
}
