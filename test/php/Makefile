
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate php \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-dir=../test/php/src/$(EXAMPLE) \
		-namespace='Cyamli\'$(EXAMPLE)
	cd ../../v2/ && go run ./internal/gen-test php ../test/php/tests/$(EXAMPLE) 'Cyamli\'$(EXAMPLE) 'Tests\'$(EXAMPLE) < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make generate EXAMPLE=demo
	make generate EXAMPLE=features
	docker compose run test composer install
	docker compose run test composer dump-autoload
	docker compose up --build --exit-code-from=test test
