package lib

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Planet struct {
	PlanetIndex        int     `json:"planetIndex"`
	PlanetName         string  `json:"planetName"`
	InitialOwner       string  `json:"initialOwner"`
	CurrentOwner       string  `json:"currentOwner"`
	PosX               float64 `json:"posX"`
	PosY               float64 `json:"posY"`
	WaypointIndices    string  `json:"waypointIndices"`
	WaypointNames      string  `json:"waypointNames"`
	Health             int     `json:"health"`
	MaxHealth          int     `json:"maxHealth"`
	Players            int64   `json:"players"`
	RegenPerSecond     float64 `json:"regenPerSecond"`
	MissionsWon        int64   `json:"missionsWon"`
	MissionsLost       int64   `json:"missionsLost"`
	MissionTime        int64   `json:"missionTime"`
	BugKills           int64   `json:"bugKills"`
	AutomatonKills     int64   `json:"automatonKills"`
	IlluminateKills    int64   `json:"illuminateKills"`
	BulletsFired       int64   `json:"bulletsFired"`
	BulletsHit         int64   `json:"bulletsHit"`
	TimePlayed         int64   `json:"timePlayed"`
	Deaths             int64   `json:"deaths"`
	Revives            int64   `json:"revives"`
	Friendlies         int64   `json:"friendlies"`
	MissionSuccessRate int     `json:"missionSuccessRate"`
	Accuracy           int     `json:"accurracy"`
	LibPercent         float64 `json:"libPercent"`
	HoursComplete      float64 `json:"hoursComplete"`
}

type PlanetName struct {
	Index int    `json:"planetIndex"`
	Name  string `json:"planetName"`
}

type RewardType int

const (
	MedalRewardType RewardType = 1
)

type Reward struct {
	Type   RewardType `json:"type"`
	Amount int        `json:"amount"`
}

type TaskType int

const (
	LiberateTaskType TaskType = 11
	ControlTaskType  TaskType = 13
)

type TaskValueType int

const (
	PlanetTaskValueType TaskValueType = 12
)

type Task struct {
	Type     TaskType   `json:"type"`
	Target   PlanetName `json:"target"`
	Progress float64    `json:"progress"`
}

type MajorOrder struct {
	Title       string    `json:"title"`
	Briefing    string    `json:"briefing"`
	Description string    `json:"description"`
	Tasks       []Task    `json:"tasks"`
	Reward      Reward    `json:"reward"`
	EndsAt      time.Time `json:"ends_at"`
}

func SplitNewsMessage(news NewsMessage) (title, message string) {
	title = "New Message"
	message = news.Message
	newsSplit := strings.Split(news.Message, "\n")
	if len(newsSplit) > 1 {
		title = newsSplit[0]
		message = strings.Join(newsSplit[1:], "\n")
	}

	return title, message
}

func (c *Client) GetNewsMessage() (NewsMessage, error) {
	resp, err := c.Request("GET", GoDiversFeedRoute, nil)
	if err != nil {
		return NewsMessage{}, err
	}

	var newsMessage NewsMessage
	err = json.Unmarshal(resp.bodyRead, &newsMessage)
	return newsMessage, err
}

func (c *Client) GetMajorOrder() (MajorOrder, error) {
	resp, err := c.Request("GET", GoDiversOrderRoute, nil)
	if err != nil {
		return MajorOrder{}, err
	}

	var order MajorOrder
	err = json.Unmarshal(resp.bodyRead, &order)
	return order, err
}

func (c *Client) GetPlanetsName() ([]PlanetName, error) {
	resp, err := c.Request("GET", GoDiversPlanetsNameRoute, nil)
	if err != nil {
		return nil, err
	}

	var planetName []PlanetName
	err = json.Unmarshal(resp.bodyRead, &planetName)
	return planetName, err
}

func (c *Client) GetPlanet(planetId int) (Planet, error) {
	resp, err := c.Request("GET", fmt.Sprintf(GoDiversPlanetRoute, planetId), nil)
	if err != nil {
		return Planet{}, err
	}

	var planet Planet
	err = json.Unmarshal(resp.bodyRead, &planet)
	return planet, err
}

func (c *Client) GetGalaxyStats() (GalaxyStats, error) {
	resp, err := c.Request("GET", GoDiversGalaxyStatsRoute, nil)
	if err != nil {
		return GalaxyStats{}, err
	}

	var stats GalaxyStats
	err = json.Unmarshal(resp.bodyRead, &stats)
	return stats, err
}

func (c *Client) GetCampaigns() ([]Campaign, error) {
	resp, err := c.Request("GET", GoDiversCampaignsRoute, nil)
	if err != nil {
		return nil, err
	}

	var campaigns []Campaign
	err = json.Unmarshal(resp.bodyRead, &campaigns)
	return campaigns, err
}
