package redisCache

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
)

func GetPlanetNames() ([]lib.PlanetName, error) {
	data, err := Client.Get(Context, "planet_names").Bytes()
	if err != nil {
		return nil, err
	}

	var planetNames []lib.PlanetName
	err = json.Unmarshal(data, &planetNames)
	if err != nil {
		return nil, err
	}

	return planetNames, nil
}

func GetNewsMessage() (lib.NewsMessage, error) {
	data, err := Client.Get(Context, "news_message").Bytes()
	if err != nil {
		return lib.NewsMessage{}, err
	}

	var message lib.NewsMessage
	err = json.Unmarshal(data, &message)
	if err != nil {
		return lib.NewsMessage{}, err
	}

	return message, nil
}

func GetLatestNewsMessage() (lib.NewsMessage, error) {
	data, err := Client.Get(Context, "latest_news_message").Bytes()
	if err != nil {
		return lib.NewsMessage{}, err
	}

	var message lib.NewsMessage
	err = json.Unmarshal(data, &message)
	if err != nil {
		return lib.NewsMessage{}, err
	}

	return message, nil
}

func GetPlanetsStats() ([]lib.PlanetStats, error) {
	data, err := Client.Get(Context, "planets_stats").Bytes()
	if err != nil {
		return nil, err
	}

	var planets []lib.PlanetStats
	err = json.Unmarshal(data, &planets)
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func GetGalaxyStats() (lib.GalaxyStats, error) {
	data, err := Client.Get(Context, "galaxy_stats").Bytes()
	if err != nil {
		return lib.GalaxyStats{}, err
	}

	var galaxy lib.GalaxyStats
	err = json.Unmarshal(data, &galaxy)
	if err != nil {
		return lib.GalaxyStats{}, err
	}

	return galaxy, nil
}

func GetAssignment() (lib.Assignment, error) {
	data, err := Client.Get(Context, "assignment").Bytes()
	if err != nil {
		return lib.Assignment{}, err
	}

	var assignment lib.Assignment
	err = json.Unmarshal(data, &assignment)
	if err != nil {
		return lib.Assignment{}, err
	}

	return assignment, nil
}

func GetStatus() (lib.Status, error) {
	data, err := Client.Get(Context, "status").Bytes()
	if err != nil {
		return lib.Status{}, err
	}

	var status lib.Status
	err = json.Unmarshal(data, &status)
	if err != nil {
		return lib.Status{}, err
	}

	return status, nil
}

func GetWarInfo() (lib.WarInfo, error) {
	data, err := Client.Get(Context, "war_info").Bytes()
	if err != nil {
		return lib.WarInfo{}, err
	}

	var warInfo lib.WarInfo
	err = json.Unmarshal(data, &warInfo)
	if err != nil {
		return lib.WarInfo{}, err
	}

	return warInfo, nil
}
