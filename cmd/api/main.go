package main

import (
	"Helldivers2Tools/pkg/api/redisCache"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/utils"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
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

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Get("/health", healthcheck.NewHealthChecker())

	app.Use(compress.New())
	apiGroup := app.Group("/api")
	apiGroup.Get(lib.GoDiversFeedRoute, feedHandler)
	apiGroup.Get(lib.GoDiversOrderRoute, orderHandler)
	apiGroup.Get(lib.GoDiversGalaxyStatsRoute, galaxyStatsHandler)
	apiGroup.Get(lib.GoDiversCampaignsRoute, campaignsHandler)
	apiGroup.Get(lib.GoDiversHomeWorldsRoute, homeWorldsHandler)
	apiGroup.Get(lib.GoDiversJointOperationsRoute, jointOperationsHandler)

	planetsGroup := apiGroup.Group("/planets")
	planetsGroup.Get("/names", planetsNameHandler)
	planetsGroup.Get("/infos", planetInfosHandler)
	planetsGroup.Get("/status", planetStatusHandler)
	planetsGroup.Get("/attacks", planetAttacksHandler)
	planetsGroup.Get("/events", planetEventsHandler)
	planetsGroup.Get("/planet/:planet_id", planetHandler)

	go startServer(app, os.Getenv("HDII__API__BIND_ADDRESS"))

	log.Println("Press Ctrl-c to shut down")
	<-c
	log.Println("Ctrl-c pressed, shutting down...")
	_ = app.Shutdown()
}
