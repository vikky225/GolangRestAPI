# Use the official Golang image as a build stage
FROM golang:1.20 as build

# Set the Current Working Directory inside the container
WORKDIR /app

# Enable CGO
ENV CGO_ENABLED=1

# Install the necessary dependencies for SQLite
RUN apt-get update && apt-get install -y gcc libc-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM debian:stable-slim

# Install SQLite runtime dependencies
RUN apt-get update && apt-get install -y libsqlite3-0

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
