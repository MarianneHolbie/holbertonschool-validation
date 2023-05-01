#!/bin/bash
# use docker image to dispose of Ubunto 18.04
docker run --rm --tty --interactive --volume=$(pwd):/app --workdir=/app ubuntu:18.04 /bin/bash
# same installation of GoHugo and Make
apt-get update && apt-get install -y hugo make
# test command
make build
# print the string to standard output
echo $?