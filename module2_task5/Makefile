.POSIX:
DESTDIR=dist
HUGO_VERSION=0.84.0


# Below are PHONY targets

help: ## Show command help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\##//'

build: ## execute Go-Hugo and Golang builds
	@make go-build
	@make hugo-build

hugo-build: ## check version of Hugo and install if not good and build website

	curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
	apt install ./hugo.deb
	rm hugo.deb
	@make check && hugo -c $(DESTDIR)

# > (redirige STDOUT) /dev/null(chemin de redirection) 2>&1 (2: STDERR redirection vers là ou est &1: STDOUT)
# build execute only if make lint success : &&
go-build: ## Compile source code of the application to a binary name awesome.api
	@go mod init github.com/MarianneHolbie/awesome.api
	@go get github.com/gorilla/mux
	@make lint >/dev/null 2>&1 && go build || exit -1

clean: ## execute hugo-clean and go-clean
	@make hugo-clean
	@make go-clean

hugo-clean: ## remove all file -r (directories and content) -f force ignore nonexistent files

	@rm -rf dist/*

go-clean: ## Stop the application and delete awesome-api and awesome.log
	@make stop && rm -rf awesome-api awesome-api.log coverage-units.out coverage-integrations.out go*

post: ## create new post with POST_NAME ans POST_TITLE given

	@hugo new posts/$(POST_NAME).md 
	@sed -i 's|title:.*|title: "$(POST_TITLE)"|' content/posts/$(POST_NAME).md

check: ## check/analyse lint of the POSTS Markdown source file
	@markdownlint ./content/posts/*.md
	@markdown-link-check ./content/posts/*.md

validate: ## check Holberton's W3C Validator (non-blocking quality indicator)
	@-./w3c_validator.py ./dist/index.html

run: ## Run the application in background and write logs in awesome-api.log
	@./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application
	@kill "$(shell pgrep awesome-api)" 2>/dev/null &


test: ## Test to ensure expected beahaviour
	@make unit-tests 
	@make integration-tests

lint: ## Test lint on the files
	@golangci-lint run

unit-tests: ## Run Golang uni tests
	@go test -v -short -coverprofile=coverage-units.out

integration-tests: ## Run Golang integration tests
	@go test -v -coverprofile=coverage-integrations.out