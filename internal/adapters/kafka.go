package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"sword-project/internal/models"
	"sword-project/pkg/exceptions"
	"sword-project/pkg/logging"
)

type KafkaService struct {
	kafkaProducer Publisher
}

func NewKafkaService(kafkaProducer Publisher) *KafkaService {
	return &KafkaService{kafkaProducer: kafkaProducer}
}

type Publisher interface {
	Publish(topic string, message []byte) error
}

func (k *KafkaService) ProduceMessage(message models.KafkaMessage, ctx context.Context, topic string) {
	go func() {
		messageBytes := new(bytes.Buffer)
		if err := json.NewEncoder(messageBytes).Encode(message); err != nil {
			logging.Logger.Error(ctx, exceptions.NewInternalHandledError(err.Error()), logging.Metadata{
				Context: "Encoding Error",
				Payload: message,
			})
			return
		}

		if err := k.kafkaProducer.Publish(topic, messageBytes.Bytes()); err != nil {
			logging.Logger.Error(ctx, exceptions.NewInternalHandledError(err.Error()), logging.Metadata{
				Context: topic,
				Payload: message,
			})
		}
	}()
}
