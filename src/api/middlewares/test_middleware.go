package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-Key")
		if apiKey == "1" {
			ctx.Next()
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"result": "Api Key is required",
		})
		return
	}
}
