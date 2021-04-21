package contracts

type MonitoringContract struct {
	RoutingKey string `json:"routing_key"`
	Data       string `json:"data"`
}
