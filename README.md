# Habit Tracker 

## Описание
Habit Tracker — это RESTful API для управления привычками пользователей. Позволяет:
- Создавать, получать, обновлять и удалять пользователей
- Добавлять, отслеживать и управлять привычками
- Вести логи по привычкам

## Технологии
- **Язык**: Go
- **Фреймворк**: Gin
- **База данных**: PostgreSQL
- **ORM**: Gorm

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

## API Эндпоинты

### Пользователи
| Метод | Эндпоинт | Описание |
|--------|----------|------------|
| `POST` | `/users` | Создать пользователя |
| `GET` | `/users` | Получить список пользователей |
| `PUT` | `/users/:id` | Обновить пользователя |
| `DELETE` | `/users/:id` | Удалить пользователя |

### Привычки
| Метод | Эндпоинт | Описание |
|--------|----------|------------|
| `POST` | `/habits` | Создать привычку |
| `GET` | `/habits` | Получить список привычек |
| `PUT` | `/habits/:id` | Обновить привычку |
| `DELETE` | `/habits/:id` | Удалить привычку |

### Логи привычек
| Метод | Эндпоинт | Описание |
|--------|----------|------------|
| `POST` | `/habitlogs` | Добавить запись в лог |
| `GET` | `/habitlogs` | Получить все записи лога |
| `PUT` | `/habitlogs/:id` | Обновить запись |
| `DELETE` | `/habitlogs/:id` | Удалить запись |

## Postman
**Создание пользователя:**
```json
POST http://localhost:8080/users
{
  "name": "Nursultan",
  "email": "nursultan@example.com"
}
```

**Создание привычки:**
```json
POST http://localhost:8080/habits
{
  "title": "Читать книги",
  "description": "Читать по 30 минут каждый день",
  "user_id": 1
}
```

Проект создан для учебных целей.

