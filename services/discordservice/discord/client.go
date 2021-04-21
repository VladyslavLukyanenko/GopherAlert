package discord

import (
	"github.com/Lukaesebrot/dgc"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/commands"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/configs"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var Client *discordgo.Session

func InitDiscord() {
	DiscordClient, err := discordgo.New("Bot " + configs.AppConfig.Discord.Token)
	if err != nil {
		log.Errorf("Error creating Discord session, %s", err)
		return
	}
	DiscordClient.Identify.Intents = discordgo.IntentsDirectMessages | discordgo.IntentsGuildMessages

	err = DiscordClient.Open()
	if err != nil {
		log.Errorf("Error opening connection, %s", err)
		return
	}

	router := dgc.Create(&dgc.Router{
		Prefixes: []string{"!"},
	})

	router.RegisterCmd(commands.MonitorCommand())

	router.Initialize(DiscordClient)
}
