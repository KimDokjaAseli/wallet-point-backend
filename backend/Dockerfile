# Build stage
FROM golang:1.24-alpine AS builder

# Install specific version of git to avoid compatibility issues
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# Disable CGO for static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /root/

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/server .

# Create directory for uploads
RUN mkdir -p public/uploads

# Expose port (default 8080)
EXPOSE 8080

# Run the binary
CMD ["./server"]
