package main

import (
	"github.com/VladyslavLukyanenko/MonitoringService/configs"
	"github.com/VladyslavLukyanenko/MonitoringService/database"
	"github.com/VladyslavLukyanenko/MonitoringService/messagebroker"
	"github.com/VladyslavLukyanenko/MonitoringService/monitors"
	"github.com/VladyslavLukyanenko/MonitoringService/routes"
	log "github.com/sirupsen/logrus"
)


func main() {
	log.SetLevel(log.DebugLevel)
	configs.ReadConfig()
	database.InitMongo()
	initAMQP()
	initMonitors()
	select {}
}

func initAMQP() {
	messagebroker.InitAMQP()
	messagebroker.BindQueueToFunction("monitor-add-task", routes.MonitorAddTask)
	messagebroker.BindQueueToFunction("monitor-remove-task", routes.MonitorRemoveTask)
}

func initMonitors() {
	database.Monitors["youtube"] = monitors.YoutubeMonitor
	database.Monitors["twitch"] = monitors.TwitchMonitor
}