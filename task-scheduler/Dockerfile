# Stage 1: Build the Go binary (compatible with ARM64)
FROM --platform=linux/arm64/v8 golang:1.23-alpine AS builder

# Install build tools (git and build-base)
RUN apk add --no-cache git build-base

# Set the working directory
WORKDIR /app

# Set environment variables for cross-compilation to ARM64
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
# Enable Go module proxy to fetch dependencies
ENV GOPROXY=https://proxy.golang.org,direct

# Copy go.mod and go.sum to leverage caching of dependencies
COPY go.mod go.sum ./

# Download dependencies and clean up unnecessary files
RUN go mod tidy

# Copy the entire project directory to the container
COPY . .

# Build the Go binary (main.go as the entry point)
RUN go build -o /app/scheduler/main ./main.go

# Stage 2: Create a minimal container for running the application
FROM --platform=linux/arm64/v8 alpine:latest

# Install necessary CA certificates and bash for the script
RUN apk add --no-cache ca-certificates bash

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/scheduler/main /app/scheduler/main

# Copy the configuration files
COPY config /app/config

# Copy the wait-for-kafka.sh script to the container
COPY wait-for-kafka.sh /app/wait-for-kafka.sh

# Set executable permissions explicitly
RUN chmod +x /app/scheduler/main /app/wait-for-kafka.sh

# Expose the port your application listens on
EXPOSE 8081

# Set environment variables (optional - update as needed)
ENV MODEL_NAME="resnet50"
ENV DATA_SOURCE="dataset1"

# Run the binary using the wait-for-kafka script
ENTRYPOINT ["/app/wait-for-kafka.sh", "kafka:9092", "/app/scheduler/main"]
