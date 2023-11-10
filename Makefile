.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


.PHONY: gen-cli
gen-cli: ## Generates Go CLI for cyamli command.
	go run ./internal/tools/build/main.go < cmd/cyamli/cli.yaml > cmd/cyamli/cli.gen.go

.PHONY: examples
examples: ## Generates Go CLI for cyamli command.
	go run ./internal/tools/build/main.go < examples/cmd/example/cli.yaml > examples/cmd/example/cli.gen.go
	go run ./internal/tools/build/main.go < examples/cmd/greet/cli.yaml > examples/cmd/greet/cli.gen.go
