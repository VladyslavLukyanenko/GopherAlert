package main

import (
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/contracts"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func main() {
	initAmqp()
}

func initAmqp() {
	var _ error

	conn, _ = amqp.Dial("AMQP_CONNECTION_URL")

	ch, _ = conn.Channel()

	_ = ch.ExchangeDeclare(
		"monitoring-service-exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	monitor := core.Monitor{
		Channel:            "togglebit",
		//Delay:              30,
		//WebhookURI:         "SLACK_WEBHOOK_URL",
		//MonitoringPlatform: core.Youtube,
		//DeliveryPlatform:   core.Slack,
	}
	mo, _ := json.Marshal(&monitor)
	contract := contracts.MonitoringContract{
		RoutingKey: "monitor-remove-task",
		Data:       string(mo),
	}
	payload, _ := json.Marshal(&contract)
	err := ch.Publish(
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
