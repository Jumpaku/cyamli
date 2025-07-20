
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate csharp \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-path=../test/csharp/CyamliCsharp/$(EXAMPLE)/Cli_Gen.cs \
		-namespace=$(EXAMPLE)
	cd ../../v2/ && go run ./internal/gen-test csharp ../test/csharp/CyamliCsharp.Tests/$(EXAMPLE) $(EXAMPLE) < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make generate EXAMPLE=demo
	make generate EXAMPLE=features
	docker compose up --build --exit-code-from=test test
