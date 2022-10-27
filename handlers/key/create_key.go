package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/core/key"
	"github.com/hnamzian/hedera-vault-plugin/entities"
	"github.com/hnamzian/hedera-vault-plugin/storage"
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

	keyEntity, err := entities.FromKeyPair(id, keypair).ToBytes()
	if err != nil {
		return nil, errwrap.Wrapf("json encoding failed: {{err}}", err)
	}

	if err = storage.
	NewStorage(req).
	WithContext(ctx).
	WithKey(req.ClientToken, path, id).
	WithValue(keyEntity).Write(); err != nil {
		return nil, errwrap.Wrapf("store key pair failed: {{err}}", err)
	}

	// h.logger.Debug("Handle Write", "data", data, "\nreq", req)

	return nil, nil
}