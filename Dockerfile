FROM golang:1.21.1-alpine3.18

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .

RUN go mod download