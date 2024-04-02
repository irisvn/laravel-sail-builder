FROM golang:1.20-alpine

WORKDIR /app

RUN apk update
RUN apk add ca-certificates && update-ca-certificates

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY *.go ./
COPY laravel.temp ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /laravel-build

EXPOSE 8080

# Run
CMD ["/laravel-build"]