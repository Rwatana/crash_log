FROM golang:1.20-alpine AS builder

WORKDIR /app

ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod ./

RUN go mod download

COPY logger.go ./

RUN go build -o logger .

RUN chmod +x /app/logger

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/logger .

RUN chmod +x /app/logger

RUN mkdir -p /app/logs

ENTRYPOINT ["./logger"]
