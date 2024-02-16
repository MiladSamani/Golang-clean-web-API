package main

import (
	"github.com/MiladSamani/Golang-clean-web-API/api"
	"github.com/MiladSamani/Golang-clean-web-API/config"
	"github.com/MiladSamani/Golang-clean-web-API/data/cache"
	"github.com/MiladSamani/Golang-clean-web-API/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	if err != nil {
		return
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		return
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
