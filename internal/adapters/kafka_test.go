package adapters

import (
	"context"
	"errors"
	"fmt"
	"sword-project/internal/models"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

const producerError = "producer error"

type produceMessageTestCase struct {
	name    string
	message models.KafkaMessage
	topic   string
	error   error
}

type FakePublisher struct {
	topics map[string][][]byte
}

func (p *FakePublisher) Publish(topic string, message []byte) error {
	p.topics[topic] = append(p.topics[topic], message)
	if topic == "error_topic" {
		return errors.New(producerError)
	}
	return nil
}

func TestProduceMessage(t *testing.T) {
	testCases := []produceMessageTestCase{
		{
			name: "valid message",
			message: models.KafkaMessage{
				Id:            ulid.Make().String(),
				SchemaVersion: "vtest",
				Source:        "vname",
				Data:          "vname",
				Timestamp:     time.Now(),
			},
			topic: "test",
			error: nil,
		},
		{
			name: producerError,
			message: models.KafkaMessage{
				Id:            ulid.Make().String(),
				SchemaVersion: "itest",
				Source:        "iname",
				Data:          "iname",
				Timestamp:     time.Now(),
			},
			topic: "error_topic",
			error: errors.New(producerError),
		},
	}
	ctx := context.TODO()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			publisher := &FakePublisher{topics: make(map[string][][]byte)}
			adapter := NewKafkaService(publisher)

			adapter.ProduceMessage(tc.message, ctx, tc.topic)
			fmt.Println(tc.error)

		})
	}
}
