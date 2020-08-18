GO = go

VERSION = $(shell git describe --dirty --broken --match 'v*')

udprelay: *.go
	$(GO) build \
		-ldflags '-X main.Version=$(VERSION)' \
		-o '$@' \
		.

%: %.scd
	scdoc < $< > $@

.PHONY: docs
docs: udprelay.1 udprelay.7

.PHONY: deps
deps:
	$(GO) get -v .
