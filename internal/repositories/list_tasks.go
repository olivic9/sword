package repositories

import (
	"context"
	"database/sql"
	"sword-project/internal/models"
)

func (t *TaskRepository) ListTasks(ctx context.Context, params *models.ListTasksParams) (*[]models.Task, error) {
	var tasks []models.Task
	var query string
	var result *sql.Rows
	var err error

	switch params.Role {
	case "Manager":
		query = `SELECT id, title, summary, status, created_at, finished_at
				FROM tasks 
				WHERE team_id = ?
				LIMIT ? 
				OFFSET ?`
		result, err = t.database.QueryContext(ctx, query, params.TeamID, params.Size, params.Page)
	default:
		query = `SELECT t.id,  t.title, t.summary, t.status, t.created_at, t.finished_at
					FROM tasks as t
					LEFT JOIN users ON t.assigned_technician_id = users.id
					WHERE t.assigned_technician_id IS NULL OR users.uuid = ?
					LIMIT ? 
					OFFSET ?`
		result, err = t.database.QueryContext(ctx, query, params.UUID, params.Size, params.Page)
	}

	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var task models.Task
		var tmp struct {
			FinishedAt sql.NullTime
		}
		if err = result.Scan(&task.ID, &task.Title, &task.Summary, &task.Status, &task.CreatedAt, &tmp.FinishedAt); err != nil {
			return &tasks, err
		}
		if tmp.FinishedAt.Valid {
			task.FinishedAt = tmp.FinishedAt.Time
		}
		tasks = append(tasks, task)
	}

	if err = result.Err(); err != nil {
		return &tasks, err
	}

	return &tasks, err
}
