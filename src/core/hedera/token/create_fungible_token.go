package token

import (
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type FTokenCreation struct {
	Name               string            `json:"name"`
	Symbol             string            `json:"symbol"`
	Decimals           uint              `json:"decimal"`
	InitSupply         uint              `json:"initSupply"`
	TreasuryAccountID  hedera.AccountID  `json:"treasuryAccountID"`
	TreasuryPublicKey  hedera.PublicKey  `json:"treasuryPublicKey"`
	TreasuryPrivateKey hedera.PrivateKey `json:"treasuryPrivateKey"`
	AdminPublicKey     hedera.PublicKey  `json:"adminPublicKey"`
	AdminPrivateKey    hedera.PrivateKey `json:"adminPrivateKey"`
	KycKey             hedera.PublicKey  `json:"kycKey"`
	FreezeKey          hedera.PublicKey  `json:"freezeKey"`
	WipeKey            hedera.PublicKey  `json:"wipeKey"`
	SupplyKey          hedera.PublicKey  `json:"supplyKey"`
	FeeScheduleKey     hedera.PublicKey  `json:"feeScheduleKey"`
	PauseKey           hedera.PublicKey  `json:"pauseKey"`
	CustomFees         []hedera.Fee      `json:"customFees"`
	MaxSupply          uint              `json:"maxSupply"`
	SupplyType         string            `json:"supplyType"`
	FreezeDefault      bool              `json:"freezeDefault"`
	ExpirationTime     time.Time         `json:"expirationTime"`
	AutoRenewAccount   hedera.AccountID  `json:"autoRenewAccount"`
	Memo               string            `json:"memo"`
}

func (t *Token) CreateFT(tokenCreation *FTokenCreation) (*hedera.TokenID, error) {
	tokenCreateTransaction := hedera.NewTokenCreateTransaction().
		SetTokenName(tokenCreation.Name).
		SetTokenSymbol(tokenCreation.Symbol).
		SetTreasuryAccountID(tokenCreation.TreasuryAccountID)

	if tokenCreation.Decimals > 0 {
		tokenCreateTransaction = tokenCreateTransaction.SetDecimals(uint(tokenCreation.Decimals))
	}
	if tokenCreation.InitSupply > 0 {
		tokenCreateTransaction = tokenCreateTransaction.SetInitialSupply(uint64(tokenCreation.InitSupply))
	}
	if tokenCreation.AdminPublicKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetAdminKey(tokenCreation.AdminPublicKey)
	}
	if tokenCreation.KycKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetKycKey(tokenCreation.KycKey)
	}
	if tokenCreation.FreezeKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetFreezeKey(tokenCreation.KycKey)
	}
	if tokenCreation.SupplyKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetSupplyKey(tokenCreation.KycKey)
	}
	if tokenCreation.FeeScheduleKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetFeeScheduleKey(tokenCreation.KycKey)
	}
	if tokenCreation.PauseKey != (hedera.PublicKey{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetPauseKey(tokenCreation.KycKey)
	}
	if tokenCreation.MaxSupply != 0 {
		tokenCreateTransaction = tokenCreateTransaction.
			SetSupplyType(hedera.TokenSupplyTypeFinite).
			SetMaxSupply(int64(tokenCreation.MaxSupply))
	}
	if len(tokenCreation.CustomFees) > 0 {
		tokenCreateTransaction = tokenCreateTransaction.SetCustomFees(tokenCreation.CustomFees)
	}
	if tokenCreation.AutoRenewAccount != (hedera.AccountID{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetAutoRenewAccount(tokenCreation.AutoRenewAccount)
	}
	if len(tokenCreation.Memo) > 0 {
		tokenCreateTransaction = tokenCreateTransaction.SetTokenMemo(tokenCreation.Memo)
	}
	if tokenCreation.ExpirationTime != (time.Time{}) {
		tokenCreateTransaction = tokenCreateTransaction.SetExpirationTime(tokenCreation.ExpirationTime)
	}
	tokenCreateTransaction = tokenCreateTransaction.SetFreezeDefault(tokenCreation.FreezeDefault)

	transaction, err := tokenCreateTransaction.FreezeWith(t.client)
	if err != nil {
		return nil, fmt.Errorf("freeze transaction failed: %s", err)
	}

	txResponse, err := transaction.
		Sign(tokenCreation.AdminPrivateKey).
		Sign(hedera.PrivateKey(tokenCreation.TreasuryPrivateKey)).
		Execute(t.client)
	if err != nil {
		return nil, fmt.Errorf("execute transaction failed: %s", err)
	}

	receipt, err := txResponse.GetReceipt(t.client)
	if err != nil {
		return nil, fmt.Errorf("get transaction receipt failed: %s", err)
	}

	tokenID := receipt.TokenID

	return tokenID, nil
}
