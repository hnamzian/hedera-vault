package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *KeyHandler) handleDelete(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	id := data.Get("id").(string)

	// Remove entry for specified path
	if err := storage.
		NewStorage(req.Storage).
		WithContext(ctx).
		WithKey(req.ClientToken, id).
		Delete(); err != nil {
		return nil, fmt.Errorf("delete key from storage failed %s", err)
	}

	// h.logger.Debug("Handle Update", "data", data, "\nreq", req)

	return nil, nil
}
