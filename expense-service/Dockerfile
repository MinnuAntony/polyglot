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
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o expense-service main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/expense-service .
EXPOSE 8080
CMD ["./expense-service"]

