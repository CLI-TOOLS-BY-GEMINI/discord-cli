package discord

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://discord.com/api/v10"
	userAgent      = "DiscordBot (https://github.com/minatosingull/discord-cli, 0.1.0)"
)

// Client is a Discord API client
type Client struct {
	token      string
	httpClient *http.Client
	BaseURL    string
}

// NewClient creates a new Discord API client
func NewClient(token string) *Client {
	return &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		BaseURL: defaultBaseURL,
	}
}

// SetHTTPClient allows custom HTTP client (for testing)
func (c *Client) SetHTTPClient(hc *http.Client) {
	c.httpClient = hc
}

func (c *Client) do(method, path string, body interface{}) (*http.Response, error) {
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("User-Agent", userAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.httpClient.Do(req)
}
