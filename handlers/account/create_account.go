package account

import (
	"context"
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	hc "github.com/hnamzian/hedera-vault-plugin/core/hedera"
	accountEntity "github.com/hnamzian/hedera-vault-plugin/entities/account"
	keyEntity "github.com/hnamzian/hedera-vault-plugin/entities/key"
	"github.com/hnamzian/hedera-vault-plugin/handlers/formatters"
	"github.com/hnamzian/hedera-vault-plugin/storage"
)

func (h *AccountHandler) handleCreate(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if req.ClientToken == "" {
		return nil, fmt.Errorf("client token empty")
	}

	// Check to make sure that kv pairs provpathed
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("data must be provided to store in secret")
	}

	id := data.Get("id").(string)
	nextId := data.Get("nextId").(string)

	operator_account_buf, err := storage.NewStorage(req.Storage).WithContext(ctx).WithKey(req.ClientToken, "", id).Read()
	if err != nil {
		return nil, errwrap.Wrapf("read account failed: {{err}}", err)
	}

	operator_account, err := accountEntity.FromBytes(operator_account_buf)
	if err != nil {
		return nil, errwrap.Wrapf("decode account failed: {{err}}", err)
	}

	operator_key_buf, err := storage.NewStorage(req.Storage).WithContext(ctx).WithKey(req.ClientToken, "", operator_account.KeyID).Read()
	if err != nil {
		return nil, errwrap.Wrapf("read account failed: {{err}}", err)
	}

	operator_key, err := keyEntity.FromBytes(operator_key_buf)
	if err != nil {
		return nil, errwrap.Wrapf("decode account failed: {{err}}", err)
	}

	client := hc.NewClient(hc.LocalTestNetClientConfig())

	operator_account_id, err := hedera.AccountIDFromString(operator_account.AccountID)
	if err != nil {
		return nil, errwrap.Wrapf("invalid account ID: {{err}}", err)
	}
	operator_private_key, err := hedera.PrivateKeyFromString(operator_key.PrivateKey)
	if err != nil {
		return nil, errwrap.Wrapf("invalid private key: {{err}}", err)
	}

	new_account_id, err := client.WithOperator(operator_account_id, operator_private_key).NewAccount(operator_private_key)
	if err != nil {
		return nil, errwrap.Wrapf("create new account failed: {{err}}", err)
	}

	new_account := accountEntity.NewAccount(nextId, new_account_id.String(), operator_account.KeyID)
	account_buf, err := new_account.ToBytes()
	if err != nil {
		return nil, errwrap.Wrapf("encode json failed: {{err}}", err)
	}

	err = storage.
		NewStorage(req.Storage).
		WithContext(ctx).
		WithKey(req.ClientToken, "", nextId).
		WithValue(account_buf).Write()
	if err != nil {
		return nil, errwrap.Wrapf("write account to stoage failed: {{err}}", err)
	}

	respData := formatters.FormatResponse(new_account)

	return &logical.Response{
		Data: respData,
	}, nil
}