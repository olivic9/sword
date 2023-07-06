package repositories

import (
	"context"
	"database/sql"
	"sword-project/internal/models"
)

func (t *TaskRepository) GetTaskByID(ctx context.Context, id int64) (*models.Task, error) {
	var err error
	var task models.Task

	query := `SELECT t.title, t.summary, t.status, t.created_at, t.finished_at, u.name 
				FROM tasks as t 
				LEFT JOIN users u ON t.assigned_technician_id = u.id 
				WHERE t.id = ?`

	result, err := t.database.QueryContext(ctx, query, id)
	if err != nil {
		return &task, err
	}
	defer result.Close()

	for result.Next() {
		var tmp struct {
			FinishedAt sql.NullTime
		}
		if err = result.Scan(&task.ID, &task.Title, &task.Summary, &task.Status, &task.CreatedAt, &tmp.FinishedAt, &task.AssignedTechnicianName); err != nil {
			return &task, err
		}
		if tmp.FinishedAt.Valid {
			task.FinishedAt = tmp.FinishedAt.Time
		}

	}

	if err = result.Err(); err != nil {
		return &task, err
	}

	return &task, err
}
