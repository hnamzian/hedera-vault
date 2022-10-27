package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	keyPair "github.com/hnamzian/hedera-vault-plugin/core/key"
	keyEntity "github.com/hnamzian/hedera-vault-plugin/entities/key"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *KeyHandler) handleImport(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	path := data.Get("path").(string)
	id := data.Get("id").(string)
	priv := data.Get("privateKey").(string)
	algo := data.Get("algo").(string)
	keypair, err := keyPair.FromPrivateKey(priv, algo)
	if err != nil {
		return nil, errwrap.Wrapf("wrap key pair failed: {{err}}:", err)
	}

	fmt.Printf("%s%s%s%s", path, id, priv, algo)

	keybuf, err := keyEntity.FromKeyPair(id, keypair).ToBytes()
	if err != nil {
		return nil, errwrap.Wrapf("json encoding failed: {{err}}", err)
	}

	if err = storage.NewStorage(req).WithContext(ctx).WithKey(req.ClientToken, path, id).WithValue(keybuf).Write(); err != nil {
		return nil, errwrap.Wrapf("write to storage failed: {{err}}", err)
	}

	return nil, nil
}
