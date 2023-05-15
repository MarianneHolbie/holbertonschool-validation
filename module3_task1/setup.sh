#!/bin/bash

# download/install minimal version to use template ananke
curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
apt install ./hugo.deb
#remove file after installation
rm hugo.deb

# download/install linter
go install github.com/golangci/golangci-lint/cmd/golangci-lint@
npm install -g markdownlint-cli
npm install -g markdown-link-check