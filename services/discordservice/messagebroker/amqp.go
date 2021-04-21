package messagebroker

import (
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/configs"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var Channel *amqp.Channel

func InitAMQP() {
	var err error
	conn, err = amqp.Dial(configs.AppConfig.RabbitMQ.URI)
	if err != nil || conn == nil {
		log.Panic("Error while creating connection")
		return
	}
	log.Debug("Created connection")

	Channel, err = conn.Channel()
	if err != nil {
		log.Panic("Error while creating channel")
		return
	}
	log.Debug("Created channel")

	err = Channel.ExchangeDeclare(
		"discord-service-exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic("Error while declaring exchange")
		return
	}
	log.Debugf("Declared exchange %s", "monitoring-service-exchange")

	log.Debug("Connected to RabbitMQ")
}
