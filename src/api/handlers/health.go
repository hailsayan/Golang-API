package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

func (h *HealthHandler) HealthPostById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, fmt.Sprintf("working post by id: %s", id))
	return
}
