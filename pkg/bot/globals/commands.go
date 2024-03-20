package globals

import "github.com/bwmarrin/discordgo"

var (
	Bot *discordgo.Session
)

var DiscordCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "feed",
		Description: "Get latest feed message",
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
