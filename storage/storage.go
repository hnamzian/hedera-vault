package storage

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/logical"
)

type Storage struct {
	req  *logical.Request
	ctx context.Context
	key string
	val []byte
}

func NewStorage(req *logical.Request) *Storage {
	return &Storage{req: req}
}

func (s *Storage) WithContext(ctx context.Context) *Storage {
	s.ctx = ctx
	return s
}

func (s *Storage) WithKey(clientToken, path, id string) *Storage {
	key := fmt.Sprintf("%s/%s/%s", clientToken, path, id)
	s.key = key
	return s
}

func (s *Storage) WithValue(buf []byte) *Storage {
	s.val = buf
	return s
}

func (s *Storage) Write() error {
	entry := &logical.StorageEntry{
		Key:      s.key,
		Value:    s.val,
		SealWrap: false,
	}
	if err := s.req.Storage.Put(s.ctx, entry); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Read() ([]byte, error) {
	entry, err := s.req.Storage.Get(s.ctx, s.key)
	if err != nil {
		return nil, err
	}
	fetchedData := entry.Value
	return fetchedData, nil
}

func (s *Storage) List() ([]string, error) {
	entries, err := s.req.Storage.List(s.ctx, s.key)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (s *Storage) Delete() error {
	return s.req.Storage.Delete(s.ctx, s.key)
}