package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/config"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EnvInitializer()
	initializers.GormInitializer()
	initializers.EnumsInitializer()
	initializers.MigrationsInitializer()
	initializers.RedisInitializer()
}

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies(strings.Split(os.Getenv("TRUSTED_PROXIES"), ","))

	router.Use(config.CorsConfig())

	routes.ApiRoutes(router)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
