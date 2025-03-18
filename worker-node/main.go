package main

import (
	"os"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting Worker Node...")

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "kafka:9092"
	}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{kafkaBroker}, config)
	if err != nil {
		logrus.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("train-tasks", 0, sarama.OffsetOldest)
	if err != nil {
		logrus.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	logrus.Info("Worker Node listening on 'train-tasks' topic...")

	for msg := range partitionConsumer.Messages() {
		task := string(msg.Value)
		logrus.Infof("Received task: %s", task)

		fields := strings.Split(task, ", ")
		if len(fields) < 5 {
			logrus.Warnf("Invalid task format: %s", task)
			continue
		}

		trainName := strings.Split(fields[0], ": ")[1]
		source := strings.Split(fields[1], ": ")[1]
		destination := strings.Split(fields[2], ": ")[1]
		seats := strings.Split(fields[3], ": ")[1]
		id := strings.Split(fields[4], ": ")[1]

		logrus.Infof("Processing train: %s (Source: %s, Destination: %s, Seats: %s, ID: %s)", trainName, source, destination, seats, id)
		time.Sleep(5 * time.Second)
		logrus.Infof("Processing completed for train: %s (ID: %s)", trainName, id)
	}
}
