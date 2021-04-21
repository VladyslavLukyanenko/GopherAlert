package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Lukaesebrot/dgc"
	localcore "github.com/VladyslavLukyanenko/GopherAlert/DiscordService/core"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/database"
	"github.com/VladyslavLukyanenko/GopherAlert/DiscordService/models"
	"github.com/VladyslavLukyanenko/GopherAlert/contracts"
	"github.com/VladyslavLukyanenko/GopherAlert/core"
	"github.com/bwmarrin/discordgo"
	"github.com/lucasjones/reggen"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"strings"
)

func MonitorCommand() *dgc.Command {
	return &dgc.Command{
		Name:        "monitor",
		Description: "monitor",
		Usage:       "monitor",
		Example:     "monitor",
		IgnoreCase:  true,
		Handler:     MonitorHandler,
	}
}

const regexExpression = "^gopheralert-[a-zA-Z]{16}$"

func MonitorHandler(ctx *dgc.Ctx) {
	channel, _ := ctx.Session.Channel(ctx.Event.ChannelID)
	if channel.Type != discordgo.ChannelTypeGuildText {
		_ = ctx.RespondText("Only allowed in guild")
		return
	}
	regex, _ := regexp.Compile(regexExpression)
	gen, _ := reggen.NewGenerator(regexExpression)
	var subcommand = strings.ToLower(ctx.Arguments.Get(0).Raw())
	switch subcommand {
	case "list":
		{
			cursor, err := database.Database.Collection("monitors").Find(context.TODO(), bson.M{"channel_id": channel.ID})
			if err != nil {
				return
			}
			objects := cursor.RemainingBatchLength()
			var results []models.Monitor
			for cursor.Next(context.TODO()) {
				var monitor models.Monitor
				err := cursor.Decode(&monitor)
				if err != nil {
					log.Fatal(err)
				}
				results = append(results, monitor)
			}
			if err := cursor.Err(); err != nil {
				log.Fatal(err)
			}
			_ = cursor.Close(context.TODO())
			//todo: Improve Embed Design
			var embed discordgo.MessageEmbed
			embed.Description = fmt.Sprintf("Monitors (%d)", objects)
			for _, result := range results {
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   result.Id,
					Value:  result.WebhookURI,
					Inline: false,
				})
			}
			_ = ctx.RespondEmbed(&embed)

			break
		}
	case "add":
		{
			if len(ctx.Arguments.Get(1).Raw()) <= 0 {
				_ = ctx.RespondText("Argument [monitor String] not provided")
				return
			}
			channelId := ctx.Arguments.Get(1).Raw()
			result := database.Database.Collection("monitors").FindOne(context.TODO(), bson.M{"id": channelId, "channel_id": channel.ID})
			err := result.Err()
			if err == nil {
				_ = ctx.RespondText(fmt.Sprintf("Channel %s is already being monitored in this channel", channelId))
				return
			}
			webhooks, err := ctx.Session.ChannelWebhooks(channel.ID)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
			var webhook discordgo.Webhook
			var found = false
			for _, localWebhook := range webhooks {
				if regex.Match([]byte(localWebhook.Name)) {
					webhook = *localWebhook
					found = true
				}
			}
			if !found {
				lw, err := ctx.Session.WebhookCreate(channel.ID, gen.Generate(1), "https://image.emojipng.com/561/5513561.jpg") //todo: check why the image doesn't work
				if err != nil {
					log.Errorf(err.Error())
					_ = ctx.RespondText("Error while processing your command")
					return
				}
				webhook = *lw
			}
			webhookURI := fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", webhook.ID, webhook.Token)
			_, err = database.Database.Collection("monitors").InsertOne(context.TODO(), models.Monitor{
				GuildId:    ctx.Event.GuildID,
				ChannelId:  ctx.Event.ChannelID,
				WebhookURI: webhookURI,
				Id:         channelId,
			})
			if err != nil {
				log.Errorf(err.Error())
				_ = ctx.RespondText("Error while processing your command")
				return
			}
			monitor := core.Monitor{
				Channel:            channelId,
				Delay:              30,
				WebhookURI:         webhookURI,
				MonitoringPlatform: core.Youtube,
				DeliveryPlatform:   core.Discord,
			}
			data, _ := json.Marshal(monitor)
			contract := contracts.MonitoringContract{
				RoutingKey: "monitor-add-task",
				Data:       string(data),
			}
			localcore.SendMessage(contract)
			err = ctx.Session.MessageReactionAdd(ctx.Event.ChannelID, ctx.Event.Message.ID, "🆗")
			if err != nil {
				log.Errorf(err.Error())
				_ = ctx.RespondText("Error while reacting, but your command was processed")
				return
			}
		}
	case "remove":
		{
			if len(ctx.Arguments.Get(1).Raw()) <= 0 {
				_ = ctx.RespondText("Argument [monitor String] not provided")
				return
			}
			channelId := ctx.Arguments.Get(1).Raw()
			result := database.Database.Collection("monitors").FindOneAndDelete(context.TODO(), bson.M{"id": channelId, "channel_id": channel.ID})
			if result.Err() != nil {
				_ = ctx.RespondText(fmt.Sprintf("Error: Monitor %s not found", channelId))
				return
			}
			monitor := core.Monitor{
				Channel: channelId,
			}
			data, _ := json.Marshal(monitor)
			contract := contracts.MonitoringContract{
				RoutingKey: "monitor-remove-task",
				Data:       string(data),
			}
			localcore.SendMessage(contract)

			err := ctx.Session.MessageReactionAdd(ctx.Event.ChannelID, ctx.Event.Message.ID, "🆗")
			if err != nil {
				log.Errorf(err.Error())
				_ = ctx.RespondText("Error while reacting, but your command was processed")
				return
			}
		}
	default:
		{
			_ = ctx.RespondText("Subcommand not known, options [list, add, remove]")
			return
		}
	}
}
