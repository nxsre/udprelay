GO = go

VERSION = $(shell git describe --dirty --broken --match 'v*')

LDFLAGS = '$(if $(VERSION),-X main.Version=$(VERSION))'

# list of binary names for every os/arch combination returned from `go tool dist list`
PLATFORMS := $(foreach plat,$(shell go tool dist list),udprelay-$(subst /,-,$(plat))$(if $(findstring windows,$(plat)),.exe))

platform = $(subst -, ,$(patsubst udprelay-%,%,$(basename $@)))
os = $(word 1, $(platform))
arch = $(word 2, $(platform))

udprelay: *.go
	$(GO) build \
		-ldflags $(LDFLAGS) \
		-o '$@' \
		.

.PHONY: all-platforms
all-platforms: $(PLATFORMS)

$(PLATFORMS): *.go
	GOOS=$(os) GOARCH=$(arch) $(GO) build \
		-ldflags $(LDFLAGS) \
		-o '$@' \
		.

.PHONY: list-binary-targets
list-binary-targets:
	@echo $(PLATFORMS)

%: %.scd
	scdoc < $< > $@

.PHONY: docs
docs: udprelay.1 udprelay.7

.PHONY: deps
deps:
	$(GO) get -v .

.PHONY: clean
clean:
	rm -f udprelay.1 udprelay.7 udprelay $(PLATFORMS)
