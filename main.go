package main

import (
	"fmt"
	"os"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/config"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EnvInitializer()
}

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(config.CorsConfig())

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
