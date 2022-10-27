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

fmt:
	go fmt $$(go list ./...)

.PHONY: build clean fmt start enable
