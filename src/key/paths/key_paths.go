package key

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	kc "github.com/hnamzian/hedera-vault-plugin/src/key/controller"
)

type KeyPaths struct {
}

func NewKeyPaths() *KeyPaths {
	return &KeyPaths{}
}

func (kp KeyPaths) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			pathKeys(),
			pathImportKeys(),
			pathSign(),
		},
	)
}

func pathKeys() *framework.Path {
	return &framework.Path{
		Pattern: "keys/?",

		Fields: map[string]*framework.FieldSchema{
			"path": {
				Type:        framework.TypeString,
				Description: "Specifies the path of the secret.",
			},
			"id": {
				Type: framework.TypeString,
			},
			"algo": {
				Type: framework.TypeString,
			},
			"curve": {
				Type: framework.TypeString,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: kc.Create,
			},
			logical.ReadOperation: &framework.PathOperation{
				Callback: kc.Get,
			},
			logical.ListOperation: &framework.PathOperation{
				Callback: kc.List,
			},
			logical.DeleteOperation: &framework.PathOperation{
				Callback: kc.Delete,
			},
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathImportKeys() *framework.Path {
	return &framework.Path{
		Pattern: "keys/import",

		Fields: map[string]*framework.FieldSchema{
			"path": {
				Type:        framework.TypeString,
				Description: "Specifies the path of the secret.",
			},
			"id": {
				Type: framework.TypeString,
			},
			"privateKey": {
				Type: framework.TypeString,
			},
			"algo": {
				Type: framework.TypeString,
			},
			"curve": {
				Type: framework.TypeString,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: kc.Import,
			},
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathSign() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("keys/%s/sign", framework.GenericNameRegex("id")),

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"message": {
				Type:     framework.TypeString,
				Required: true,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: kc.Sign,
			},
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	out, err := req.Storage.Get(ctx, req.Path)
	if err != nil {
		return false, fmt.Errorf("existence check failed: %s", err)
	}

	return out != nil, nil
}
