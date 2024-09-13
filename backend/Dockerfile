# Use a specific Go image that matches your requirements
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code and build the Go application
COPY . ./
RUN go build -o main .

# Use a compatible base image for the runtime environment
FROM debian:bookworm-slim

# Install necessary libraries
RUN apt-get update && apt-get install -y \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 3000

# Run the application
CMD ["./main"]
