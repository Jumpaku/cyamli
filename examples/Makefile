.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


.PHONY: gen
gen: ## Generates Go CLI for cyamli command. Example: make gen TARGET=greet
	@mkdir -p "cmd/$(TARGET)" && ./../main golang < cmd/$(TARGET)/cli.yaml > cmd/$(TARGET)/cli.gen.go

.PHONY: run
run: gen ## Run example app with ARGS. Example: make gen TARGET=greet ARGS="hello -target-name=Bob Alice"
	@go run ./cmd/$(TARGET) $(ARGS)
