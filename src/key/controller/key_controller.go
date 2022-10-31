package key_controller

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hnamzian/hedera-vault-plugin/src/key/service"
)

type KeyController struct {
	service *key_service.KeyService
}

func New(ctx context.Context, req *logical.Request) *KeyController {
	return &KeyController{
		key_service.New(ctx, req.Storage, req.ClientToken),
	}
}
