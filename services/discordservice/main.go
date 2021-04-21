package main

import (
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/configs"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/database"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/discord"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/messagebroker"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	log.SetLevel(log.DebugLevel)
	configs.ReadConfig()
	messagebroker.InitAMQP()
	database.InitMongo()
	discord.InitDiscord()
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
	}()
	_ = discord.Client.Close()
}
