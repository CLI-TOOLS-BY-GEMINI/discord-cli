package discord

import "time"

// User represents a Discord user
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
}

// Channel represents a Discord channel
type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

// Guild represents a Discord guild (server)
type Guild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	OwnerID     string `json:"owner_id"`
	Description string `json:"description"`
}

// Message represents a Discord message
type Message struct {
	ID        string    `json:"id"`
	ChannelID string    `json:"channel_id"`
	Author    User      `json:"author"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// MessageCreateRequest is the payload to send a message
type MessageCreateRequest struct {
	Content string `json:"content"`
}

// MessageEditRequest is the payload to edit a message
type MessageEditRequest struct {
	Content string `json:"content"`
}

// ChannelModifyRequest is the payload to modify a channel
type ChannelModifyRequest struct {
	Name string `json:"name,omitempty"`
}
