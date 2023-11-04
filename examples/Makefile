.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


.PHONY: gen
gen: ## Generates Go CLI for cliautor command.
	mkdir -p "cmd/example" && go run "github.com/Jumpaku/cliautor/cmd/cliautor" < cli.yaml > cmd/cliautor/cli.gen.go

.PHONY: run
run: gen ## Run example app with ARGS.
	go run cmd/exammple $(ARGS)
