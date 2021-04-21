package monitors

import (
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/MonitoringService/messagebroker"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Notify(webhook core.Webhook) {
	payload, err := json.Marshal(&webhook)
	if err != nil {
		log.Error("Failed Marshaling Webhook")
		return
	}
	err = messagebroker.Channel.Publish("publisher-service-exchange", "publisher-service", false, false, amqp.Publishing{
		DeliveryMode: amqp.Transient,
		ContentType:  "application/json",
		Body:         payload,
	})
	if err != nil {
		log.Error("Failed sending message to publisher")
		return
	}
}
