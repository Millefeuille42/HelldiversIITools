package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Response struct {
	*http.Response
	bodyRead []byte
}

type Client struct {
	Url    url.URL
	client *http.Client
	jar    *cookiejar.Jar

	WarSeasons WarSeasons
}

func New(scheme, host string, port int) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		Url: url.URL{
			Scheme: scheme,
			Host:   fmt.Sprintf("%s:%d", host, port),
			Path:   "/api",
		},
		client: &http.Client{
			Jar: jar,
		},
	}, nil
}

func (c *Client) generateRequest(method string, endpoint string, data []byte) (*http.Request, error) {
	if data == nil {
		data = []byte("")
	}
	req, err := http.NewRequest(method, c.Url.JoinPath(endpoint).String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return req, err
}

// TODO add 429 management

func (c *Client) Request(method, endpoint string, body []byte) (Response, error) {
	req, err := c.generateRequest(method, endpoint, body)
	if err != nil {
		return Response{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Response{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return Response{resp, respBody}, err
}

func (c *Client) GetWarSeasons() (WarSeasons, error) {
	resp, err := c.Request("GET", "", nil)
	if err != nil {
		return c.WarSeasons, err
	}

	err = json.Unmarshal(resp.bodyRead, &c.WarSeasons)
	return c.WarSeasons, err
}
