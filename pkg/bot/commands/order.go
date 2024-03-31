package commands

import (
	"Helldivers2Tools/pkg/bot/components"
	"Helldivers2Tools/pkg/bot/embeds"
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"github.com/bwmarrin/discordgo"
	"log"
)

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
		Components: components.BuildOrderComponents(order),
		Embeds: []*discordgo.MessageEmbed{
			embeds.BuildOrderEmbed(order),
		},
	})

	if err != nil {
		log.Println(err)
	}
}
