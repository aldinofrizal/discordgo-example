package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var (
	TextGeneralId string = os.Getenv("CHANNEL_GENERAL_ID")
	Token         string = os.Getenv("BOT_TOKEN")
)

var DiscordBot *discordgo.Session

func Init() {
	d, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err, "-failed to generate bot")
	}

	DiscordBot = d
}
