package pubsub

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaConsumerHandler interface {
	HandleKafkaMessage(message kafka.Message) error
}

type Consumer interface {
	Consume(topic string, handler KafkaConsumerHandler)
}
