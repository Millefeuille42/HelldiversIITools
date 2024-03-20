package commands

import (
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var colorMap = map[string]int{
	"Automaton": 6684929,
	"Terminids": 10521697,
	"Humans":    30646,
}

var imageMap = map[string]string{
	"Automaton": "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040096477085716.png",
	"Terminids": "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040867453337720.png",
	"Humans":    "https://cdn.discordapp.com/app-assets/1219964573231091713/1220040867881156618.png",
}

func buildPlanetEmbed(planet lib.Planet) *discordgo.MessageEmbed {
	status, err := helldivers.Client.GetCurrentWarPlanetStatus(planet.Index)
	if err != nil {
		log.Println(err)
		return nil
	}

	ret := &discordgo.MessageEmbed{
		Type: "rich",
		// TODO set color depending on controlling faction
		Color: colorMap[status.Owner],
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: imageMap[status.Owner],
		},
		Title:  planet.Name,
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Sector",
		Value:  planet.Sector,
		Inline: false,
	})

	if planet.InitialOwner != "Humans" {
		ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
			Name:   "Menace",
			Value:  planet.InitialOwner,
			Inline: false,
		})
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Under control of",
		Value:  status.Owner,
		Inline: true,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Helldivers",
		Value:  fmt.Sprintf("%d in mission", status.Players),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Liberation status",
		Value:  fmt.Sprintf("%f%% liberated", status.Liberation),
		Inline: true,
	})

	return ret
}

func buildPlanetsChoices(planets []lib.Planet) []*discordgo.ApplicationCommandOptionChoice {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0)
	for _, planet := range planets {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  planet.Name,
			Value: planet.Name,
		})
	}

	return choices
}

func planetCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["name"]; !ok {
		interactionSendError(s, i, "No planet provided", discordgo.MessageFlagsEphemeral)
		return
	}

	warSeason, err := helldivers.Client.GetWarSeasons()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting feed", 0)
		return
	}

	planets, err := helldivers.Client.GetPlanets(warSeason.Current)
	if i.Type != discordgo.InteractionApplicationCommand {
		if err != nil {
			log.Println(err)
			return
		}
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
		if planet.Name == optionMap["name"].StringValue() {
			selectedPlanet = planet
			break
		}
	}
	if selectedPlanet.Name == "" {
		interactionSendError(s, i, "Planet not found", discordgo.MessageFlagsEphemeral)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    "",
			Components: nil,
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
