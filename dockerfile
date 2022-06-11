# FROM        golang:1.17.3-alpine3.13 AS stage1
# ENV         RUN_PATH=/app PROJ_PATH=/build
# RUN         mkdir -p $RUN_PATH
# WORKDIR     $RUN_PATH
# ENV         GO111MODULE=on
# COPY        go.mod .
# COPY        go.sum .
# RUN         go mod download

# FROM        stage1 AS stage2
# RUN         apk update && apk add make
# USER        root
# ADD         . $PROJ_PATH
# WORKDIR     $PROJ_PATH
# RUN         make build pack unpack path=/app

# FROM        alpine
# USER        root
# ENV         RUN_PATH=/app
# RUN         mkdir -p $RUN_PATH
# RUN         apk add tzdata && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime && echo "Asia/Taipei" >  /etc/timezone
# RUN         date
# RUN         apk del tzdata
# COPY        --from=stage2 ${RUN_PATH} ${RUN_PATH}
# WORKDIR     ${RUN_PATH}
# EXPOSE      80
# ENTRYPOINT  ["./app"]
FROM    golang:1.17.3-alpine3.13 AS stage1
WORKDIR /app
ENV     GO111MODULE=on
COPY    . .
RUN     go mod download
RUN     CGO_ENABLED=0 GOOS=linux go build -a -o main

FROM    alpine:latest
RUN     apk --no-cache add ca-certificates
WORKDIR /app
COPY    --from=builder /app/main .
ENTRYPOINT ["./main"]