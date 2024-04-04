package embeds

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func SplitNewsMessage(news lib.NewsMessage) (title, message string) {
	title = "New Message"
	message = news.Message
	newsSplit := strings.Split(news.Message, "\n")
	if len(newsSplit) > 1 {
		title = newsSplit[0]
		message = strings.Join(newsSplit[1:], "\n")
	}

	return title, message
}

func BuildFeedEmbed(news lib.NewsMessage, color int) *discordgo.MessageEmbed {
	title, desc := SplitNewsMessage(news)
	return &discordgo.MessageEmbed{
		Type:        "rich",
		Color:       color,
		Title:       "Incoming message: " + utils.StripMarkup(title, "**"),
		Description: utils.StripMarkup(desc, "**"),
	}
}
