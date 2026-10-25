# Winthor data platform — common commands.
# Recipes use '>' as prefix (no tabs) and run under bash (Git for Windows on the dev server).
.RECIPEPREFIX = >
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

ifeq ($(OS),Windows_NT)
  export PATH := C:\Program Files\Git\bin;C:\Program Files\Git\usr\bin;$(PATH)
  # Keep Go caches off the small C: drive when D: exists.
  ifneq ($(wildcard D:/),)
    export GOMODCACHE := D:/mint/go/mod
    export GOCACHE := D:/mint/go/cache
  endif
  POWERSHELL := powershell -NoProfile -ExecutionPolicy Bypass
  EXE := .exe
else
  POWERSHELL := pwsh -NoProfile
  EXE :=
endif

# Version is computed inside the recipe (bash), not with $(shell), because
# GNU make on Windows may hand $(shell) to cmd.exe and choke on /dev/null.
VERSION_CMD := git describe --tags --always --dirty 2>/dev/null || echo dev
LDFLAGS = -s -w -X main.version=$${VERSION:-$$($(VERSION_CMD))}
DIST := dist
E ?=

.PHONY: help build build-windows test test-integration lint fmt tidy doctor sqlcheck verify clean

help: ## List targets
> @grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-18s %s\n", $$1, $$2}'

build: ## Build all binaries for the current platform into dist/
> mkdir -p $(DIST)
> go build -trimpath -ldflags "$(LDFLAGS)" -o $(DIST)/gateway$(EXE) ./gateway/cmd/gateway

build-windows: ## Build gateway.exe for Windows amd64 (Windows Server 2016+)
> mkdir -p $(DIST)
> GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o $(DIST)/gateway.exe ./gateway/cmd/gateway
> @ls -l $(DIST)/gateway.exe

test: ## Unit tests (mocks only; no Oracle)
> go vet ./...
> go test -count=1 ./...

test-integration: ## Integration tests against the test Oracle (requires ORACLE_DSN, ORACLE_USER, ORACLE_PASSWORD)
> @test -n "$${ORACLE_DSN:-}" || { echo "ORACLE_DSN is not set"; exit 2; }
> go test -count=1 -tags integration -run Integration ./...

lint: ## gofmt and go vet
> @out=$$(gofmt -l . 2>/dev/null); if [ -n "$$out" ]; then echo "gofmt needed:"; echo "$$out"; exit 1; fi
> go vet ./...

fmt: ## Format Go code
> gofmt -w .

tidy: ## go mod tidy
> go mod tidy

doctor: build ## Run gateway doctor with config.yaml + config.local.yaml
> $(DIST)/gateway$(EXE) doctor

sqlcheck: build ## Run the read-only guard over scripts/verify/testdata/forbidden.sql (expected to be refused)
> ! $(DIST)/gateway$(EXE) sqlcheck scripts/verify/testdata/forbidden.sql

verify: ## Run the verification script of a deliverable: make verify E=1.1
> @test -n "$(E)" || { echo "usage: make verify E=1.1"; exit 2; }
> $(POWERSHELL) -File scripts/verify/E$(E).ps1

clean: ## Remove build output
> rm -rf $(DIST)
