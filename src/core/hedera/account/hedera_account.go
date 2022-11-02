package account

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type Account struct {
	client *hedera.Client
}

func New(client *hedera.Client) *Account {
	return &Account{client}
}
