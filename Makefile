UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
# This controls the default version of buf to install and use.
BUF_VERSION ?= 1.0.0-rc8
# Bufmak
CACHE_BIN := bin
# Update the $PATH so we can use buf directly
export PATH := $(PATH):$(abspath $(CACHE_BIN))
export GOBIN := $(abspath $(CACHE_BIN))

.PHONY: buf/*

buf/install:
	@rm -f $(CACHE_BIN)/buf
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_OS)-$(UNAME_ARCH)" \
		-o "$(CACHE_BIN)/buf"
	chmod +x "$(CACHE_BIN)/buf"

.PHONY: buf/init
buf/init:
	@echo $(PATH)
	@ls ./bin
	@./bin/buf config init