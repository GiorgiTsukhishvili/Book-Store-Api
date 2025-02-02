package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/translations"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"github.com/gin-gonic/gin"
)

func Me(ctx *gin.Context) {
	claims := scripts.GetUserClaims(ctx)

	var user models.User

	if err := initializers.DB.Select("id", "name", "email", "phone_number", "image", "type", "created_at").First(&user, "id = ?", claims.UserID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": gin.H{
		"id":           user.ID,
		"name":         user.Name,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"image":        user.Image,
		"type":         user.Type,
		"created_at":   user.CreatedAt,
	}})
}

func PutUserPassword(ctx *gin.Context) {
	var req requests.UserPasswordPutRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.Password != req.RepeatPassword {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password must be same as repeat password",
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		log.Fatal(err)
	}

	if err := initializers.DB.Model(models.User{}).Where("id = ?", claims.UserID).Updates(models.User{Password: hashedPassword}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User password updated successfully"})

}

func PutUser(ctx *gin.Context) {
	var req requests.UserPutRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var image string

	if req.ImagePath != "" {
		image = req.ImagePath
	} else {
		image = scripts.SaveImage(ctx)
	}

	claims := scripts.GetUserClaims(ctx)

	if err := initializers.DB.Model(models.User{}).Where("id = ?", claims.UserID).Updates(models.User{Name: req.Name, Image: image}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := initializers.DB.Delete(&models.User{}, "id = ?", userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	scripts.InvalidateJwtCookies(ctx)

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func UserEmailUpdate(ctx *gin.Context) {
	var req requests.UserEmailPutRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	code := scripts.RandomNumber()

	if err := initializers.Redis.Set(context.Background(), code, req.Email, 30*time.Minute).Err(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	var user models.User

	if err := initializers.DB.Select("id", "name", "email", "phone_number", "image", "type", "created_at").First(&user, "id = ?", claims.UserID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	utils.SendEmail(user.Email, "Email Verification", "en", code, user.Name, translations.GetTranslation("en", "email-verification-text"), translations.GetTranslation("en", "email-verification"))

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Verification email was sent",
	})
}

func UserEmailUpdateVerify(ctx *gin.Context) {
	var req requests.UserVerifyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := initializers.Redis.Get(context.Background(), req.Code).Result()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid code",
			"err":   err,
		})
		return
	}

	var email string

	if err := json.Unmarshal([]byte(data), &email); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse user email: " + err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	if err := initializers.DB.Model(models.User{}).Where("id = ?", claims.UserID).Updates(models.User{Email: email}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User email updated successfully"})

}
