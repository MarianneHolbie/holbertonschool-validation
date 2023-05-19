# parent image
FROM golang:1.15.8-buster
# To set the working directory for any commands that follow in the Dockerfile.
WORKDIR /module4_task1
# To copy over files or directories from a specific location.
COPY . .
# To install any applications and packages required for your container.
RUN apt-get update
ENV NODE_VERSION=16.13.0
RUN apt install -y curl zip
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN node --version
RUN npm --version
RUN curl -L https://github.com/gohugoio/hugo/releases/download/v0.84.0/hugo_extended_0.84.0_Linux-64bit.deb -o hugo.deb
RUN apt install ./hugo.deb
RUN apt-get update -y && apt-get install -y docker.io
RUN npm install -g markdownlint-cli
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.29.0
# Cleanup
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*