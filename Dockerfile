FROM golang:1.18-alpine AS builder

WORKDIR /device-service

COPY . /device-service/

ARG GOOS=linux
ARG GOARCH=amd64

RUN apk add git make && make build

FROM alpine

WORKDIR /

COPY --from=builder /device-service/build/device-service /device-service

ENV PORT=5050
ENV DB_HOST=
ENV DB_PORT=6379

ENTRYPOINT ["/device-service"]
