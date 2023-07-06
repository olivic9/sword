package handlers

import (
	"context"
	"fmt"
	"sword-project/internal/core/domain/task/use_case/get"
	"sword-project/internal/models"
	"sword-project/internal/repositories"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type FinishedTaskNotificationHandler struct {
	taskRepository *repositories.TaskRepository
}

func NewFinishedTaskNotificationHandler(taskRepository *repositories.TaskRepository) *FinishedTaskNotificationHandler {
	return &FinishedTaskNotificationHandler{
		taskRepository: taskRepository,
	}
}

func (f FinishedTaskNotificationHandler) HandleKafkaMessage(message kafka.Message) error {

	useCase := get.NewGetTaskUseCase(f.taskRepository)

	finishedMessageModel := models.FinishedTaskMessage{}

	finishedMessageModel, err := finishedMessageModel.From(message)

	if err != nil {
		return err
	}

	task, err := useCase.Get(context.Background(), finishedMessageModel.Data.ID)

	if err == nil {
		fmt.Printf("The tech %s performed the task %s on date %s \n", task.AssignedTechnicianName, task.Title, task.FinishedAt)
	}

	return err
}
