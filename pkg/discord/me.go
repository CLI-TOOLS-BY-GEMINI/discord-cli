package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
