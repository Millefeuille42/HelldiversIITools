package embeds

import (
	"Helldivers2Tools/pkg/bot/globals"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var NameMap = map[lib.Faction]string{
	lib.Automatons: "Automatons",
	lib.Terminids:  "Terminids",
	lib.Humans:     "Humans",
}

var ColorMap = map[lib.Faction]int{
	lib.Automatons: 6684929,
	lib.Terminids:  10521697,
	lib.Humans:     30646,
}

var ImageMap = map[lib.Faction]string{
	lib.Automatons: "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040096477085716.png",
	lib.Terminids:  "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040867453337720.png",
	lib.Humans:     "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040867881156618.png",
}

func BuildPlanetEmbed(planet lib.Planet, names []lib.PlanetName) *discordgo.MessageEmbed {
	ret := &discordgo.MessageEmbed{
		Type:  "rich",
		Color: ColorMap[planet.Owner],
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ImageMap[planet.Owner],
		},
		Title:  planet.Name,
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	if planet.InitialOwner != lib.Humans {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Menace",
			Value:  NameMap[planet.InitialOwner],
			Inline: false,
		})
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Under control of",
		Value:  NameMap[planet.InitialOwner],
		Inline: true,
	})

	if planet.LiberationPercent > 0 {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Liberation status",
			Value:  fmt.Sprintf("%f%% liberated", planet.LiberationPercent),
			Inline: false,
		})
	}

	if planet.Players > 0 || planet.Deaths > 0 {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Helldivers",
			Value:  globals.NumberPrinter.Sprintf("%d in mission, %d KIA", planet.Players, planet.Deaths),
			Inline: false,
		})
	}

	if planet.MissionsWon > 0 || planet.MissionsLost > 0 {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Missions",
			Value:  globals.NumberPrinter.Sprintf("%d won, %d lost", planet.MissionsWon, planet.MissionsLost),
			Inline: false,
		})
	}

	if planet.Waypoints != nil && len(planet.Waypoints) > 0 {
		var waypointNames []string
		for _, waypoint := range planet.Waypoints {
			for _, name := range names {
				if name.Index == waypoint {
					waypointNames = append(waypointNames, name.Name)
				}
			}
		}
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Related planets",
			Value:  strings.Join(waypointNames, ", "),
			Inline: false,
		})
	}

	return ret
}
