# Kafka (Docker + Bitnami) с автоинициализацией топиков

## 📌 Описание
Этот стенд поднимает **Apache Kafka** (в Zookeeper-режиме) через Docker Compose и автоматически создаёт топики, указанные в `.env.kafka`.  
В комплект входит:
- **docker-compose.kafka.yml** — описание сервисов Kafka, Zookeeper и скрипта инициализации.
- **.env.kafka** — переменные окружения для настройки брокера и списка топиков.
- **scripts/init-topics.sh** — скрипт, создающий топики при старте.
- **README.md** — этот файл.

## 📂 Структура
kafka/<br>
├─ docker-compose.kafka.yml # docker-compose конфигурация<br>
├─ .env.kafka # переменные окружения для брокера и топиков<br>
├─ scripts/init-topics.sh # скрипт создания топиков<br>
└─ README.md # документация<br>

## 🚀 Запуск
```bash
docker compose -f docker-compose.kafka.yml up -d
```
- Будут подняты zookeeper, kafka и topics-init.
- Скрипт init-topics.sh создаст топики из переменной KAFKA_INIT_TOPICS.

## 🛑 Остановка
```bash
docker compose -f docker-compose.kafka.yml down -v
```

## ⚙️ Настройка топиков
В .env.kafka укажи:
```dotenv
KAFKA_INIT_TOPICS="report.ready:3:1,report.failed:6:1"
```
Формат:
```dotenv
<имя>:<число_партиций>:<replication_factor>
```
Пример:
```dotenv
orders:4:1 → топик orders с 4 партициями и фактором репликации 1.
```

## 📜 Основные команды
Список топиков
```bash
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-topics.sh \
  --bootstrap-server kafka:9092 --list
```
Создание топика вручную
```bash
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-topics.sh \
  --bootstrap-server kafka:9092 \
  --create --topic test --partitions 3 --replication-factor 1
```
Продюсер (отправка сообщений)
```bash
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-console-producer.sh \
  --bootstrap-server kafka:9092 --topic test
```
Консьюмер (чтение сообщений)
```bash
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-console-consumer.sh \
  --bootstrap-server kafka:9092 --topic test --from-beginning
```