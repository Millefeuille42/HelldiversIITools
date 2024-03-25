package commands

import (
	"Helldivers2Tools/pkg/bot/embeds"
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"github.com/bwmarrin/discordgo"
	"log"
)

func galaxyCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := interactionSendDefer(s, i)
	if err != nil {
		interactionSendError(s, i, "An error ocurred while sending message", 0)
		return
	}

	guild := models.GuildModel{}
	_, err = guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendFollowupError(s, i, "Error getting galaxy stats", 0)
		return
	}

	stats, err := helldivers.GoDiversClient.GetGalaxyStats()
	if err != nil {
		log.Println(err)
		interactionSendFollowupError(s, i, "Error getting galaxy stats", 0)
		return
	}

	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			embeds.BuildGalaxyEmbed(stats),
		},
	})

	if err != nil {
		log.Println(err)
	}
}
