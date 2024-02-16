package main

import (
	"github.com/MiladSamani/Golang-clean-web-API/api"
	"github.com/MiladSamani/Golang-clean-web-API/config"
	"github.com/MiladSamani/Golang-clean-web-API/data/cache"
	"github.com/MiladSamani/Golang-clean-web-API/data/db"
	"github.com/MiladSamani/Golang-clean-web-API/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()

	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
