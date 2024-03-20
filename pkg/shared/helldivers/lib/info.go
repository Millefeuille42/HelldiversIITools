package lib

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetInfo(warSeasonId string) (Info, error) {
	resp, err := c.Request("GET", fmt.Sprintf(InfoRoute, warSeasonId), nil)
	if err != nil {
		return Info{}, err
	}

	var info Info
	err = json.Unmarshal(resp.bodyRead, &info)
	return info, err
}

func (c *Client) GetCurrentWarInfo() (Info, error) {
	return c.GetInfo(c.WarSeasons.Current)
}
