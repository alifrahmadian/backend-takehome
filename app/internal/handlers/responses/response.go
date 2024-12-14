package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Success{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Error{
		Status:       "error",
		ErrorCode:    statusCode,
		ErrorMessage: message,
	})
}
