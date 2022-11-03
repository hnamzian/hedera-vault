GOARCH = amd64

UNAME = $(shell uname -s)

ifndef OS
	ifeq ($(UNAME), Linux)
		OS = linux
	else ifeq ($(UNAME), Darwin)
		OS = darwin
	endif
endif

.DEFAULT_GOAL := all

all: fmt build start

build:
	GOOS=$(OS) GOARCH="$(GOARCH)" go build -o vault/plugins/hedera-vault-plugin ./main.go

start:
	vault server -dev -dev-root-token-id=root -dev-plugin-dir=./vault/plugins

enable:
	vault secrets enable -path=hedera hedera-vault-plugin

clean:
	rm -f ./vault/plugins/hedera-vault-plugin

test:
	vault write hedera/keys algo="ED25519" curve="" id="1"
	vault read hedera/keys id="1"
	vault list hedera/keys
	vault delete hedera/keys id="1"
	vault write hedera/keys/import id="2" algo="ED25519" curve="" privateKey="302e020100300506032b65700422042091132178e72057a1d7528025956fe39b0b847f200ab59b2fdd367017f3087137"
	vault read hedera/keys id="2"
	vault write hedera/keys/2/sign message="123"
	vault write hedera/accounts/import id="1" keyId="2" accountId="0.0.2"
	vault read hedera/accounts id="1"
	vault write hedera/accounts id="1" newId="2"
	vault write hedera/token \
		operatorId="1" \
		adminId="1" \
		treasuryId="1" \
		supplyId="1" \
		kycId="1" \
		freezeId="1" \
		pauseId="1" \
		wipeId="1" \
		name="test" \
		symbol="tst" \
		decimal="3" \
		initSupply=100000 \
		maxSupply=10000000 
	vault write hedera/tokens/pause tokenId="0.0.1168" pauseId="1" operatorId="1"
	vault write hedera/tokens/unpause tokenId="0.0.1168" pauseId="1" operatorId="1"
	vault write hedera/tokens/mint  tokenId="0.0.1168" supplyId="1" amount="10" operatorId="1"
	vault write hedera/tokens/burn  tokenId="0.0.1168" supplyId="1" amount="10" operatorId="1"
	vault write hedera/tokens/associate  tokenId="0.0.1168" userId="2" operatorId="1"
	vault write hedera/tokens/wipe  tokenId="0.0.1168" userId="2" amount="10" wipeId="1" operatorId="1"
	vault write hedera/tokens/grant_kyc  tokenId="0.0.1168" userId="2" kycId="1" operatorId="1"
	vault write hedera/tokens/revoke_kyc  tokenId="0.0.1168" userId="2" kycId="1" operatorId="1"
	vault write hedera/tokens/freeze  tokenId="0.0.1168" userId="2" kycId="1" operatorId="1"
	vault write hedera/tokens/unfreeze  tokenId="0.0.1168" userId="2" kycId="1" operatorId="1"
	vault write hedera/tokens/dissociate  tokenId="0.0.1168" userId="2" operatorId="1"
	vault write hedera/tokens/delete tokenId="0.0.1168" adminId=1 operatorId=1

fmt:
	go fmt $$(go list ./...)

.PHONY: build clean fmt start enable
