# Start from the base Golang image
FROM golang:1.21

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
