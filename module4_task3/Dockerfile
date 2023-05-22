# official image
FROM alpine:3.13

# define workdir
WORKDIR /app

# copy file
COPY ./dist/ ./dist/
COPY ./awesome-api

# entrypoint
ENTRYPOINT ["/app/awesome-api"]