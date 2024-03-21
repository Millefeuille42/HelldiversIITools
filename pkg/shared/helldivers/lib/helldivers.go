package lib

import (
	"encoding/json"
	"errors"
	"fmt"
)

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
