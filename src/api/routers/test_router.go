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

	r.POST("/binder/headerno1", h.HeaderBinderNo1)
	r.POST("/binder/headerno2", h.HeaderBinderN02)

	r.POST("binder/query1", h.QueryBinderN01)
}
