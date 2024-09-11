package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/Golang-API/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTesthandler()
	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/users/:id/", h.UserById)
}
