include helpers.mk

PROJECT := errorBudgetBurn-investigator
GIT_HASH := $(shell git rev-parse --short=7 HEAD)

unexport GOFLAGS 

GOOS?=linux
GOARCH?=amd64
GOENV=GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 GOFLAGS= 
GOBUILDFLAGS=-gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}" 

GORELEASER_SINGLE_TARGET ?= true

default: all

.PHONY: license vet mod fmt lint test build
all: license vet mod fmt lint test build

vet:
	go vet ${BUILDFLAGS} ./... 

mod: 
	go mod tidy 
	@git diff --exit-code -- go.mod

fmt:
	gofmt -w -s . 
	@git diff --exit-code . 

test:
	go test ${BUILDFLAGS} ./... -covermode=atomic -coverpkg=./...

#Requires golangci-lint
# `make install-golangci-lint` or `make install-tools` to install
lint:
	golangci-lint run

# Requires goreleaser
# `make install-goreleaser` or `make install-tools` to install
build:
	goreleaser build --clean --snapshot --single-target=${GORELEASER_SINGLE_TARGET}

# Ensure license text
# Requires https://github.com/google/addlicense 
# `make install-license` or `make install-tools` to install
.PHONY: license 
license: 
	@addlicense -c "Red Hat, Inc."  -l apache -v -y 2021-2024 *.go **/*.go **/**/*.go
