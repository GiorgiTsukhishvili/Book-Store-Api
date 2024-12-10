package scripts

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"github.com/gin-gonic/gin"
)

func GetUserClaims(ctx *gin.Context) *utils.CustomClaims {
	userInfo, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil
	}

	claims, ok := userInfo.(*utils.CustomClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return nil
	}

	return claims
}
