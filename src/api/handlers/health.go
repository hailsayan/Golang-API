package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/Golang-API/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, helper.GenerateBaseResponse("working...", true, 0))
	return
}
