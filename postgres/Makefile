.DEFAULT_GOAL:=help

# make -f postgres/Makefile test-postgres

.PHONY: help
help: ## Show this help.
	@grep -E '^[0-9a-zA-Z_%-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%s\033[0m\n\t%s\n", $$1, $$2}'


.PHONY: test-postgres
test-postgres: ## run in container work
	docker compose -f postgres/docker-compose.yaml up -d --build
	docker compose -f postgres/docker-compose.yaml exec work \
		go test ./postgres/... -data-source="postgres://test:password@postgres:5432/test?sslmode=disable"
