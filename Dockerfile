# Step 1: Build the Go plugin using the official Go image
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the plugin source code and main program
COPY . .
# COPY plugin.go ./plugin.go
# COPY main.go ./main.go
# COPY common.go ./common.go

# Build the plugin (plugin.so)
RUN go build -o plugin.so -buildmode=plugin ./plugin.go ./common.go

# Build the main program
RUN go build -o main main.go ./common.go

# Step 2: Create the runtime image using Ubuntu 22.04 (which has the required GLIBC version)
FROM ubuntu:22.04

# Install necessary libraries (including libc6)
RUN apt-get update && apt-get install -y libc6

# Set the working directory
WORKDIR /app

# Copy the built plugin and main program from the builder stage
COPY --from=builder /app/plugin.so /app/
COPY --from=builder /app/main /app/

# Run the main program
CMD ["./main"]
