package main

import (
	"Helldivers2Tools/pkg/bot/commands"
	"Helldivers2Tools/pkg/bot/globals"
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/redisEvent"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var eventMap = map[redisEvent.EventType]func([]byte) error{
	redisEvent.NewMessageEventType:      handleNewMessage,
	redisEvent.NewOrderEventType:        handleNewOrder,
	redisEvent.PlanetLiberatedEventType: handlePlanetLiberated,
	redisEvent.PlanetLostEventType:      handlePlanetLost,
}

func streamEmbed(embed *discordgo.MessageEmbed) error {
	guilds, err := models.GetGuilds()
	if err != nil {
		return err
	}
	for _, guild := range guilds {
		if guild.AnnouncementChannel == "" {
			continue
		}
		_, err = globals.Bot.ChannelMessageSendEmbed(guild.AnnouncementChannel, embed)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func handleNewMessage(event []byte) error {
	var newMessage lib.NewsMessage
	err := json.Unmarshal(event, &newMessage)
	if err != nil {
		return errors.New("invalid data type")
	}

	newsTitle, message := lib.SplitNewsMessage(newMessage)
	err = setBotStatus(newsTitle)
	if err != nil {
		log.Println(err)
	}

	return streamEmbed(&discordgo.MessageEmbed{
		Type:        "rich",
		Color:       15616811,
		Title:       "Incoming message: " + newsTitle,
		Description: message,
	})
}

func handleNewOrder(event []byte) error {
	var newOrder lib.MajorOrder
	err := json.Unmarshal(event, &newOrder)
	if err != nil {
		return errors.New("invalid data type")
	}

	embed := commands.BuildOrderEmbed(newOrder)
	embed.Title = "NEW " + embed.Title
	embed.Color = 15616811
	return streamEmbed(embed)
}

func handlePlanetLiberated(event []byte) error {
	var planet lib.Planet
	err := json.Unmarshal(event, &planet)
	if err != nil {
		return errors.New("invalid data type")
	}

	embed := commands.BuildPlanetEmbed(planet)
	embed.Title = fmt.Sprintf("✅ %s liberated", embed.Title)
	return streamEmbed(embed)
}

func handlePlanetLost(event []byte) error {
	var planet lib.Planet
	err := json.Unmarshal(event, &planet)
	if err != nil {
		return errors.New("invalid data type")
	}

	embed := commands.BuildPlanetEmbed(planet)
	embed.Title = fmt.Sprintf("❌ %s lost to the %s", embed.Title, planet.CurrentOwner)
	return streamEmbed(embed)
}
