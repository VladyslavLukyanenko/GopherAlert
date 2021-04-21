package main

import (
	"context"
	core2 "github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/VladyslavLukyanenko/MonitoringService/configs"
	"github.com/VladyslavLukyanenko/MonitoringService/monitors"
	"github.com/VladyslavLukyanenko/MonitoringService/routes"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Database *mongo.Database
var Monitors map[string]func(monitor core2.Monitor, ctx context.Context)
var ActiveMonitors map[string]context.CancelFunc

func main() {
	log.SetLevel(log.DebugLevel)
	configs.ReadConfig()
	initMongo()
	initAMQP()
	initMonitors()
}
func initMongo() {
	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.AppConfig.MongoDB.URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	Database = client.Database("monitorService")
	log.Debug("Connected to MongoDB")
}

func initAMQP() {
	InitAMQP()
	BindQueueToFunction("monitor-add-task", routes.MonitorAddTask)
	BindQueueToFunction("monitor-remove-task", routes.MonitorRemoveTask)
}

func initMonitors() {
	Monitors["youtube"] = monitors.YoutubeMonitor
	Monitors["twitch"] = monitors.TwitchMonitor
}