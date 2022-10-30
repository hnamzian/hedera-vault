package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *KeyHandler) handleList(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	path := data.Get("path").(string)

	keys, err := storage.
		NewStorage(req.Storage).
		WithContext(ctx).
		WithKey(req.ClientToken, "").
		List()
	if err != nil {
		return nil, err
	}
	if keys == nil {
		resp := logical.ErrorResponse("No value at %v%v", req.MountPoint, path)
		return resp, nil
	}

	// Generate the response
	resp := logical.ListResponse(keys)

	return resp, nil
}
