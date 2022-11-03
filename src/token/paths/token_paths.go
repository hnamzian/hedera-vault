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
			pathTokenBurn(),
			pathTokenDelete(),
			pathTokenAssociate(),
			pathTokenDissociate(),
			pathTokenFreeze(),
			pathTokenUnfreeze(),
			pathTokenGrantKyc(),
			pathTokenRevokeKyc(),
			pathTokenPause(),
			pathTokenUnpause(),
			pathTokenWipe(),
		},
	)
}

func pathToken() *framework.Path {
	return &framework.Path{
		Pattern: "token/?",

		Fields: map[string]*framework.FieldSchema{
			"type": {
				Type: framework.TypeString,
				Required: true,
			},
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
			"kycId": {
				Type: framework.TypeString,
			},
			"freezeId": {
				Type: framework.TypeString,
			},
			"wipeId": {
				Type: framework.TypeString,
			},
			"supplyId": {
				Type: framework.TypeString,
			},
			"feeScheduleId": {
				Type: framework.TypeString,
			},
			"pauseId": {
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
			"autoRenewId": {
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
			"supplyId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
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
func pathTokenBurn() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/burn",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"amount": {
				Type: framework.TypeString,
				Required: true,
			},
			"supplyId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.BurnToken,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathTokenDelete() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/delete",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"adminId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.DeleteToken,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathTokenAssociate() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/associate",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"userId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.AssociateWithToken,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}
func pathTokenDissociate() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/dissociate",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"userId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.DissociateWithToken,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathTokenFreeze() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/freeze",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"userId": {
				Type: framework.TypeString,
				Required: true,
			},
			"kycId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.FreezeAccount,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}
func pathTokenUnfreeze() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/unfreeze",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"userId": {
				Type: framework.TypeString,
				Required: true,
			},
			"kycId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.UnfreezeAccount,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathTokenGrantKyc() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/grant_kyc",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"userId": {
				Type: framework.TypeString,
				Required: true,
			},
			"kycId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.GrantKyc,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}
func pathTokenRevokeKyc() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/revoke_kyc",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"userId": {
				Type: framework.TypeString,
				Required: true,
			},
			"kycId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.RevokeKyc,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathTokenPause() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/pause",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"amount": {
				Type: framework.TypeString,
				Required: true,
			},
			"pauseId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.PauseToken,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}
func pathTokenUnpause() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/unpause",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"amount": {
				Type: framework.TypeString,
				Required: true,
			},
			"pauseId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.UnpauseToken,
		},

		ExistenceCheck: handleExistenceCheck,
	}
}

func pathTokenWipe() *framework.Path {
	return &framework.Path{
		Pattern: "tokens/wipe",

		Fields: map[string]*framework.FieldSchema{
			"tokenId": {
				Type: framework.TypeString,
				Required: true,
			},
			"userId": {
				Type: framework.TypeString,
				Required: true,
			},
			"amount": {
				Type: framework.TypeString,
				Required: true,
			},
			"wipeId": {
				Type: framework.TypeString,
				Required: true,
			},
			"operatorId": {
				Type: framework.TypeString,
				Required: true,
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: tc.WipeToken,
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
