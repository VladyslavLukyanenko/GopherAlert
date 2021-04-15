package core

type DeliveryPlatformType string

const (
	Discord DeliveryPlatformType = DeliveryPlatformType(rune(iota))
	Telegram
	Slack
)