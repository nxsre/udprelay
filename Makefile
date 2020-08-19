GO = go

VERSION := $(shell git describe --dirty --broken --match 'v*' 2>/dev/null)

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

%.txt: %
	groff -Tutf8 -man < $< | col -bx > $@

%: %.scd
ifdef VERSION
	sed '1 s/v\([0-9]\+\.\)\{2\}[0-9]\+/$(VERSION)/' $< | scdoc > $@
else
	scdoc < $< > $@
endif

.PHONY: docs
docs: udprelay.1 udprelay.7

.PHONY: docs-text
docs-text: udprelay.1.txt udprelay.7.txt

.PHONY: deps
deps:
	@echo 'no dependencies :)'

.PHONY: clean
clean:
	rm -f udprelay.1 udprelay.1.txt udprelay.7 udprelay.7.txt udprelay $(PLATFORMS)
