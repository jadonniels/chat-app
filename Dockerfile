# Stage 1: Build the Go binary
FROM golang:1.23.0-alpine as builder

# Install required dependencies (ca-certificates for HTTPS)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod tidy

# Copy the entire source code into the container
COPY . .

# Build the Go binary from the cmd/main.go file
RUN go build -o chatapp ./cmd

# Stage 2: Create a minimal image for the final app
FROM alpine:latest

# Install the CA certificates (for making HTTPS requests)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/chatapp .
COPY --from=builder /app/templates /root/templates
COPY ./public /root/public

# Expose the port the app will run on (example: 8080)
EXPOSE 8080

# Run the binary
CMD ["./chatapp"]
