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

func (h *AccountHandler) handleImport(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	id := data.Get("id").(string)
	keyID := data.Get(("keyId")).(string)
	accountID := data.Get("accountId").(string)

	newAccount := accountEntity.NewAccount(id, accountID, keyID)
	account_buf, err := newAccount.ToBytes()
	if err != nil {
		return nil, errwrap.Wrapf("json encoding account failed: {{err}}", err)
	}

	err = storage.
		NewStorage(req.Storage).
		WithContext(ctx).
		WithKey(req.ClientToken, "", id).
		WithValue(account_buf).
		Write()
	if err != nil {
		return nil, errwrap.Wrapf("storing account failed: {{err}}", err)
	}

	respData := formatters.FormatResponse(newAccount)

	return &logical.Response{
		Data: respData,
	}, nil
}
