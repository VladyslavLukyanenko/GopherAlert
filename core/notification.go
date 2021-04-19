package core

type Notification struct {
	ChannelName string `json:"channel_name"`
	Title       string `json:"title"`
	Avatar      string `json:"avatar"`
	Game        string `json:"game"`
	Viewers     int    `json:"viewers"`
	URI         string `json:"uri"`
}
