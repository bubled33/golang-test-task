## Описание

Простой микросервис на Go с REST эндпоинтом, который принимает число, сохраняет его в PostgreSQL и возвращает отсортированный список всех сохранённых чисел.

## Требования

- Go 1.24+
- Docker и Docker Compose
- PostgreSQL (вплоть до запуска через Docker)

## Запуск

1. Соберите и запустите сервис вместе с базой через Docker Compose:

```bash
make up
```

2. Сервер будет доступен по адресу `http://localhost:8080/number`

## Тестирование

1. Убедитесь, что запущен PostgreSQL (например, через Docker):

```bash
docker run -d --name test-postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=numbers -p 5432:5432 postgres:15-alpine
```

2. Установите переменную окружения для тестов:

```bash
export TEST_DATABASE_URL=postgres://user:password@localhost:5432/numbers?sslmode=disable
```

3. Запустите тесты:

```bash
TEST_DATABASE_URL=$TEST_DATABASE_URL go test ./... -v
```

## Пример запроса

```bash
curl -X POST http://localhost:8080/number \
  -H "Content-Type: application/json" \
  -d '{"number": 3}'
```

Результат:

```json
[3]
```

Отправка следующего числа:

```bash
curl -X POST http://localhost:8080/number \
  -H "Content-Type: application/json" \
  -d '{"number": 2}'
```

Результат:

```json
[2, 3]
```

## Структура проекта

- `cmd/api/` – точка входа и запуск сервера
- `internal/handler/` – HTTP обработчики
- `internal/repository/` – работа с базой данных
- `migrations/` – SQL миграции для базы
- `Dockerfile` и `docker-compose.yml` – для контейнеризации
- `Makefile` – удобные команды запуска и остановки
