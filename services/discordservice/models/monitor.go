package models

type Monitor struct {
	GuildId    string `bson:"guild_id"`
	ChannelId  string `bson:"channel_id"`
	WebhookURI string `bson:"webhook_id"`
	Id         string `bson:"id"`
}
