package components

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func BuildOrderComponents(order lib.MajorOrder) []discordgo.MessageComponent {
	var planets []lib.PlanetName

	for _, task := range order.Tasks {
		if task.Progress <= 0 {
			planets = append(planets, task.Target)
		}
	}

	if len(planets) <= 0 {
		return nil
	}

	var buttons []discordgo.MessageComponent

	for _, planet := range planets {
		buttons = append(buttons, discordgo.Button{
			Label:    planet.Name,
			Style:    0,
			Disabled: false,
			Emoji: discordgo.ComponentEmoji{
				Name: "🌎",
			},
			CustomID: fmt.Sprintf("planet_button-%d", planet.Index),
		})
	}

	return []discordgo.MessageComponent{discordgo.ActionsRow{Components: buttons}}
}
