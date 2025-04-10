FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/docs ./docs 
COPY --from=builder /app/.env.docker ./.env

EXPOSE 8080

CMD ["./myapp"]