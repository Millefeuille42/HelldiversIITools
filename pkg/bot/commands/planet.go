package commands

import (
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var colorMap = map[string]int{
	"Automatons": 6684929,
	"Terminids":  10521697,
	"Humans":     30646,
}

var imageMap = map[string]string{
	"Automatons": "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040096477085716.png",
	"Terminids":  "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040867453337720.png",
	"Humans":     "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040867881156618.png",
}

var planets []lib.PlanetName

func buildPlanetEmbed(planet lib.Planet) *discordgo.MessageEmbed {
	ret := &discordgo.MessageEmbed{
		Type:  "rich",
		Color: colorMap[planet.CurrentOwner],
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: imageMap[planet.CurrentOwner],
		},
		Title:  planet.PlanetName,
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	if planet.InitialOwner != "Humans" {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Menace",
			Value:  planet.InitialOwner,
			Inline: false,
		})
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Under control of",
		Value:  planet.CurrentOwner,
		Inline: true,
	})

	if planet.LibPercent > 0 {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Liberation status",
			Value:  fmt.Sprintf("%f%% liberated", planet.LibPercent),
			Inline: false,
		})

		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Liberated in",
			Value:  fmt.Sprintf("%f hours", planet.HoursComplete),
			Inline: true,
		})
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Helldivers",
		Value:  fmt.Sprintf("%d in mission, %d KIA", planet.Players, planet.Deaths),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Missions",
		Value:  fmt.Sprintf("%d won, %d lost", planet.MissionsWon, planet.MissionsLost),
		Inline: false,
	})

	if planet.WaypointNames != "" {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Related planets",
			Value:  planet.WaypointNames,
			Inline: false,
		})
	}

	return ret
}

func buildPlanetsChoices(planets []lib.PlanetName) []*discordgo.ApplicationCommandOptionChoice {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0)
	for _, planet := range planets {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  planet.Name,
			Value: planet.Name,
		})
	}

	return choices
}

func buildPlanetComponent(planet lib.Planet) []discordgo.MessageComponent {
	if planet.WaypointNames == "" {
		return nil
	}

	var buttons []discordgo.MessageComponent
	waypointsNames := strings.Split(planet.WaypointNames, ", ")
	waypointsIndices := strings.Split(planet.WaypointIndices, ", ")

	for index, waypointsIndex := range waypointsIndices {
		buttons = append(buttons, discordgo.Button{
			Label:    waypointsNames[index],
			Style:    0,
			Disabled: false,
			Emoji: discordgo.ComponentEmoji{
				Name: "🌎",
			},
			CustomID: fmt.Sprintf("planet_button-%s", waypointsIndex),
		})
	}

	return []discordgo.MessageComponent{discordgo.ActionsRow{Components: buttons}}
}

func planetCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["name"]; !ok {
		interactionSendError(s, i, "No planet provided", discordgo.MessageFlagsEphemeral)
		return
	}

	var err error
	if planets == nil {
		planets, err = helldivers.GoDiversClient.GetPlanetsName()
		if err != nil {
			log.Println(err)
			return
		}
	}
	if i.Type != discordgo.InteractionApplicationCommand {
		choices := buildPlanetsChoices(planets)
		handleAutocomplete(s, i, optionMap["name"].StringValue(), choices)
		return
	}

	if err != nil {
		interactionSendError(s, i, "Error fetching planets", discordgo.MessageFlagsEphemeral)
		return
	}

	selectedPlanet := lib.Planet{}
	for _, planet := range planets {
		if strings.ToLower(planet.Name) == strings.ToLower(optionMap["name"].StringValue()) {
			selectedPlanet, err = helldivers.GoDiversClient.GetPlanet(planet.Index)
			if err != nil {
				interactionSendError(s, i, "Planet not found", discordgo.MessageFlagsEphemeral)
				return
			}
			break
		}
	}

	planets = nil

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    "",
			Components: buildPlanetComponent(selectedPlanet),
			Embeds: []*discordgo.MessageEmbed{
				buildPlanetEmbed(selectedPlanet),
			},
			AllowedMentions: nil,
			Choices:         nil,
			CustomID:        "",
			Title:           "",
		},
	})

	if err != nil {
		log.Println(err)
	}
}

func planetComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := strings.Split(i.MessageComponentData().CustomID, "-")[1]
	planet, err := helldivers.GoDiversClient.GetPlanet(utils.SafeAtoi(id))
	if err != nil {
		interactionSendError(s, i, "Planet not found", discordgo.MessageFlagsEphemeral)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    "",
			Components: buildPlanetComponent(planet),
			Embeds: []*discordgo.MessageEmbed{
				buildPlanetEmbed(planet),
			},
			AllowedMentions: nil,
			Choices:         nil,
			CustomID:        "",
			Title:           "",
		},
	})

	if err != nil {
		log.Println(err)
	}
}
