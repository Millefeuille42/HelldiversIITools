package lib

import (
	"bytes"
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
	Url    *url.URL
	client *http.Client
	jar    *cookiejar.Jar
}

func New(address string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	parsed, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	return &Client{
		Url: parsed,
		client: &http.Client{
			Jar: jar,
		},
	}, nil
}

func (c *Client) generateRequest(method string, endpoint string, data []byte) (*http.Request, error) {
	if data == nil {
		data = []byte("")
	}
	req, err := http.NewRequest(method, c.Url.String()+endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", "Helldivers II Tools")

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
