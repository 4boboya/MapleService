FROM golang:1.17.3-alpine3.13
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go get github.com/gorilla/mux
RUN         go build -o app
ENTRYPOINT  ["./app"]