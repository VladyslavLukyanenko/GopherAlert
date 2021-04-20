package publishers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	localcore "github.com/VladyslavLukyanenko/GopherAlert/services/publisherservice/core"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func PublishToSlack(webhook *core.Webhook) {
	log.Debugf("Received webhook into publisher: %s", webhook.URI)
	payload := localcore.Payload{
		Blocks: []localcore.Block{},
	}
	payload.Blocks = append(payload.Blocks, localcore.Block{
		Type: "section",
		Text: &localcore.Text{
			Type: "plain_text",
			Text: func() string {
				if webhook.NotificationType == core.Live {
					return fmt.Sprintf("%s just went live on %s", webhook.Notification.ChannelName, webhook.MonitoringPlatformType)
				} else {
					return fmt.Sprintf("%s just uploaded a video on %s", webhook.Notification.ChannelName, webhook.MonitoringPlatformType)
				}
			}(),
			Emoji: func() *bool { b := true; return &b }(),
		},
		Accessory: nil,
		Elements:  nil,
	})
	payload.Blocks = append(payload.Blocks, localcore.Block{
		Type:      "divider",
		Text:      nil,
		Accessory: nil,
		Elements:  nil,
	})
	payload.Blocks = append(payload.Blocks, localcore.Block{
		Type: "section",
		Text: &localcore.Text{
			Type: "mrkdwn",
			Text: func() string {
				main := fmt.Sprintf("*<%s|%s>*\n", webhook.Notification.URI, webhook.Notification.Title)
				if webhook.NotificationType == core.Live {
					if webhook.MonitoringPlatformType == core.Twitch {
						main += fmt.Sprintf("Playing: %s for %d viewers\n", webhook.Notification.Game, webhook.Notification.Viewers)
					}
					main += fmt.Sprintf("*<%s|Watch Stream>*", webhook.Notification.URI)
				}else {
					main += fmt.Sprintf("*<%s|Watch Video>*", webhook.Notification.URI)
				}
				return main
			}(),
			Emoji: nil,
		},
		Accessory: &localcore.Accessory{
			Type:     "image",
			ImageURL: webhook.Notification.Avatar,
			AltText:  "avatar",
		},
		Elements: nil,
	})
	payload.Blocks = append(payload.Blocks, localcore.Block{
		Type:      "divider",
		Text:      nil,
		Accessory: nil,
		Elements:  nil,
	})
	payload.Blocks = append(payload.Blocks, localcore.Block{
		Type:      "context",
		Text:      nil,
		Accessory: nil,
		Elements: []localcore.Text{{
			Type:  "plain_text",
			Text:  "GopherAlert",
			Emoji: func() *bool { b := true; return &b }(),
		}},
	})
	marshal, err := JSONMarshal(payload)
	if err != nil {
		log.Errorf("Error unmarshalling SlackWebhook")
		return
	}

	res, err := http.Post(webhook.URI, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Errorf("Error when sending request publish@slack: %s", err.Error())
		return
	}
	if res.StatusCode != 200 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}
		log.Errorf("Status code not 200 [%s]", b)
		return
	}
	log.Debugf("Sent webhook publish@slack")
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
