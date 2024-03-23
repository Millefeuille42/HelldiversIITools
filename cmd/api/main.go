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
	helldivers.HDClient, err = lib.New(os.Getenv("HDII__API__HELLDIVERS_API"))
	if err != nil {
		log.Fatal(err)
	}

	helldivers.DiveHarderClient, err = lib.New(os.Getenv("HDII__API__DIVEHARDER_API"))
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
	defer redisCache.Client.Close()

	app := fiber.New()
	app.Use(logger.New())

	apiGroup := app.Group("/api")
	apiGroup.Get("/feed", feedHandler)
	apiGroup.Get("/order", orderHandler)
	apiGroup.Get("/galaxy", galaxyStatsHandler)
	apiGroup.Get("/campaigns", campaignsHandler)

	planetsGroup := apiGroup.Group("/planets")
	planetsGroup.Get("/", planetsNameHandler)
	planetsGroup.Get("/:planet_id", planetHandler)

	go startServer(app, os.Getenv("HDII__API__BIND_ADDRESS"))

	log.Println("Press Ctrl-c to shut down")
	<-c
	log.Println("Ctrl-c pressed, shutting down...")
	_ = app.Shutdown()
}
