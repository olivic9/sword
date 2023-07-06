package models

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/oklog/ulid/v2"
)

const (
	schemaVersion = "v1.0"
	source        = "sword-test"
)

type KafkaMessage struct {
	Id            string      `json:"id"`
	SchemaVersion string      `json:"schema_version"`
	Source        string      `json:"source"`
	Data          interface{} `json:"data"`
	Timestamp     time.Time   `json:"timestamp"`
}

func CreateFinishedTaskKafkaMessage(model *FinishedTask) KafkaMessage {
	return KafkaMessage{
		Id:            ulid.Make().String(),
		SchemaVersion: schemaVersion,
		Source:        source,
		Data:          model,
		Timestamp:     time.Now(),
	}
}

type FinishedTaskMessage struct {
	Id            string `json:"id"`
	SchemaVersion string `json:"schema_version"`
	Source        string `json:"source"`
	Data          struct {
		ID         int64     `json:"ID"`
		FinishedAt time.Time `json:"FinishedAt"`
	} `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func (f *FinishedTaskMessage) From(message kafka.Message) (FinishedTaskMessage, error) {

	err := json.Unmarshal(message.Value, &f)

	if err != nil {
		return *f, err
	}

	return *f, nil
}
