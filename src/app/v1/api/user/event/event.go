package event

import (
	"github.com/google/uuid"
	"github.com/sofyan48/nemo/src/app/v1/api/user/entity"
	"github.com/sofyan48/nemo/src/app/v1/utility/kafka"
	"github.com/sofyan48/nemo/src/app/v1/utility/mongodb"
)

// USEREVENT ...
const USEREVENT = "user_event"

// UserEvent ...
type UserEvent struct {
	Kafka kafka.KafkaLibraryInterface
	Mongo mongodb.MongoDBInterface
}

// UserEventHandler ...
func UserEventHandler() *UserEvent {
	return &UserEvent{
		Kafka: kafka.KafkaLibraryHandler(),
		Mongo: mongodb.MongoDBHandler(),
	}
}

// UserEventInterface ...
type UserEventInterface interface {
	UserCreateEvent(data *entity.UserEvent) (*entity.UserEvent, error)
}

// UserCreateEvent ...
func (event *UserEvent) UserCreateEvent(data *entity.UserEvent) (*entity.UserEvent, error) {
	format := event.Kafka.GetStateFull()
	format.Action = data.Action
	format.CreatedAt = data.CreatedAt
	format.Data = data.Data
	format.UUID = uuid.New().String()
	data.UUID = format.UUID
	go event.Kafka.SendEvent(USEREVENT, format)
	return data, nil
}
