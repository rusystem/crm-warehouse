# Этап 1: сборка приложения
FROM golang:1.22 AS builder

WORKDIR /app

# Копируем go.mod и go.sum отдельно для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Сборка приложения
RUN go build -o crm-warehouse ./cmd/main.go

# Этап 2: создание минимального образа
FROM debian:bullseye-slim

WORKDIR /app

# Установка зависимостей (если нужны, например, psql)
RUN apt-get update && apt-get install -y libpq-dev && rm -rf /var/lib/apt/lists/*

# Копируем собранное приложение из первого этапа
COPY --from=builder /app/crm-warehouse /app/crm-warehouse

# Настройка прав доступа
RUN chmod +x crm-warehouse

# Запуск приложения
CMD ["./crm-warehouse"]