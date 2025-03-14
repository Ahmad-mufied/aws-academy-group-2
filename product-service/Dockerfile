# Stage 1: Install dependencies
FROM golang:1.24.0-bookworm AS deps

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Stage 2: Build the application
FROM golang:1.24.0-bookworm AS builder

WORKDIR /app
COPY --from=deps /go/pkg/mod /go/pkg/mod
COPY . .

# Build flags for better optimization
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build with additional security flags
RUN go build -ldflags="-w -s" \
    -trimpath \
    -o main cmd/main.go

# Final stage: Run the application
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Copy the binary
COPY --from=builder /app/main .

# Use unprivileged user
USER nonroot

ENTRYPOINT ["/app/main"]