package lib

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPlanets(warSeasonId string) ([]Planet, error) {
	resp, err := c.Request("GET", fmt.Sprintf(PlanetsRoute, warSeasonId), nil)
	if err != nil {
		return nil, err
	}

	var planets []Planet
	err = json.Unmarshal(resp.bodyRead, &planets)
	return planets, err
}

func (c *Client) GetCurrentWarPlanets() ([]Planet, error) {
	return c.GetPlanets(c.WarSeasons.Current)
}

func (c *Client) GetPlanet(warSeasonId string, planetId int) (Planet, error) {
	resp, err := c.Request("GET", fmt.Sprintf(PlanetRoute, warSeasonId, planetId), nil)
	if err != nil {
		return Planet{}, err
	}

	var planet Planet
	err = json.Unmarshal(resp.bodyRead, &planet)
	return planet, err
}

func (c *Client) GetCurrentWarPlanet(planetId int) (Planet, error) {
	return c.GetPlanet(c.WarSeasons.Current, planetId)
}

func (c *Client) GetPlanetStatus(warSeasonId string, planetId int) (PlanetStatus, error) {
	resp, err := c.Request("GET", fmt.Sprintf(PlanetStatusRoute, warSeasonId, planetId), nil)
	if err != nil {
		return PlanetStatus{}, err
	}

	var planetStatus PlanetStatus
	err = json.Unmarshal(resp.bodyRead, &planetStatus)
	return planetStatus, err
}

func (c *Client) GetCurrentWarPlanetStatus(planetId int) (PlanetStatus, error) {
	return c.GetPlanetStatus(c.WarSeasons.Current, planetId)
}
