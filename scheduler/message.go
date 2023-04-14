package scheduler

import (
	"time"

	"github.com/aldinofrizal/discord-bot/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func InitMessageScheduler(d *discordgo.Session) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("00:00").Do(func() { // example time hh:mm
		d.ChannelMessageSend(bot.TextGeneralId, "example message")
	})

	s.StartAsync()
}
