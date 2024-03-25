package main

import (
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/influx"
	"Helldivers2Tools/pkg/shared/redisEvent"
	"Helldivers2Tools/pkg/shared/utils"
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func checkMessage(previous lib.NewsMessage) lib.NewsMessage {
	message, err := helldivers.GoDiversClient.GetNewsMessage()
	if err != nil {
		log.Println(err)
		return previous
	}

	if message.Id == previous.Id {
		return previous
	}

	err = redisEvent.SendEvent(redisEvent.NewMessageEvent{NewsMessage: message})
	if err != nil {
		log.Println(err)
		return previous
	}

	return message
}

func checkOrder(previous lib.MajorOrder) lib.MajorOrder {
	order, err := helldivers.GoDiversClient.GetMajorOrder()
	if err != nil {
		log.Println(err)
		return previous
	}

	if order.Briefing == previous.Briefing {
		return previous
	}

	err = redisEvent.SendEvent(redisEvent.NewOrderEvent{MajorOrder: order})
	if err != nil {
		log.Println(err)
		return previous
	}

	return order
}

func checkPlanets(previous []lib.Planet) []lib.Planet {
	for index, prevPlanet := range previous {
		planet, err := helldivers.GoDiversClient.GetPlanet(prevPlanet.Index)

		if planet.Owner == 0 {
			continue
		}

		err = influx.WritePlanet(planet)
		if err != nil {
			log.Println(err)
		}

		if planet.Owner == prevPlanet.Owner {
			continue
		}

		err = redisEvent.SendEvent(redisEvent.NewPlanetEvent{Planet: planet})
		if err != nil {
			log.Println(err)
			return previous
		}
		previous[index] = planet
	}

	return previous
}

func checkGalaxy() {
	stats, err := helldivers.GoDiversClient.GetGalaxyStats()
	if err != nil {
		log.Println(err)
		return
	}

	err = influx.WriteGalaxy(stats)
	if err != nil {
		log.Println(err)
	}
}

func generatePlanets() ([]lib.Planet, error) {
	planets, err := helldivers.GoDiversClient.GetPlanetsName()
	if err != nil {
		return nil, err
	}

	var fullPlanets []lib.Planet
	for _, planet := range planets {
		fullPlanet, err := helldivers.GoDiversClient.GetPlanet(planet.Index)
		if err != nil {
			return nil, err
		}
		fullPlanets = append(fullPlanets, fullPlanet)
	}

	return fullPlanets, nil
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

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

	influx.Client, err = influxdb3.New(influxdb3.ClientConfig{
		Host:  os.Getenv("HDII__UPDATER__INFLUXDB__URL"),
		Token: os.Getenv("HDII__UPDATER__INFLUXDB__TOKEN"),
	})
	if err != nil {
		log.Println(err)
	}
	defer influx.Client.Close()

	newsTicker := time.Tick(time.Minute * 1)
	orderTicker := time.Tick(time.Minute * 1)
	planetTicker := time.Tick(time.Minute * 4)
	galaxyTicker := time.Tick(time.Minute * 1)

	log.Println("Getting message")
	message, err := helldivers.GoDiversClient.GetNewsMessage()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Getting order")
	order, err := helldivers.GoDiversClient.GetMajorOrder()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Getting planets")
	planets, err := generatePlanets()
	if err != nil {
		log.Fatal(err)
	}

	checkOrder(order)
	checkMessage(message)

	log.Println("Press Ctrl-c to shut down")
	for {
		select {
		case <-c:
			log.Println("Ctrl-c pressed, shutting down...")
			return
		case <-newsTicker:
			log.Println("checking message")
			message = checkMessage(message)
		case <-orderTicker:
			log.Println("checking order")
			order = checkOrder(order)
		case <-planetTicker:
			log.Println("checking planets")
			planets = checkPlanets(planets)
		case <-galaxyTicker:
			log.Println("checking galaxy")
			checkGalaxy()
		}
	}
}
