package lib

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Faction int
type RewardType int
type TaskType int
type TaskValueType int
type EventType int
type CampaignType int

const (
	Humans     Faction = 1
	Terminids  Faction = 2
	Automatons Faction = 3
	Illuminate Faction = 4

	MedalRewardType RewardType = 1

	LiberateTaskType TaskType = 11
	ControlTaskType  TaskType = 13

	PlanetTaskValueType TaskValueType = 12

	DefenseEventType EventType = 1

	MajorOrderCampaignType CampaignType = 2
)

type Planet struct {
	Index     int          `json:"index"`
	Name      string       `json:"name"`
	Position  Position     `json:"position"`
	Waypoints []PlanetName `json:"waypoints"`
	Sector    int          `json:"sector"`
	Disabled  bool         `json:"disabled"`

	MaxHealth         int     `json:"max_health"`
	Health            int     `json:"health"`
	RegenPerSecond    float64 `json:"regen_per_second"`
	LiberationPercent float64 `json:"liberation_percent"`

	Owner        Faction `json:"owner"`
	InitialOwner Faction `json:"initial_owner"`

	Players            int64 `json:"players"`
	MissionsWon        int64 `json:"missions_won"`
	MissionsLost       int64 `json:"missions_lost"`
	MissionTime        int64 `json:"mission_time"`
	TerminidKills      int64 `json:"terminid_kills"`
	AutomatonKills     int64 `json:"automaton_kills"`
	IlluminateKills    int64 `json:"illuminate_kills"`
	BulletsFired       int64 `json:"bullets_fired"`
	BulletsHit         int64 `json:"bullets_hit"`
	TimePlayed         int64 `json:"time_played"`
	Deaths             int64 `json:"deaths"`
	Revives            int64 `json:"revives"`
	Friendlies         int64 `json:"friendlies"`
	MissionSuccessRate int   `json:"mission_success_rate"`
	Accuracy           int   `json:"accuracy"`

	Attacks         []PlanetAttack   `json:"attacks"`
	Events          []PlanetEvent    `json:"events"`
	JointOperations []JointOperation `json:"joint_operations"`
	HomeWorlds      []HomeWorld      `json:"home_worlds"`
	Campaigns       []Campaign       `json:"campaigns"`
}

type PlanetName struct {
	Index int    `json:"planet_index"`
	Name  string `json:"planet_name"`
}

type Reward struct {
	Type   RewardType `json:"type"`
	Amount int        `json:"amount"`
}

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

func (c *Client) GetPlanetsInfo() ([]PlanetInfo, error) {
	resp, err := c.Request("GET", GoDiversPlanetInfosRoute, nil)
	if err != nil {
		return nil, err
	}

	var planetInfo []PlanetInfo
	err = json.Unmarshal(resp.bodyRead, &planetInfo)
	return planetInfo, err
}

func (c *Client) GetPlanetsStatus() ([]PlanetStatus, error) {
	resp, err := c.Request("GET", GoDiversPlanetStatusRoute, nil)
	if err != nil {
		return nil, err
	}

	var planetStatus []PlanetStatus
	err = json.Unmarshal(resp.bodyRead, &planetStatus)
	return planetStatus, err
}

func (c *Client) GetPlanetsAttack() ([]PlanetAttack, error) {
	resp, err := c.Request("GET", GoDiversPlanetAttacksRoute, nil)
	if err != nil {
		return nil, err
	}

	var planetAttack []PlanetAttack
	err = json.Unmarshal(resp.bodyRead, &planetAttack)
	return planetAttack, err
}

func (c *Client) GetPlanetsEvent() ([]PlanetEvent, error) {
	resp, err := c.Request("GET", GoDiversPlanetEventsRoute, nil)
	if err != nil {
		return nil, err
	}

	var planetEvent []PlanetEvent
	err = json.Unmarshal(resp.bodyRead, &planetEvent)
	return planetEvent, err
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
