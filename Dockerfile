# Первый этап: сборка приложения
FROM golang:1.23.0-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код приложения
COPY . .

# Загружаем зависимости
RUN go mod download

# Собираем приложение в бинарный файл
RUN go build -o ./bin/vet ./cmd/main.go

# Второй этап: создание минимального образа
FROM alpine:latest

# Устанавливаем bash и другие необходимые утилиты
RUN apk add --no-cache bash curl

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарник, конфиг и wait-for-it.sh из стадии сборки
COPY --from=builder /app/bin/vet .
COPY --from=builder /app/prod.yaml .
COPY wait-for-it.sh .

# Делаем бинарный файл исполняемым
RUN chmod +x ./vet wait-for-it.sh

# Открываем нужный порт (если приложение слушает определенный порт)
EXPOSE 3000

# Команда для запуска приложения с ожиданием базы данных
CMD ["bash", "./wait-for-it.sh", "db:5432", "--", "./vet"]
