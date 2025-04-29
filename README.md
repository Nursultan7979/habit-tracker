# Habit Tracker 

## Описание
**Habit Tracker** — это REST API на Go для отслеживания привычек пользователей. Позволяет:

- Регистрировать и аутентифицировать пользователей (JWT)
- Управлять профилем пользователя
- Создавать и отслеживать привычки
- Вести и анализировать логи привычек

## Технологии
- Язык: Go
- Фреймворк: Gin
- БД: PostgreSQL
- ORM: GORM
- JWT: Аутентификация
- Миграции: `golang-migrate`

## Установка и запуск
### 1. Клонировать репозиторий
```
git clone https://github.com/Nursultan7979/habit-tracker.git
cd habit-tracker
```

### 2. Установить зависимости
```
go mod tidy
```

### 3. Настроить базу данных (PostgreSQL)
Создать базу `habit_tracker` и указать данные в `config/database.go`:
```go
var dsn = "user=postgres password=yourpassword dbname=habit_tracker port=5432 sslmode=disable"
```

### 4. Запустить сервер
```
go run main.go
```
Сервер запустится на `http://localhost:8080`

## Применить миграции (если используете миграции)
### Применить все миграции
.\migrate.windows-amd64\migrate.exe -path db/migrations -database "postgres://postgres:yourpassword@localhost:5432/habit_tracker?sslmode=disable" up

### Откатить последнюю миграцию
.\migrate.windows-amd64\migrate.exe -path db/migrations -database "postgres://postgres:yourpassword@localhost:5432/habit_tracker?sslmode=disable" down 1

## Аутентификация
Регистрация: POST /register

Вход: POST /login

После входа вы получите JWT-токен.

Для доступа к защищённым эндпоинтам используйте:

Authorization: Bearer <ваш токен>

## API Эндпоинты

### Пользователи
| Метод | Эндпоинт      | Описание             |
|-------|---------------|----------------------|
| POST  | `/register`   | Регистрация          |
| POST  | `/login`      | Вход                 |
| GET   | `/profile`    | Получить профиль     |
| PUT   | `/profile`    | Обновить профиль     |
| DELETE| `/profile`    | Удалить профиль      |

### Привычки
| Метод | Эндпоинт      | Описание             |
|-------|---------------|----------------------|
| POST  | `/habits`     | Создать привычку     |
| GET   | `/habits`     | Получить список      |
| PUT   | `/habits/:id` | Обновить привычку    |
| DELETE| `/habits/:id` | Удалить привычку     |

### Логи привычек
| Метод | Эндпоинт         | Описание               |
|-------|------------------|------------------------|
| POST  | `/habitlogs`     | Добавить запись        |
| GET   | `/habitlogs`     | Получить все записи    |
| PUT   | `/habitlogs/:id` | Обновить запись        |
| DELETE| `/habitlogs/:id` | Удалить запись         |

## Примеры (Postman)
**Регистрация:**
```json
POST http://localhost:8080/register
{
"name": "Nursultan",
"email": "nursultan@example.com",
"password": "123456"
}

```

**Создание привычки:**
```json
POST http://localhost:8080/habits
Authorization: Bearer <ваш токен>
{
"title": "Читать книги",
"description": "Читать по 30 минут каждый день"
}

```


Проект создан для учебных целей.

