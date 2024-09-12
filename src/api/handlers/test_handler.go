package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type header struct {
	UserId  string
	Browser string
}
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

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}

func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}

func (h *TestHandler) HeaderBinderNo1(c *gin.Context) {
	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinderNo1",
		"userId": userId,
	})
}

func (h *TestHandler) HeaderBinderN02(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinderN02",
		"header": header,
	})
}
