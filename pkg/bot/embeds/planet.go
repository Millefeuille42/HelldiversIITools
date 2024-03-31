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
	lib.Humans:     "Super-Earth",
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

func buildPlanetEventEmbedFields(planet lib.Planet) (string, []*discordgo.MessageEmbedField) {
	var fields []*discordgo.MessageEmbedField
	emoji := ":shield:"
	action := "defended"
	percentName := "Defense"

	switch lib.EventType(planet.Event.EventType) {
	}

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("Must be %s from", action),
		Value:  NameMap[lib.Faction(planet.Event.Race)],
		Inline: false,
	})

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Under control of",
		Value:  NameMap[planet.Owner],
		Inline: false,
	})

	eventPercent := 100.0 - 100.0*float64(planet.Event.Health)/float64(planet.Event.MaxHealth)

	if eventPercent > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s status", percentName),
			Value:  fmt.Sprintf("%f%% %s", eventPercent, action),
			Inline: false,
		})
	}

	return emoji, fields
}

func buildPlanetNoEventEmbedFields(planet lib.Planet) []*discordgo.MessageEmbedField {
	var fields []*discordgo.MessageEmbedField

	if planet.InitialOwner != lib.Humans && planet.Owner != lib.Humans {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Menace",
			Value:  NameMap[planet.InitialOwner],
			Inline: false,
		})
	}

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Under control of",
		Value:  NameMap[planet.Owner],
		Inline: true,
	})

	if planet.LiberationPercent > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Liberation status",
			Value:  fmt.Sprintf("%f%% liberated", planet.LiberationPercent),
			Inline: false,
		})
	}

	return fields
}

func BuildPlanetEmbed(planet lib.Planet) *discordgo.MessageEmbed {
	ret := &discordgo.MessageEmbed{
		Type:  "rich",
		Color: ColorMap[planet.Owner],
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ImageMap[planet.Owner],
		},
		Title:  planet.Name,
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	if planet.Event != nil {
		emoji, fields := buildPlanetEventEmbedFields(planet)
		ret.Title = fmt.Sprintf("%s %s", emoji, ret.Title)
		ret.Fields = append(ret.Fields, fields...)
	} else {
		ret.Fields = append(ret.Fields, buildPlanetNoEventEmbedFields(planet)...)
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
			waypointNames = append(waypointNames, waypoint.Name)
		}
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Related planets",
			Value:  strings.Join(waypointNames, ", "),
			Inline: false,
		})
	}

	return ret
}
