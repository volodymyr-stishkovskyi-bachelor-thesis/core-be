# Используем официальный Go-образ
FROM golang:1.20

WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем код в контейнер
COPY . .

# Компилируем приложение
RUN go build -o core-be ./cmd/main.go

# Указываем порт
EXPOSE 8080

# Запускаем сервер
CMD ["/app/core-be"]