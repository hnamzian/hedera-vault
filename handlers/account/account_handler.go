package account

import (
	"context"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type AccountHandler struct {}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			{
				Pattern: "accounts/?$",

				Fields: map[string]*framework.FieldSchema{
					"id": {
						Type: framework.TypeString,
						Required: true,
					},
					"nextId": {
						Type: framework.TypeString,
						Required: true,
					},
				},

				Operations: map[logical.Operation]framework.OperationHandler{
					logical.ReadOperation: &framework.PathOperation{
						Callback: h.handleRead,
					},
					logical.CreateOperation: &framework.PathOperation{
						Callback: h.handleCreate,
					},
				},

				ExistenceCheck: h.handleExistenceCheck,
			},
			{
				Pattern: "accounts/import",

				Fields: map[string]*framework.FieldSchema{
					"id": {
						Type: framework.TypeString,
						Required: true,
					},
					"accountId": {
						Type: framework.TypeString,
						Required: true,
					},
					"keyId": {
						Type: framework.TypeString,
					},
				},

				Operations: map[logical.Operation]framework.OperationHandler{
					logical.CreateOperation: &framework.PathOperation{
						Callback: h.handleImport,
					},
				},

				ExistenceCheck: h.handleExistenceCheck,
			},
		},
	)
}

func (h *AccountHandler) handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	out, err := req.Storage.Get(ctx, req.Path)
	if err != nil {
		return false, errwrap.Wrapf("existence check failed: {{err}}", err)
	}

	return out != nil, nil
}
