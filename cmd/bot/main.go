package main

import (
	"Helldivers2Tools/pkg/bot/commands"
	"Helldivers2Tools/pkg/bot/database"
	"Helldivers2Tools/pkg/bot/discord"
	"Helldivers2Tools/pkg/bot/globals"
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func setUpBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + globals.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discordBot.AddHandler(discord.BotConnected)
	discordBot.AddHandler(discord.GuildJoined)
	discordBot.AddHandler(discord.GuildLeft)

	discordBot.AddHandler(commands.CommandManager)

	err = discordBot.Open()
	if err != nil {
		log.Fatal(err)
	}

	discord.SetUpCloseHandler(discordBot)
	return discordBot
}

func populateDatabase() error {
	guild := models.GuildModel{}
	err := guild.CreateTable()
	if err != nil {
		return err
	}

	return nil
}

// Post V1

// TODO add language support

// Events

// TODO reverse helldivers API for event-based elements

// TODO write an updater program that will notify of new information
//  - new feed element
//  - change of planet status

func main() {
	var err error

	globals.SetGlobals()

	database.Database, err = database.NewDatabase(globals.DatabaseDriver, globals.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	err = populateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	globals.Bot = setUpBot()
	commands.PopulateCommandMap()

	helldivers.Client, err = lib.New(
		globals.ApiScheme,
		globals.ApiHost,
		globals.ApiPort,
	)
	if err != nil {
		log.Fatal(err)
	}

	routine()
}

func routine() {
	for {
		warSeasons, err := helldivers.Client.GetWarSeasons()
		if err != nil {
			log.Println(err)
		}
		if warSeasons.Current == "" {
			log.Println("Could not fetch war seasons")
			time.Sleep(time.Second * 5)
			continue
		}
		log.Println("current warId is:", warSeasons.Current)

		time.Sleep(time.Hour * 24)
	}
}
