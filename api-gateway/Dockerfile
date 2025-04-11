# Build stage
FROM golang:1.24-alpine AS build
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy everything
COPY . .

# Build the binary from cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway ./cmd

# Final stage
FROM alpine:3.18
COPY --from=build /gateway /gateway
EXPOSE 8090
ENTRYPOINT ["/gateway"]
