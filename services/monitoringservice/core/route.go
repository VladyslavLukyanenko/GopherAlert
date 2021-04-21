package core

import "github.com/streadway/amqp"

type Route struct {
	RoutingKey    string
	RouteFunction func(amqp.Delivery, *interface{})
}

type Route_ interface {
	GetRoutingKey() string
	RouteFunction(amqp.Delivery, *interface{})
	GetDataStructure() *interface{}
}
