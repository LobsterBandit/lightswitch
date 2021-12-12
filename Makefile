VERSION=$(shell git describe --tags --long)
GO_BUILDFLAGS=-v -o lightswitch
GO_LDFLAGS="-X main.Version=$(VERSION)"

.PHONY: build
build:
	go build $(GO_BUILDFLAGS) -ldflags $(GO_LDFLAGS)
