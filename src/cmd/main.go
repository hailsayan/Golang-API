package main

import (
	"log"

	"github.com/hailsayan/Golang-API/api"
	"github.com/hailsayan/Golang-API/config"

	"github.com/hailsayan/Golang-API/data/cache"
	"github.com/hailsayan/Golang-API/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
