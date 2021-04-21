package core

type Route struct {
	RoutingKey    string
	RouteFunction func(string)
}
