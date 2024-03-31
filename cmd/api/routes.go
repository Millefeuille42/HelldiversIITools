package main

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"Helldivers2Tools/pkg/shared/utils"
	"github.com/gofiber/fiber/v3"
	"log"
	"time"
)

// feedHandler: returns: lib.NewsMessage
func feedHandler(c fiber.Ctx) error {
	message, err := getLastMessage()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(message)
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
	planetNames, err := getPlanetNames()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(planetNames)
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
	stats, err := getGalaxyStats()
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

// planetEventsHandler: returns: []lib.PlanetEvent
func planetEventsHandler(c fiber.Ctx) error {
	status, err := getStatus()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(status.PlanetEvents)
}

// planetStatusHandler: returns: []lib.PlanetStatus
func planetStatusHandler(c fiber.Ctx) error {
	status, err := getStatus()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(status.PlanetStatus)
}

// planetAttacksHandler: returns: []lib.PlanetAttack
func planetAttacksHandler(c fiber.Ctx) error {
	status, err := getStatus()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(status.PlanetAttacks)
}

// jointOperationsHandler: returns: []lib.JointOperation
func jointOperationsHandler(c fiber.Ctx) error {
	status, err := getStatus()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(status.JointOperations)
}

// planetInfosHandler: returns: []lib.PlanetInfo
func planetInfosHandler(c fiber.Ctx) error {
	warInfo, err := getWarInfo()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(warInfo.PlanetInfos)
}

// homeWorldsHandler: returns: []lib.HomeWorld
func homeWorldsHandler(c fiber.Ctx) error {
	warInfo, err := getWarInfo()
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(warInfo.HomeWorlds)
}
