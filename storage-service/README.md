## GH Analyzer Storage Service

### Перед началом создайте конфиг файл .env

```
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gh-analyzer
DB_PORT=5432

KAFKA_BROKER=kafka:9092
```

### Запуск Kafka

```
make kafka-up
```

### Остановка Kafka

```
make kafka-down
```

### Запуск сервиса (включая сборку образа)

```
make docker-up
```

### Остановка сервиса

```
make docker-down
```

### Полное очищение (с удалением томов)

```
make docker-clear
```

### Просмотр логов сервиса

```
make logs
```

### Проверка содержимого БД

#### Подключись к базе (например, через pgadmin) введя значения из .env файла, чтобы удобно смотреть обновления данных, используйте скрипт:

```
\x
SELECT * FROM git_hub_analyses;
\x
```

### Отправка тестового сообщения в Kafka

```
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-console-producer.sh --bootstrap-server kafka:9092 --topic analysis.results
```

### После запуска введи JSON с анализом, например:
```
{"task_id":"task123","user_id":"user456","username":"deimossy","repo_count":5,"total_stars":42,"languages":{"Go":15000,"Python":8000},"last_commit":"2025-08-09T12:00:00Z","computed_at":"2025-08-09T12:05:00Z"}
```
### Что слушает консюмер

#### Сервис слушает Kafka-топик analysis.results. Он ожидает сообщения с анализом репозиториев и сохраняет их в таблицу git_hub_analyses в базе данных.