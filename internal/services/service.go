package services

type Services struct {
	CreateTask *CreateTaskService
	ListTasks  *ListTasksService
	FinishTask *FinishTaskService
}

func NewServices(
	createTaskService *CreateTaskService,
	listTasks *ListTasksService,
	finishTasks *FinishTaskService,
) Services {
	return Services{
		CreateTask: createTaskService,
		ListTasks:  listTasks,
		FinishTask: finishTasks,
	}
}
