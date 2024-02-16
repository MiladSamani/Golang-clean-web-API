package handlers

import (
	"net/http"

	"github.com/MiladSamani/Golang-clean-web-API/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthCheckHandler) HealthCheckResponse(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("i'm alive", true, 0))
	return
}
