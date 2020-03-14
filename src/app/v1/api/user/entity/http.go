package entity

import "time"

// UserRequest ...
type UserRequest struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Region  string `json:"region" form:"region"`
}

// UserResponses ...
type UserResponses struct {
	UUID      string     `json:"uuid"`
	CreatedAt *time.Time `json:"created_at"`
	Event     *UserEvent `json:"event"`
}
