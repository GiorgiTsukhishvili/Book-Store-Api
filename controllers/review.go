package controllers

import (
	"net/http"
	"strconv"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
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
	var req requests.ReviewGetRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	paginate := scripts.Paginate(req.Page, req.Size, ctx)

	var reviews []models.Review

	if err := initializers.DB.Scopes(paginate).Preload("User").Preload("Book").Find(&reviews).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "reviews not found"})
		return
	}

	var totalRecords int64
	if err := initializers.DB.Model(&models.Review{}).Count(&totalRecords).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	size, _ := strconv.Atoi(req.Size)

	ctx.JSON(http.StatusOK, gin.H{
		"data": reviews,
		"pagination": gin.H{
			"current_page": req.Page,
			"first_page":   1,
			"last_page":    int(totalRecords) / size,
			"total":        totalRecords,
		},
	})
}

func PostReview(ctx *gin.Context) {
	var req requests.ReviewPostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	bookId, err := strconv.Atoi(req.BookID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var review = models.Review{
		Rating:  req.Rating,
		BookID:  uint(bookId),
		Comment: req.Comment,
		UserID:  claims.UserID,
	}

	if err := initializers.DB.Create(&review).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"review": review,
	})
}

func PutReview(ctx *gin.Context) {
	var req requests.ReviewPutRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)
}
