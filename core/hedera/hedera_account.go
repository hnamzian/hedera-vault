package hedera_client

import "github.com/hashgraph/hedera-sdk-go/v2"

func (hc *HederaClient) NewAccount(priv hedera.PrivateKey) (*hedera.AccountID, error) {
	tr, err := hedera.NewAccountCreateTransaction().SetKey(priv).Execute(hc.client)
	if err != nil {
		return nil, err
	}

	receipt, err := tr.GetReceipt(hc.client)
	if err != nil {
		return nil, err
	}

	acc := receipt.AccountID

	return acc, nil
}