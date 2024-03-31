package components

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func BuildPlanetComponent(planet lib.Planet) []discordgo.MessageComponent {
	if planet.Waypoints == nil || len(planet.Waypoints) <= 0 {
		return nil
	}

	var buttons []discordgo.MessageComponent

	for _, waypoint := range planet.Waypoints {
		buttons = append(buttons, discordgo.Button{
			Label:    waypoint.Name,
			Style:    0,
			Disabled: false,
			Emoji: discordgo.ComponentEmoji{
				Name: "🌎",
			},
			CustomID: fmt.Sprintf("planet_button-%d", waypoint.Index),
		})
	}

	return []discordgo.MessageComponent{discordgo.ActionsRow{Components: buttons}}
}
