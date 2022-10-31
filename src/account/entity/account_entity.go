package entity

import "encoding/json"

type Account struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	KeyID     string `json:"keyId"`
}

func New(id, accountID, keyID string) *Account {
	return &Account{
		ID: id,
		AccountID: accountID,
		KeyID: keyID,
	}
}

func FromBytes(ab []byte) (*Account, error) {
	var account Account
	if err := json.Unmarshal(ab, &account); err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *Account) ToBytes() ([]byte, error) {
	ab, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return ab, nil
}