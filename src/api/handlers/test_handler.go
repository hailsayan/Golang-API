package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hailsayan/Golang-API/api/helper"
)

type header struct {
	UserId  string
	Browser string
}
type personDataa struct {
	FirstName    string `json:"FirstName" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"LastName" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"MobileNumber" binding:"required,mobile,min=11,max=11"`
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
	c.JSON(200, helper.GenerateBaseResponse("Test", true, 0))
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(200, helper.GenerateBaseResponse("Users", true, 0))
}

func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) HeaderBinderNo1(c *gin.Context) {
	userId := c.GetHeader("UserId")
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinderNo1",
		"userId": userId,
	}, true, 0))
}

func (h *TestHandler) HeaderBinderN02(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinderN02",
		"header": header,
	}, true, 0))
}

func (h *TestHandler) QueryBinderN01(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinderN01",
		"ids":    ids,
		"name":   name,
	}, true, 0))
}

func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	}, true, 0))
}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personDataa{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personDataa{}
	c.ShouldBind(&p)
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
	}, true, 0))
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	}, true, 0))
}
