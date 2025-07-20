
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate dart3 \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-path=../test/dart3/lib/$(EXAMPLE)/cli.g.dart
	cd ../../v2/ && go run ./internal/gen-test dart3 ../test/dart3/test/$(EXAMPLE) "package:cyamli_dart/$(EXAMPLE)/cli.g.dart" < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make generate EXAMPLE=demo
	make generate EXAMPLE=features
	docker compose up --build --exit-code-from=test test
