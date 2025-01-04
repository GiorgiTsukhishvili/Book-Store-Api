# Stage 1: Build
FROM golang:1.20-alpine AS build

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Copy Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application source code
COPY . .

# Build the application
RUN go build -o app ./main.go

# Stage 2: Runtime
FROM alpine:latest AS runtime

# Set working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./app"]
