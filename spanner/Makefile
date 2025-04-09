.DEFAULT_GOAL:=help

# make -f spanner/Makefile test-spanner

.PHONY: help
help: ## Show this help.
	@grep -E '^[0-9a-zA-Z_%-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%s\033[0m\n\t%s\n", $$1, $$2}'


.PHONY: test-spanner
test-spanner: ## initialize Spanner emulator database for develop. run in service work
	docker compose -f spanner/docker-compose.yaml up -d --build
	docker compose -f spanner/docker-compose.yaml exec work \
		make -f spanner/Makefile test-spanner-docker


.PHONY: test-spanner-docker
test-spanner-docker: ## initialize Spanner emulator database for develop. run in service work
	gcloud config set project schenerate
	gcloud config set auth/disable_credentials true
	yes | gcloud config set api_endpoint_overrides/spanner http://spanner:9020/
	SPANNER_EMULATOR_HOST=spanner:9010 \
	yes | gcloud spanner instances delete test || true
	SPANNER_EMULATOR_HOST=spanner:9010 \
	gcloud spanner instances create test --config=emulator-config --description="Instance for integration test"

	SPANNER_EMULATOR_HOST=spanner:9010 \
	go test ./spanner/... -project=schenerate -instance=test
