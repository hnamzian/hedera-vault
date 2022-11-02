package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (ht *Token) DissociateWithToken(tokenIDString, accountIDString, accountKeyString string) (*hedera.Status, error) {
	tokenID, err := hedera.TokenIDFromString(tokenIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid tokenID: %s", err)
	}

	accountID, err := hedera.AccountIDFromString(accountIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	accountKey, err := hedera.PrivateKeyFromString(accountKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	transaction, err := hedera.
		NewTokenDissociateTransaction().
		SetTokenIDs(tokenID).
		SetAccountID(accountID).
		FreezeWith(ht.client)
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(accountKey).
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
