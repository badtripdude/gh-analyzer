#!/usr/bin/env bash
set -euo pipefail

bootstrap="kafka:9092"

attempts="${KAFKA_INIT_MAX_ATTEMPTS:-30}"
sleep_sec="${KAFKA_INIT_SLEEP_SEC:-2}"
topics="${KAFKA_INIT_TOPICS:-}"

echo "Waiting for Kafka at $bootstrap ..."
for i in $(seq 1 "$attempts"); do
  if /opt/bitnami/kafka/bin/kafka-metadata-quorum.sh --help >/dev/null 2>&1; then
    # просто проверка, что бинари на месте
    :
  fi
  if /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server "$bootstrap" --list >/dev/null 2>&1; then
    echo "Kafka is up."
    break
  fi
  echo "  attempt $i/$attempts: not ready, sleeping ${sleep_sec}s..."
  sleep "$sleep_sec"
  if [ "$i" -eq "$attempts" ]; then
    echo "Kafka is not ready after $attempts attempts" >&2
    exit 1
  fi
done

if [ -z "$topics" ]; then
  echo "No topics to create. Set KAFKA_INIT_TOPICS to create topics."
  exit 0
fi

IFS=',' read -ra arr <<< "$topics"
for spec in "${arr[@]}"; do
  name="${spec%%:*}"
  rest="${spec#*:}"
  partitions="${rest%%:*}"
  repl="${rest#*:}"

  echo "Ensuring topic: $name (partitions=$partitions, repl=$repl)"
  if /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server "$bootstrap" --topic "$name" --describe >/dev/null 2>&1; then
    echo "  Topic '$name' already exists."
  else
    /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server "$bootstrap" \
      --create --topic "$name" --partitions "$partitions" --replication-factor "$repl"
    echo "  Created."
  fi
done
