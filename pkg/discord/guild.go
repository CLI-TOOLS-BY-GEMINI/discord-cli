package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
