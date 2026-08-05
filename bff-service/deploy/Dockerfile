FROM golang:1.25 AS dev
WORKDIR /app

RUN go install github.com/air-verse/air@v1.65.3
RUN go install golang.org/x/tools/gopls@v0.21.0
RUN go install github.com/go-delve/delve/cmd/dlv@v1.27.0

FROM golang:1.25 AS builder 
WORKDIR /app

COPY ./bff-service ./bff-service
COPY ./postgres ./postgres
COPY ./redis ./redis
COPY ./rabbitmq ./rabbitmq
COPY ./config ./config

WORKDIR /app/bff-service
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:3.20 AS prod
RUN adduser -D appuser

WORKDIR /app
COPY --from=builder /app/bff-service .

USER root

RUN mv /app/deploy/entrypoint.sh /usr/local/bin/entrypoint.sh && \
    chmod +x /usr/local/bin/entrypoint.sh

USER appuser

EXPOSE 3000
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

FROM golang:1.25 AS test

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

