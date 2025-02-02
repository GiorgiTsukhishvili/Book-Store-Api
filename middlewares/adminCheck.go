package middlewares

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/gin-gonic/gin"
)

func AdminCheck(ctx *gin.Context) {
	claims := scripts.GetUserClaims(ctx)

	var user models.User

	if err := initializers.DB.Select("id", "name", "email", "image", "type", "created_at").First(&user, "id = ?", claims.UserID).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Type != models.UserTypeAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is has no valid credentials"})
		return
	}

	ctx.Next()
}
