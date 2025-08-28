# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o login-api ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/login-api .
COPY --from=builder /app/configs ./configs

EXPOSE 8080
CMD ["./login-api"]