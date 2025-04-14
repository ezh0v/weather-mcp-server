.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

##@ Tests
generate-mocks: ## generate mock files using mockgen
	mockgen -source=./internal/server/services/services.go \
		-destination=internal/server/services/mock/mock.go -package=mock

run-tests: ## run unit tests
	go test -v -race ./...