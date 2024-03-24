package commands

import (
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func feedCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := interactionSendDefer(s, i)
	if err != nil {
		interactionSendError(s, i, "An error ocurred while sending message", 0)
		return
	}

	guild := models.GuildModel{}
	_, err = guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendFollowupError(s, i, "Error getting feed", 0)
		return
	}

	newsMessage, err := helldivers.GoDiversClient.GetNewsMessage()
	if err != nil {
		log.Println(err)
		interactionSendFollowupError(s, i, "Error getting feed", 0)
		return
	}

	newsTitle, message := lib.SplitNewsMessage(newsMessage)
	newsSplit := strings.Split(newsMessage.Message, "\n")
	if len(newsSplit) > 1 {
		newsTitle = newsSplit[0]
		message = strings.Join(newsSplit[1:], "\n")
	}

	// TODO Add language choice
	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			{
				Type:        "rich",
				Title:       newsTitle,
				Description: message,
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}
