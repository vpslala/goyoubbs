ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go generate
RUN go mod tidy
RUN go build -o ./app ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app .

EXPOSE 8082

ENTRYPOINT ["./app", "-addr=:8082", "-sdbDir=/data/goyoubbs/onlinedb"]
