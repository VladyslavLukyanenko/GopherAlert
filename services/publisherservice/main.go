package main

import (
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/PublisherService/configs"
	"github.com/VladyslavLukyanenko/GopherAlert/PublisherService/publishers"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var messages <-chan amqp.Delivery

func main() {
	log.SetLevel(log.DebugLevel)
	configs.ReadConfig()
	initAmqp()
	webhook := core.Webhook{
		Notification:              core.Notification{
			ChannelName: "togglebit",
			Title: "Rust | Vim | Making a pixel editor | Command input, then some networking malarkey",
			Avatar: "https://static-cdn.jtvnw.net/jtv_user_pictures/0bb9c502-ab5d-4440-9c9d-14e5260ebf86-profile_image-70x70.png",
			Game: "Science & Technology",
			Viewers: 163,
			URI: "https://www.twitch.tv/togglebit",
		},
		URI:                    "https://hooks.slack.com/services/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",
		DeliveryPlatform:       core.Slack,
		NotificationType:       core.Live,
		MonitoringPlatformType: core.Twitch,
	}
	payload, err := json.Marshal(webhook)
	for true {
		err = ch.Publish(
			"publisher-service-exchange", // exchange
			"publisher-service",      // routing key
			false,              // mandatory
			false,              // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Transient,
				ContentType:  "application/json",
				Body:         payload,
			})
	}
	if err != nil {

	}
	//go handler(messages)
	select {}
}

func initAmqp() {
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
		"publisher-service-exchange",
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
	log.Debugf("Declared exchange %s", "publisher-service-exchange")
	//q, err := ch.QueueDeclare(
	//	"publisher-service-queue",
	//	true,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//if err != nil {
	//	log.Error(err.Error())
	//	log.Panic("Error while declaring queue")
	//	return
	//}
	//log.Debugf("Declared queue %s", q.Name)
	//err = ch.QueueBind(
	//	q.Name,
	//	"publisher-service",
	//	"publisher-service-exchange",
	//	false,
	//	nil,
	//)
	//if err != nil {
	//	log.Panic("Error while binding queue")
	//	return
	//}
	//log.Debugf("Binded to queue %s", q.Name)
	//messages, err = ch.Consume(
	//	q.Name,
	//	"pubisher-service",
	//	true,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)

	log.Debug("Connected to RabbitMQ")
}

func handler(channel <-chan amqp.Delivery) {
	for message := range channel {
		log.Debugf("Received message: %s", message.Body)
		webhook := core.Webhook{}
		err := json.Unmarshal(message.Body, &webhook)
		if err != nil {
			log.Errorf("Message %s couldn't be unmarshaled", message.Body)
			continue
		}
		log.Debugf("Received webhook message: %s", webhook.NotificationType)

		switch webhook.DeliveryPlatform {
		case core.Discord:
			go publishers.PublishToDiscord(&webhook)
			break
		case core.Slack:
			go publishers.PublishToSlack(&webhook)
			break

		}
	}
}
