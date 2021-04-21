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
				//Notify(webhook)
				time.Sleep(time.Duration(monitor.Delay) * time.Second)
			}
		}

	}
}
