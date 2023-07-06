package models

import "time"

type Task struct {
	ID                     int64
	Title                  string
	Summary                string
	Status                 string
	TeamID                 int
	AssignedTechnicianID   int64
	AssignedTechnicianName string
	CreatedAt              time.Time
	FinishedAt             time.Time
}

type FinishedTask struct {
	ID         int64
	FinishedAt time.Time
}

type ListTasksParams struct {
	Page   int
	Size   int
	TeamID int
	UUID   string
	Role   string
}

type FinishTaskParams struct {
	TaskID int64
}
