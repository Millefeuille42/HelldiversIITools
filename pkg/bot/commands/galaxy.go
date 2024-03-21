package commands

import (
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func buildGalaxyEmbed(stats lib.GalaxyStats) *discordgo.MessageEmbed {
	ret := &discordgo.MessageEmbed{
		Type:   "rich",
		Title:  "Galaxy",
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Kills",
		Value:  fmt.Sprintf("%d terminids", stats.BugKills),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Value:  fmt.Sprintf("%d automatons", stats.AutomatonKills),
		Inline: true,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Value:  fmt.Sprintf("%d illuminates", stats.IlluminateKills),
		Inline: true,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Missions",
		Value:  fmt.Sprintf("%d won, %d lost", stats.MissionsWon, stats.MissionsLost),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Bullets",
		Value:  fmt.Sprintf("%d fired, %d hits", stats.BulletsFired, stats.BulletsHit),
		Inline: false,
	})

	ret.Fields = append(ret.Fields, &discordgo.MessageEmbedField{
		Name:   "Accidentals",
		Value:  fmt.Sprintf("%d accidental deaths", stats.Friendlies),
		Inline: false,
	})

	return ret
}

func galaxyCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(i.GuildID)
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting galaxy stats", 0)
		return
	}

	stats, err := helldivers.GoDiversClient.GetGalaxyStats()
	if err != nil {
		log.Println(err)
		interactionSendError(s, i, "Error getting galaxy stats", 0)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:    "",
			Components: nil,
			Embeds: []*discordgo.MessageEmbed{
				buildGalaxyEmbed(stats),
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
