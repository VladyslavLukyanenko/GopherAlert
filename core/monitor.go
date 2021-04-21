package core

type Monitor struct {
	Channel            string                 `json:"channel"`
	Delay              int                    `json:"delay,omitempty"`
	WebhookURI         string                 `json:"webhook_uri,omitempty"`
	MonitoringPlatform MonitoringPlatformType `json:"monitoring_platform,omitempty"`
	DeliveryPlatform   DeliveryPlatformType   `json:"delivery_platform,omitempty"`
}
