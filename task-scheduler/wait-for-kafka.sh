#!/bin/bash

KAFKA_HOST=$1

echo "Waiting for Kafka at $KAFKA_HOST..."

while ! nc -z $(echo $KAFKA_HOST | cut -d':' -f1) $(echo $KAFKA_HOST | cut -d':' -f2); do
  echo "Kafka is not available yet, sleeping..."
  sleep 1
done

echo "Kafka is up - executing command"
exec /app/scheduler/main "${@:2}"
