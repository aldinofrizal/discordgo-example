package scheduler

import (
	"time"

	"github.com/aldinofrizal/discord-bot/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func InitMessageScheduler(d *discordgo.Session) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("14:00").Do(func() { // example time
		d.ChannelMessageSend(bot.TextGeneralId, "example message send to discord")
	})

	s.StartAsync()
}
