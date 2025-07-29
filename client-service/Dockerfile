# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o client-service .

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/client-service .

# Expose port
EXPOSE 8081

# Command to run the executable
CMD ["./client-service"]
