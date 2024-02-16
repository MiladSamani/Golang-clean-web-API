package api

import (
	"fmt"
	"github.com/MiladSamani/Golang-clean-web-API/docs"

	"github.com/MiladSamani/Golang-clean-web-API/api/middlewares"
	"github.com/MiladSamani/Golang-clean-web-API/api/routers"
	"github.com/MiladSamani/Golang-clean-web-API/api/validations"
	"github.com/MiladSamani/Golang-clean-web-API/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterValidators()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())
	r.Use(middlewares.Cors(cfg))
	RegisterRoutes(r)
	RegisterSwagger(r, cfg)
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		HealthCheck := v1.Group("/health")
		routers.HealthCheckRoute(HealthCheck)

		Test := v1.Group("/test")
		routers.TestRouter(Test)
	}
	v2 := api.Group("/v2")
	{
		HealthCheck := v2.Group("/health-check")
		routers.HealthCheckRoute(HealthCheck)
	}
}

func RegisterValidators() {
	//custom validation add to gin
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
