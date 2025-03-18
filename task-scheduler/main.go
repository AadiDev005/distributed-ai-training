package main

import (
	"task-scheduler/scheduler"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/") // Updated to use relative path for config directory
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Error reading config file: %v", err)
	}

	// Set logging level
	level, err := logrus.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		logrus.Fatalf("Invalid log level: %v", err)
	}
	logrus.SetLevel(level)

	logrus.Info("Configuration loaded successfully.")
}

func main() {
	initConfig()
	logrus.Info("Starting Task Scheduler...")
	scheduler.StartScheduler()
}
