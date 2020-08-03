# API чат-сервера

Реализация чат-сервера, предоставляющего HTTP API для работы с чатами и сообщениями пользователя.

## Запуск
- `docker-compose up --build` - запуск сервера на 9000 порту, базы данных со схемой
- `docker-compose up mysql` - запускает только базу со схемой, но без приложения, приложение можно запустить `go run main.go`
- `docker-compose down` - остановить и удалить все контейнеры
- `docker run --rm -v $(pwd):/test -w /test golang:1.14 go test ./...` - запуск тестов

## Конфигурация

- `PORT` - порт HTTP сервера (по умолчанию 9000)
- `MYSQL_DSN` - dsn базы данных (по умолчанию настроено на бд в докере)

Схема базы данных в файле `schema.sql`.

## Примеры API методов для тестирования

### Добавить нового пользователя

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add
```

Перед добавлением пользователя проверяется непустой ли usermane в запросе.

### Создать новый чат между пользователями

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": [1, 2]}' \
  http://localhost:9000/chats/add
```

Перед созданием чата проверяется сущетствуют ли users.

### Отправить сообщение в чат от лица пользователя

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1, "author": 2, "text": "hi"}' \
  http://localhost:9000/messages/add
```

Перед отправкой сообщения проверяется находится ли author в данном чате.

### Получить список чатов конкретного пользователя

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": 1}' \
  http://localhost:9000/chats/get
```

### Получить список сообщений в конкретном чате

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1}' \
  http://localhost:9000/messages/get
```