package publishers

import (
	"bytes"
	"fmt"
	localcore "github.com/VladyslavLukyanenko/GopherAlert/PublisherService/core"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func PublishToDiscord(webhook *core.Webhook) {
	log.Debugf("Received webhook into publisher: %s", webhook.URI)
	var discordWebhook = localcore.DiscordWebhook{
		Username:  "GopherAlert",
		AvatarUrl: "",
		Embeds:    []localcore.Embed{},
	}

	var embed = localcore.Embed{
		Title:       "",
		Description: "",
		Color:       0,
		Footer: localcore.Footer{
			Text:    "GopherAlert",
			IconURL: "",
		},
		Timestamp: time.Now(),
	}
	switch webhook.NotificationType {
	case core.Live:
		{
			embed.Author = localcore.Author{
				Name:    fmt.Sprintf("%s just went live on %s", webhook.Notification.ChannelName, webhook.MonitoringPlatformType),
				URL:     "",
				IconURL: webhook.Notification.Avatar,
			}
			embed.Title = webhook.Notification.Title
			embed.URL = webhook.Notification.URI
			if webhook.MonitoringPlatformType == core.Twitch {
				embed.Description = fmt.Sprintf("Playing: %s for %d viewers\n", webhook.Notification.Game, webhook.Notification.Viewers)
			}
			embed.Description += fmt.Sprintf("[Watch Stream](%s)", webhook.Notification.URI)
			break
		}
	case core.Video:
		{
			embed.Author = localcore.Author{
				Name:    fmt.Sprintf("%s just uploaded a video on %s", webhook.Notification.ChannelName, webhook.MonitoringPlatformType),
				URL:     "",
				IconURL: webhook.Notification.Avatar,
			}
			embed.Title = webhook.Notification.Title
			embed.URL = webhook.Notification.URI
			embed.Description += fmt.Sprintf("[Watch video](%s)", webhook.Notification.URI)
			break
		}
	}
	discordWebhook.Embeds = append(discordWebhook.Embeds, embed)

	marshal, err := discordWebhook.Marshal()
	if err != nil {
		log.Errorf("Error unmarshalling DiscordWebhook")
		return
	}
	res, err := http.Post(webhook.URI, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Errorf("Error when sending request publish@discord: %s", err.Error())
		return
	}
	if res.StatusCode != 204 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}

		log.Errorf("Status code not 204, [%s]", b)
		return
	}
	log.Debugf("Sent webhook publish@discord")
}
