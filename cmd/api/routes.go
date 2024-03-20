package main

import (
	"Helldivers2Tools/pkg/api/redisCache"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/utils"
	"github.com/gofiber/fiber/v3"
	"log"
)

func getWarSeasons(c fiber.Ctx) error {
	warSeasons, err := redisCache.GetWarSeasons()
	if err != nil {
		warSeasons, err = helldivers.Client.GetWarSeasons()
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
		err := redisCache.SetWarSeasons(warSeasons)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
	}

	return c.JSON(warSeasons)
}

func getFeed(c fiber.Ctx) error {
	warSeasonId := c.Params("war_id")

	feed, err := redisCache.GetFeed(warSeasonId)
	if err != nil {
		feed, err = helldivers.Client.GetFeed(warSeasonId)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
		err := redisCache.SetFeed(warSeasonId, feed)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
	}

	return c.JSON(feed)
}

func getPlanets(c fiber.Ctx) error {
	warSeasonId := c.Params("war_id")

	planets, err := redisCache.GetPlanets(warSeasonId)
	if err != nil {
		planets, err = helldivers.Client.GetPlanets(warSeasonId)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
		err := redisCache.SetPlanets(warSeasonId, planets)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
	}

	return c.JSON(planets)
}

func getPlanet(c fiber.Ctx) error {
	warSeasonId := c.Params("war_id")
	planetId := utils.SafeAtoi(c.Params("planet_id"))

	planet, err := redisCache.GetPlanet(warSeasonId, planetId)
	if err != nil {
		planet, err = helldivers.Client.GetPlanet(warSeasonId, planetId)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
		err := redisCache.SetPlanet(warSeasonId, planet)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
	}

	return c.JSON(planet)
}

func getPlanetStatus(c fiber.Ctx) error {
	warSeasonId := c.Params("war_id")
	planetId := utils.SafeAtoi(c.Params("planet_id"))

	planetStatus, err := redisCache.GetPlanetStatus(warSeasonId, planetId)
	if err != nil {
		planetStatus, err = helldivers.Client.GetPlanetStatus(warSeasonId, planetId)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
		err := redisCache.SetPlanetStatus(warSeasonId, planetStatus)
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
	}

	return c.JSON(planetStatus)
}
