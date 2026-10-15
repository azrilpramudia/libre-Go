# ---- build stage ----
FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

# ---- run stage ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/api .
COPY .env.example .env

EXPOSE 8080

CMD ["./api"]