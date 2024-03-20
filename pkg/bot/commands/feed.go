package commands

import (
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func feedCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting feed", 0)
		return
	}

	warSeason, err := helldivers.Client.GetWarSeasons()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting feed", 0)
		return
	}
	feed, err := helldivers.Client.GetFeed(warSeason.Current)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting feed", 0)
		return
	}

	if len(feed) <= 0 && feed[len(feed)-1].Message.En != "" {
		log.Println(err)
		interactionSendError(s, i, "No message in feed", 0)
		return
	}

	feedElement := feed[len(feed)-1]

	// TODO Add language choice
	message := fmt.Sprintf("%s", feedElement.Message.En)
	message = fmt.Sprintf("```%s```", message)

	interactionSendResponse(s, i, message, 0)
}
