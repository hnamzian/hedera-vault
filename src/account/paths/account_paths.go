package paths

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hnamzian/hedera-vault-plugin/src/account/controller"
)

func Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			pathAccounts(),
			pathImportAccounts(),
		},
	)
}

func pathAccounts() *framework.Path {
	return &framework.Path{
		Pattern: "accounts/?$",

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"newId": {
				Type:     framework.TypeString,
				Required: true,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: &framework.PathOperation{
				Callback: controller.GetAccount,
			},
			logical.CreateOperation: &framework.PathOperation{
				Callback: controller.CreateAccount,
			},
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathImportAccounts() *framework.Path {
	return &framework.Path{
		Pattern: "accounts/import",

		Fields: map[string]*framework.FieldSchema{
			"id": {
				Type:     framework.TypeString,
				Required: true,
			},
			"accountId": {
				Type:     framework.TypeString,
				Required: true,
			},
			"keyId": {
				Type: framework.TypeString,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: controller.ImportAccount,
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
