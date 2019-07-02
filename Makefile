.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: install
install: ## Install all dependencies
	go mod tidy

.PHONY: serve
serve: ## Run the server
	go run main.go

.PHONY: test
test: ## Run the tests
	go test ./...

