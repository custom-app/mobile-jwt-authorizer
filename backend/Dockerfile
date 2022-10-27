FROM golang:1.18 AS build-deps

ENV GO111MODULE=on

WORKDIR /usr/src/backend
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o launcher main/launcher.go
RUN ls -hl

FROM ubuntu:18.04

RUN apt-get -y update && apt-get -y install ca-certificates

WORKDIR /usr/src/app

COPY --from=build-deps /usr/src/backend/launcher main

ENTRYPOINT ["./main", "-db", "host=lev-postgres port=5432 user=penachett password=samplesapp908365 dbname=jwt_sample sslmode=disable"]