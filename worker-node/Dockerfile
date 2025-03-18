# Build Stage
FROM --platform=linux/arm64/v8 golang:1.23-alpine AS builder

WORKDIR /app

# Install git and build tools to fetch modules and build the binary
RUN apk update && apk add --no-cache git build-base

# Set environment variables for cross-compilation to ARM64
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

# Copy go.mod and go.sum files first to leverage caching
COPY go.mod go.sum ./

# Force download dependencies
RUN go env -w GO111MODULE=on && go mod download

# Copy the application source code
COPY main.go ./

# Build the binary for ARM64 architecture
RUN go build -o worker main.go

# Run Stage
FROM --platform=linux/arm64/v8 alpine:latest

WORKDIR /app

# Install necessary CA certificates and bash for the script
RUN apk add --no-cache ca-certificates bash

# Copy the worker binary from the builder stage
COPY --from=builder /app/worker .

# Copy the wait-for-kafka script
COPY wait-for-kafka.sh /app/wait-for-kafka.sh
RUN chmod +x /app/wait-for-kafka.sh

# Run the binary with Kafka readiness check
CMD ["/app/wait-for-kafka.sh", "kafka:9092", "./worker"]
