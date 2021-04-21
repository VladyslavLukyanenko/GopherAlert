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

	conn, _ = amqp.Dial("amqps://psoxlllh:CGbU2fNW9gixVF0s28W1AATVJ-mpw29T@cow.rmq2.cloudamqp.com/psoxlllh")

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
		//WebhookURI:         "https://hooks.slack.com/services/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",
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
