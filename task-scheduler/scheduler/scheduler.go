package scheduler

import (
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartScheduler() {
	broker := viper.GetString("kafka.broker")
	topic := viper.GetString("kafka.topic")

	logrus.Infof("Connecting to Kafka at %s", broker)

	// Kafka producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		logrus.Fatalf("Failed to start Kafka producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			logrus.Errorf("Failed to close Kafka producer: %v", err)
		}
	}()

	logrus.Info("Task Scheduler is running...")

	// Get model name and data source from environment variables
	model := os.Getenv("MODEL_NAME")
	data := os.Getenv("DATA_SOURCE")

	if model == "" || data == "" {
		logrus.Fatalf("Error: MODEL_NAME or DATA_SOURCE environment variable not set")
	}

	message := fmt.Sprintf("Model: %s, Data: %s", model, data)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	for {
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logrus.Errorf("Failed to send message: %v", err)
		} else {
			logrus.Infof("Message sent to partition %d with offset %d", partition, offset)
		}

		// Sleep for a while before sending the next message
		time.Sleep(5 * time.Second)
	}
}
