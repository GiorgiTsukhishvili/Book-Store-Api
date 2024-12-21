package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/gin-gonic/gin"
)

func GetAuthor(ctx *gin.Context) {
	authorID := ctx.Param("id")

	var author models.Author

	if err := initializers.DB.Preload("Books").First(&author, "id = ?", authorID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"author": author})
}

func GetAuthors(ctx *gin.Context) {}

func PostAuthor(ctx *gin.Context) {}

func PutAuthor(ctx *gin.Context) {}

func DeleteAuthor(ctx *gin.Context) {}
