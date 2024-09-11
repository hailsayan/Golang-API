package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTesthandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}
