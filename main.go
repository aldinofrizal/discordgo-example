package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aldinofrizal/discord-bot/bot"
	"github.com/aldinofrizal/discord-bot/scheduler"
)

func main() {
	bot.Init()

	err := bot.DiscordBot.Open()
	if err != nil {
		fmt.Println(err)
	}

	scheduler.InitMessageScheduler(bot.DiscordBot)

	defer bot.DiscordBot.Close()
	fmt.Println("bot is running, ctrl + c to terminate")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
