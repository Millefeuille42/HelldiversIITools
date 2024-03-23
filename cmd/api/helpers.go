package main

import (
	"Helldivers2Tools/pkg/api/redisCache"
	"Helldivers2Tools/pkg/shared/helldivers"
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"errors"
)

const warId = "801"

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

func getDiveHarderPlanets() (lib.DiveHarderPlanetsResponse, error) {
	planets, err := redisCache.GetDiveHarderPlanets()
	if err != nil {
		planets, err = helldivers.DiveHarderClient.GetDiveHarderPlanets()
		if err != nil {
			return lib.DiveHarderPlanetsResponse{}, err
		}
		err = redisCache.SetDiveHarderPlanets(planets)
	}
	return planets, err
}

func getDiveHarderPlanetsActive() (lib.DiveHarderPlanetsActiveResponse, error) {
	planets, err := redisCache.GetDiveHarderPlanetsActive()
	if err != nil {
		planets, err = helldivers.DiveHarderClient.GetDiveHarderPlanetsActive()
		if err != nil {
			return lib.DiveHarderPlanetsActiveResponse{}, err
		}
		err = redisCache.SetDiveHarderPlanetsActive(planets)
	}
	return planets, err
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

func searchInDiveHarderPlanets(planets lib.DiveHarderPlanetsResponse, id int) (lib.DiveHarderPlanet, error) {
	for i := range planets.Planets {
		if planets.Planets[i].PlanetIndex == id {
			return planets.Planets[i], nil
		}
	}
	return lib.DiveHarderPlanet{}, errors.New("not found")
}

func searchInDiveHarderPlanetActive(planets lib.DiveHarderPlanetsActiveResponse, id int) (lib.DiveHarderPlanetsActive, error) {
	for i := range planets.Planets {
		if planets.Planets[i].PlanetIndex == id {
			return planets.Planets[i], nil
		}
	}
	return lib.DiveHarderPlanetsActive{}, errors.New("not found")
}

func searchInDiveHarderPlanetStats(planets []lib.PlanetStats, id int) (lib.PlanetStats, error) {
	for i := range planets {
		if planets[i].PlanetIndex == id {
			return planets[i], nil
		}
	}
	return lib.PlanetStats{}, errors.New("not found")
}

func constructPlanet(planetId int) (lib.Planet, error) {
	dhPlanets, err := getDiveHarderPlanets()
	if err != nil {
		return lib.Planet{}, err
	}
	dhPlanetsActive, err := getDiveHarderPlanetsActive()
	if err != nil {
		return lib.Planet{}, err
	}
	dhPlanetsStats, err := getDiveHarderPlanetStats()
	if err != nil {
		return lib.Planet{}, err
	}

	planet := lib.Planet{}

	if dhPlanet, err := searchInDiveHarderPlanets(dhPlanets, planetId); err == nil {
		planet.PlanetIndex = dhPlanet.PlanetIndex
		planet.PlanetName = dhPlanet.PlanetName
		planet.InitialOwner = dhPlanet.InitialOwner
		planet.CurrentOwner = dhPlanet.CurrentOwner
		planet.PosX = dhPlanet.PosX
		planet.PosY = dhPlanet.PosY
		planet.WaypointIndices = dhPlanet.WaypointIndices
		planet.WaypointNames = dhPlanet.WaypointNames
		planet.Health = dhPlanet.Health
		planet.MaxHealth = dhPlanet.MaxHealth
		planet.Players = dhPlanet.Players
		planet.RegenPerSecond = dhPlanet.RegenPerSecond
	}

	if dhPlanetActive, err := searchInDiveHarderPlanetActive(dhPlanetsActive, planetId); err == nil {
		planet.PlanetIndex = dhPlanetActive.PlanetIndex
		planet.PlanetName = dhPlanetActive.PlanetName
		planet.Health = dhPlanetActive.Health
		planet.MaxHealth = dhPlanetActive.MaxHealth
		planet.LibPercent = dhPlanetActive.LibPercent
		planet.Players = dhPlanetActive.Players
		planet.HoursComplete = dhPlanetActive.HoursComplete
	}

	if dhPlanetStats, err := searchInDiveHarderPlanetStats(dhPlanetsStats, planetId); err == nil {
		planet.PlanetIndex = dhPlanetStats.PlanetIndex
		planet.MissionsWon = dhPlanetStats.MissionsWon
		planet.MissionsLost = dhPlanetStats.MissionsLost
		planet.MissionTime = dhPlanetStats.MissionTime
		planet.BugKills = dhPlanetStats.BugKills
		planet.AutomatonKills = dhPlanetStats.AutomatonKills
		planet.IlluminateKills = dhPlanetStats.IlluminateKills
		planet.BulletsFired = dhPlanetStats.BulletsFired
		planet.BulletsHit = dhPlanetStats.BulletsHit
		planet.TimePlayed = dhPlanetStats.TimePlayed
		planet.Deaths = dhPlanetStats.Deaths
		planet.Revives = dhPlanetStats.Revives
		planet.Friendlies = dhPlanetStats.Friendlies
		planet.MissionSuccessRate = dhPlanetStats.MissionSuccessRate
		planet.Accuracy = dhPlanetStats.Accuracy
	}

	return planet, nil
}

func constructTasks(assignment lib.Assignment) ([]lib.Task, error) {
	var tasks []lib.Task
	for taskIndex, t := range assignment.Setting.Tasks {
		task := lib.Task{
			Type:     lib.TaskType(t.Type),
			Progress: assignment.Progress[taskIndex],
		}
		for index, val := range t.ValueTypes {
			if lib.TaskValueType(val) == lib.PlanetTaskValueType {
				planet, err := constructPlanet(t.Values[index])
				if err != nil {
					return nil, err
				}
				task.Target.Name = planet.PlanetName
				task.Target.Index = planet.PlanetIndex
			}
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
