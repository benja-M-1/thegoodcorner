DIST_ENV_FILE=".env.dist"
ENV_FILE=".env"

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: check-dependencies ## Install all dependencies
	go mod tidy
	cp -n ${DIST_ENV_FILE} ${ENV_FILE}

.PHONY: check-dependencies
check-dependencies: ## Ensure dependencies are isntalled
	@docker --version > /dev/null 2>&1 || (echo -e 'docker is not installed\n' && exit 42)
	@docker ps > /dev/null 2>&1 || (echo -e 'docker is not installed properly, consult the docker installation process\n' && exit 42)

	@docker-compose --version > /dev/null 2>&1 || (echo -e 'docker-compose is not installed\n' && exit 42)

.PHONY: serve
serve: start ## Run the server
	go run main.go

.PHONY: test
test: ## Run the tests with benchmark
	go test ./... -bench=.

.PHONY: start
start: ## Start Docker container
	docker-compose up -d

.PHONY: clean
clean: ## Stop Docker containers and remove Docker volumes
	docker-compose stop
	docker-compose rm -v
