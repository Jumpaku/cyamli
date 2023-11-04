.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


.PHONY: gen-cli
gen-cli: ## Generates Go CLI for cliautor command.
	rm -rf cmd/tmp
	mkdir -p cmd/tmp && cp -r cmd/cliautor/* cmd/tmp/
	go run ./cmd/tmp golang < cmd/tmp/cli.yaml > cmd/cliautor/cli.gen.go
	cp cmd/tmp/main.go cmd/cliautor/main.go 
	cp cmd/tmp/cli.yaml cmd/cliautor/cli.yaml
