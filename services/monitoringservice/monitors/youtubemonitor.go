package monitors

import (
	"context"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"time"
)

func YoutubeMonitor(monitor core.Monitor, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			{
				//logic
				Notify(core.Webhook{
					Notification: core.Notification{
						ChannelName: monitor.Channel,
						Title:       "testing",
						Avatar:      "https://static-cdn.jtvnw.net/jtv_user_pictures/0bb9c502-ab5d-4440-9c9d-14e5260ebf86-profile_image-70x70.png",
						Game:        "Science & Technology",
						Viewers:     100,
						URI:         "https://www.twitch.tv/togglebit",
					},
					URI:                    monitor.WebhookURI,
					DeliveryPlatform:       monitor.DeliveryPlatform,
					NotificationType:       core.Live,
					MonitoringPlatformType: monitor.MonitoringPlatform,
				})
				time.Sleep(time.Duration(monitor.Delay) * time.Second)
			}
		}

	}
}
