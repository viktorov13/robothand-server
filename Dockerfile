# Stage 1: Build Go binary
FROM golang:1.25.2-bookworm AS builder

WORKDIR /app

# Устанавливаем зависимости для сборки go-sqlite3
RUN apt-get update && apt-get install -y \
    build-essential \
    libsqlite3-dev \
    git \
    && rm -rf /var/lib/apt/lists/*

# Копируем модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник для Linux amd64 с включённым cgo
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o robot-server ./cmd/server

# Stage 2: минимальный образ
FROM debian:bookworm-slim

WORKDIR /app

# Устанавливаем sqlite runtime
RUN apt-get update && apt-get install -y \
    sqlite3 \
    libsqlite3-0 \
    && rm -rf /var/lib/apt/lists/*

# Копируем собранный бинарник
COPY --from=builder /app/robot-server .

# Копируем директории для данных и загрузок
COPY data/ ./data/
COPY uploads/ ./uploads/

EXPOSE 8080

CMD ["./robot-server"]