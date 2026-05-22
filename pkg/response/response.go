package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Body{Success: true, Message: "success", Data: data})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Body{Success: true, Message: "success", Data: data})
}

func Error(c *gin.Context, status int, message string, code string) {
	c.JSON(status, Body{Success: false, Message: message, Error: code})
}
