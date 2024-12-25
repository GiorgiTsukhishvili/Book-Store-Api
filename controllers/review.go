package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/gin-gonic/gin"
)

func GetReview(ctx *gin.Context) {
	reviewID := ctx.Param("id")

	var Review models.Review

	if err := initializers.DB.First(&Review, "id = ?", reviewID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Review": Review})
}

func GetReviews(ctx *gin.Context) {

}

func PostReview(ctx *gin.Context) {

}

func PutReview(ctx *gin.Context) {

}
