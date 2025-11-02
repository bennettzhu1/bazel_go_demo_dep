# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY main.go .

# Build the binary
RUN go build -o hello-k8s main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/hello-k8s .

# Default port
ENV PORT=8080
ENV APP_NAME=hello-k8s

EXPOSE 8080

CMD ["./hello-k8s"]


