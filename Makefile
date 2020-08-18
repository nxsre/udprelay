GO = go

VERSION = $(shell git describe --dirty --broken --match 'v*')

udprelay: *.go
	$(GO) build \
		-ldflags '-X main.Version=$(VERSION)' \
		-o '$@' \
		.

.PHONY: deps
deps:
	$(GO) get -v .
