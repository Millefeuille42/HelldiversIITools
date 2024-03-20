package discord

import (
	"Helldivers2Tools/pkg/bot/models"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func BotConnected(s *discordgo.Session, r *discordgo.Ready) {
	r = nil
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

func GuildJoined(s *discordgo.Session, g *discordgo.GuildCreate) {
	log.Printf("Joined guild: %v", g.Name)

	guild := models.GuildModel{
		Name:    g.Name,
		GuildId: g.ID,
	}
	for _, channel := range g.Channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			permissions, err := s.UserChannelPermissions(s.State.User.ID, channel.ID)
			if err != nil {
				log.Println(err)
				continue
			}

			if (permissions & discordgo.PermissionSendMessages) != 0 {
				guild.AnnouncementChannel = channel.ID
				break
			}
		}
	}

	err := guild.CreateGuild()
	if err != nil && !strings.Contains(err.Error(), "UNIQUE constraint failed") {
		log.Println(err)
	}

	RegisterCommands(s, g)
}

func GuildLeft(s *discordgo.Session, g *discordgo.GuildDelete) {
	log.Printf("Left guild: %v", g.BeforeDelete.Name)
	DeleteCommands(s, g.ID)
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(g.ID)
	if err != nil {
		log.Println(err)
		return
	}
	err = guild.DeleteGuild()
	if err != nil {
		log.Println(err)
		return
	}
}
