
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate typescript \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-path=../test/typescript/$(EXAMPLE)/cli_gen.ts \
	cd ../../v2/ && go run ./internal/gen-test typescript ../test/typescript/$(EXAMPLE) "./cli_gen.js" < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make generate EXAMPLE=demo
	make generate EXAMPLE=features
	docker compose run test npm install
	docker compose run test npx tsc
	docker compose up --build --exit-code-from=test test
