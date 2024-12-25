package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("id")

	var Book models.Book

	if err := initializers.DB.Preload("Reviews").Preload("Genres").Preload("Author").Preload("User").First(&Book, "id = ?", bookID).Error; err != nil {
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
	bookId := ctx.Param("id")

	if err := initializers.DB.Delete(models.Book{}, "id = ?", bookId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
