package key_controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hnamzian/hedera-vault-plugin/src/core/formatters"
)

func Get(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	id := data.Get("id").(string)

	kc := New(ctx, req)

	key, err := kc.service.Get(id)
	if err != nil {
		return nil, fmt.Errorf("Read Key form storage failed: %s", err)
	}

	return &logical.Response{
		Data: formatters.FormatResponse(key),
	}, nil
}
