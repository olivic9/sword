package responses

import (
	"net/http"
	"sword-project/pkg/logging"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, status int, response gin.H) {
	c.JSON(status, response)
}

func OkResponse(c *gin.Context, response gin.H) {
	SuccessResponse(c, http.StatusOK, response)
}

func CreatedResponse(c *gin.Context) {
	SuccessResponse(c, http.StatusCreated, gin.H{})
}

func NoContentResponse(c *gin.Context) {
	SuccessResponse(c, http.StatusNoContent, gin.H{})
}

func ErrorResponse(c *gin.Context, err error, status int, message interface{}) {
	logging.Logger.Error(c.Request.Context(), err, logging.GinContextToMetadata(c))
	c.JSON(status, gin.H{
		"error": gin.H{
			"message": message,
		},
	})
}

// nolint
func ConflictErrorResponse(c *gin.Context, err error, message interface{}) {
	ErrorResponse(c, err, http.StatusConflict, message)
}

// nolint
func BadRequestErrorResponse(c *gin.Context, err error, message interface{}) {
	ErrorResponse(c, err, http.StatusBadRequest, message)
}

// nolint
func NotFoundErrorResponse(c *gin.Context, err error, message interface{}) {
	ErrorResponse(c, err, http.StatusNotFound, message)
}

// nolint
func UnauthorizedErrorResponse(c *gin.Context, err error, message interface{}) {
	ErrorResponse(c, err, http.StatusUnauthorized, message)
}

// nolint
func PreconditionFailedErrorResponse(c *gin.Context, err error, message interface{}) {
	ErrorResponse(c, err, http.StatusPreconditionFailed, message)
}

// nolint
func ServerErrorResponse(c *gin.Context, err error) {
	message := "Um problema inesperado aconteceu"
	ErrorResponse(c, err, http.StatusInternalServerError, message)
}
