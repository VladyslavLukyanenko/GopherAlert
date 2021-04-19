package main

import (
	"github.com/VladyslavLukyanenko/GopherAlert/configs"
	"github.com/VladyslavLukyanenko/GopherAlert/twitchAPI"
	log "github.com/sirupsen/logrus"
)

func main() {
	configs.ReadConfig()
	setupLogger()
	twitchAPI.GetTwitchChannelStatus("ArQuel")
}

func setupLogger() {
	lvl, err := log.ParseLevel(configs.AppConfig.Logger.Level)
	if err != nil {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(lvl)
	}
}
