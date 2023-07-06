package repositories

import (
	"context"
	"errors"
	"sword-project/internal/models"
)

func (t *TaskRepository) FinishTask(ctx context.Context, params *models.FinishTaskParams) error {

	query := `UPDATE tasks SET status = "finished", finished_at = NOW() WHERE id = ? and finished_at is null`
	result, err := t.database.ExecContext(ctx, query, params.TaskID)
	if err != nil {
		return err
	}

	rowCnt, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCnt == 0 {
		return errors.New("not found")
	} else if rowCnt > 1 {
		return errors.New("fatal error")
	}

	return nil
}
