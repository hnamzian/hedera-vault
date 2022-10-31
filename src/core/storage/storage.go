package storage

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"
)

type ContextStorage struct {
	storage logical.Storage
	ctx context.Context
}

func New(ctx context.Context, storage logical.Storage) *ContextStorage {
	return &ContextStorage{
		storage,
		ctx,
	}
}