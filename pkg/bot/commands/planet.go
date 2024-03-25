package commands

import (
	"Helldivers2Tools/pkg/bot/embeds"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

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

func buildPlanetComponent(planet lib.Planet, names []lib.PlanetName) []discordgo.MessageComponent {
	if planet.Waypoints == nil || len(planet.Waypoints) <= 0 {
		return nil
	}

	var buttons []discordgo.MessageComponent

	for _, waypoint := range planet.Waypoints {
		waypointName := ""
		for _, name := range names {
			if name.Index == waypoint {
				waypointName = name.Name
			}
		}
		buttons = append(buttons, discordgo.Button{
			Label:    waypointName,
			Style:    0,
			Disabled: false,
			Emoji: discordgo.ComponentEmoji{
				Name: "🌎",
			},
			CustomID: fmt.Sprintf("planet_button-%d", waypoint),
		})
	}

	return []discordgo.MessageComponent{discordgo.ActionsRow{Components: buttons}}
}

func planetCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionMap := parseOptions(i.ApplicationCommandData().Options)

	if _, ok := optionMap["name"]; !ok {
		interactionSendFollowupError(s, i, "No planet provided", discordgo.MessageFlagsEphemeral)
		return
	}

	planets, err := helldivers.GoDiversClient.GetPlanetsName()
	if err != nil {
		interactionSendFollowupError(s, i, "Error fetching planets", discordgo.MessageFlagsEphemeral)
		return
	}
	if i.Type != discordgo.InteractionApplicationCommand {
		choices := buildPlanetsChoices(planets)
		handleAutocomplete(s, i, optionMap["name"].StringValue(), choices)
		return
	}

	selectedPlanet := lib.Planet{}
	for _, planet := range planets {
		if strings.ToLower(planet.Name) == strings.ToLower(optionMap["name"].StringValue()) {
			selectedPlanet, err = helldivers.GoDiversClient.GetPlanet(planet.Index)
			if err != nil {
				interactionSendFollowupError(s, i, "Planet not found", discordgo.MessageFlagsEphemeral)
				return
			}
			break
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    "",
			Components: buildPlanetComponent(selectedPlanet, planets),
			Embeds: []*discordgo.MessageEmbed{
				embeds.BuildPlanetEmbed(selectedPlanet, planets),
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
	err := interactionSendDefer(s, i)
	if err != nil {
		interactionSendError(s, i, "An error ocurred while sending message", 0)
		return
	}

	id := strings.Split(i.MessageComponentData().CustomID, "-")[1]
	planet, err := helldivers.GoDiversClient.GetPlanet(utils.SafeAtoi(id))
	if err != nil {
		interactionSendFollowupError(s, i, "Planet not found", discordgo.MessageFlagsEphemeral)
		return
	}

	planets, err := helldivers.GoDiversClient.GetPlanetsName()
	if err != nil {
		interactionSendFollowupError(s, i, "Error fetching planets", discordgo.MessageFlagsEphemeral)
		return
	}

	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Components: buildPlanetComponent(planet, planets),
		Embeds: []*discordgo.MessageEmbed{
			embeds.BuildPlanetEmbed(planet, planets),
		},
	})

	if err != nil {
		log.Println(err)
	}
}
