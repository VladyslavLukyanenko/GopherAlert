package core

import "github.com/streadway/amqp"


type Route struct {
	RoutingKey string
	RouteFunction func(amqp.Delivery)
	DataStructure struct{}
}
