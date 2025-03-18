package main

import (
	"task-scheduler/scheduler"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")   // Local run
	viper.AddConfigPath("/app/config") // Docker container path

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Error reading config file: %v", err)
	}

	// Set logging level with a default value
	level, err := logrus.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		logrus.Warnf("Invalid or missing log level in config, defaulting to 'info': %v", err)
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	logrus.Info("Configuration loaded successfully.")
}

func main() {
	initConfig()
	logrus.Info("Starting Task Scheduler...")
	scheduler.StartScheduler()
}
