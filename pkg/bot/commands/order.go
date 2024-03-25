package commands

import (
	"Helldivers2Tools/pkg/bot/embeds"
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func buildOrderComponents(order lib.MajorOrder) []discordgo.MessageComponent {
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

func orderCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := interactionSendDefer(s, i)
	if err != nil {
		interactionSendError(s, i, "An error ocurred while sending message", 0)
		return
	}

	guild := models.GuildModel{}
	_, err = guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendFollowupError(s, i, "Error getting order", 0)
		return
	}

	order, err := helldivers.GoDiversClient.GetMajorOrder()
	if err != nil {
		log.Println(err)
		interactionSendFollowupError(s, i, "Error getting order", 0)
		return
	}

	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Components: buildOrderComponents(order),
		Embeds: []*discordgo.MessageEmbed{
			embeds.BuildOrderEmbed(order),
		},
	})

	if err != nil {
		log.Println(err)
	}
}
