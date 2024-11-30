package main

import (
	"github.com/GiorgiTsukhishvili/BookShelf-Api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(config.CorsConfig())

	router.Run(":3000")
}
