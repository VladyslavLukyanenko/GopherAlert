package core

type DeliveryPlatformType string

const(
	Twitch DeliveryPlatformType = DeliveryPlatformType(rune(iota))
	Youtube
)