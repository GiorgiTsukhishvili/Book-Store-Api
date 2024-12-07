package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/translations"
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
		Image:    "default.png",
	}

	code := scripts.RandomNumber()

	fmt.Println(code)

	userData, err := json.Marshal(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to serialize user data: " + err.Error(),
		})
		return
	}

	if err := initializers.Redis.Set(context.Background(), code, userData, 30*time.Minute).Err(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	utils.SendEmail(user.Email, "Account verification", "en", code, user.Name, translations.GetTranslation("en", "joining-text"), translations.GetTranslation("en", "account-verification"))

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Verification email was sent",
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("refreshToken", "", -1, "/", "", true, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User logged out",
	})
}

func VerifyUser(ctx *gin.Context) {

	var req requests.UserVerifyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	data, err := initializers.Redis.Get(context.Background(), req.Code).Result()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid code",
			"err":   err,
		})
		return
	}

	if err := json.Unmarshal([]byte(data), &user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse user data: " + err.Error(),
		})
		return
	}

	if err := initializers.DB.Create(&user); err.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	initializers.Redis.Del(context.Background(), req.Code)

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
