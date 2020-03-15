package entity

import (
	"time"
)

// UserEvent ...
type UserEvent struct {
	UUID      string            `json:"__uuid" bson:"__uuid"`
	Action    string            `json:"__action" bson:"__action"`
	Offset    int64             `json:"__offset" bson:"__offset"`
	Data      map[string]string `json:"data" bson:"data"`
	History   map[string]string `json:"history" bson:"history"`
	Status    string            `json:"status" bson:"status"`
	CreatedAt *time.Time        `json:"created_at" bson:"created_at"`
}
