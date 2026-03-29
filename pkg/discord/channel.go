package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
