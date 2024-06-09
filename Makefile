.PHONY: help wire run test migrate

help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

wire: ## Generate wire
	@go generate ./internal/provider/.

run: ## Run the app
	@go run ./cmd/.

test: ## Run unit test with coverage info
	@go test ./... -v --cover

GOOSE_DBSTRING="host=127.0.0.1 port=5432 user=dealls password=password dbname=dealls sslmode=disable"

migrate: ## Run db migration up
	@GOOSE_DRIVER=postgres \
 	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
 	goose -dir migrations up

migrate-down: ## Run db migration down
	@GOOSE_DRIVER=postgres \
 	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
 	goose -dir migrations down

gen-mock: ## Generate mocks
	@go generate ./mocks/.