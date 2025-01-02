package routes

import (
	"github.com/GiorgiTsukhishvili/BookShelf-Api/middlewares"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"github.com/gin-gonic/gin"
)

func ChannelRoutes(router *gin.Engine) {
	router.Use(middlewares.AuthCheck)
	router.GET("/notifications", utils.HandleWebSocket)
}
