package middlewares

import "github.com/gin-gonic/gin"

func AuthCheck(ctx *gin.Context) {
	ctx.Next()
}
