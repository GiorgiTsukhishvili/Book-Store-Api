package scripts

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/types"
	"github.com/gin-gonic/gin"
)

func GetUserClaims(ctx *gin.Context) *types.CustomClaims {
	userInfo, exists := ctx.Get("user")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil
	}

	claims, ok := userInfo.(*types.CustomClaims)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return nil
	}

	return claims
}
