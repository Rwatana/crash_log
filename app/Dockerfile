
FROM golang:1.20-alpine AS builder

WORKDIR /app

ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod ./

RUN go mod download

COPY app.go ./

RUN go build -o app .

RUN chmod +x /app/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .


RUN chmod +x /app/app

RUN mkdir -p /app/logs

ENTRYPOINT ["./app"]
