package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct {
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, "working...")
	return
}
