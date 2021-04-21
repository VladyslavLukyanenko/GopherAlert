package monitors

import (
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/VladyslavLukyanenko/MonitoringService/messagebroker"
	"github.com/streadway/amqp"
)

func Notify(webhook core.Webhook) {
	payload, err := json.Marshal(&webhook)
	err = messagebroker.Channel.Publish("publisher-service-exchange", "publisher-service", false, false, amqp.Publishing{
		DeliveryMode: amqp.Transient,
		ContentType:  "application/json",
		Body:         payload,
	})
	if err != nil {
		return
	}
}
