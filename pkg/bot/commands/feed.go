package commands

import (
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func feedCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := interactionSendDefer(s, i)
	if err != nil {
		return
	}

	guild := models.GuildModel{}
	_, err = guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting feed", 0)
		return
	}

	newsMessage, err := helldivers.GoDiversClient.GetNewsMessage()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting feed", 0)
		return
	}

	newsTitle := "New Message"
	message := newsMessage.Message
	newsSplit := strings.Split(newsMessage.Message, "\n")
	if len(newsSplit) > 1 {
		newsTitle = newsSplit[0]
		message = strings.Join(newsSplit[1:], "\n")
	}

	// TODO Add language choice
	_, err = s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
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
