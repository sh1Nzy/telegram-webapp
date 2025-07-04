# TelegramShop Backend API

Backend API для приложения TelegramShop, построенное на Go с использованием Fiber фреймворка.

## Swagger Документация

API документация автоматически генерируется с помощью Swagger и доступна по адресу:

```
http://localhost:8080/swagger/index.html
```

### Генерация документации

Для обновления Swagger документации после изменения аннотаций используйте:

```bash
make swagger
```

Или напрямую:

```bash
$(HOME)/go/bin/swag init -g cmd/main.go
```

### Запуск приложения

```bash
# Генерация документации и запуск приложения
make dev

# Только запуск приложения
make run

# Только генерация документации
make swagger
```

## API Endpoints

### Основные разделы:

- **Users** (`/api/v1/users/*`) - управление пользователями
- **Basket** (`/api/v1/basket/*`) - корзина пользователя
- **Favorites** (`/api/v1/favorites/*`) - избранные товары
- **Orders** (`/api/v1/orders/*`) - заказы

### Структура ответов

Все API возвращают стандартизированную структуру:

```json
{
    "status": "success_operation_name",
    "data": { /* данные */ }
}
```

Или в случае ошибки:

```json
{
    "status": "error_operation_name", 
    "data": "сообщение об ошибке"
}
```

## Разработка

### Требования

- Go 1.24.3+
- PostgreSQL
- Swaggo/Swag для генерации документации

### Установка зависимостей

```bash
go mod download
```

### Миграции базы данных

```bash
# Применить миграции
make migrate

# Откатить миграции  
make migrate-down
```

## Swagger Аннотации

Все API методы документированы с помощью аннотаций Swagger:

```go
// CreateUser creates a new user
// @Summary Create new user
// @Description Creates a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "User creation data"
// @Success 200 {object} models.UserResponse "User successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Router /api/v1/users [post]
```

Документация автоматически обновляется при изменении аннотаций в коде. 