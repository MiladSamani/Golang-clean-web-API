package handlers

import (
	"net/http"

	"github.com/MiladSamani/Golang-clean-web-API/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {}

func NewHealthCheckHandler()  *HealthCheckHandler{
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheckResponse(c *gin.Context)  {
	c.JSON(http.StatusOK , helper.GenerateBaseResponse("i'm alive",true,0))
	return
}