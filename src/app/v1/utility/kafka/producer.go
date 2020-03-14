package kafka

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

// InitProducer ...
func (kafka *KafkaLibrary) InitProducer() (sarama.SyncProducer, error) {
	configKafka := kafka.Init("", "")
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_HOST")
	return sarama.NewSyncProducer([]string{kafkaHost + ":" + kafkaPort}, configKafka)
}

// GetMessageInput ...
func (kafka *KafkaLibrary) GetMessageInput() *ProducersMessageFormat {
	return &ProducersMessageFormat{}
}

// SendMessages ...
func (kafka *KafkaLibrary) SendEvent(topic string, payload *ProducersMessageFormat) (*ProducersMessageFormat, int64, error) {
	now := time.Now()
	fixPayload := &ProducersMessageFormat{}
	fixPayload.Action = payload.Action
	fixPayload.CreatedAt = &now
	fixPayload.Data = payload.Data
	fixPayload.UUID = uuid.New().String()
	// producers, err := kafka.InitProducer()
	// if err != nil {
	// 	return nil,0, err
	// }
	// data, err := json.Marshal(payload)
	// if err != nil {
	// 	return nil,0, err
	// }
	// kafkaMsg := &sarama.ProducerMessage{
	// 	Topic: topic,
	// 	Value: sarama.StringEncoder(data),
	// }
	// _, offset, err := producers.SendMessage(kafkaMsg)
	// fixPayload.Offset = offset
	// return offset, err
	return fixPayload, 1, nil
}
