package core

import (
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/messagebroker"
	"github.com/VladyslavLukyanenko/GopherAlert/contracts"
	"github.com/streadway/amqp"
)

func SendMessage(contract contracts.MonitoringContract) {
	payload, _ := json.Marshal(&contract)
	err := messagebroker.Channel.Publish(
		"monitoring-service-exchange",
		"monitoring-service",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
		})
	if err != nil {
		print(err)
	}
}
