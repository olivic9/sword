package pubsub

import (
	"context"
	"sword-project/pkg/configs"
	"sword-project/pkg/exceptions"
	"sword-project/pkg/logging"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	context  context.Context
}

func NewKafkaPublisher(ctx context.Context) *KafkaProducer {
	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": configs.KafkaCfg.BrokersHost,
		},
	)
	if err != nil {
		logging.Logger.Fatal(ctx, err, logging.Metadata{})
	}

	return &KafkaProducer{
		producer: producer,
		context:  ctx,
	}
}

func (k *KafkaProducer) Shutdown() {
	_ = k.producer.Flush(15 * 1000)
	k.producer.Close()
}

func (k *KafkaProducer) Publish(topic string, message []byte) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}

	err := k.producer.Produce(msg, deliveryChan)
	if err != nil {
		logging.Logger.Error(k.context, err, logging.Metadata{})
		return exceptions.NewInternalHandledError(err.Error())
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		logging.Logger.Error(k.context, m.TopicPartition.Error, logging.Metadata{})
		return exceptions.NewInternalHandledError(m.TopicPartition.Error.Error())
	}

	return nil
}
