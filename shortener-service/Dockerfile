FROM golang:alpine AS builder

WORKDIR /shortener-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /shortener-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /shortener-service/shortener-service .

COPY .env .

EXPOSE 8001

CMD ["./shortener-service"]
