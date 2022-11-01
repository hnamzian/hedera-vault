package paths

import (
	"context"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	tc "github.com/hnamzian/hedera-vault-plugin/src/token/controller"
)

type KeyPaths struct {
}

func NewKeyPaths() *KeyPaths {
	return &KeyPaths{}
}

func (kp KeyPaths) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			pathToken(),
		},
	)
}

func pathToken() *framework.Path {
	return &framework.Path{
		Pattern: "token/?",

		Fields: map[string]*framework.FieldSchema{
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
			"adminId": {
				Type: framework.TypeString,
				Required: true,
			},
			"treasuryId": {
				Type: framework.TypeString,
				Required: true,
			},
			"name": {
				Type: framework.TypeString,
				Required: true,
			},
			"symbol": {
				Type: framework.TypeString,
				Required: true,
			},
			"decimals": {
				Type: framework.TypeInt,
			},
			"initSupply": {
				Type: framework.TypeInt,
			},
			"treasuryAccountId": {
				Type: framework.TypeString,
				Required: true,
			},
			"treasuryKey": {
				Type: framework.TypeString,
			},
			"adminKey": {
				Type: framework.TypeString,
			},
			"kycKey": {
				Type: framework.TypeString,
			},
			"freezeKey": {
				Type: framework.TypeString,
			},
			"wipeKey": {
				Type: framework.TypeString,
			},
			"supplyKey": {
				Type: framework.TypeString,
			},
			"feeScheduleKey": {
				Type: framework.TypeString,
			},
			"pauseKey": {
				Type: framework.TypeString,
			},
			"customFees": {
				Type: framework.TypeString,
			},
			"maxSupply": {
				Type: framework.TypeInt,
			},
			"supplyType": {
				Type: framework.TypeString,
			},
			"freezeDefault": {
				Type: framework.TypeBool,
			},
			"expirationTime": {
				Type: framework.TypeString,
			},
			"autoRenewAccount": {
				Type: framework.TypeString,
			},
			"memo": {
				Type: framework.TypeString,
			},
		},

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: tc.CreateToken,
			},
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func handleExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	out, err := req.Storage.Get(ctx, req.Path)
	if err != nil {
		return false, errwrap.Wrapf("existence check failed: {{err}}", err)
	}

	return out != nil, nil
}