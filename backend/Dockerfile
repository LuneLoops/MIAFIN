# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiar go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar binario
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8080

# Ejecutar
CMD ["./main"]
