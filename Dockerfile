FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o lamoda ./cmd/lamoda_api/main.go
CMD ["/app/lamoda"]
