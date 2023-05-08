.POSIX:
DESTDIR=dist
HUGO_VERSION=0.84.0


# Below are PHONY targets

help: ## Show command help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/\##//'

build: ## check version of Hugo and install if not good and build website

	curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
	apt install ./hugo.deb
	rm hugo.deb
	@make check && hugo -c $(DESTDIR)


clean: ## remove all file -r (directories and content) -f force ignore nonexistent files

	@cd $(DESTDIR) && rm -rf *


post: ## create new post with POST_NAME ans POST_TITLE given

	@hugo new posts/$(POST_NAME).md 
	@sed -i 's|title:.*|title: "$(POST_TITLE)"|' content/posts/$(POST_NAME).md

check: ## check/analyse lint of the POSTS Markdown source file
	@markdownlint ./content/posts/*.md
	@markdown-link-check ./content/posts/*.md

validate: ## check Holberton's W3C Validator (non-blocking quality indicator)
	@-./w3c_validator.py ./dist/index.html