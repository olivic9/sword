package pubsub

import (
	"context"
	"log"
	"sword-project/pkg/configs"
	"sword-project/pkg/logging"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	context  context.Context
}

func NewKafkaConsumer(ctx context.Context) (*KafkaConsumer, error) {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KafkaCfg.BrokersHost,
		"group.id":          configs.KafkaCfg.ConsumerGroupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return &KafkaConsumer{consumer: nil, context: ctx}, ErrInvalidConsumer
	}

	return &KafkaConsumer{
		consumer: consumer,
		context:  ctx,
	}, nil
}

func (k *KafkaConsumer) Consume(topic string, handler KafkaConsumerHandler) {

	err := k.consumer.Subscribe(topic, nil)
	if err != nil {
		logging.Logger.Error(k.context, err, logging.Metadata{})
		log.Fatalf("topic subscribe error. topic: %s error %s", topic, err.Error())
	}

	defer func(consumer *kafka.Consumer) {
		err = consumer.Close()
		if err != nil {
			logging.Logger.Error(k.context, err, logging.Metadata{})
			log.Fatalf("close consumer error. topic: %s error %s", topic, err)
		}
	}(k.consumer)

	for {
		msg, msgError := k.consumer.ReadMessage(-1)

		if msgError == nil {

			handlerHasError := handler.HandleKafkaMessage(*msg)

			if handlerHasError != nil {
				logging.Logger.Error(k.context, err, logging.Metadata{})
				continue
			}

		}
	}
}
