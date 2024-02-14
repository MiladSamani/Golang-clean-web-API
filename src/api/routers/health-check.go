package routers

import (
	"github.com/MiladSamani/Golang-clean-web-API/api/handlers"
	"github.com/gin-gonic/gin"
)

func HealthCheckRoute(r *gin.RouterGroup)  {
	handlers:= handlers.NewHealthCheckHandler()
	r.GET("/", handlers.HealthCheckResponse)
}