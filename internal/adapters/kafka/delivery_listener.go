package kafka

import (
	"context"
	"log"

	"github.com/KidPudel/client-service/internal/infrastructure/kafka"
)

type DeliveryListener struct {
	KafkaClient *kafka.KafkaClient
}

func NewDeliveryListener(kafkaClient *kafka.KafkaClient) *DeliveryListener {
	return &DeliveryListener{
		KafkaClient: kafkaClient,
	}
}

func (deliveryListener *DeliveryListener) ListenOnDeliveries(ctx context.Context) error {
	for {
		message, err := deliveryListener.KafkaClient.Reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		log.Printf("message from delivery: %s\n", message.Value)
	}
}
