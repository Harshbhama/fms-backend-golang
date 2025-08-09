# Build stage
FROM golang:1.23-alpine AS builder


WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download


# Install fresh for hot reloading
RUN go install github.com/pilu/fresh@latest

# Copy the source code
COPY . .

# Final stage
FROM golang:1.23-alpine

WORKDIR /app

# Install necessary packages
RUN apk --no-cache add ca-certificates \
    && go install github.com/pilu/fresh@latest

# Copy the source code
COPY . .

# Expose port
EXPOSE 8081

# Command to run the application with fresh
CMD ["fresh"]
