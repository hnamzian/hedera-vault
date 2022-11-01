package service

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"
	account_service "github.com/hnamzian/hedera-vault-plugin/src/account/service"
	key_service "github.com/hnamzian/hedera-vault-plugin/src/key/service"
)

type TokenService struct {
	k_svc *key_service.KeyService
	a_svc *account_service.AccountService
}

func New(ctx context.Context, storage logical.Storage, clientToken string) *TokenService {
	return &TokenService{
		k_svc: key_service.New(ctx, storage, clientToken),
		a_svc: account_service.New(ctx, storage, clientToken),
	}
}
