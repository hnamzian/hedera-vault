package store

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/account/entity"
)

type AccountStore struct {
	storage     logical.Storage
	ctx         context.Context
	clientToken string
}

func New(ctx context.Context, storage logical.Storage, clientToken string) *AccountStore {
	return &AccountStore{
		storage,
		ctx,
		clientToken,
	}
}

func (as *AccountStore) Write(id string, account *entity.Account) error {
	account_buf, err := account.ToBytes()
	if err != nil {
		return err
	}

	entry := &logical.StorageEntry{
		Key:      as.getKey(id),
		Value:    account_buf,
		SealWrap: false,
	}
	if err := as.storage.Put(as.ctx, entry); err != nil {
		return err
	}
	return nil
}

func (as *AccountStore) Read(id string) (*entity.Account, error) {
	entry, err := as.storage.Get(as.ctx, as.getKey(id))
	if err != nil {
		return nil, err
	}

	account_buf := entry.Value
	account, err := entity.FromBytes(account_buf)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (as *AccountStore) List() ([]string, error) {
	entries, err := as.storage.List(as.ctx, as.getKey(""))
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (as *AccountStore) Delete(id string) error {
	return as.storage.Delete(as.ctx, as.getKey(id))
}

func (as *AccountStore) getKey(id string) string {
	return fmt.Sprintf("%s/%s/%s", as.clientToken, "accounts", id)
}
