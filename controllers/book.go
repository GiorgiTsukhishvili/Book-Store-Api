package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {
	BookID := ctx.Param("id")

	var Book models.Book

	if err := initializers.DB.Preload("Reviews").Preload("Genres").Preload("Author").Preload("User").First(&Book, "id = ?", BookID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Book": Book})
}

func GetBooks(ctx *gin.Context) {

}

func PostBook(ctx *gin.Context) {

}

func PutBook(ctx *gin.Context) {

}

func DeleteBook(ctx *gin.Context) {

}
