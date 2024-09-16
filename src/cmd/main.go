package main

import (
	"github.com/hailsayan/Golang-API/api"
	"github.com/hailsayan/Golang-API/config"
	"github.com/hailsayan/Golang-API/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
