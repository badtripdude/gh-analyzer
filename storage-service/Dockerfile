FROM golang:1.24-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o storage-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/storage-service .
COPY .env .

CMD ["./storage-service"]

# multi-stage build for smaller final image.
# consider further optimizations for build speed and image size.