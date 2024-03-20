package redisCache

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
	"fmt"
)

func GetWarSeasons() (lib.WarSeasons, error) {
	data, err := Client.Get(Context, "war_seasons").Bytes()
	if err != nil {
		return lib.WarSeasons{}, err
	}

	var warSeasons lib.WarSeasons
	err = json.Unmarshal(data, &warSeasons)
	if err != nil {
		return lib.WarSeasons{}, err
	}

	return warSeasons, nil
}

func GetPlanets(warSeasonId string) ([]lib.Planet, error) {
	data, err := Client.Get(Context, fmt.Sprintf("%s-planets", warSeasonId)).Bytes()
	if err != nil {
		return nil, err
	}

	var planets []lib.Planet
	err = json.Unmarshal(data, &planets)
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func GetPlanet(warSeasonId string, planetId int) (lib.Planet, error) {
	data, err := Client.Get(Context, fmt.Sprintf("%s-planet-%d", warSeasonId, planetId)).Bytes()
	if err != nil {
		return lib.Planet{}, err
	}

	var planet lib.Planet
	err = json.Unmarshal(data, &planet)
	if err != nil {
		return lib.Planet{}, err
	}

	return planet, nil
}

func GetPlanetStatus(warSeasonId string, planetId int) (lib.PlanetStatus, error) {
	data, err := Client.Get(Context, fmt.Sprintf("%s-planet_status-%d", warSeasonId, planetId)).Bytes()
	if err != nil {
		return lib.PlanetStatus{}, err
	}

	var planetStatus lib.PlanetStatus
	err = json.Unmarshal(data, &planetStatus)
	if err != nil {
		return lib.PlanetStatus{}, err
	}

	return planetStatus, nil
}

func GetFeed(warSeasonId string) ([]lib.FeedMessage, error) {
	data, err := Client.Get(Context, fmt.Sprintf("%s-feed", warSeasonId)).Bytes()
	if err != nil {
		return nil, err
	}

	var feed []lib.FeedMessage
	err = json.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
