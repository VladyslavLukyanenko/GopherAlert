package core

type MonitoringPlatformType string

const (
	Twitch MonitoringPlatformType = MonitoringPlatformType(rune(iota))
	Youtube
)
