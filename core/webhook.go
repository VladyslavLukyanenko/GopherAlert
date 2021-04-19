package core

type Webhook struct {
	Notification           Notification           `json:"notification"`
	URI                    string                 `json:"uri"`
	DeliveryPlatform       DeliveryPlatformType   `json:"deliveryPlatform"`
	NotificationType       NotificationType       `json:"notification_type"`
	MonitoringPlatformType MonitoringPlatformType `json:"monitoring_platform_type"`
}
