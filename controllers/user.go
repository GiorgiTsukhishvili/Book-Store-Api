package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"github.com/gin-gonic/gin"
)

func Me(ctx *gin.Context) {
	userInfo, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	claims, ok := userInfo.(*utils.CustomClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return
	}

	var user models.User

	if err := initializers.DB.Select("id", "name", "email", "image", "type", "created_at").First(&user, "id = ?", claims.UserID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"image":      user.Image,
		"type":       user.Type,
		"created_at": user.CreatedAt,
	}})
}
