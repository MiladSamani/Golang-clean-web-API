package api

import (
	"fmt"

	"github.com/MiladSamani/Golang-clean-web-API/api/middlewares"
	"github.com/MiladSamani/Golang-clean-web-API/api/routers"
	"github.com/MiladSamani/Golang-clean-web-API/api/validations"
	"github.com/MiladSamani/Golang-clean-web-API/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()

	//gin.New() is a function that returns a new Gin Engine instance.
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery() , middlewares.LimitByRequest())
	r.Use(middlewares.Cors(cfg))

	//custom validation add to gin
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		HealthCheck := v1.Group("/health-check")
		routers.HealthCheckRoute(HealthCheck)

		Test := v1.Group("/test")
		routers.TestRouter(Test)

	}
	v2 := api.Group("/v2")
	{
		HealthCheck := v2.Group("/health-check")
		routers.HealthCheckRoute(HealthCheck)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
