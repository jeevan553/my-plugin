# Use official golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application (plugin.go) inside the container
RUN go build -o plugin ./plugin.go

# Define the command that runs when the container starts
CMD ["./plugin"]