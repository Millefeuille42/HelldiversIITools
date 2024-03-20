package redisCache

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
	"fmt"
	"time"
)

func Set(key string, value []byte, ttl time.Duration) error {
	return Client.Set(Context, key, value, ttl).Err()
}

func SetWarSeasons(warSeasons lib.WarSeasons) error {
	val, err := json.Marshal(warSeasons)
	if err != nil {
		return err
	}
	return Set("war_seasons", val, time.Hour*24)
}

func SetPlanets(warSeasonId string, planets []lib.Planet) error {
	val, err := json.Marshal(planets)
	if err != nil {
		return err
	}
	return Set(fmt.Sprintf("%s-planets", warSeasonId), val, time.Hour*24)
}

func SetPlanet(warSeasonId string, planet lib.Planet) error {
	val, err := json.Marshal(planet)
	if err != nil {
		return err
	}
	return Set(fmt.Sprintf("%s-planet-%d", warSeasonId, planet.Index), val, time.Minute*5)
}

func SetPlanetStatus(warSeasonId string, planetStatus lib.PlanetStatus) error {
	val, err := json.Marshal(planetStatus)
	if err != nil {
		return err
	}
	return Set(fmt.Sprintf("%s-planet_status-%d", warSeasonId, planetStatus.Planet.Index), val, time.Minute*5)
}

func SetFeed(warSeasonId string, feed []lib.FeedMessage) error {
	val, err := json.Marshal(feed)
	if err != nil {
		return err
	}
	return Set(fmt.Sprintf("%s-feed", warSeasonId), val, time.Minute*5)
}
