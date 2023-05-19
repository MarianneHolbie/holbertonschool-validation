# parent image
FROM golang:1.15.8-alpine

# Env
ENV PATH="/usr/local/go/bin:${PATH}"

SHELL ["/bin/ash", "-o", "pipefail", "-c"]

# To install any applications and packages required for your container.
RUN apk update && \
rm -rf /var/lib/apt/lists/*

ENV NODE_VERSION=16.13.0
# hadolint ignore=DL3018
RUN apk add --no-cache curl zip &&\
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash && \
. ~/.bashrc
ENV NODE_VERSION=18.16.0
RUN nvm install ${NODE_VERSION} && \
nvm use v${NODE_VERSION} && \
nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN node --version && \
npm --version && \
curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
# hadolint ignore=DL3018
RUN apk add --no-cache ./hugo.deb
# hadolint ignore=DL3018
RUN apk update && \
apk add --no-cache docker && \
RUN npm install -g markdownlint-cli@*

# hadolint ignore=DL3059
# Cleanup
RUN apk clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

CMD ["/bin/sh"]