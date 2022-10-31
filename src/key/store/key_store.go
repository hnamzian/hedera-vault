package key_store

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/key/entity"
)

type KeyStore struct {
	storage     logical.Storage
	ctx         context.Context
	clientToken string
}

func New(ctx context.Context, storage logical.Storage, clientToken string) *KeyStore {
	return &KeyStore{
		storage,
		ctx,
		clientToken,
	}
}

func (ks *KeyStore) Write(id string, key *key_entity.Key) error {
	key_buf, err := key.ToBytes()
	if err != nil {
		return err
	}

	entry := &logical.StorageEntry{
		Key:      ks.getKey(id),
		Value:    key_buf,
		SealWrap: false,
	}
	if err := ks.storage.Put(ks.ctx, entry); err != nil {
		return err
	}
	return nil
}

func (ks *KeyStore) Read(id string) (*key_entity.Key, error) {
	entry, err := ks.storage.Get(ks.ctx, ks.getKey(id))
	if err != nil {
		return nil, err
	}

	key_buf := entry.Value
	key, err := key_entity.FromBytes(key_buf)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (ks *KeyStore) List() ([]string, error) {
	entries, err := ks.storage.List(ks.ctx, ks.getKey(""))
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (ks *KeyStore) Delete(id string) error {
	return ks.storage.Delete(ks.ctx, ks.getKey(id))
}

func (ks *KeyStore) getKey(id string) string {
	return fmt.Sprintf("%s/%s/%s", ks.clientToken, "key", id)
}
