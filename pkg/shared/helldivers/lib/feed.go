package lib

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetFeed(warSeasonId string) ([]FeedMessage, error) {
	resp, err := c.Request("GET", fmt.Sprintf(FeedRoute, warSeasonId), nil)
	if err != nil {
		return nil, err
	}

	var feedMessages []FeedMessage
	err = json.Unmarshal(resp.bodyRead, &feedMessages)
	return feedMessages, err
}

func (c *Client) GetCurrentWarFeed() ([]FeedMessage, error) {
	return c.GetFeed(c.WarSeasons.Current)
}
