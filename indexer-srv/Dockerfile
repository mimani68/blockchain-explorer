ARG GOLANG_VERSION=1.18
ARG ALPINE_TAG=3.15

#
# Build server
#
FROM golang:$GOLANG_VERSION AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct

WORKDIR /app

RUN apt update

COPY go.sum go.mod /app/

RUN go get ./...

COPY . /app

RUN go build -o ./build/indexer_restful main.go

#
# Serve binary file
#
FROM alpine:$ALPINE_TAG

LABEL maintainer="Mahdi Imani <imani.mahdi@gmail.com>"

ENV TZ=Asia/Istanbul

RUN apk update && \
    apk add --update --no-cache tzdata openssl && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

COPY --from=build /app/build/app /app/

EXPOSE 3000

CMD ["./indexer_restful"]