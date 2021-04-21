package main

import (
	"github.com/VladyslavLukyanenko/GopherAlert/MonitoringService/configs"
	"github.com/VladyslavLukyanenko/GopherAlert/MonitoringService/database"
	"github.com/VladyslavLukyanenko/GopherAlert/MonitoringService/messagebroker"
	"github.com/VladyslavLukyanenko/GopherAlert/MonitoringService/monitors"
	"github.com/VladyslavLukyanenko/GopherAlert/MonitoringService/routes"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	log.SetLevel(log.DebugLevel)
	configs.ReadConfig()
	database.InitMongo()
	initAMQP()
	initMonitors()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
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