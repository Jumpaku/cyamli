
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate kotlin \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-path=../test/kotlin/src/main/kotlin/$(EXAMPLE)/Cli.kt \
		-package=$(EXAMPLE)
	cd ../../v2/ && go run ./internal/gen-test kotlin ../test/kotlin/src/test/kotlin/$(EXAMPLE) $(EXAMPLE) < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make generate EXAMPLE=demo
	make generate EXAMPLE=features
	docker compose down
	docker compose up --build --exit-code-from=test test
