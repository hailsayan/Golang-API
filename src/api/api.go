package api

import "github.com/gin-gonic/gin"

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, "working...")
			return
		})
	}
	r.Run(":8000")
}
