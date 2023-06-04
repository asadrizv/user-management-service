# Use the official Go image as the base image
FROM golang:1.20.0-alpine3.16 AS build

# Set the working directory to /app
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o user-service ./cmd/server/main.go

# Create a new image from scratch
FROM alpine:3.14

# Install PostgreSQL client
RUN apk add --no-cache postgresql-client

# Set the working directory to /app
WORKDIR /app

# Copy the built Go application to the working directory
COPY --from=build /app/user-service .

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Set the command to run the HTTP server
CMD ["./user-service", "serve"]
