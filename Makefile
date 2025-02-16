.DEFAULT_TARGET=help

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
#
# More info on how the awk command works in this blog post:
# https://www.padok.fr/en/blog/beautiful-makefile-awk

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Validation

.PHONY: fmt
fmt: ## Format source code.
	go fmt ./...

.PHONY: vet
vet: ## Vet source code.
	go vet ./...

##@ Testing

.PHONY: test
test: fmt vet test-unit test-integration ## Run all tests.

.PHONY: test-unit
test-unit: ## Run unit tests.
	go test -cover ./... -unit

.PHONY: test-integration
test-integration: ## Run integration tests.
	go test -cover ./... -unit=false -integration

##@ Build

.PHONY: build
build: ## Build jumpstart-decklists binary.
	go build -o bin/jumpstart-decklists

##@ Usage

.PHONY: export
export: build ## Export decklists as PNG images.
	bin/jumpstart-decklists export

.PHONY: serve
serve: build ## Serve decklists as HTML.
	bin/jumpstart-decklists serve
