BINARY        = terraform-provider-asgardeo
VERSION      ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
GOFMT_FILES  ?= $$(find . -name '*.go' | grep -v vendor)
OS_ARCH      := $(shell go env GOOS)_$(shell go env GOARCH)

# ── Paths ──────────────────────────────────────────────────────────────────────
TF_PLUGIN_DIR := ~/.terraform.d/plugins/registry.terraform.io/asgardeo/asgardeo/$(VERSION)/$(OS_ARCH)

default: build

# ── Build ─────────────────────────────────────────────────────────────────────
.PHONY: build
build:
	go build -o $(BINARY) .

.PHONY: install
install: build
	@echo "Installing provider to $(TF_PLUGIN_DIR)"
	@mkdir -p $(TF_PLUGIN_DIR)
	@mv $(BINARY) $(TF_PLUGIN_DIR)/$(BINARY)

# ── Testing ───────────────────────────────────────────────────────────────────
.PHONY: test
test:
	go test ./... -v -count=1 -timeout 120s

.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v -count=1 -timeout 600s -run "TestAcc"

# ── Quality ───────────────────────────────────────────────────────────────────
.PHONY: fmt
fmt:
	gofmt -w $(GOFMT_FILES)

.PHONY: fmtcheck
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: vet
vet:
	go vet ./...

# ── Documentation ─────────────────────────────────────────────────────────────
.PHONY: docs
docs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name asgardeo

.PHONY: docs-lint
docs-lint:
	@misspell -error docs/

# ── Tools ─────────────────────────────────────────────────────────────────────
.PHONY: tools
tools:
	go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# ── Tidying ───────────────────────────────────────────────────────────────────
.PHONY: tidy
tidy:
	go mod tidy

# ── Misc ──────────────────────────────────────────────────────────────────────
.PHONY: todo
todo:
	@grep -rn "TODO\|FIXME\|HACK" --include="*.go" .

.PHONY: clean
clean:
	rm -f $(BINARY)

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build      - Compile the provider binary"
	@echo "  install    - Build and install to ~/.terraform.d/plugins"
	@echo "  test       - Run unit tests"
	@echo "  testacc    - Run acceptance tests (requires ASGARDEO_* env vars)"
	@echo "  fmt        - Format Go source files"
	@echo "  lint       - Run golangci-lint"
	@echo "  vet        - Run go vet"
	@echo "  docs       - Regenerate provider documentation"
	@echo "  tools      - Install development tooling"
	@echo "  tidy       - Run go mod tidy"
	@echo "  clean      - Remove built binary"
