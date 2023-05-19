# parent image
FROM golang:1.15.8-alpine

# Env
ENV PATH="/usr/local/go/bin:${PATH}"

# ash = default for alpine
SHELL ["/bin/sh", "-o", "pipefail", "-c"]

# To install any applications and packages required for your container.
RUN apk update && \
rm -rf /var/lib/apt/lists/*

ENV NODE_VERSION=16.13.0
# hadolint ignore=DL3018
RUN apk add --no-cache curl zip bash

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install hadolint
RUN curl -sSfL -o /usr/local/bin/hadolint https://github.com/hadolint/hadolint/releases/download/v2.6.0/hadolint-Linux-x86_64 \
    && chmod +x /usr/local/bin/hadolint

# Install hugo
# hadolint ignore=DL3018
RUN curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb && \
apt-get -y install --no-install-recommends ./hugo.deb

# container-structure-test
# hadolint ignore=DL3018
RUN curl -sSfL -o /usr/local/bin/container-structure-test https://storage.googleapis.com/container-structure-test/latest/container-structure-test-linux-amd64 \
    && chmod +x /usr/local/bin/container-structure-test

# hadolint ignore=DL3059
# Cleanup
RUN apk clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
