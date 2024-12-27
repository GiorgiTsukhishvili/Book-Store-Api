package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/gin-gonic/gin"
)

func GetUserFavorites(ctx *gin.Context) {
	var req requests.UserFavoriteGetRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	paginate := scripts.Paginate(req.Page, req.Size, ctx)

	var favorites []models.Favorite

	if err := initializers.DB.Scopes(paginate).Preload("User").Preload("Book").Find(&favorites).Where("user_id = ?", claims.UserID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "favorites not found"})
		return
	}

	var totalRecords int64
	if err := initializers.DB.Model(&models.Favorite{}).Count(&totalRecords).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	size := scripts.ConvertStringToInt(req.Size, ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"data": favorites,
		"pagination": gin.H{
			"current_page": req.Page,
			"first_page":   1,
			"last_page":    int(totalRecords) / size,
			"total":        totalRecords,
		},
	})
}

func PostFavorite(ctx *gin.Context) {
	var req requests.FavoritePostRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	BookID := scripts.ConvertStringToInt(req.BookID, ctx)

	favorite := models.Favorite{
		BookID: uint(BookID),
		UserID: claims.UserID,
	}

	if err := initializers.DB.Create(&favorite).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"favorite": favorite,
	})
}

func DeleteFavorite(ctx *gin.Context) {}
