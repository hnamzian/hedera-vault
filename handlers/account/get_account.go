package account

import (
	"context"
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	accountEntity "github.com/hnamzian/hedera-vault-plugin/entities/account"
	"github.com/hnamzian/hedera-vault-plugin/handlers/formatters"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *AccountHandler) handleRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	id := data.Get("id").(string)

	account_buf, err := storage.
		NewStorage(req.Storage).
		WithContext(ctx).
		WithKey(req.ClientToken, storage.Repository_Account, id).
		Read()
	if err != nil {
		return nil, errwrap.Wrapf("read account from storage failed: {{err}}", err)
	}
	account_vault, err := accountEntity.FromBytes(account_buf)
	if err != nil {
		return nil, errwrap.Wrapf("json decoding account buffer failed: {{err}}", err)
	}

	respData := formatters.FormatResponse(account_vault)

	return &logical.Response{
		Data: respData,
	}, nil
}
