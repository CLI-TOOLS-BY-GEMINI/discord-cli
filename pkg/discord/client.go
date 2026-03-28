package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// GetMe returns the current user (Bot)
func (c *Client) GetMe() (*User, error) {
	resp, err := c.do(http.MethodGet, "/users/@me", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetChannel returns a channel info
func (c *Client) GetChannel(channelID string) (*Channel, error) {
	resp, err := c.do(http.MethodGet, "/channels/"+channelID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var channel Channel
	if err := json.NewDecoder(resp.Body).Decode(&channel); err != nil {
		return nil, err
	}
	return &channel, nil
}

// CreateMessage creates a message in a channel
func (c *Client) CreateMessage(channelID string, content string) (*Message, error) {
	reqBody := MessageCreateRequest{Content: content}
	resp, err := c.do(http.MethodPost, "/channels/"+channelID+"/messages", reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, err
	}
	return &message, nil
}

// GetUser returns info about a specific user
func (c *Client) GetUser(userID string) (*User, error) {
	resp, err := c.do(http.MethodGet, "/users/"+userID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetGuild returns info about a specific guild
func (c *Client) GetGuild(guildID string) (*Guild, error) {
	resp, err := c.do(http.MethodGet, "/guilds/"+guildID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var guild Guild
	if err := json.NewDecoder(resp.Body).Decode(&guild); err != nil {
		return nil, err
	}
	return &guild, nil
}

// GetMessage returns info about a specific message
func (c *Client) GetMessage(channelID, messageID string) (*Message, error) {
	resp, err := c.do(http.MethodGet, "/channels/"+channelID+"/messages/"+messageID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, err
	}
	return &message, nil
}

// DeleteMessage deletes a specific message
func (c *Client) DeleteMessage(channelID, messageID string) error {
	resp, err := c.do(http.MethodDelete, "/channels/"+channelID+"/messages/"+messageID, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// ModifyChannel modifies a channel
func (c *Client) ModifyChannel(channelID string, name string) (*Channel, error) {
	reqBody := ChannelModifyRequest{Name: name}
	resp, err := c.do(http.MethodPatch, "/channels/"+channelID, reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var channel Channel
	if err := json.NewDecoder(resp.Body).Decode(&channel); err != nil {
		return nil, err
	}
	return &channel, nil
}

// DeleteChannel deletes a channel
func (c *Client) DeleteChannel(channelID string) error {
	resp, err := c.do(http.MethodDelete, "/channels/"+channelID, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// EditMessage edits a message
func (c *Client) EditMessage(channelID, messageID, content string) (*Message, error) {
	reqBody := MessageEditRequest{Content: content}
	resp, err := c.do(http.MethodPatch, "/channels/"+channelID+"/messages/"+messageID, reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, err
	}
	return &message, nil
}

// GetGuildChannels returns channels in a guild
func (c *Client) GetGuildChannels(guildID string) ([]Channel, error) {
	resp, err := c.do(http.MethodGet, "/guilds/"+guildID+"/channels", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var channels []Channel
	if err := json.NewDecoder(resp.Body).Decode(&channels); err != nil {
		return nil, err
	}
	return channels, nil
}

// GetMeGuilds returns guilds the current user is in
func (c *Client) GetMeGuilds() ([]Guild, error) {
	resp, err := c.do(http.MethodGet, "/users/@me/guilds", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var guilds []Guild
	if err := json.NewDecoder(resp.Body).Decode(&guilds); err != nil {
		return nil, err
	}
	return guilds, nil
}
