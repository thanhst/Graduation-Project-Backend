# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first, then download deps (cache hiệu quả)
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source vào container
COPY . .

# Build binary
RUN go build -o apiserver ./cmd/server/main.go

# Final image
FROM alpine:latest

WORKDIR /app

# Copy binary từ builder stage
COPY --from=builder /app/apiserver .
COPY --from=builder /app/cmd ./cmd
COPY --from=builder /app/internal/db/migrations/sql /app/internal/db/migrations/sql
COPY --from=builder /app/uploads /app/uploads
# Mở cổng cho API (ví dụ 8080)
EXPOSE 8080

CMD ["./apiserver"]
