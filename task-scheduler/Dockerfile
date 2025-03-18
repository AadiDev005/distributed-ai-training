# Stage 1: Build the Go binary (compatible with ARM64)
FROM --platform=linux/arm64 golang:1.23-alpine AS builder

# Install build tools
RUN apk add --no-cache git build-base

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project directory to the container
COPY . .

# Set environment variables for cross-compilation to ARM64
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

# Build the Go binary (main.go as the entry point)
RUN go build -o /app/task-scheduler ./main.go

# Stage 2: Create a minimal container for running the application
FROM --platform=linux/arm64 alpine:latest

# Install necessary CA certificates (for HTTPS communication)
RUN apk add --no-cache ca-certificates bash

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/task-scheduler .

# Copy the configuration files
COPY config ./config

# Copy the wait-for-kafka.sh script to the container
COPY wait-for-kafka.sh /app/wait-for-kafka.sh

# Set executable permissions explicitly
RUN chmod +x /app/task-scheduler /app/wait-for-kafka.sh

# Expose the port your application listens on
EXPOSE 8081

# Run the binary
ENTRYPOINT ["/app/wait-for-kafka.sh"]
CMD ["/app/task-scheduler"]
