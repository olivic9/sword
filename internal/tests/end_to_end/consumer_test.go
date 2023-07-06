package end_to_end

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"sword-project/internal/handlers"
	"sword-project/internal/repositories"
	"sword-project/pkg/configs"
	"sword-project/pkg/logging"
	"sword-project/pkg/pubsub"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/mock"
)

type Consumer interface {
	Subscribe(string, kafka.RebalanceCb) error
	Close() error
	ReadMessage(int) (*kafka.Message, error)
}

type MockConsumer struct {
	mock.Mock
}

func (m *MockConsumer) Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error {
	args := m.Called(topic, rebalanceCb)
	return args.Error(0)
}

func (m *MockConsumer) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConsumer) ReadMessage(timeout int) (*kafka.Message, error) {
	args := m.Called(timeout)
	return args.Get(0).(*kafka.Message), args.Error(1)
}

type MockKafkaConsumerHandler struct {
	mock.Mock
}

func (m *MockKafkaConsumerHandler) HandleKafkaMessage(message kafka.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

type KafkaConsumer struct {
	consumer *MockConsumer
}

func (c KafkaConsumer) Consume(handler pubsub.KafkaConsumerHandler) {

	msg, _ := c.consumer.ReadMessage(-1)

	_ = handler.HandleKafkaMessage(*msg)

}

const testMessage = `{"id":"01H4HJJGTYBD70SA822H1EM6QN","schema_version":"v1.0","source":"sword-test","data":{"ID":1,"FinishedAt":"2023-07-04T20:07:34.110781-03:00"},"timestamp":"2023-07-04T20:07:34.110851-03:00"}
`

func TestConsumer(t *testing.T) {

	configs.InitializeConfigs()

	logging.InitializeApplicationLogger()
	defer logging.Logger.Sync()

	mockConsumer := new(MockConsumer)
	mockHandler := new(MockKafkaConsumerHandler)
	kafkaConsumer := &KafkaConsumer{
		consumer: mockConsumer,
	}

	expectedMessage := "The tech John Doe performed the task test on date 2023-01-01 00:00:00 +0000 UTC \n"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repository := repositories.NewTaskRepository(db)
	handler := handlers.NewFinishedTaskNotificationHandler(repository)
	asOf, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")

	msg := &kafka.Message{
		Value: []byte(testMessage),
	}

	rows := sqlmock.NewRows([]string{"id", "title", "summary", "status", "created_at", "finished_at", "name"}).AddRow(1, "test", "test", "pending", asOf, asOf, "John Doe")
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT t.title, t.summary, t.status, t.created_at, t.finished_at, u.name 
			FROM tasks as t 
			LEFT JOIN users u ON t.assigned_technician_id = u.id 
			WHERE t.id = ?`,
	)).WithArgs(1).WillReturnRows(rows)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	mockConsumer.On("ReadMessage", -1).Return(msg, nil)

	kafkaConsumer.Consume(handler)

	mockConsumer.AssertExpectations(t)
	mockHandler.AssertExpectations(t)

	w.Close()
	os.Stdout = old
	out := <-outC
	if out != expectedMessage {
		t.Errorf("Expected '%s', but got '%s'", expectedMessage, out)
	}

}
