FROM golang:1.17.3-alpine3.13
RUN         mkdir -p /app
WORKDIR     /app
ENV         GO111MODULE=on
COPY        . .
RUN         go mod download
RUN         go build -o app
ENTRYPOINT  ./app