APP ?= gq
IMAGE_TAG ?= "quay.io/otaviof/$(APP):latest"

GO_FLAGS ?= -v -mod=vendor
GO_TEST_FLAGS ?= -failfast

RUN_ARGS ?=

default: build

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: clean
clean:
	@rm -rf $(APP)

.PHONY: $(APP)
$(APP):
	go build $(GO_FLAGS) .

build: vendor $(APP)

image:
	docker build --tag="$(IMAGE_TAG)" .

run:
	go run $(GO_FLAGS) . $(RUN_ARGS)

test: test-unit test-integration

.PHONY: test-unit
test-unit:
	go test $(GO_FLAGS) $(GO_TEST_FLAGS) .

.PHONY: test-integration
test-integration: build
	bats test/bats/integration.bats

install:
	go install $(GO_FLAGS) .