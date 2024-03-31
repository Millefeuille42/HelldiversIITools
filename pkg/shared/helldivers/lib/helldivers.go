package lib

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TimeSinceStart struct {
	SecondsSinceStart int64 `json:"secondsSinceStart"`
}

type WarTime struct {
	Time int64 `json:"time"`
}

type PlanetAttack struct {
	Source int `json:"source"`
	Target int `json:"target"`
}

type Campaign struct {
	Id          int `json:"id"`
	PlanetIndex int `json:"planetIndex"`
	Type        int `json:"type"`
	Count       int `json:"count"`
}

type PlanetStatus struct {
	Index          int     `json:"index"`
	Owner          int     `json:"owner"`
	Health         int     `json:"health"`
	RegenPerSecond float64 `json:"regenPerSecond"`
	Players        int64   `json:"players"`
}

type JointOperation struct {
	Id          int `json:"id"`
	PlanetIndex int `json:"planetIndex"`
	HqNodeIndex int `json:"hqNodeIndex"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type HomeWorld struct {
	Race          int   `json:"race"`
	PlanetIndices []int `json:"planetIndices"`
}

type PlanetEvent struct {
	Id                int   `json:"id"`
	PlanetIndex       int   `json:"planetIndex"`
	EventType         int   `json:"eventType"`
	Race              int   `json:"race"`
	Health            int   `json:"health"`
	MaxHealth         int   `json:"maxHealth"`
	StartTime         int   `json:"startTime"`
	ExpireTime        int   `json:"expireTime"`
	CampaignId        int   `json:"campaignId"`
	JointOperationIds []int `json:"jointOperationIds"`
}

type NewsMessage struct {
	Id        int           `json:"id"`
	Published int           `json:"published"`
	Type      int           `json:"type"`
	TagIds    []interface{} `json:"tagIds"`
	Message   string        `json:"message"`
}

type Status struct {
	WarId                       int              `json:"warId"`
	Time                        int              `json:"time"`
	ImpactMultiplier            float64          `json:"impactMultiplier"`
	StoryBeatId32               int              `json:"storyBeatId32"`
	PlanetStatus                []PlanetStatus   `json:"planetStatus"`
	PlanetAttacks               []PlanetAttack   `json:"planetAttacks"`
	Campaigns                   []Campaign       `json:"campaigns"`
	JointOperations             []JointOperation `json:"jointOperations"`
	PlanetEvents                []PlanetEvent    `json:"planetEvents"`
	CommunityTargets            []interface{}    `json:"communityTargets"` // TODO Get format of inline and below
	PlanetActiveEffects         []interface{}    `json:"planetActiveEffects"`
	ActiveElectionPolicyEffects []interface{}    `json:"activeElectionPolicyEffects"`
	GlobalEvents                []interface{}    `json:"globalEvents"`
	SuperEarthWarResults        []interface{}    `json:"superEarthWarResults"`
}

type PlanetInfo struct {
	Index        int      `json:"index"`
	SettingsHash int64    `json:"settingsHash"`
	Position     Position `json:"position"`
	Waypoints    []int    `json:"waypoints"`
	Sector       int      `json:"sector"`
	MaxHealth    int      `json:"maxHealth"`
	Disabled     bool     `json:"disabled"`
	InitialOwner int      `json:"initialOwner"`
}

type WarInfo struct {
	WarId                  int           `json:"warId"`
	StartDate              int           `json:"startDate"`
	EndDate                int           `json:"endDate"`
	MinimumClientVersion   string        `json:"minimumClientVersion"`
	PlanetInfos            []PlanetInfo  `json:"planetInfos"`
	HomeWorlds             []HomeWorld   `json:"homeWorlds"`
	CapitalInfos           []interface{} `json:"capitalInfos"`
	PlanetPermanentEffects []interface{} `json:"planetPermanentEffects"`
}

type Assignment struct {
	Id32      int       `json:"id32"`
	Progress  []float64 `json:"progress"`
	ExpiresIn int       `json:"expiresIn"`
	Setting   struct {
		Type            int    `json:"type"`
		OverrideTitle   string `json:"overrideTitle"`
		OverrideBrief   string `json:"overrideBrief"`
		TaskDescription string `json:"taskDescription"`
		Tasks           []struct {
			Type       int   `json:"type"`
			Values     []int `json:"values"`
			ValueTypes []int `json:"valueTypes"`
		} `json:"tasks"`
		Reward struct {
			Type   int `json:"type"`
			Id32   int `json:"id32"`
			Amount int `json:"amount"`
		} `json:"reward"`
		Flags int `json:"flags"`
	} `json:"setting"`
}

type GenericStats struct {
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

type Summary struct {
	GalaxyStats  GalaxyStats   `json:"galaxy_stats"`
	PlanetsStats []PlanetStats `json:"planets_stats"`
}

func (c *Client) GetHelldiversNewsFeed(warId string, timestamp int) ([]NewsMessage, error) {
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversNewsFeedRoute, warId, timestamp), nil)
	if err != nil {
		return nil, err
	}

	var newsFeed []NewsMessage
	err = json.Unmarshal(resp.bodyRead, &newsFeed)
	return newsFeed, err
}

func (c *Client) GetHelldiversLatestNewsFeed(warId string) (NewsMessage, error) {
	newsFeed, err := c.GetHelldiversNewsFeed(warId, 0)
	if err != nil {
		return NewsMessage{}, err
	}
	for len(newsFeed) != 1 {
		latest := newsFeed[len(newsFeed)-1]
		newNewsFeed, err := c.GetHelldiversNewsFeed(warId, latest.Published)
		if err != nil {
			return NewsMessage{}, err
		}
		if len(newNewsFeed) == 0 {
			break
		}
		newsFeed = newNewsFeed
	}

	return newsFeed[len(newsFeed)-1], nil
}

func (c *Client) GetHelldiversAssignment(warId string) (Assignment, error) {
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversAssignmentRoute, warId), nil)
	if err != nil {
		return Assignment{}, err
	}

	var assignment []Assignment
	err = json.Unmarshal(resp.bodyRead, &assignment)
	if len(assignment) <= 0 {
		return Assignment{}, errors.New("no assignment")
	}
	return assignment[0], err
}

func (c *Client) GetHelldiversStatus(warId string) (Status, error) {
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversWarStatusRoute, warId), nil)
	if err != nil {
		return Status{}, err
	}

	var status Status
	err = json.Unmarshal(resp.bodyRead, &status)
	return status, err
}

func (c *Client) GetHelldiversWarInfo(warId string) (WarInfo, error) {
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversWarInfoRoute, warId), nil)
	if err != nil {
		return WarInfo{}, err
	}

	var warInfo WarInfo
	err = json.Unmarshal(resp.bodyRead, &warInfo)
	return warInfo, err
}

func (c *Client) GetHelldiversTimeSinceStart(warId string) (TimeSinceStart, error) {
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversTimeSinceStartRoute, warId), nil)
	if err != nil {
		return TimeSinceStart{}, err
	}

	var timeSinceStart TimeSinceStart
	err = json.Unmarshal(resp.bodyRead, &timeSinceStart)
	return timeSinceStart, err
}

func (c *Client) GetHelldiversWarTime(warId string) (WarTime, error) {
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversWarTimeRoute, warId), nil)
	if err != nil {
		return WarTime{}, err
	}

	var warTime WarTime
	err = json.Unmarshal(resp.bodyRead, &warTime)
	return warTime, err
}

func (c *Client) GetHelldiversSummary(warId string) (Summary, error) {
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversSummaryRoute, warId), nil)
	if err != nil {
		return Summary{}, err
	}

	var summary Summary
	err = json.Unmarshal(resp.bodyRead, &summary)
	return summary, err
}
