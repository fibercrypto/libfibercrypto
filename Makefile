.DEFAULT_GOAL := help
.PHONY: test-libc test-lint build-libc check build build-skyapi test-skyapi
.PHONY: install-linters format clean-libc format-libc lint-libc docs
.PHONY: build-skyhwd test-skyhwd

COIN ?= skycoin

# Resource paths
# --- Absolute path to repository root
LIBSRC_ABS_PATH        = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
# --- Relative path to repository root
LIBSRC_REL_PATH        = .
# --- Relative path to repository root from local vendor directory
LIBSRC_VENDORREL_PATH  = ..
# --- Relative path to repository root from Skycoin source root directory
LIBSRC_SKYSRCREL_PATH  = ../../../..
# --- Relative path to libfibercrypto vendor directory
LIBVENDOR_REL_PATH     = vendor
# --- Relative path to Skycoin source code submodule
SKYSRC_REL_PATH        = $(LIBVENDOR_REL_PATH)/github.com/fibercryto/fibercryptowallet
# --- Relative path to Skycoin vendor directory
SKYVENDOR_REL_PATH     = $(SKYSRC_REL_PATH)/vendor

# Compilation output for libfibercrypto
BUILD_DIR = build
BUILDLIB_DIR = $(BUILD_DIR)/libfibercrypto
LIB_DIR = lib
BIN_DIR = bin
DOC_DIR = docs
INCLUDE_DIR = include
LIBSRC_DIR = lib/cgo
LIBFCWDOC_DIR = $(DOC_DIR)/libc

# Source files
LIB_FILES = $(shell find $(LIBSRC_DIR) -type f -name "*.go")
HEADER_FILES = $(shell find $(INCLUDE_DIR) -type f -name "*.h")

# Compilation flags for libfibercrypto
CC_VERSION = $(shell $(CC) -dumpversion)
STDC_FLAG = $(python -c "if tuple(map(int, '$(CC_VERSION)'.split('.'))) < (6,): print('-std=C99'")
LIBC_LIBS = `pkg-config --cflags --libs check`
LIBC_FLAGS = -I$(LIBSRC_DIR) -I$(INCLUDE_DIR) -I$(BUILD_DIR)/usr/include -L $(BUILDLIB_DIR) -L$(BUILD_DIR)/usr/lib
# Platform specific checks
OSNAME  = $(TRAVIS_OS_NAME)
UNAME_S = $(shell uname -s)
CGO_ENABLED=1

PKG_CLANG_FORMAT = clang-format
PKG_CLANG_LINTER = clang-tidy
PKG_LIB_TEST = check

ifeq ($(UNAME_S),Linux)
  LDLIBS=$(LIBC_LIBS) -lpthread
  LDPATH=$(shell printenv LD_LIBRARY_PATH)
  LDPATHVAR=LD_LIBRARY_PATH
  LDFLAGS=$(LIBC_FLAGS) $(STDC_FLAG)
  OSNAME ?= linux
else ifeq ($(UNAME_S),Darwin)
  OSNAME ?= osx
  LDLIBS = $(LIBC_LIBS)
  LDPATH=$(shell printenv DYLD_LIBRARY_PATH)
  LDPATHVAR=DYLD_LIBRARY_PATH
  LDFLAGS=$(LIBC_FLAGS) -framework CoreFoundation -framework Security
else
  LDLIBS = $(LIBC_LIBS)
  LDPATH=$(shell printenv LD_LIBRARY_PATH)
  LDPATHVAR=LD_LIBRARY_PATH
  LDFLAGS=$(LIBC_FLAGS)
endif

configure-build:
	mkdir -p $(BUILD_DIR)/usr/tmp $(BUILD_DIR)/usr/lib $(BUILD_DIR)/usr/include
	mkdir -p $(BUILDLIB_DIR) $(BIN_DIR) $(INCLUDE_DIR)

## Update links to dependency packages
dep:
	git submodule update --init --recursive

$(BUILDLIB_DIR)/libfibercrypto.so: $(LIB_FILES) $(SRC_FILES) $(HEADER_FILES)
	rm -Rf $(BUILDLIB_DIR)/libfibercrypto.so
	go build -buildmode=c-shared  -o $(BUILDLIB_DIR)/libfibercrypto.so $(LIB_FILES)
	mv $(BUILDLIB_DIR)/libfibercrypto.h $(INCLUDE_DIR)/

$(BUILDLIB_DIR)/libfibercrypto.a: $(LIB_FILES) $(SRC_FILES) $(HEADER_FILES)
	rm -Rf $(BUILDLIB_DIR)/libfibercrypto.a
	go build -buildmode=c-archive -o $(BUILDLIB_DIR)/libfibercrypto.a  $(LIB_FILES)
	mv $(BUILDLIB_DIR)/libfibercrypto.h $(INCLUDE_DIR)/

build-libc-static: $(BUILDLIB_DIR)/libfibercrypto.a ## Build libfibercrypto C static library

build-libc-shared: $(BUILDLIB_DIR)/libfibercrypto.so ## Build libfibercrypto C shared library

build-libc: configure-build build-libc-static build-libc-shared ## Build libfibercrypto C client libraries

build: build-libc ## Build libraries

## Build libfibercrypto C client library and executable C test suites
## with debug symbols. Use this target to debug the source code
## with the help of an IDE
build-libc-dbg: configure-build build-libc-static build-libc-shared
	$(CC) -g -o $(BIN_DIR)/test_libfibercrypto_shared $(LIB_DIR)/cgo/tests/*.c                     $(LDLIBS) $(LDFLAGS)
	$(CC) -g -o $(BIN_DIR)/test_libfibercrypto_static $(LIB_DIR)/cgo/tests/*.c $(BUILDLIB_DIR)/libfibercrypto.a $(LDLIBS) $(LDFLAGS)

test-libc: build-libc ## Run tests for libfibercrypto C client library
	echo "Compiling with $(CC) $(CC_VERSION) $(STDC_FLAG)"
	$(eval TESTS_SRC := $(shell ls $(LIB_DIR)/cgo/tests/*.c))
	$(CC) -o $(BIN_DIR)/test_libfibercrypto_shared $(TESTS_SRC) $(LIB_DIR)/cgo/tests/testutils/*.c -lfibercrypto                    $(LDLIBS) $(LDFLAGS)
	$(CC) -o $(BIN_DIR)/test_libfibercrypto_static $(TESTS_SRC) $(LIB_DIR)/cgo/tests/testutils/*.c $(BUILDLIB_DIR)/libfibercrypto.a $(LDLIBS) $(LDFLAGS)
	$(LDPATHVAR)="$(LDPATH):$(BUILD_DIR)/usr/lib:$(BUILDLIB_DIR)" $(BIN_DIR)/test_libfibercrypto_shared -lfibercrypto
	$(LDPATHVAR)="$(LDPATH):$(BUILD_DIR)/usr/lib"         $(BIN_DIR)/test_libfibercrypto_static


test: test-libc ## Run all test for libfibercrypto


docs-libc: ## Generate libfibercrypto documentation
	doxygen ./.Doxyfile
	moxygen -o $(LIBFCWDOC_DIR)/API.md $(LIBFCWDOC_DIR)/xml/

docs: docs-libc ## Generate documentation for all libraries

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	# lib/cgo needs separate linting rules
	golangci-lint run -c .golangci.yml ./lib/cgo/...
	# The govet version in golangci-lint is out of date and has spurious warnings, run it separately
	go vet -all ./...

lint-libc: format-libc
	# Linter LIBC
	clang-tidy  lib/cgo/tests/*.c -- $(LIBC_FLAGS) -Wincompatible-pointer-types


check: lint test-libc  ## Run tests and linters

install-linters-Linux: ## Install linters on GNU/Linux
	sudo apt-get update
	sudo apt-get install $(PKG_CLANG_FORMAT) -y
	sudo apt-get install $(PKG_CLANG_LINTER) -y

install-linters-Darwin: ## Install linters on Mac OSX
	# brew install $(PKG_CLANG_FORMAT)
	brew install llvm
	ln -s "/usr/local/opt/llvm/bin/clang-format" "/usr/local/bin/clang-format"
	ln -s "/usr/local/opt/llvm/bin/clang-tidy" "/usr/local/bin/clang-tidy"

install-deps-Linux: ## Install deps on GNU/Linux
	sudo apt-get install $(PKG_LIB_TEST)

install-deps-Darwin: ## Install deps on Mac OSX
	brew install $(PKG_LIB_TEST)

install-linters: install-linters-$(UNAME_S) ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	cat ./ci-scripts/install-golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.10.2

install-deps-libc: install-deps-libc-$(UNAME_S) ## Install deps for libc

install-deps-libc-Linux: configure-build check-0.12.0/src/.libs/libcheck.so ## Install locally dependencies for testing libfibercrypto

check-0.12.0/src/.libs/libcheck.so: ## Install libcheck
	wget -c https://github.com/libcheck/check/releases/download/0.12.0/check-0.12.0.tar.gz
	tar -xzf check-0.12.0.tar.gz
	cd check-0.12.0 && ./configure --prefix=/usr --disable-static && make && sudo make install

install-deps-libc-Darwin: configure-build ## Install locally dependencies for testing libfibercrypto
	brew install check

install-deps: install-deps-libc ## Install deps for libc and skyapi

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -local github.com/fibercryto/FiberCryptoWallet ./lib

clean-libc: ## Clean files generated by libc
	rm -rfv $(BUILDLIB_DIR)
	rm -rfv bin
	rm -rfv qemu_test_libfibercrypto*
	rm -rfv qemu_*
	rm -rfv include/libfibercrypto.h

clean: clean-libc ## Clean all files generated by libraries
	rm -rfv $(BUILD_DIR)
format-libc:
	$(PKG_CLANG_FORMAT) -sort-includes -i -assume-filename=.clang-format lib/cgo/tests/*.c
	$(PKG_CLANG_FORMAT) -sort-includes -i -assume-filename=.clang-format include/*.h

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
