# ===============================
# Stage 1 — build
# ===============================
FROM golang:1.26-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем статический бинарник (без libc)
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o tg-bot .

# ===============================
# Stage 2 — runtime
# ===============================
FROM alpine:3.22

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/tg-bot ./tg-bot

# ВАЖНО: твой конфиг читает texts.yaml из текущей директории
COPY texts.yaml ./texts.yaml

ENTRYPOINT ["./tg-bot"]