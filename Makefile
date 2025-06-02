.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## make version VERSION=v2.0.0
	sed -e 's|^version: .*|version: $(VERSION)|g' v2/cmd/cyamli/cli.cyamli.yaml \
		> v2/cmd/cyamli/cli.cyamli.yaml.new \
		&& mv v2/cmd/cyamli/cli.cyamli.yaml.new v2/cmd/cyamli/cli.cyamli.yaml
	sed -e 's|^$id: .*|$id: $id: https://github.com/Jumpaku/cyamli/raw/$(VERSION)/v2/docs/cyamli-cli.schema.json|g' v2/schema/cyamli-cli.schema.yaml \
		> v2/schema/cyamli-cli.schema.yaml.new \
		&& mv v2/schema/cyamli-cli.schema.yaml.new v2/schema/cyamli-cli.schema.yaml
	make install
	make docs
	cyamli version

.PHONY: install
install: ## Installs cyamli CLI tool locally.
	cd v2 && go generate ./...
	cd v2 && go install ./cmd/cyamli

.PHONY: docs
docs: install ## Generates documentation for cyamli CLI tool.
	cd v2 && go run ./internal/yaml2json/main.go < schema/cyamli-cli.schema.yaml > schema/cyamli-cli.schema.json
	cp v2/schema/cyamli-cli.schema.yaml docs/cyamli-cli.schema.yaml
	cp v2/schema/cyamli-cli.schema.json docs/cyamli-cli.schema.json
	cyamli generate docs -format=text -schema-path=v2/cmd/cyamli/cli.cyamli.yaml -out-path=docs/cyamli-docs.text
	cyamli generate docs -format=html -schema-path=v2/cmd/cyamli/cli.cyamli.yaml -out-path=docs/cyamli-docs.html
	cyamli generate docs -format=markdown -schema-path=v2/cmd/cyamli/cli.cyamli.yaml -out-path=docs/cyamli-docs.md
