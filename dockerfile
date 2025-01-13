# Stage 1: Build the Go application
FROM golang:1.21 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Create a smaller runtime container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose the port that the Gin application will run on
EXPOSE 8080

# Run the application
CMD ["./main"]
