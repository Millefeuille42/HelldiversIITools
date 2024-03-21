package redisCache

import (
	"Helldivers2Tools/pkg/shared/helldivers/lib"
	"encoding/json"
	"time"
)

func Set(key string, value []byte, ttl time.Duration) error {
	return Client.Set(Context, key, value, ttl).Err()
}

func SetNewsMessage(message lib.NewsMessage) error {
	val, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return Set("news_message", val, time.Minute*30)
}

func SetLatestNewsMessage(message lib.NewsMessage) error {
	val, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return Set("latest_news_message", val, 0)
}

func SetDiveHarderPlanets(planets lib.DiveHarderPlanetsResponse) error {
	val, err := json.Marshal(planets)
	if err != nil {
		return err
	}
	return Set("diveharder_planets", val, time.Minute*5)
}

func SetDiveHarderPlanetsActive(planets lib.DiveHarderPlanetsActiveResponse) error {
	val, err := json.Marshal(planets)
	if err != nil {
		return err
	}
	return Set("diveharder_planets_active", val, time.Minute*5)
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

func SetAssignment(assignment lib.Assignment) error {
	val, err := json.Marshal(assignment)
	if err != nil {
		return err
	}
	return Set("assignment", val, time.Minute*30)
}
