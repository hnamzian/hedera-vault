package key

import (
	"context"
	"encoding/json"
	"fmt"

	// "github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hnamzian/hedera-vault-plugin/core/key"
	key_entity "github.com/hnamzian/hedera-vault-plugin/entities/key"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *KeyHandler) handleSign(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	id := data.Get("id").(string)
	message := data.Get("message").(string)

	key_buf, err := storage.NewStorage(req.Storage).WithContext(ctx).WithKey(req.ClientToken, "", id).Read()
	if err != nil {
		return nil, errwrap.Wrapf("read key by ID failed: {{err}}", err)
	}
	if key_buf == nil {
		resp := logical.ErrorResponse("key is empty")
		return resp, nil
	}

	var key_vault key_entity.Key
	if err := json.Unmarshal(key_buf, &key_vault); err != nil {
		return nil, errwrap.Wrapf("parse key from vault failed: {{err}}", err)
	}

	signature, err := key.Sign(key_vault.PrivateKey, []byte(message))
	if err != nil {
		return nil, errwrap.Wrapf("signing message failed: {{err}}", err)
	}

	respData := make(map[string]interface{})
	respData["id"] = id
	respData["publicKey"] = key_vault.Publickey
	respData["message"] = message
	respData["signature"] = signature

	return &logical.Response{
		Data: respData,
	}, nil
}
