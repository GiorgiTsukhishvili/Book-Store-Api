package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	token, err := jwt.ParseWithClaims(tokenString, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(*utils.CustomClaims)

	if !ok || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		ctx.Abort()
		return
	}

	ctx.Set("user", claims)

	ctx.Next()
}
