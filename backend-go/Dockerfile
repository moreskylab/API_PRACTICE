# Stage 1: Build environment
FROM golang:1.27.1-alpine AS builder

WORKDIR /usr/src/app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build a statically linked binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/app .

# Stage 2: Minimal runtime environment
FROM alpine:3.24

WORKDIR /root/
COPY --from=builder /usr/local/bin/app .

EXPOSE 8081

CMD ["./app"]
