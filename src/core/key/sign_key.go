package key

import "github.com/hashgraph/hedera-sdk-go/v2"

func Sign(priv string, message []byte) ([]byte, error) {
	privateKey, err := hedera.PrivateKeyFromString(priv)
	if err != nil {
		return nil, nil
	}
	signature := privateKey.Sign([]byte(message))

	return signature, nil
}
