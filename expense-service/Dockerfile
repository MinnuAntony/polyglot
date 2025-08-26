# FROM golang:1.21-alpine
# WORKDIR /app
# COPY go.mod . 
# COPY go.sum .
# RUN go mod download
# COPY . .
# RUN go build -o expense-service main.go
# EXPOSE 8080
# CMD ["./expense-service"]

# Build stage
#FROM golang:1.21-alpine AS build
#WORKDIR /app
#COPY go.mod go.sum ./
#RUN go mod download
#COPY . .
#RUN go build -o expense-service main.go

# Run stage
#FROM alpine:latest
#WORKDIR /app
#COPY --from=build /app/expense-service .
#EXPOSE 8080
#CMD ["./expense-service"]


# Use the official Go image as the base image for building
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Use `go mod download` to download all the dependencies listed in go.mod
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application, specifying the output name
RUN go build -o expense-service main.go

# --- Run Stage ---
# Use a minimal Alpine image for the final, smaller container
FROM alpine:3.18

# Set the working directory for the running application
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/expense-service .

# Expose the port the Go application will run on
EXPOSE 8080

# The command to run the application when the container starts
CMD ["./expense-service"]

