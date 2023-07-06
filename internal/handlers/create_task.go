package handlers

import (
	"net/http"
	"strconv"
	"sword-project/internal/models"
	"sword-project/pkg/responses"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title   string `json:"title" binding:"required"`
	Summary string `json:"summary" binding:"required" validate:"max=2500"`
}

func buildCreateTaskModel(request *CreateTaskRequest, user *models.Token) *models.Task {
	teamID, _ := strconv.Atoi(user.Team)

	return &models.Task{
		Title:   request.Title,
		Summary: request.Summary,
		TeamID:  teamID,
	}
}

func (ah *ApiHandler) NewTask(c *gin.Context) {

	var request *CreateTaskRequest

	ctx := c.Request.Context()

	userToken := ctx.Value("user").(*models.Token)

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(c, err, http.StatusBadRequest, buildValidationErrorMessage(err))
		return
	}

	err := ah.Services.CreateTask.Execute(ctx, buildCreateTaskModel(request, userToken))

	if err != nil {
		responses.ErrorResponse(c, err, http.StatusInternalServerError, "internal error")
		return
	}

	responses.CreatedResponse(c)
}
