package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/Golang-API/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/health", handler.Health)
}
