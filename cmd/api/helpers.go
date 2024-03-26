package main

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"errors"
	"log"
)

const warId = "801"

func searchInDiveHarderPlanetStats(planets []lib.PlanetStats, id int) (lib.PlanetStats, error) {
	for i := range planets {
		if planets[i].PlanetIndex == id {
			return planets[i], nil
		}
	}
	return lib.PlanetStats{}, errors.New("not found")
}

func searchInPlanetInfo(planets []lib.PlanetInfo, id int) (lib.PlanetInfo, error) {
	for i := range planets {
		if planets[i].Index == id {
			return planets[i], nil
		}
	}
	return lib.PlanetInfo{}, errors.New("not found")
}

func searchInPlanetNames(planets []lib.PlanetName, id int) (lib.PlanetName, error) {
	for i := range planets {
		if planets[i].Index == id {
			return planets[i], nil
		}
	}
	return lib.PlanetName{}, errors.New("not found")
}

func searchInPlanetStatus(planets []lib.PlanetStatus, id int) (lib.PlanetStatus, error) {
	for i := range planets {
		if planets[i].Index == id {
			return planets[i], nil
		}
	}
	return lib.PlanetStatus{}, errors.New("not found")
}

func searchHomeWorlds(homeWorlds []lib.HomeWorld, planetId int) []lib.HomeWorld {
	if homeWorlds == nil {
		return nil
	}
	var planetHomeWorlds []lib.HomeWorld
	for i := range homeWorlds {
		for i2 := range homeWorlds[i].PlanetIndices {
			if homeWorlds[i].PlanetIndices[i2] == planetId {
				planetHomeWorlds = append(planetHomeWorlds, homeWorlds[i])
				break
			}
		}
	}
	return planetHomeWorlds
}

func searchPlanetEvents(planets []lib.PlanetEvent, planetId int) []lib.PlanetEvent {
	if planets == nil {
		return nil
	}
	var events []lib.PlanetEvent
	for i := range planets {
		if planets[i].PlanetIndex == planetId {
			events = append(events, planets[i])
		}
	}
	return events
}

func searchPlanetAttacks(planets []lib.PlanetAttack, planetId int) []lib.PlanetAttack {
	if planets == nil {
		return nil
	}
	var attacks []lib.PlanetAttack
	for i := range planets {
		if planets[i].Source == planetId || planets[i].Target == planetId {
			attacks = append(attacks, planets[i])
		}
	}
	return attacks
}

func searchJointOperations(operations []lib.JointOperation, planetId int) []lib.JointOperation {
	if operations == nil {
		return nil
	}
	var planetOperations []lib.JointOperation
	for i := range operations {
		if operations[i].PlanetIndex == planetId {
			planetOperations = append(planetOperations, operations[i])
		}
	}
	return planetOperations
}

func searchCampaigns(campaigns []lib.Campaign, planetId int) []lib.Campaign {
	var planetCampaigns []lib.Campaign
	for i := range campaigns {
		if campaigns[i].PlanetIndex == planetId {
			planetCampaigns = append(planetCampaigns, campaigns[i])
		}
	}
	return planetCampaigns
}

func constructPlanet(planetId int) (lib.Planet, error) {
	dhPlanetsStats, err := getDiveHarderPlanetStats()
	if err != nil {
		log.Println(err)
	}

	warInfo, err := getWarInfo()
	if err != nil {
		return lib.Planet{}, err
	}

	status, err := getStatus()
	if err != nil {
		return lib.Planet{}, err
	}

	planetNames, err := getPlanetNames()
	if err != nil {
		return lib.Planet{}, err
	}

	planet := lib.Planet{}

	if planetName, err := searchInPlanetNames(planetNames, planetId); err == nil {
		planet.Name = planetName.Name
	}

	if planetInfo, err := searchInPlanetInfo(warInfo.PlanetInfos, planetId); err == nil {
		planet.Index = planetInfo.Index
		planet.Position = planetInfo.Position
		for _, waypointId := range planetInfo.Waypoints {
			planetName, err := searchInPlanetNames(planetNames, waypointId)
			if err != nil {
				continue
			}
			planet.Waypoints = append(planet.Waypoints, planetName)
		}
		planet.Sector = planetInfo.Sector
		planet.MaxHealth = planetInfo.MaxHealth
		planet.Disabled = planetInfo.Disabled
		planet.InitialOwner = lib.Faction(planetInfo.InitialOwner)
	}

	if planetStatus, err := searchInPlanetStatus(status.PlanetStatus, planetId); err == nil {
		planet.Index = planetStatus.Index
		planet.Owner = lib.Faction(planetStatus.Owner)
		planet.Health = planetStatus.Health
		planet.RegenPerSecond = planetStatus.RegenPerSecond
		planet.Players = planetStatus.Players
	}

	if planet.MaxHealth != 0 && planet.Health != 0 {
		planet.LiberationPercent = 100.0 - 100.0*float64(planet.Health)/float64(planet.MaxHealth)
	}

	planet.Attacks = searchPlanetAttacks(status.PlanetAttacks, planetId)
	planet.Events = searchPlanetEvents(status.PlanetEvents, planetId)
	planet.JointOperations = searchJointOperations(status.JointOperations, planetId)
	planet.Campaigns = searchCampaigns(status.Campaigns, planetId)
	planet.HomeWorlds = searchHomeWorlds(warInfo.HomeWorlds, planetId)

	if dhPlanetsStats != nil {
		if dhPlanetStats, err := searchInDiveHarderPlanetStats(dhPlanetsStats, planetId); err == nil {
			planet.Index = dhPlanetStats.PlanetIndex
			planet.MissionsWon = dhPlanetStats.MissionsWon
			planet.MissionsLost = dhPlanetStats.MissionsLost
			planet.MissionTime = dhPlanetStats.MissionTime
			planet.TerminidKills = dhPlanetStats.BugKills
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
				task.Target.Index = planet.Index
				task.Target.Name = planet.Name
			}
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
