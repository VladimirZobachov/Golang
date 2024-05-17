# Hostess Service
Микросервис для хостесов. Сервис бронирования столиков в ресторане.

## Установка (Linux)
У вас должны быть установлено окружение: Git, Docker, Go 1.20

1. Клонирование репозитория

```git clone https://gl.eda1.ru/go/hostess-service.git```

2. Переход в директорию hostess-service

```cd hostess-service```

3. Установка базы данных в докере

```docker compose up```

4. Установка зависимостей

```go mod tidy```

6. Запуск сервиса для демонстрации возможностей

```go run cmd/hostess-service/main.go```

7. Cервис будет доступен по адресу

```http://localhost:8080/swagger/index.html```