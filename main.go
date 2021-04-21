package main

import (
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/contracts"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/streadway/amqp"
	"time"
)

var conn *amqp.Connection
var ch *amqp.Channel

func main() {
	initAmqp()
}

func initAmqp() {
	var _ error

	conn, _ = amqp.Dial("amqps://tvygwybl:sAlaFyRMjc5Vn9CEz7B4mAy_mofHlxSo@cow.rmq2.cloudamqp.com/tvygwybl")

	ch, _ = conn.Channel()

	_ = ch.ExchangeDeclare(
		"test-exchange", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // noWait
		nil,             // arguments
	)
	monitor := core.Monitor{
		Channel:            "togglebit",
		Delay:              30000,
		WebhookURI:         "https://discord.com/api/webhooks/804199560963948544/wSu1HmcVfDw_djm0zFqaiDerfBo1b08E3KvMeXjfFVvR7f-7xhTzfVdEi9xM9K8kCkKH",
		MonitoringPlatform: core.Youtube,
		DeliveryPlatform:   core.Discord,
	}
	mo, _ := json.Marshal(&monitor)
	contract := contracts.MonitoringContract{
		RoutingKey: "monitor-add-task",
		Data:       string(mo),
	}
	payload, _ := json.Marshal(&contract)
	err := ch.Publish(
		"publisher-service-exchange",
		"monitoring-service",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})
	if err != nil {
		print(err)
	}
}
