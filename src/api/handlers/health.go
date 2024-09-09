package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, "working...")
	return
}

func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(200, "post method")
	return
}
