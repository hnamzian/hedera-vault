package handlers

import "github.com/hnamzian/hedera-vault-plugin/handlers/key"

type Handlers struct {
	Key *key.KeyHandler
}

func NewHandler() *Handlers {
	return &Handlers{
		Key: key.NewKeyHandler(),
	}
}