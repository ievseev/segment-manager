# Build the Go app
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o segment-manager ./cmd/service/

# Runner stage
FROM alpine AS runner

WORKDIR /app

# Устанавливаем bash и netcat
RUN apk add --no-cache bash netcat-openbsd

# Копируем бинарный файл приложения
COPY --from=builder /app/segment-manager ./

# Копируем директорию миграций
COPY --from=builder /app/db/migrations /app/db/migrations

# Копируем файл .env
COPY .env ./

# Копируем скрипт wait-for-it.sh
COPY wait-for-it.sh ./

# Делаем скрипт исполняемым
RUN chmod +x wait-for-it.sh

# Точка входа с ожиданием базы данных
CMD ["./wait-for-it.sh", "postgres:5432", "--", "./segment-manager"]
