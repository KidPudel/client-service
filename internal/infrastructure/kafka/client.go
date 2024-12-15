package kafka

import (
	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	Reader *kafka.Reader
}

func NewKafkaClient() *KafkaClient {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "quickstart-events",
		Partition: 0,
		MaxBytes:  10e6,
	})
	return &KafkaClient{
		Reader: reader,
	}
}
