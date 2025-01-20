# Устанавливаем базовый образ для стадии сборки
FROM golang:1.20 AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .

# Загружаем зависимости и собираем проект
RUN go mod download
RUN go build -o proxy-server

# Устанавливаем базовый образ для финального контейнера
FROM alpine:latest

# Устанавливаем certs для поддержки HTTPS
RUN apk --no-cache add ca-certificates

# Задаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированное приложение из стадии сборки
COPY --from=builder /app/proxy-server .

# Экспортируем порт, на котором будет работать приложение
EXPOSE 8080

# Указываем команду запуска контейнера
CMD ["./proxy-server"]
