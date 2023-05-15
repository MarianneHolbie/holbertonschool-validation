#!/bin/bash

# download/install minimal version to use template ananke
sudo curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
sudo apt install ./hugo.deb
sudo mv ./hugo /usr/bin/
#remove file after installation
sudo rm hugo.deb

# download/install linter
sudo go install github.com/golangci/golangci-lint/cmd/golangci-lint@
sudo npm install -g markdownlint-cli
sudo npm install -g markdown-link-check

