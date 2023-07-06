package main

import (
	"context"
	"sword-project/internal/handlers"
	"sword-project/internal/repositories"
	"sword-project/pkg/configs"
	"sword-project/pkg/databases"
	"sword-project/pkg/exceptions"
	"sword-project/pkg/logging"
	"sword-project/pkg/pubsub"
)

var ctx = context.Background()

func main() {
	configs.InitializeConfigs()

	logging.InitializeApplicationLogger()

	consumerConfig := configs.NotificationConsumerConfig()
	defer logging.Logger.Sync()

	kafkaConsumer, err := pubsub.NewKafkaConsumer(ctx)

	if err != nil {
		logging.Logger.Fatal(ctx, exceptions.NewInternalHandledError(err.Error()), logging.Metadata{})
	}

	db := databases.GetMysqlDatabase(ctx)
	defer db.Close()

	repository := repositories.NewTaskRepository(db)
	handler := handlers.NewFinishedTaskNotificationHandler(repository)

	kafkaConsumer.Consume(consumerConfig.Topic, handler)

}
