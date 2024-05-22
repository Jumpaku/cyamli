.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: check
check: ## Checks version, runs tests.
	grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' < version.txt
	$(eval VERSION := $(shell head -n 1 version.txt))
	grep -E '^version: $(VERSION)$$' < cmd/cyamli/cli.yaml
	go test ./...

.PHONY: gen-cli
gen-cli: ## Generates Go CLI for cyamli command.
	go run ./internal/cmd/gen-golang < cmd/cyamli/cli.yaml > cmd/cyamli/cli.gen.go

.PHONY: version-apply
version-apply: ## Generates Go CLI for cyamli command.
	grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' < version.txt
	grep -E '^version: v[0-9]+\.[0-9]+\.[0-9]+$$' < cmd/cyamli/cli.yaml
	$(eval VERSION := $(shell head -n 1 version.txt))
	sed -E -i.backup "s/^version: v[0-9]+\.[0-9]+\.[0-9]+$$/version: $(VERSION)/g" cmd/cyamli/cli.yaml
	rm cmd/cyamli/cli.yaml.backup
	make gen-cli
	make examples

.PHONY: examples
examples: ## Generates Go CLI for cyamli command.
	go run ./internal/cmd/gen-golang/main.go < examples/cmd/example/cli.yaml > examples/cmd/example/cli.gen.go
	go run ./internal/cmd/gen-golang/main.go < examples/cmd/demo-app/cli.yaml > examples/cmd/demo-app/cli.gen.go

	go run ./internal/cmd/gen-python3/main.go < examples/cmd/example/cli.yaml > examples/cmd/example/cli_gen.py
	go run ./internal/cmd/gen-python3/main.go < examples/cmd/demo-app/cli.yaml > examples/cmd/demo-app/cli_gen.py
