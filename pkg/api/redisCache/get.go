package redisCache

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
)

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

func GetDiveHarderPlanets() (lib.DiveHarderPlanetsResponse, error) {
	data, err := Client.Get(Context, "diveharder_planets").Bytes()
	if err != nil {
		return lib.DiveHarderPlanetsResponse{}, err
	}

	var planets lib.DiveHarderPlanetsResponse
	err = json.Unmarshal(data, &planets)
	if err != nil {
		return lib.DiveHarderPlanetsResponse{}, err
	}

	return planets, nil
}

func GetDiveHarderPlanetsActive() (lib.DiveHarderPlanetsActiveResponse, error) {
	data, err := Client.Get(Context, "diveharder_planets_active").Bytes()
	if err != nil {
		return lib.DiveHarderPlanetsActiveResponse{}, err
	}

	var planets lib.DiveHarderPlanetsActiveResponse
	err = json.Unmarshal(data, &planets)
	if err != nil {
		return lib.DiveHarderPlanetsActiveResponse{}, err
	}

	return planets, nil
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
