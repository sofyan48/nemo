package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

// KafkaLibrary ...
type KafkaLibrary struct{}

// ProducersMessageFormat ...
type StateFullFormat struct {
	UUID      string            `json:"__id"`
	Action    string            `json:"__action"`
	Data      map[string]string `json:"data"`
	Offset    int64             `json:"offset"`
	CreatedAt *time.Time        `json:"created_at"`
}

// KafkaLibraryHandler ...
func KafkaLibraryHandler() *KafkaLibrary {
	return &KafkaLibrary{}
}

// KafkaLibraryInterface ...
type KafkaLibraryInterface interface {
	GetStateFull() *StateFullFormat
	SendEvent(topic string, payload *StateFullFormat) (*StateFullFormat, int64, error)
	InitConsumer() (sarama.Consumer, error)
}

// GetStateFull ...
func (kafka *KafkaLibrary) GetStateFull() *StateFullFormat {
	return &StateFullFormat{}
}

// Init ...
func (kafka *KafkaLibrary) init(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
