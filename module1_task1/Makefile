.POSIX:
DESTDIR=dist
HUGO_VERSION=0.84.0
HUGO_EXTENDED=1


# Below are PHONY targets

help:
# help command
	@echo "Usage: make <command>"
	@echo "	make build		Builds the website with Hugo"
	@echo " make clean		Clean all build files"
	@echo " make POST_NAME=    POST_TITLE=     post		Create a new blog post"



build:
# check version of Hugo and install if not good
# build website in the DESTDIR directory
	@echo "Check Hugo version"
	curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
	apt install ./hugo.deb
	rm hugo.deb
	@echo "Generating site"
	hugo -c $(DESTDIR)


clean:
# move to the directory
# remove all file -r (directories and content) -f force ignore nonexistent files
	@echo "Cleaning old build"
	cd $(DESTDIR) && rm -rf *


post:
# create new post with POST_NAME as name of file
# replace title in the file create just before with sed
# sed 's|<search regex>|<replacement>|' <input file>

	@echo "Create new blog post"
	hugo new ./content/posts/$(POST_NAME).md 
	sed 's|title:*|title: "$(POST_TITLE)"|' ./content/posts/$(POST_NAME).md