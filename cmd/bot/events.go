package main

import (
	"Helldivers2Tools/pkg/bot/components"
	"Helldivers2Tools/pkg/bot/embeds"
	"Helldivers2Tools/pkg/bot/globals"
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/redisEvent"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var eventMap = map[redisEvent.EventType]func([]byte) error{
	redisEvent.NewMessageEventType:      handleNewMessage,
	redisEvent.NewOrderEventType:        handleNewOrder,
	redisEvent.PlanetLiberatedEventType: handlePlanetLiberated,
	redisEvent.PlanetLostEventType:      handlePlanetLost,
}

func streamComplex(complex *discordgo.MessageSend) error {
	guilds, err := models.GetGuilds()
	if err != nil {
		return err
	}
	for _, guild := range guilds {
		if guild.AnnouncementChannel == "" {
			continue
		}
		_, err = globals.Bot.ChannelMessageSendComplex(guild.AnnouncementChannel, complex)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func streamEmbed(embed *discordgo.MessageEmbed) error {
	return streamComplex(&discordgo.MessageSend{Embeds: []*discordgo.MessageEmbed{embed}})
}

func handleNewMessage(event []byte) error {
	var newMessage lib.NewsMessage
	err := json.Unmarshal(event, &newMessage)
	if err != nil {
		return errors.New("invalid data type")
	}

	newsTitle, _ := embeds.SplitNewsMessage(newMessage)
	if !strings.HasPrefix(newsTitle, "BOT: ") {
		err = setBotStatus(newsTitle)
		if err != nil {
			log.Println(err)
		}
	}

	return streamEmbed(embeds.BuildFeedEmbed(newMessage, 15616811))
}

func handleNewOrder(event []byte) error {
	var newOrder lib.MajorOrder
	err := json.Unmarshal(event, &newOrder)
	if err != nil {
		return errors.New("invalid data type")
	}

	embed := embeds.BuildOrderEmbed(newOrder)
	embed.Title = "NEW MAJOR ORDER"
	embed.Color = 15616811
	return streamComplex(&discordgo.MessageSend{
		Components: components.BuildOrderComponents(newOrder),
		Embeds:     []*discordgo.MessageEmbed{embeds.BuildOrderEmbed(newOrder)},
	})
}

func handlePlanetLiberated(event []byte) error {
	var planet lib.Planet
	err := json.Unmarshal(event, &planet)
	if err != nil {
		return errors.New("invalid data type")
	}

	planet.LiberationPercent = 0
	embed := embeds.BuildPlanetEmbed(planet)
	embed.Title = fmt.Sprintf("✅ %s liberated", embed.Title)
	return streamComplex(&discordgo.MessageSend{
		Components: components.BuildPlanetComponent(planet),
		Embeds:     []*discordgo.MessageEmbed{embed},
	})
}

func handlePlanetLost(event []byte) error {
	var planet lib.Planet
	err := json.Unmarshal(event, &planet)
	if err != nil {
		return errors.New("invalid data type")
	}

	embed := embeds.BuildPlanetEmbed(planet)
	embed.Title = fmt.Sprintf("❌ %s lost to the %s", embed.Title, embeds.NameMap[planet.Owner])
	return streamComplex(&discordgo.MessageSend{
		Components: components.BuildPlanetComponent(planet),
		Embeds:     []*discordgo.MessageEmbed{embed},
	})
}
