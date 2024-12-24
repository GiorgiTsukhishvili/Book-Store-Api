package middlewares

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/gin-gonic/gin"
)

func BusinessCheck(ctx *gin.Context) {
	claims := scripts.GetUserClaims(ctx)

	var user models.User

	if err := initializers.DB.Select("id", "name", "email", "image", "type", "created_at").First(&user, "id = ?", claims.UserID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Type != models.UserTypeBusiness {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is has no valid credentials"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
