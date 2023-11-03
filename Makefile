.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


.PHONY: gen-cli
gen-cli: ## Generates Go CLI for cliautor command.
	mkdir -p "cmd/cliautor" && go run playground/main.go < playground/cli.yaml > cmd/cliautor/cli.gen.go
