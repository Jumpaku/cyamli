
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate golang \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-path=../test/golang/$(EXAMPLE)/cli.gen.go \
		-package=$(EXAMPLE)
	cd ../../v2/ && go run ./internal/gen-test golang ../test/golang/$(EXAMPLE) $(EXAMPLE) < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make generate EXAMPLE=demo
	make generate EXAMPLE=features
	docker compose up --build --exit-code-from=test test
