package main

import (
	"Helldivers2Tools/pkg/api/redisCache"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func startServer(app *fiber.App, bindAddress string) {
	err := app.Listen(bindAddress)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO make better config loader (like in bot)
// TODO add missing routes

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	var err error
	helldivers.Client, err = lib.New(
		os.Getenv("HDII__API__REMOTE_SCHEME"),
		os.Getenv("HDII__API__REMOTE_HOST"),
		utils.SafeAtoi(os.Getenv("HDII__API__REMOTE_PORT")),
	)
	if err != nil {
		log.Fatal(err)
	}

	redisCache.Context = redisCache.NewContext()
	redisCache.Client = redisCache.New(&redis.Options{
		Addr:       os.Getenv("HDII__API__REDIS_HOST") + ":" + os.Getenv("HDII__API__REDIS_PORT"),
		Password:   os.Getenv("HDII__API__REDIS_PASSWORD"),
		DB:         utils.SafeAtoi(os.Getenv("HDII__API__REDIS_DB")),
		ClientName: "HDII-API",
	})

	app := fiber.New()
	app.Use(logger.New())

	apiGroup := app.Group("/api")
	apiGroup.Get("/", getWarSeasons)

	warGroup := apiGroup.Group("/:war_id")
	warGroup.Get("/feed", getFeed)

	planetsGroup := warGroup.Group("/planets")
	planetsGroup.Get("/", getPlanets)
	planetsGroup.Get("/:planet_id", getPlanet)
	planetsGroup.Get("/:planet_id/status", getPlanetStatus)

	go startServer(app, os.Getenv("HDII__API__BIND_ADDRESS"))

	log.Println("Press Ctrl-c to shut down")
	<-c
	log.Println("Ctrl-c pressed, shutting down...")
	_ = app.Shutdown()
}
