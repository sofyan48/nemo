package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

// KafkaLibrary ...
type KafkaLibrary struct{}

// KafkaLibraryHandler ...
func KafkaLibraryHandler() *KafkaLibrary {
	return &KafkaLibrary{}
}

// KafkaLibraryInterface ...
type KafkaLibraryInterface interface {
	Init(username, password string) *sarama.Config
}

// Init ...
func (kafka *KafkaLibrary) Init(username, password string) *sarama.Config {
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
