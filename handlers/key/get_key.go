package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/helper/jsonutil"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *KeyHandler) handleRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	path := data.Get("path").(string)
	id := data.Get("id").(string)

	// Decode the data
	key_buf, err := storage.NewStorage(req).WithContext(ctx).WithKey(req.ClientToken, path, id).Read()
	if err != nil {
		return nil, err
	}
	if key_buf == nil {
		resp := logical.ErrorResponse("No value at %v%v", req.MountPoint, path)
		return resp, nil
	}

	var key map[string]interface{}
	if err := jsonutil.DecodeJSON(key_buf, &key); err != nil {
		return nil, errwrap.Wrapf("json decoding failed: {{err}}", err)
	}

	// Generate the response
	resp := &logical.Response{
		Data: key,
	}

	return resp, nil
}
