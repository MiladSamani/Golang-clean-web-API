package api

import (
	"fmt"

	"github.com/MiladSamani/Golang-clean-web-API/api/routers"
	"github.com/MiladSamani/Golang-clean-web-API/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	//gin.New() is a function that returns a new Gin Engine instance.
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		HealthCheck := v1.Group("/health-check")
		routers.HealthCheckRoute(HealthCheck)
	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
