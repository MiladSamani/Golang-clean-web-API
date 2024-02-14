package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {}

func NewHealthCheckHandler()  *HealthCheckHandler{
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheckResponse(c *gin.Context)  {
	c.JSON(http.StatusOK , "I'm Working...")
	return
}