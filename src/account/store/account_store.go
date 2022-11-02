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
		return fmt.Errorf("encode accounts to bytes failed: %s", err)
	}

	entry := &logical.StorageEntry{
		Key:      as.getKey(id),
		Value:    account_buf,
		SealWrap: false,
	}
	if err := as.storage.Put(as.ctx, entry); err != nil {
		return fmt.Errorf("write account to storage failed: %s", err)
	}
	return nil
}

func (as *AccountStore) Read(id string) (*entity.Account, error) {
	entry, err := as.storage.Get(as.ctx, as.getKey(id))
	if err != nil {
		return nil, fmt.Errorf("fetch account from storage failed: %s", err)
	}

	if entry == nil {
		return nil, fmt.Errorf("account not found")
	}

	account_buf := entry.Value
	account, err := entity.FromBytes(account_buf)
	if err != nil {
		return nil, fmt.Errorf("encode account to byte failed: %s", err)
	}

	return account, nil
}

func (as *AccountStore) List() ([]string, error) {
	entries, err := as.storage.List(as.ctx, as.getKey(""))
	if err != nil {
		return nil, fmt.Errorf("fetch accounts from storage failed: %s", err)
	}
	return entries, nil
}

func (as *AccountStore) Delete(id string) error {
	return as.storage.Delete(as.ctx, as.getKey(id))
}

func (as *AccountStore) getKey(id string) string {
	return fmt.Sprintf("%s/%s/%s", as.clientToken, "accounts", id)
}
