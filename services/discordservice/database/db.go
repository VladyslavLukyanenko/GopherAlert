package database

import (
	"context"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/configs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Database *mongo.Database

func InitMongo() {
	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.AppConfig.MongoDB.URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	Database = client.Database("discordService")
	log.Debug("Connected to MongoDB")
}