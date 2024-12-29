package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/gin-gonic/gin"
)

func GetNotifications(ctx *gin.Context) {
	var req requests.NotificationGetRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	paginate := scripts.Paginate(req.Page, req.Size, ctx)

	var notifications []models.Notification

	if err := initializers.DB.Scopes(paginate).Preload("User").Preload("Book").Find(&notifications).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "notifications not found"})
		return
	}

	var totalRecords int64
	if err := initializers.DB.Model(&models.Notification{}).Count(&totalRecords).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	size := scripts.ConvertStringToInt(req.Size, ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"data": notifications,
		"pagination": gin.H{
			"current_page": req.Page,
			"first_page":   1,
			"last_page":    int(totalRecords) / size,
			"total":        totalRecords,
		},
	})
}

func PutNotification(ctx *gin.Context) {

}
