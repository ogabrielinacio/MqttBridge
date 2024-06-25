# First stage: build the Go application
FROM golang:1.19-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o mqttbridge .

# Second stage: create a minimal image with the compiled Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/mqttbridge .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the binary
CMD ["./mqttbridge"]

