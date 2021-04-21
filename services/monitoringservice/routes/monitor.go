package routes

import (
	"context"
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	main "github.com/VladyslavLukyanenko/MonitoringService"
	"github.com/VladyslavLukyanenko/MonitoringService/models"
	log "github.com/sirupsen/logrus"
)

func MonitorAddTask(data string) {
	var monitor core.Monitor
	err := json.Unmarshal([]byte(data), &monitor)
	if err != nil {
		log.Error("Error while unmarshalling")
		return
	}
	for name, fn := range main.Monitors {
		if name == monitor.MonitoringPlatform {
			result, err := main.Database.Collection("monitors").InsertOne(context.TODO(), models.Monitor{
				Channel:          monitor.Channel,
				MonitorDelay:     monitor.Delay,
				WebhookURI:       monitor.WebhookURI,
				DeliveryPlatform: monitor.DeliveryPlatform,
			})
			if err != nil {
				log.Errorf("Error while inserting monitor to database, %s", result)
			}
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			main.ActiveMonitors[monitor.Channel] = cancel
			go fn(monitor, ctx)
		}
	}
}
func MonitorRemoveTask(data string) {
	var monitor core.Monitor
	err := json.Unmarshal([]byte(data), &monitor)
	if err != nil {
		log.Error("Error while unmarshalling")
		return
	}
	res := main.Database.Collection("monitors").FindOneAndDelete(context.TODO(), models.Monitor{Channel: monitor.Channel})
	if res.Err() != nil {
		log.Error("Error while deleting entry in the database")
		return
	}
	cancelFn := main.ActiveMonitors[monitor.Channel]
	if cancelFn != nil {
		cancelFn()
	}
}
