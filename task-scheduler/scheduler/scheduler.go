package scheduler

import (
	"bufio"
	"fmt"
	"os"

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

	// Listen for incoming tasks
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter model name: ")
		model, err := reader.ReadString('\n')
		if err != nil {
			logrus.Errorf("Error reading model name: %v", err)
			continue
		}
		model = model[:len(model)-1] // Trim the newline character

		fmt.Print("Enter data source: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			logrus.Errorf("Error reading data source: %v", err)
			continue
		}
		data = data[:len(data)-1] // Trim the newline character

		message := fmt.Sprintf("Model: %s, Data: %s", model, data)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logrus.Errorf("Failed to send message: %v", err)
		} else {
			logrus.Infof("Message sent to partition %d with offset %d", partition, offset)
		}
	}
}
