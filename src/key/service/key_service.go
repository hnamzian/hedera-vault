package key_service

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/key/store"
)

type KeyService struct {
	storage *key_store.KeyStore
}

func New(ctx context.Context, storage logical.Storage, clientToken string) *KeyService {
	return &KeyService{
		storage: key_store.New(ctx, storage, clientToken),
	}
}

