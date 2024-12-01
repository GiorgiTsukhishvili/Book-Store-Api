package controllers

import (
	"log"
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {}

func Register(ctx *gin.Context) {
	var req requests.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Type:     req.Type,
	}

	if err := initializers.DB.Create(&user); err.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	tokensInfo, err := utils.GenerateJWTTokens(user.ID, user.Email)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"type":  user.Type,
			"image": user.Image,
		},
		"jwt": tokensInfo,
	})
}

func Logout(ctx *gin.Context) {}
