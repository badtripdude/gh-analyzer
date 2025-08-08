import logging

from confluent_kafka import Producer

producer = Producer({"bootstrap.servers": "localhost:9092"})

logger = logging.getLogger("kafka_producer")


def delivery_report(err, msg):
    if err:
        logger.error(f"Delivery failed: {err}")
    else:
        logger.info(f"Message delivered to {msg.topic()} [{msg.partition()}]")


def send_event(topic: str, data: dict):
    producer.produce(topic, value=str(data), callback=delivery_report)
    print(data)
    producer.flush()
