.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: cli
cli: ## Generates CLIs to be tested
	for cli_yaml in cli/*.yaml; do \
  		group=$$(basename "$${cli_yaml}" .yaml); \
		cyamli generate dart < "cli/$${group}.yaml" > "dart/$${group}/cli.g.dart"; \
		cyamli generate golang -package=main < "cli/$${group}.yaml" > "golang/$${group}/cli.gen.go"; \
		cyamli generate python3 < "cli/$${group}.yaml" > "python3/$${group}/cli_gen.py"; \
	done

.PHONY: testcases
testcases: ## Generates testcases
	go run ./internal/gen-testcases .

.PHONY: run_test
run_test: ## Run tests
	docker compose up test-dart
	docker compose up test-golang
	docker compose up test-python3
