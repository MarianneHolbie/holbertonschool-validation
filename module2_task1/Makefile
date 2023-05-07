

help: ## Show command help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\##//'

# > (redirige STDOUT) /dev/null(chemin de redirection) 2>&1 (2: STDERR redirection vers là ou est &1: STDOUT)
# build execute only if make lint success : &&
build: ## Compile source code of the application to a binary name awesome.api
	@go get github.com/gorilla/mux
	@make lint >/dev/null 2>&1 && @go build

run: ## Run the application in background and write logs in awesome-api.log
	@./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application
	@kill "$(shell pgrep awesome-api)"

clean: ## Stop the application and delete awesome-api and awesome-api.log
	@make stop
	@rm -rf awesome-api

test: ## Test to ensure expected beahaviour
	@curl http://localhost:9999
	@curl http://localhost:9999/health

lint: ## Test lint on the files
	@golangci-lint run