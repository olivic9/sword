package repositories

import (
	"sword-project/internal/models"

	"golang.org/x/net/context"
)

func (t *TaskRepository) NewTask(ctx context.Context, task *models.Task) error {
	var err error
	query := "INSERT INTO `tasks` (`title`, `summary`, `team_id`) VALUES (?, ?, ?)"
	_, err = t.database.ExecContext(ctx, query, task.Title, task.Summary, task.TeamID)
	return err
}
