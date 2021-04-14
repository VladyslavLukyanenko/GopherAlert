package main

import (
	"github.com/VladyslavLukyanenko/twitch-discord-bot/configs"
	log "github.com/sirupsen/logrus"
)


func main() {
	configs.ReadConfig()
	setupLogger()
}

func setupLogger() {
	lvl, err := log.ParseLevel(configs.AppConfig.Logger.Level)
	if err != nil {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(lvl)
	}
}