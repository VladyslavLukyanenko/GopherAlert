package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	RabbitMQ rabbitMQ
	Logger   logger
	MongoDB  mongodb
	Discord discord
}

type rabbitMQ struct {
	URI string
}
type logger struct {
	Level string
}
type mongodb struct {
	URI string
}
type discord struct {
	Token string
}

var AppConfig config

func ReadConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")

	if os.Getenv("dev") != "" {
		viper.SetConfigName("config.development")
	} else {
		viper.SetConfigName("config.production")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Failed to read config: %s", err.Error())
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Error("Failed to unmarshal config to struct")
	}
}