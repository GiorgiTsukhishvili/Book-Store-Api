package routes

import "github.com/gin-gonic/gin"

func ApiRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", func(ctx *gin.Context) {})
		v1.POST("/register", func(ctx *gin.Context) {})
	}
}
