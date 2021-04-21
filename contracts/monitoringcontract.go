package contracts

import "github.com/VladyslavLukyanenko/GopherAlert/core"

type MonitoringContract struct {
	RoutingKey         string                      `json:"routing_key"`
	Data               string                      `json:"data"`
	MonitoringPlatform core.MonitoringPlatformType `json:"monitoring_platform"`
}
