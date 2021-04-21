package models

import (
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monitor struct {
	ID               primitive.ObjectID        `bson:"_id,omitempty"`
	Channel          string                    `bson:"channel"`
	MonitorDelay     int                       `bson:"monitor_delay"`
	WebhookURI       string                    `bson:"webhook_uri"`
	DeliveryPlatform core.DeliveryPlatformType `bson:"delivery_platform"`
}
