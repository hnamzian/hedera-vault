package paths

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	tc "github.com/hnamzian/hedera-vault-plugin/src/token/controller"
)

type TokenPaths struct {
}

func NewTokenPaths() *TokenPaths {
	return &TokenPaths{}
}

func (tp TokenPaths) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			pathToken(),
			pathTokenMint(),
		},
	)
}

func pathToken() *framework.Path {
	return &framework.Path{
		Pattern: "token/?",

		Fields: map[string]*framework.FieldSchema{
			"operatorId": {
				Type:     framework.TypeString,
				Required: true,
			},
			"adminId": {
				Type:     framework.TypeString,
				Required: true,
			},
			"treasuryId": {
				Type:     framework.TypeString,
				Required: true,
			},
			"name": {
				Type:     framework.TypeString,
				Required: true,
			},
			"symbol": {
				Type:     framework.TypeString,
				Required: true,
			},
			"decimals": {
				Type: framework.TypeInt,
			},
			"initSupply": {
				Type: framework.TypeInt,
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
			"supplyId": {
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

func pathTokenMint() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/mint",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"amount": {
				Type: framework.TypeString,
				Required: true,
			},
			"supplyAccountId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorAccountId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.MintToken,
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
