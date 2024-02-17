# выбираем образ для сборки контейнера, здесь подойдет легковесный гошный alpine
FROM golang:1.21-alpine AS builder

# устанавливаем рабочую директорию внутри контейнера. При запуске окажемся там
WORKDIR /app

# копируем исходники
COPY . .

# подтягиваем зависимости сервиса
RUN go mod download

# собираем бинарь
RUN CGO_ENABLED=0 go build -o segment-manager ./cmd/service/

# в итоге приложение запустится в еще более легком образе alpine - восхитительно
FROM alpine AS runner

COPY --from=builder ["/app/segment-manager","/app/.env", "./"]

# точка входа
CMD ["./segment-manager"]