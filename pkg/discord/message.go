package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
