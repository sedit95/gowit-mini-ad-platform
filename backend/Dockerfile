# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the API binary
RUN go build -o api ./cmd/api

# Runtime stage
FROM alpine:3.20

WORKDIR /app

# Install ca-certificates in case outbound HTTPS is required
RUN apk add --no-cache ca-certificates

# Copy only the compiled binary
COPY --from=builder /app/api .

# Expose API port
EXPOSE 8080

# Command to run the executable
CMD ["./api"]
