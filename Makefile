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
	vault write hedera/accounts id="1" nextId="2"


fmt:
	go fmt $$(go list ./...)

.PHONY: build clean fmt start enable
