package main

import (
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/redisEvent"
	"Helldivers2Tools/pkg/shared/utils"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func forceEvent(eventType redisEvent.EventType) {
	var event redisEvent.Event
	switch eventType {
	case redisEvent.NewOrderEventType:
		order, err := helldivers.GoDiversClient.GetMajorOrder()
		if err != nil {
			fmt.Println(err)
			return
		}
		event = redisEvent.NewOrderEvent{MajorOrder: order}
	case redisEvent.NewMessageEventType:
		message, err := helldivers.GoDiversClient.GetNewsMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		event = redisEvent.NewMessageEvent{NewsMessage: message}
	default:
		fmt.Println("Invalid event_type")
		return
	}
	err := redisEvent.SendEvent(event)
	if err != nil {
		fmt.Println(err)
	}
}

func sendMessage(title, body string) {
	err := redisEvent.SendEvent(redisEvent.NewMessageEvent{NewsMessage: lib.NewsMessage{
		Id:        0,
		Published: 0,
		Type:      0,
		TagIds:    nil,
		Message:   fmt.Sprintf("BOT: %s\n%s", title, body),
	}})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	forceCmd := flag.NewFlagSet("force", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	forceEventType := forceCmd.String("event_type", "", "Specify the event type")

	sendTitle := sendCmd.String("title", "", "Specify the title")
	sendBody := sendCmd.String("body", "", "Specify the body")

	if len(os.Args) < 2 {
		fmt.Println("missing subcommand")
		os.Exit(1)
	}

	var err error
	helldivers.GoDiversClient, err = lib.New(os.Getenv("HDII__BOT__API"))
	if err != nil {
		log.Fatal(err)
	}

	redisEvent.Context = redisEvent.NewContext()
	redisEvent.Client = redisEvent.New(&redis.Options{
		Addr:       os.Getenv("HDII__API__REDIS_HOST") + ":" + os.Getenv("HDII__API__REDIS_PORT"),
		Password:   os.Getenv("HDII__API__REDIS_PASSWORD"),
		DB:         utils.SafeAtoi(os.Getenv("HDII__API__REDIS_DB")),
		ClientName: "HDII-UPDATER",
	})
	defer redisEvent.Client.Close()

	// Parse the command
	switch os.Args[1] {
	case "force":
		err = forceCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		if *forceEventType == "" {
			fmt.Println("Event type is required for force command")
			forceCmd.PrintDefaults()
			os.Exit(1)
		}
		forceEvent(redisEvent.EventType(*forceEventType))
	case "send":
		err = sendCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		if *sendTitle == "" || *sendBody == "" {
			fmt.Println("Title and body are required for send command")
			sendCmd.PrintDefaults()
			os.Exit(1)
		}
		sendMessage(*sendTitle, *sendBody)
	default:
		fmt.Println("force or send subcommand is required")
		os.Exit(1)
	}
}
