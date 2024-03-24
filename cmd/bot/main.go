package main

import (
	"Helldivers2Tools/pkg/bot/commands"
	"Helldivers2Tools/pkg/bot/database"
	"Helldivers2Tools/pkg/bot/discord"
	"Helldivers2Tools/pkg/bot/globals"
	"Helldivers2Tools/pkg/bot/models"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/redisEvent"
	"Helldivers2Tools/pkg/shared/utils"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func setBotStatus(status string) error {
	return globals.Bot.UpdateListeningStatus(status)
}

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

	helldivers.GoDiversClient, err = lib.New(globals.ApiUrl)
	if err != nil {
		log.Fatal(err)
	}

	newsMessage, err := helldivers.GoDiversClient.GetNewsMessage()
	newsTitle, _ := lib.SplitNewsMessage(newsMessage)
	err = setBotStatus(newsTitle)
	if err != nil {
		log.Println(err)
	}

	redisEvent.Context = redisEvent.NewContext()
	redisEvent.Client = redisEvent.New(&redis.Options{
		Addr:       os.Getenv("HDII__API__REDIS_HOST") + ":" + os.Getenv("HDII__API__REDIS_PORT"),
		Password:   os.Getenv("HDII__API__REDIS_PASSWORD"),
		DB:         utils.SafeAtoi(os.Getenv("HDII__API__REDIS_DB")),
		ClientName: "HDII-BOT",
	})
	defer redisEvent.Client.Close()

	routine()
}

func routine() {
	sub := redisEvent.Client.Subscribe(redisEvent.Context, "events")
	defer sub.Close()
	for {
		data, err := sub.Receive(redisEvent.Context)
		if err != nil || data == nil {
			log.Println(err)
			continue
		}

		switch msg := data.(type) {
		case *redis.Subscription:
			fmt.Println("subscribed to", msg.Channel)
		case *redis.Message:
			fmt.Println("received", msg.Payload, "from", msg.Channel)
			var event redisEvent.Generic
			err = json.Unmarshal([]byte(msg.Payload), &event)
			if err != nil {
				log.Println(err)
				continue
			}

			err = eventMap[event.Type]([]byte(event.Data))
			if err != nil {
				log.Println(err)
			}
		}
	}
}
