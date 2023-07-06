package handlers

import (
	"net/http"
	"sword-project/internal/models"
	"sword-project/pkg/responses"

	"github.com/gin-gonic/gin"
)

type FinishTaskRequest struct {
	TaskID int64 `json:"task_id" binding:"required"`
}

func buildFinishTaskParams(request *FinishTaskRequest) *models.FinishTaskParams {
	return &models.FinishTaskParams{
		TaskID: request.TaskID,
	}
}

func (ah *ApiHandler) FinishTask(c *gin.Context) {
	var request *FinishTaskRequest

	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(c, err, http.StatusBadRequest, buildValidationErrorMessage(err))
		return
	}

	err := ah.Services.FinishTask.Finish(ctx, buildFinishTaskParams(request))

	if err != nil {
		responses.ErrorResponse(c, err, http.StatusInternalServerError, "internal error")
		return
	}

	responses.OkResponse(c, gin.H{})
}
