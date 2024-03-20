package lib

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetStatus(warSeasonId string) (Status, error) {
	resp, err := c.Request("GET", fmt.Sprintf(StatusRoute, warSeasonId), nil)
	if err != nil {
		return Status{}, err
	}

	var status Status
	err = json.Unmarshal(resp.bodyRead, &status)
	return status, err
}

func (c *Client) GetCurrentWarStatus() (Status, error) {
	return c.GetStatus(c.WarSeasons.Current)
}
