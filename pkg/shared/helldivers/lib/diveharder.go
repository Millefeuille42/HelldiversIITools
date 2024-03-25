package lib

import (
	"encoding/json"
)

type GenericStats struct {
	PlanetIndex        int   `json:"planetIndex"`
	MissionsWon        int64 `json:"missionsWon"`
	MissionsLost       int64 `json:"missionsLost"`
	MissionTime        int64 `json:"missionTime"`
	BugKills           int64 `json:"bugKills"`
	AutomatonKills     int64 `json:"automatonKills"`
	IlluminateKills    int64 `json:"illuminateKills"`
	BulletsFired       int64 `json:"bulletsFired"`
	BulletsHit         int64 `json:"bulletsHit"`
	TimePlayed         int64 `json:"timePlayed"`
	Deaths             int64 `json:"deaths"`
	Revives            int64 `json:"revives"`
	Friendlies         int64 `json:"friendlies"`
	MissionSuccessRate int   `json:"missionSuccessRate"`
	Accuracy           int   `json:"accurracy"`
}

type GalaxyStats GenericStats

type PlanetStats struct {
	PlanetIndex int `json:"planetIndex"`
	GenericStats
}

type Stats struct {
	GalaxyStats  GalaxyStats   `json:"galaxy_stats"`
	PlanetsStats []PlanetStats `json:"planets_stats"`
}

func (c *Client) GetDiveHarderStats() (Stats, error) {
	resp, err := c.Request("GET", DiveHarderStatsRoute, nil)
	if err != nil {
		return Stats{}, err
	}

	var stats Stats
	err = json.Unmarshal(resp.bodyRead, &stats)
	return stats, err
}
