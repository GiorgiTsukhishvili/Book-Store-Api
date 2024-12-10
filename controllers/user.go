package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/gin-gonic/gin"
)

func Me(ctx *gin.Context) {
	claims := scripts.GetUserClaims(ctx)

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

func PutUser(ctx *gin.Context) {
	var req requests.UserPutRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	if err := initializers.DB.Model(models.User{}).Where("id = ?", claims.UserID).Updates(models.User{Name: req.Name, Image: req.Image}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(ctx *gin.Context) {}
