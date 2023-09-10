FROM golang:1.21 AS build-stage

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download 

# Copy source code
COPY cmd cmd/
COPY pkg pkg/

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o csvapi cmd/10x-csv-api/10x-csv-api.go

# Production image
FROM gcr.io/distroless/base-debian12 AS run-stage

WORKDIR /

# Copy binary from build stage
COPY --from=build-stage /app/csvapi csvapi
COPY seattle-weather.csv seattle-weather.csv

USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT [ "./csvapi" ]
CMD [ "seattle-weather.csv" ]
