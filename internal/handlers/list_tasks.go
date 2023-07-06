package handlers

import (
	"net/http"
	"strconv"
	"sword-project/internal/models"
	"sword-project/pkg/responses"

	"github.com/gin-gonic/gin"
)

type ListTasksRequest struct {
	Page int `form:"page,default=1"`
	Size int `form:"size,default=10"`
}

func buildListTaskParams(request *ListTasksRequest, user *models.Token) *models.ListTasksParams {

	teamID, _ := strconv.Atoi(user.Team)

	return &models.ListTasksParams{
		Page:   request.Page,
		Size:   request.Size,
		TeamID: teamID,
		UUID:   user.UUID,
		Role:   user.Role,
	}
}

func (ah *ApiHandler) ListTasks(c *gin.Context) {
	var params *ListTasksRequest
	ctx := c.Request.Context()

	userToken := ctx.Value("user").(*models.Token)

	if err := c.ShouldBindQuery(&params); err != nil {
		responses.ErrorResponse(c, err, http.StatusBadRequest, err.Error())
		return
	}
	result, err := ah.Services.ListTasks.List(ctx, buildListTaskParams(params, userToken))

	if err != nil {
		responses.ServerErrorResponse(c, err)
		return
	}

	responses.OkResponse(c, gin.H{
		"data": result,
	})
}
