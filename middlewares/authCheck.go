package middlewares

import (
	"net/http"
	"strings"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"github.com/gin-gonic/gin"
)

func AuthCheck(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		ctx.Abort()
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		ctx.Abort()
		return
	}

	tokenString := parts[1]

	claims := utils.ParseJwtToken(tokenString, ctx, "ACCESS_TOKEN_SECRET")

	ctx.Set("user", claims)

	ctx.Next()
}
