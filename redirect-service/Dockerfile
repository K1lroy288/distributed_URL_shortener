FROM golang:alpine AS builder

WORKDIR /redirect-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /redirect-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /redirect-service/redirect-service .

COPY .env .

EXPOSE 8002

CMD ["./redirect-service"]
