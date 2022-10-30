package key

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	key_entity "github.com/hnamzian/hedera-vault-plugin/entities/key"
	"github.com/hnamzian/hedera-vault-plugin/handlers/formatters"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *KeyHandler) handleRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	id := data.Get("id").(string)

	// Decode the data
	key_buf, err := storage.
		NewStorage(req.Storage).
		WithContext(ctx).
		WithKey(req.ClientToken, storage.Repository_Key, id).
		Read()
	if err != nil {
		return nil, err
	}
	if key_buf == nil {
		resp := logical.ErrorResponse("No value at %v", req.MountPoint)
		return resp, nil
	}

	var key_vault key_entity.Key
	if err := json.Unmarshal(key_buf, &key_vault); err != nil {
		return nil, errwrap.Wrapf("parse key from vault failed: {{err}}", err)
	}
	response_data := formatters.FormatResponse(key_vault)

	// Generate the response
	resp := &logical.Response{
		Data: response_data,
	}

	return resp, nil
}
