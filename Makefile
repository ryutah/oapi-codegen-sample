.PHONY: help
help: ## Prints help for targets with comments
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## initialize for project
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: generate/openapi
generate/openapi: ## Generate code from openapi.yaml
	oapi-codegen --config ./configs/oapi_codegen.yaml ./docs/api/openapi.yaml
	go mod tidy

.PHONY: serve
serve: ## start local server
	go run .
