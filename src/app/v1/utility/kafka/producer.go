package kafka

import (
	"encoding/json"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

// ProducersMessageFormat ...
type ProducersMessageFormat struct {
	UUID      string            `json:"__id"`
	Action    string            `json:"__action"`
	Data      map[string]string `json:"data"`
	CreatedAt *time.Time        `json:"created_at"`
}

// InitProducer ...
func (kafka *KafkaLibrary) InitProducer() (sarama.SyncProducer, error) {
	configKafka := kafka.Init("", "")
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_HOST")
	return sarama.NewSyncProducer([]string{kafkaHost + ":" + kafkaPort}, configKafka)
}

// SendMessages ...
func (kafka *KafkaLibrary) SendMessages(topic string, payload map[string]string) (int64, error) {
	producers, err := kafka.InitProducer()
	if err != nil {
		return 0, err
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}
	_, offset, err := producers.SendMessage(kafkaMsg)
	return offset, err
}
