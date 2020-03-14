package event

import (
	"github.com/sofyan48/nemo/src/app/v1/api/user/entity"
	"github.com/sofyan48/nemo/src/app/v1/utility/kafka"
)

// UserEvent ...
type UserEvent struct {
	Kafka kafka.KafkaLibraryInterface
}

// UserEventHandler ...
func UserEventHandler() *UserEvent {
	return &UserEvent{
		Kafka: kafka.KafkaLibraryHandler(),
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
	go event.Kafka.SendEvent("test_topic", format)
	data.UUID = format.UUID
	data.Offset = format.Offset
	return data, nil
}
