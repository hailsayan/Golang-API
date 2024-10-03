package main

import (
	"github.com/hailsayan/Golang-API/api"
	"github.com/hailsayan/Golang-API/config"
	"github.com/hailsayan/Golang-API/pkg/logging"

	"github.com/hailsayan/Golang-API/data/cache"
	"github.com/hailsayan/Golang-API/data/db"
	"github.com/hailsayan/Golang-API/data/db/migrations"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migrations.Up_1()

	api.InitServer(cfg)
}
