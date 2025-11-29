FROM golang:alpine AS builder

WORKDIR /auth-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /auth-service/auth-service .

COPY .env .

EXPOSE 8000

CMD ["./main"]
