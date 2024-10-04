package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/Golang-API/api/handlers"
	"github.com/hailsayan/Golang-API/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", h.SendOtp)
}
