# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy everythin into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Set the command to run when the container starts
CMD ["./main"]