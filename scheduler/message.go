package scheduler

import (
	"time"

	"github.com/aldinofrizal/discord-bot/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func InitMessageScheduler(d *discordgo.Session) {
	s := gocron.NewScheduler(time.UTC)

	// s.Every(5).Second().Do(func() {
	// 	d.ChannelMessageSend(bot.TextGeneralId, "hello world")
	// })

	s.Every(1).Day().At("02:25").Do(func() {
		d.ChannelMessageSend(bot.TextGeneralId, "hello world 02:23UTC")
	})

	s.Every(1).Day().At("21:25").Do(func() {
		d.ChannelMessageSend(bot.TextGeneralId, "hello world 21:23WIB")
	})

	s.StartAsync()
}
