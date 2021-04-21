package core

type Monitor struct {
	Channel              string
	Delay                int
	WebhookURI           string
	MonitoringPlatform   MonitoringPlatformType
	DeliveryPlatform DeliveryPlatformType
}
