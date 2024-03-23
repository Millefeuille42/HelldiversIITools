package lib

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Campaign struct {
	Id          int `json:"id"`
	PlanetIndex int `json:"planetIndex"`
	Type        int `json:"type"`
	Count       int `json:"count"`
}

type Status struct {
	WarId            int     `json:"warId"`
	Time             int     `json:"time"`
	ImpactMultiplier float64 `json:"impactMultiplier"`
	StoryBeatId32    int     `json:"storyBeatId32"`
	PlanetStatus     []struct {
		Index          int     `json:"index"`
		Owner          int     `json:"owner"`
		Health         int     `json:"health"`
		RegenPerSecond float64 `json:"regenPerSecond"`
		Players        int     `json:"players"`
	} `json:"planetStatus"`
	PlanetAttacks []struct {
		Source int `json:"source"`
		Target int `json:"target"`
	} `json:"planetAttacks"`
	Campaigns                   []Campaign    `json:"campaigns"`
	CommunityTargets            []interface{} `json:"communityTargets"` // TODO Get format of inline and below
	JointOperations             []interface{} `json:"jointOperations"`
	PlanetEvents                []interface{} `json:"planetEvents"`
	PlanetActiveEffects         []interface{} `json:"planetActiveEffects"`
	ActiveElectionPolicyEffects []interface{} `json:"activeElectionPolicyEffects"`
	GlobalEvents                []interface{} `json:"globalEvents"`
	SuperEarthWarResults        []interface{} `json:"superEarthWarResults"`
}

type NewsMessage struct {
	Id        int           `json:"id"`
	Published int           `json:"published"`
	Type      int           `json:"type"`
	TagIds    []interface{} `json:"tagIds"`
	Message   string        `json:"message"`
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
	resp, err := c.Request("GET", fmt.Sprintf(HelldiversStatusRoute, warId), nil)
	if err != nil {
		return Status{}, err
	}

	var status Status
	err = json.Unmarshal(resp.bodyRead, &status)
	return status, err
}
