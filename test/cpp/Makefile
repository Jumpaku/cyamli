
.PHONY: generate
generate: ## make generate EXAMPLE=features
	cd ../../v2/ && go run ./cmd/cyamli generate cpp \
		-schema-path=examples/$(EXAMPLE)/cli.cyamli.yaml \
		-out-source-path=../test/cpp/$(EXAMPLE)/cli.gen.cpp \
		-out-header-path=../test/cpp/$(EXAMPLE)/cli.gen.h \
		-include-header=cli.gen.h \
		-namespace=$(EXAMPLE)
	cd ../../v2/ && go run ./internal/gen-test cpp ../test/cpp/$(EXAMPLE) $(EXAMPLE) < examples/$(EXAMPLE)/cli.cyamli.yaml

.PHONY: build
build: ## make build EXAMPLE=features
	docker compose run test g++ --std=c++11 -o ./$(EXAMPLE)/test-main.bin ./$(EXAMPLE)/cli.gen.cpp ./$(EXAMPLE)/cli_test.gen.cpp

.PHONY: test
test:
	make generate EXAMPLE=cyamli
	make build EXAMPLE=cyamli
	docker compose run --build test ./cyamli/test-main.bin

	make generate EXAMPLE=demo
	make build EXAMPLE=demo
	docker compose run --build test ./demo/test-main.bin

	make generate EXAMPLE=features
	make build EXAMPLE=features
	docker compose run --build test ./features/test-main.bin
