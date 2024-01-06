FROM golang:1.21.4-alpine3.18

RUN mkdir /app

WORKDIR /app

COPY . .

RUN apk add --update --no-cache bash

RUN go mod tidy
