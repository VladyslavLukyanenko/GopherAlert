package core

import "encoding/json"


func (r *Payload) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


type Payload struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type      string     `json:"type"`
	Text      *Text      `json:"text,omitempty"`
	Accessory *Accessory `json:"accessory,omitempty"`
	Elements  []Text     `json:"elements,omitempty"`
}

type Accessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji *bool  `json:"emoji,omitempty"`
}

