package main

import (
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/contracts"
	"github.com/VladyslavLukyanenko/MonitoringService/configs"
	localcore "github.com/VladyslavLukyanenko/MonitoringService/core"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var channel <-chan amqp.Delivery
var amqproutes []localcore.Route

func InitAMQP()  {
	var err error
	conn, err = amqp.Dial(configs.AppConfig.RabbitMQ.URI)
	if err != nil || conn == nil {
		log.Panic("Error while creating connection")
		return
	}
	log.Debug("Created connection")

	ch, err = conn.Channel()
	if err != nil {
		log.Panic("Error while creating channel")
		return
	}
	log.Debug("Created channel")

	err = ch.ExchangeDeclare(
		"monitoring-service-exchange",
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

	queue, err := ch.QueueDeclare(
		"monitoring-service-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error(err.Error())
		log.Panic("Error while declaring queue")
		return
	}
	log.Debugf("Declared queue %s", queue.Name)
	err = ch.QueueBind(
		queue.Name,
		"monitoring-service",
		"monitoring-service-exchange",
		false,
		nil,
	)
	if err != nil {
		log.Panic("Error while binding queue")
		return
	}
	log.Debugf("Binded to queue %s", queue.Name)
	channel, err = ch.Consume(
		queue.Name,
		"monitoring-service",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("Error while consuming queue to function, %s", err.Error())
		return
	}
	log.Debugf("Consuming queue %s", queue.Name)
	log.Debug("Connected to RabbitMQ")
	go handler(channel)
}

func BindQueueToFunction(routingKey string, function func(amqp.Delivery))  {
	amqproutes = append(amqproutes, localcore.Route{
		RoutingKey:    routingKey,
		RouteFunction: function,
	})
}
func handler(channel <-chan amqp.Delivery) {
	for message := range channel {
		log.Debugf("Received message: %s", message.Body)
		webhook := &contracts.PublishContract{}
		err := json.Unmarshal(message.Body, &webhook)
		if err != nil {

		}
	}
}