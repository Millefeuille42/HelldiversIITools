package lib

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetEvents(warSeasonId string) ([]GlobalEvent, error) {
	resp, err := c.Request("GET", fmt.Sprintf(EventsRoute, warSeasonId), nil)
	if err != nil {
		return nil, err
	}

	var events []GlobalEvent
	err = json.Unmarshal(resp.bodyRead, &events)
	return events, err
}

func (c *Client) GetCurrentWarEvents() ([]GlobalEvent, error) {
	return c.GetEvents(c.WarSeasons.Current)
}

func (c *Client) GetLatestEvent(warSeasonId string) (GlobalEvent, error) {
	resp, err := c.Request("GET", fmt.Sprintf(LatestEventRoute, warSeasonId), nil)
	if err != nil {
		return GlobalEvent{}, err
	}

	var event GlobalEvent
	err = json.Unmarshal(resp.bodyRead, &event)
	return event, err
}

func (c *Client) GetCurrentWarLatestEvent() (GlobalEvent, error) {
	return c.GetLatestEvent(c.WarSeasons.Current)
}
