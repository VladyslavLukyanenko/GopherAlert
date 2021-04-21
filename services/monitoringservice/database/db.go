package database

import (
	"context"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/VladyslavLukyanenko/MonitoringService/configs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Database *mongo.Database
var Monitors map[string]func(monitor core.Monitor, ctx context.Context)
var ActiveMonitors map[string]context.CancelFunc

func InitMongo() {
	Monitors = make(map[string]func(monitor core.Monitor, ctx context.Context))
	ActiveMonitors = make(map[string]context.CancelFunc)
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