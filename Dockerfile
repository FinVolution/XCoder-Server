FROM golang:1.20-alpine  as builder
ENV GOOS=linux
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
ENV GOSUMDB=off
ENV TZ=Asia/Shanghai

RUN apk add build-base

WORKDIR /build/xcoder/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o /app/main main.go

FROM alpine

RUN apk add --no-cache --virtual .build-deps \
         build-base \
         gcc \
         tzdata \
         wget \
         git \
         && apk add --no-cache libstdc++ \
         && apk del .build-deps

WORKDIR /app

COPY --from=builder /app/main /app/main

COPY manifest/config /app/config

EXPOSE 8081

RUN echo $'#!/bin/sh\n\
exec /app/main\n\
' > entrypoint.sh

RUN chmod +x /app/entrypoint.sh
ENTRYPOINT ["/app/entrypoint.sh"]