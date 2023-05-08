#!/bin/bash

# same installation of GoHugo and Make
apt-get update && apt-get install -y hugo make
# download/install minimal version to use template ananke
curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
apt install ./hugo.deb
#remove file after installation
rm hugo.deb
# test command
make build
echo $?