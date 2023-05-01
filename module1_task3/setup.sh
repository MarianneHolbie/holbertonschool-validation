#!/bin/bash

# same installation of GoHugo and Make
apt-get update && apt-get install -y hugo make
# test command
make build
# print the string to standard output
echo $?