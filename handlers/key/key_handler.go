package key

import (
	"context"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type KeyHandler struct {
}

func NewKeyHandler() *KeyHandler {
	return &KeyHandler{}
}

func (h KeyHandler) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			{
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
						Callback: h.handleWrite,
					},
					logical.ReadOperation: &framework.PathOperation{
						Callback: h.handleRead,
					},
					logical.ListOperation: &framework.PathOperation{
						Callback: h.handleList,
					},
					logical.DeleteOperation: &framework.PathOperation{
						Callback: h.handleDelete,
					},
				},

				ExistenceCheck: h.handleExistenceCheck,
			},
			{
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
						Callback: h.handleImport,
					},
				},

				ExistenceCheck: h.handleExistenceCheck,
			},
		},
	)
}

func (h *KeyHandler) handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	// b.Logger().Debug("Handle Check", req.Path)
	out, err := req.Storage.Get(ctx, req.Path)
	if err != nil {
		return false, errwrap.Wrapf("existence check failed: {{err}}", err)
	}

	return out != nil, nil
}
