package globals

import "github.com/bwmarrin/discordgo"

var (
	Bot *discordgo.Session
)

var DiscordCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "news",
		Description: "Get latest news message",
	},
	{
		Name:        "galaxy",
		Description: "Get stats about the galaxy",
	},
	{
		Name:        "order",
		Description: "Get latest major order",
	},
	{
		Name:        "planet",
		Description: "Get information about a planet",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         "name",
				Description:  "Name of the planet",
				Required:     true,
				Autocomplete: true,
			},
		},
	},
}
