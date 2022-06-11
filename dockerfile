FROM golang:1.17.3-alpine3.13
WORKDIR /app
ADD . /app
RUN cd /app && go build
EXPOSE 8080
ENTRYPOINT ./app