package main

import (
	"Helldivers2Tools/pkg/api/redisCache"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/utils"
	"github.com/gofiber/fiber/v3"
	"log"
	"time"
)

// feedHandler: returns: lib.NewsMessage
func feedHandler(c fiber.Ctx) error {
	feed, err := redisCache.GetNewsMessage()
	if err != nil {
		feed, err = getLastMessage()
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}
	}

	return c.JSON(feed)
}

// orderHandler: returns: lib.MajorOrder
func orderHandler(c fiber.Ctx) error {
	assignment, err := getAssignment()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	tasks, err := constructTasks(assignment)
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(lib.MajorOrder{
		Title:       assignment.Setting.OverrideTitle,
		Briefing:    assignment.Setting.OverrideBrief,
		Description: assignment.Setting.TaskDescription,
		Tasks:       tasks,
		Reward: lib.Reward{
			Type:   lib.RewardType(assignment.Setting.Reward.Type),
			Amount: assignment.Setting.Reward.Amount,
		},
		EndsAt: time.Now().Add(time.Second * time.Duration(assignment.ExpiresIn)),
	})
}

// planetsNameHandler: returns: lib.PlanetName
func planetsNameHandler(c fiber.Ctx) error {
	planets, err := getDiveHarderPlanets()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	var planetsName []lib.PlanetName
	for _, planet := range planets.Planets {
		planetsName = append(planetsName, lib.PlanetName{
			Index: planet.PlanetIndex,
			Name:  planet.PlanetName,
		})
	}

	return c.JSON(planetsName)
}

// planetHandler: returns: lib.Planet
func planetHandler(c fiber.Ctx) error {
	planetId := utils.SafeAtoi(c.Params("planet_id"))
	planet, err := constructPlanet(planetId)
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(planet)
}

// galaxyStatsHandler: returns: lib.GalaxyStats
func galaxyStatsHandler(c fiber.Ctx) error {
	stats, err := getDiveHarderGalaxyStats()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(stats)
}

// campaignsHandler: returns: []lib.Campaign
func campaignsHandler(c fiber.Ctx) error {
	status, err := getStatus()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(status.Campaigns)
}
