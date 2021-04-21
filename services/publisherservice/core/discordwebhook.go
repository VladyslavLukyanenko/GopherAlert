package core

import (
	"encoding/json"
	"time"
)

func (r *DiscordWebhook) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DiscordWebhook struct {
	Content   string  `json:"content"`
	Embeds    []Embed `json:"embeds"`
	Username  string  `json:"username"`
	AvatarUrl string  `json:"avatar_url"`
}

type Embed struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url"`
	Color       int64     `json:"color,omitempty"`
	Fields      []DiscordField   `json:"fields,omitempty"`
	Author      Author    `json:"author,omitempty"`
	Footer      Footer    `json:"footer,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Image       Image     `json:"image"`
	Thumbnail   Image     `json:"thumbnail,omitempty"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type DiscordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type Image struct {
	URL string `json:"url"`
}
