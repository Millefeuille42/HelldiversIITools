package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var commandMap = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
var componentMap = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

func PopulateCommandMap() {
	commandMap["feed"] = feedCommandHandler
	commandMap["planet"] = planetCommandHandler
}

func CommandManager(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		log.Printf("Received command: %v", i.ApplicationCommandData().Name)
	}

	if i.Type == discordgo.InteractionMessageComponent {
		log.Printf("Received component: %v", i.MessageComponentData().CustomID)
		id := strings.Split(i.MessageComponentData().CustomID, "-")[0]
		if handler, ok := componentMap[id]; ok {
			handler(s, i)
		}
		return
	}

	if handler, ok := commandMap[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}
