# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy Go module files
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy backend source
COPY backend/ .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server

# ---------- Runtime stage ----------
FROM alpine:latest

WORKDIR /app

# Copy compiled binary
COPY --from=builder /app/server .

# Copy frontend static files
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]
