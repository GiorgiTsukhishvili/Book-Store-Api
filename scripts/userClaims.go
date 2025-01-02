package scripts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}

func GetUserClaims(ctx *gin.Context) *UserClaims {
	userInfo, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil
	}

	claims, ok := userInfo.(*UserClaims)

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return nil
	}

	return claims
}
