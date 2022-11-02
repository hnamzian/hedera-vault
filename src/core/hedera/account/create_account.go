package account

import "github.com/hashgraph/hedera-sdk-go/v2"

func (ha *Account) NewAccount(priv hedera.PrivateKey) (*hedera.AccountID, error) {
	tr, err := hedera.NewAccountCreateTransaction().SetKey(priv).Execute(ha.client)
	if err != nil {
		return nil, err
	}

	receipt, err := tr.GetReceipt(ha.client)
	if err != nil {
		return nil, err
	}

	acc := receipt.AccountID

	return acc, nil
}