package lib

import (
	"encoding/json"
)

type diveHarderResponse struct {
	Msg     string `json:"msg"`
	ApiTime string `json:"apiTime"`
}

type DiveHarderPlanet struct {
	PlanetIndex     int     `json:"planetIndex"`
	PlanetName      string  `json:"planetName"`
	InitialOwner    string  `json:"initialOwner"`
	CurrentOwner    string  `json:"currentOwner"`
	PosX            float64 `json:"posX"`
	PosY            float64 `json:"posY"`
	WaypointIndices string  `json:"waypointIndices"`
	WaypointNames   string  `json:"waypointNames"`
	Health          int     `json:"health"`
	MaxHealth       int     `json:"maxHealth"`
	Players         int     `json:"players"`
	RegenPerSecond  float64 `json:"regenPerSecond"`
}

type DiveHarderPlanetsResponse struct {
	diveHarderResponse
	Planets []DiveHarderPlanet `json:"planets"`
}

type DiveHarderPlanetsActive struct {
	PlanetIndex    int     `json:"planetIndex"`
	PlanetName     string  `json:"planetName"`
	RegenPerSecond int     `json:"regenPerSecond"`
	RegenPerHour   int     `json:"regenPerHour"`
	Health         int     `json:"health"`
	MaxHealth      int     `json:"maxHealth"`
	LibPercent     float64 `json:"libPercent"`
	Players        int     `json:"players"`
	HEC            float64 `json:"HEC"`
	MinPlayers     int     `json:"minPlayers"`
	HoursComplete  float64 `json:"hoursComplete"`
}

type DiveHarderPlanetsActiveResponse struct {
	diveHarderResponse
	Planets []DiveHarderPlanetsActive `json:"planets"`
}

type GalaxyStats struct {
	MissionsWon        int   `json:"missionsWon"`
	MissionsLost       int   `json:"missionsLost"`
	MissionTime        int64 `json:"missionTime"`
	BugKills           int64 `json:"bugKills"`
	AutomatonKills     int   `json:"automatonKills"`
	IlluminateKills    int   `json:"illuminateKills"`
	BulletsFired       int64 `json:"bulletsFired"`
	BulletsHit         int64 `json:"bulletsHit"`
	TimePlayed         int64 `json:"timePlayed"`
	Deaths             int   `json:"deaths"`
	Revives            int   `json:"revives"`
	Friendlies         int   `json:"friendlies"`
	MissionSuccessRate int   `json:"missionSuccessRate"`
	Accuracy           int   `json:"accurracy"`
}

type PlanetStats struct {
	PlanetIndex        int   `json:"planetIndex"`
	MissionsWon        int   `json:"missionsWon"`
	MissionsLost       int   `json:"missionsLost"`
	MissionTime        int64 `json:"missionTime"`
	BugKills           int   `json:"bugKills"`
	AutomatonKills     int   `json:"automatonKills"`
	IlluminateKills    int   `json:"illuminateKills"`
	BulletsFired       int64 `json:"bulletsFired"`
	BulletsHit         int64 `json:"bulletsHit"`
	TimePlayed         int64 `json:"timePlayed"`
	Deaths             int   `json:"deaths"`
	Revives            int   `json:"revives"`
	Friendlies         int   `json:"friendlies"`
	MissionSuccessRate int   `json:"missionSuccessRate"`
	Accuracy           int   `json:"accurracy"`
}

type Stats struct {
	GalaxyStats  GalaxyStats   `json:"galaxy_stats"`
	PlanetsStats []PlanetStats `json:"planets_stats"`
}

func (c *Client) GetDiveHarderPlanets() (DiveHarderPlanetsResponse, error) {
	resp, err := c.Request("GET", DiveHarderPlanetsRoute, nil)
	if err != nil {
		return DiveHarderPlanetsResponse{}, err
	}

	var response DiveHarderPlanetsResponse
	err = json.Unmarshal(resp.bodyRead, &response)
	return response, err
}

func (c *Client) GetDiveHarderPlanetsActive() (DiveHarderPlanetsActiveResponse, error) {
	resp, err := c.Request("GET", DiveHarderPlanetsActiveRoute, nil)
	if err != nil {
		return DiveHarderPlanetsActiveResponse{}, err
	}

	var response DiveHarderPlanetsActiveResponse
	err = json.Unmarshal(resp.bodyRead, &response)
	return response, err
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
