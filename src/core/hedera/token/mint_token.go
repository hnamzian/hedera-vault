package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (ht *Token) MintToken(tokenIDString string, amount uint64, supplyKeyString string) (*hedera.Status, error) {
	tokenID, err := hedera.TokenIDFromString(tokenIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid tokenID: %s", err)
	}

	supplyKey, err := hedera.PrivateKeyFromString(supplyKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	transaction, err := hedera.
		NewTokenMintTransaction().
		SetTokenID(tokenID).
		SetAmount(amount).
		FreezeWith(ht.client)
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(supplyKey).
		Execute(ht.client)
	if err != nil {
		return nil, fmt.Errorf("execute transaction failed: %s", err)
	}

	receipt, err := response.GetReceipt(ht.client)
	if err != nil {
		return nil, fmt.Errorf("retreive transaction response failed: %s", err)
	}

	return &receipt.Status, nil
}
