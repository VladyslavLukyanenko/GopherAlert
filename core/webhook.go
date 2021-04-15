package core

type Webhook struct {
	JSON string `json:"json"`
	URI  string `json:"uri"`
	DeliveryPlatform DeliveryPlatformType `json:"deliveryPlatform"`
}
