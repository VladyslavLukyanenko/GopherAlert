package monitors

import (
	"context"
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/VladyslavLukyanenko/MonitoringService/messagebroker"
	"github.com/streadway/amqp"
	"time"
)

func YoutubeMonitor(monitor core.Monitor, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			{
				webhook := core.Webhook{
					Notification: core.Notification{
						ChannelName: "togglebit",
						Title:       "Rust | Vim | Making a pixel editor | Command input, then some networking malarkey",
						Avatar:      "https://static-cdn.jtvnw.net/jtv_user_pictures/0bb9c502-ab5d-4440-9c9d-14e5260ebf86-profile_image-70x70.png",
						Game:        "Science & Technology",
						Viewers:     163,
						URI:         "https://www.twitch.tv/togglebit",
					},

					URI:                    monitor.WebhookURI,
					DeliveryPlatform:       monitor.DeliveryPlatform,
					NotificationType:       core.Live,
					MonitoringPlatformType: monitor.MonitoringPlatform,
				}
				payload, err := json.Marshal(&webhook)
				err = messagebroker.Channel.Publish("", "", false, false, amqp.Publishing{
					DeliveryMode: amqp.Transient,
					ContentType:  "application/json",
					Body:         payload,
				})
				if err != nil {
					return
				}
				time.Sleep(time.Duration(monitor.Delay * 30000))
			}
		}

	}
}
