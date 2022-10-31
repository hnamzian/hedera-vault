package service

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/account/store"
	key_service "github.com/hnamzian/hedera-vault-plugin/src/key/service"
)

type AccountService struct {
	storage *store.AccountStore
	k_svc   *key_service.KeyService
}

func New(ctx context.Context, storage logical.Storage, clientToken string) *AccountService {
	return &AccountService{
		storage: store.New(ctx, storage, clientToken),
		k_svc: key_service.New(ctx, storage, clientToken),
	}
}
