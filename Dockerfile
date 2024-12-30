# Use the official Golang image
FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod file
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Print Go environment and list files for debugging
RUN go env
RUN ls -la /app

# Build the Go app
RUN go build -o main ./cmd/gocache

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]