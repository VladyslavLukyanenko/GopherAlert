package routes

import (
	"context"
	"encoding/json"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/VladyslavLukyanenko/MonitoringService/database"
	"github.com/VladyslavLukyanenko/MonitoringService/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func MonitorAddTask(data string) {
	var monitor core.Monitor
	err := json.Unmarshal([]byte(data), &monitor)
	if err != nil {
		log.Error("Error while unmarshalling")
		return
	}
	for name, fn := range database.Monitors {
		if name == string(monitor.MonitoringPlatform) {
			_, err := database.Database.Collection("monitors").InsertOne(context.TODO(), models.Monitor{
				Channel:          monitor.Channel,
				MonitorDelay:     monitor.Delay,
				WebhookURI:       monitor.WebhookURI,
				DeliveryPlatform: monitor.DeliveryPlatform,
			})
			if err != nil {
				log.Errorf("Error while inserting monitor to database, %s", err)
			}
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			database.ActiveMonitors[monitor.Channel] = cancel
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
	res := database.Database.Collection("monitors").FindOneAndDelete(context.TODO(), bson.D{{"channel", monitor.Channel}})
	if res.Err() != nil {
		log.Errorf("Error while deleting entry in the database, %s", res.Err())
		return
	}
	cancelFn := database.ActiveMonitors[monitor.Channel]
	if cancelFn != nil {
		cancelFn()
	}
}
