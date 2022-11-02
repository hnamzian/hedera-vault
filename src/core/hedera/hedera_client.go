package hedera_client

import (
	hedera "github.com/hashgraph/hedera-sdk-go/v2"
)

type HederaClient struct {
	client  *hedera.Client
	configs *HederaClientConfigs
}

type HederaClientConfigs struct {
	NetworkNodeAddress   string
	MirrorNodeAddress    string
	NetworkNodeAccountID hedera.AccountID
}

func LocalTestNetClientConfig() *HederaClientConfigs {
	acc, _ := hedera.AccountIDFromString("0.0.3")
	return &HederaClientConfigs{
		NetworkNodeAddress:   "127.0.0.1:50211",
		MirrorNodeAddress:    "127.0.0.1:5600",
		NetworkNodeAccountID: acc,
	}
}

func NewClient(configs *HederaClientConfigs) *HederaClient {
	node := make(map[string]hedera.AccountID, 1)
	node[configs.NetworkNodeAddress] = configs.NetworkNodeAccountID

	client := hedera.ClientForNetwork(node)
	client.SetMirrorNetwork([]string{configs.MirrorNodeAddress})

	return &HederaClient{client, configs}
}

func (hc *HederaClient) GetClient() *hedera.Client {
	return hc.client
}

func (hc *HederaClient) WithOperator(operatorAccountID hedera.AccountID, operatorPrivateKey hedera.PrivateKey) *HederaClient {
	hc.client.SetOperator(operatorAccountID, operatorPrivateKey)
	return hc
}