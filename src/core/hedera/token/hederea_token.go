package token

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type Token struct {
	client *hedera.Client
}

func New(client *hedera.Client) *Token {
	return &Token{client}
}
