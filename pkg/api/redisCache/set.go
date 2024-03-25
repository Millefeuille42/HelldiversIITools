package redisCache

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
	"time"
)

func Set(key string, value []byte, ttl time.Duration) error {
	return Client.Set(Context, key, value, ttl).Err()
}

func SetPlanetNames(planetNames []lib.PlanetName) error {
	val, err := json.Marshal(planetNames)
	if err != nil {
		return err
	}
	return Set("planet_names", val, 0)
}

func SetLatestNewsMessage(message lib.NewsMessage) error {
	val, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return Set("latest_news_message", val, 0)
}

func SetNewsMessage(message lib.NewsMessage) error {
	val, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return Set("news_message", val, time.Minute*30)
}

func SetAssignment(assignment lib.Assignment) error {
	val, err := json.Marshal(assignment)
	if err != nil {
		return err
	}
	return Set("assignment", val, time.Minute*30)
}

func SetPlanetsStats(stats []lib.PlanetStats) error {
	val, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	return Set("planets_stats", val, time.Minute*5)
}

func SetGalaxyStats(stats lib.GalaxyStats) error {
	val, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	return Set("galaxy_stats", val, time.Minute*5)
}

func SetStatus(status lib.Status) error {
	val, err := json.Marshal(status)
	if err != nil {
		return err
	}
	return Set("status", val, time.Minute*5)
}

func SetWarInfo(warInfo lib.WarInfo) error {
	val, err := json.Marshal(warInfo)
	if err != nil {
		return err
	}
	return Set("war_info", val, time.Minute*5)
}
