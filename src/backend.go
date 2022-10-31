package backend

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"

	"github.com/hashicorp/vault/sdk/logical"
	kp "github.com/hnamzian/hedera-vault-plugin/src/key/paths"
	ap "github.com/hnamzian/hedera-vault-plugin/src/account/paths"
	// accountHandler "github.com/hnamzian/hedera-vault-plugin/handlers/account"
)

// backend wraps the backend framework and adds a map for storing key value pairs
type backend struct {
	*framework.Backend
}

var _ logical.Factory = Factory

// Factory configures and returns Mock backends
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b, err := newBackend()
	if err != nil {
		return nil, err
	}

	if conf == nil {
		return nil, fmt.Errorf("configuration passed into backend is nil")
	}

	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return b, nil
}

func newBackend() (*backend, error) {
	b := &backend{}

	b.Backend = &framework.Backend{
		Help:        strings.TrimSpace(mockHelp),
		BackendType: logical.TypeLogical,
		Paths: framework.PathAppend(
			kp.NewKeyPaths().Paths(),
			ap.Paths(),
		),
	}

	return b, nil
}

const mockHelp = `
The Mock backend is a dummy secrets backend that stores kv pairs in a map.
`
