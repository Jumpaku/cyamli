.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: check
check: ## Checks version, runs tests.
	grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$'  ./info/version.txt
	$(eval VERSION := $(shell head -n 1 ./info/version.txt))
	grep -E '^version: $(VERSION)$$'  ./cyamli/cli.yaml
	go test ./...

.PHONY: install
install: ## Install cyamli built in present status.
	go run ./internal/cmd/gen-golang < cyamli/cli.yaml > cyamli/cli.gen.go
	go install .
	cyamli generate golang -package=cyamli < cyamli/cli.yaml > cyamli/cli.gen.go
	go install .

.PHONY: version-apply
version-apply: ## Generates Go CLI for cyamli command.
	grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' < ./info/version.txt
	grep -E '^version: v[0-9]+\.[0-9]+\.[0-9]+$$' < ./cyamli/cli.yaml
	$(eval VERSION := $(shell head -n 1 ./info/version.txt))
	sed -E -i.backup "s/^version: v[0-9]+\.[0-9]+\.[0-9]+$$/version: $(VERSION)/g" ./cyamli/cli.yaml
	rm ./cyamli/cli.yaml.backup
	make install
	make examples

.PHONY: examples
examples: install ## Generates Go CLI for cyamli command.
	go run . generate golang < examples/cmd/example/cli.yaml > examples/cmd/example/cli.gen.go
	go run . generate golang < examples/cmd/demo-app/cli.yaml > examples/cmd/demo-app/cli.gen.go

	go run . generate python3 < examples/cmd/example/cli.yaml > examples/cmd/example/cli_gen.py
	go run . generate python3 < examples/cmd/demo-app/cli.yaml > examples/cmd/demo-app/cli_gen.py
