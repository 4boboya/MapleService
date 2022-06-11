FROM        golang:1.17.3-alpine3.13 AS stage1
ENV         RUN_PATH=/app PROJ_PATH=/build
RUN         mkdir -p $RUN_PATH
WORKDIR     $RUN_PATH
ENV         GO111MODULE=on
COPY        go.mod .
COPY        go.sum .
RUN         go mod download

FROM        stage1 AS stage2
RUN         apk update && apk add make
USER        root
ADD         . $PROJ_PATH
WORKDIR     $PROJ_PATH
RUN         make build pack unpack path=/app

FROM        alpine
USER        root
ENV         RUN_PATH=/app
RUN         mkdir -p $RUN_PATH
RUN         apk add tzdata && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime && echo "Asia/Taipei" >  /etc/timezone
RUN         date
RUN         apk del tzdata
COPY        --from=stage2 ${RUN_PATH} ${RUN_PATH}
WORKDIR     ${RUN_PATH}
ENTRYPOINT  ["./app"]

# 官方版本
# FROM golang:1.17-alpine
# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
# COPY . ./
# RUN go build -o /app
# ENTRYPOINT [ "/app" ]

# 網路
# FROM        golang
# RUN         mkdir -p /app
# WORKDIR     /app
# COPY        . .
# RUN         go mod download
# RUN         go build -o app
# ENTRYPOINT  ["./app"]