package main

import (
	"Helldivers2Tools/pkg/api/redisCache"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
	"os"
)

func getPlanetNames() ([]lib.PlanetName, error) {
	planetNames, err := redisCache.GetPlanetNames()
	if err != nil {
		file, err := os.ReadFile("./data/static/planet_names.json")
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(file, &planetNames)
		if err != nil {
			return nil, err
		}
		err = redisCache.SetPlanetNames(planetNames)
	}
	return planetNames, err
}

func getLastMessage() (lib.NewsMessage, error) {
	message, err := redisCache.GetLatestNewsMessage()
	if err != nil {
		message, err = helldivers.HDClient.GetHelldiversLatestNewsFeed(warId)
		if err != nil {
			return lib.NewsMessage{}, err
		}
		err = redisCache.SetLatestNewsMessage(message)
		if err != nil {
			return lib.NewsMessage{}, err
		}
		err = redisCache.SetNewsMessage(message)
	}
	return message, err
}

func getDiveHarderPlanetStats() ([]lib.PlanetStats, error) {
	planetStats, err := redisCache.GetPlanetsStats()
	if err != nil {
		err = nil
		stats, err := helldivers.DiveHarderClient.GetDiveHarderStats()
		if err != nil {
			return nil, err
		}
		planetStats = stats.PlanetsStats
		err = redisCache.SetPlanetsStats(planetStats)
		if err != nil {
			return planetStats, err
		}
		err = redisCache.SetGalaxyStats(stats.GalaxyStats)
	}
	return planetStats, err
}

func getDiveHarderGalaxyStats() (lib.GalaxyStats, error) {
	galaxyStats, err := redisCache.GetGalaxyStats()
	if err != nil {
		err = nil
		stats, err := helldivers.DiveHarderClient.GetDiveHarderStats()
		if err != nil {
			return lib.GalaxyStats{}, err
		}
		galaxyStats = stats.GalaxyStats
		err = redisCache.SetGalaxyStats(galaxyStats)
		if err != nil {
			return galaxyStats, err
		}
		err = redisCache.SetPlanetsStats(stats.PlanetsStats)
	}
	return galaxyStats, err
}

func getAssignment() (lib.Assignment, error) {
	assignment, err := redisCache.GetAssignment()
	if err != nil {
		assignment, err = helldivers.HDClient.GetHelldiversAssignment(warId)
		if err != nil {
			return lib.Assignment{}, err
		}
		err = redisCache.SetAssignment(assignment)
	}
	return assignment, err
}

func getStatus() (lib.Status, error) {
	status, err := redisCache.GetStatus()
	if err != nil {
		status, err = helldivers.HDClient.GetHelldiversStatus(warId)
		if err != nil {
			return lib.Status{}, err
		}
		err = redisCache.SetStatus(status)
	}
	return status, err
}

func getWarInfo() (lib.WarInfo, error) {
	warInfo, err := redisCache.GetWarInfo()
	if err != nil {
		warInfo, err = helldivers.HDClient.GetHelldiversWarInfo(warId)
		if err != nil {
			return lib.WarInfo{}, err
		}
		err = redisCache.SetWarInfo(warInfo)
	}
	return warInfo, err
}
