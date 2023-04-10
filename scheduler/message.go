package scheduler

import (
	"time"

	"github.com/aldinofrizal/discord-bot/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func InitMessageScheduler(d *discordgo.Session) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("01:55").Do(func() {
		d.ChannelMessageSend(bot.TextGeneralId, "absen cuy")
	})

	s.Every(1).Day().At("05:00").Do(func() {
		d.ChannelMessageSend(bot.TextGeneralId, "istirahat siang cuy")
	})

	s.Every(1).Day().At("06:05").Do(func() {
		d.ChannelMessageSend(bot.TextGeneralId, "beres cuy istirahat siang nya")
	})

	s.Every(1).Day().At("10:00").Do(func() {
		d.ChannelMessageSend(bot.TextGeneralId, "balik cuy")
	})

	s.Every(1).Day().At("14:00").Do(func() {
		d.ChannelMessageSend(bot.TextGeneralId, "tidur cuy")
	})

	s.StartAsync()
}
