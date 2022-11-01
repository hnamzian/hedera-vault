package key_controller

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func List(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	kc := New(ctx, req)

	keys, err := kc.service.List()
	if err != nil {
		return nil, fmt.Errorf("list key failed: %s", err)
	}

	return logical.ListResponse(keys), err
}
