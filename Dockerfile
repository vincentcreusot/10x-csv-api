FROM golang:1.21 AS build

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download 

# Copy source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production image
FROM alpine:latest

WORKDIR /app

# Copy binary from build stage
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"] 
