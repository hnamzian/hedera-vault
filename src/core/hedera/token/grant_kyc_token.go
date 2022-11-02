package token

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func (ht *Token) GrantKyc(tokenIDString, accountIDString, kycKeyString string) (*hedera.Status, error) {
	tokenID, err := hedera.TokenIDFromString(tokenIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid tokenID: %s", err)
	}

	accountID, err := hedera.AccountIDFromString(accountIDString)
	if err != nil {
		return nil, fmt.Errorf("invalid accountID")
	}

	kycKey, err := hedera.PrivateKeyFromString(kycKeyString)
	if err != nil {
		return nil, fmt.Errorf("invalid supply key: %s", err)
	}

	transaction, err := hedera.
		NewTokenGrantKycTransaction().
		SetTokenID(tokenID).
		SetAccountID(accountID).
		FreezeWith(ht.client)
	if err != nil {
		return nil, fmt.Errorf("prepare transaction failed: %s", err)
	}

	response, err := transaction.
		Sign(kycKey).
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
