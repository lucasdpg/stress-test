FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/stress-test ./cmd/cli/main.go

FROM alpine:3.18 AS release

WORKDIR /app

RUN apk add --no-cache curl ca-certificates

COPY --from=builder /app/stress-test /app/stress-test

ENTRYPOINT ["/app/stress-test", "start"]