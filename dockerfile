FROM golang:1.17.3-alpine3.13
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go build -o app
EXPOSE  32001
ENTRYPOINT  ["./app"]