#dependency image
FROM golang as dep

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY /util /app/util

RUN go mod download