package client_tool

import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

type Client struct {
	host string
	httpClient *http.Client
}

func NewClient(host string, timeout time.Duration) *Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Timeout: timeout,
		Transport:  tr,
	}
	return &Client{
		host:       host,
		httpClient: client,
	}
}

func (c *Client) Do(method, endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const jsonMockUrl = "/posts"

func (c *Client) GetPosts() (resp []Post, err error) {
	res, err := c.Do(http.MethodGet, jsonMockUrl, nil)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return resp, err
	}
	return
}