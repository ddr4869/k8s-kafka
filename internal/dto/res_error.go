package dto

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, code int, err error, message string) {
	c.AbortWithStatusJSON(code, ErrorResponse{
		Error:   err.Error(),
		Message: message,
	})
}
