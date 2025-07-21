
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate python3 \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-path=../test/python3/$(EXAMPLE)/cli_gen.py \
	cd ../../v2/ && go run ./internal/gen-test python3 ../test/python3/$(EXAMPLE) $(EXAMPLE).cli_gen < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make generate EXAMPLE=demo
	make generate EXAMPLE=features
	docker compose up --build --exit-code-from=test test
