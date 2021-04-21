package main

import (
	"github.com/VladyslavLukyanenko/MonitoringService/configs"
	"github.com/VladyslavLukyanenko/MonitoringService/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	configs.ReadConfig()
	initDatabase()
	initAMQP()
}
func initDatabase() {

}
func initAMQP() {
	InitAMQP()
	BindQueueToFunction("discord-add-task", routes.DiscordAddTask)
	BindQueueToFunction("discord-remove-task", routes.DiscordRemoveTask)
}
