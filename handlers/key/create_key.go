package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/core/key"
	keyEntity "github.com/hnamzian/hedera-vault-plugin/entities/key"
	"github.com/hnamzian/hedera-vault-plugin/storage"
	"github.com/hnamzian/hedera-vault-plugin/handlers/formatters"
)

func (h *KeyHandler) handleWrite(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {

	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	path := data.Get("path").(string)
	id := data.Get("id").(string)
	algo := data.Get("algo").(string)
	curve := data.Get("curve").(string)

	keypair, err := key.CreateKey(algo, curve)
	if err != nil {
		return nil, errwrap.Wrapf("generate key pair failed: {{err}}", err)
	}

	key_vault := keyEntity.FromKeyPair(id, keypair)
	keybuf, err := key_vault.ToBytes()
	if err != nil {
		return nil, errwrap.Wrapf("json encoding failed: {{err}}", err)
	}

	if err = storage.
		NewStorage(req).
		WithContext(ctx).
		WithKey(req.ClientToken, path, id).
		WithValue(keybuf).
		Write(); err != nil {
		return nil, errwrap.Wrapf("store key pair failed: {{err}}", err)
	}

	response_data := formatters.FormatResponse(key_vault)

	return &logical.Response{
		Data: response_data,
	}, nil
}
