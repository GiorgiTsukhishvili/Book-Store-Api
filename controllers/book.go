package controllers

import (
	"net/http"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("id")

	var Book models.Book

	if err := initializers.DB.Preload("Reviews").Preload("Genres").Preload("Favorites").Preload("Author").Preload("User").First(&Book, "id = ?", bookID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Book": Book})
}

func GetBooks(ctx *gin.Context) {
	var req requests.BookGetRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	paginate := scripts.Paginate(req.Page, req.Size, ctx)

	var books []models.Book

	query := initializers.DB

	if req.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+req.Keyword+"%")
	}

	if err := query.Scopes(paginate).Preload("Reviews").Preload("Genres").Preload("Favorites").Preload("Author").Preload("User").Find(&books).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "books not found"})
		return
	}

	var totalRecords int64
	if err := query.Model(&models.Book{}).Count(&totalRecords).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	size := scripts.ConvertStringToInt(req.Size, ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
		"pagination": gin.H{
			"current_page": req.Page,
			"first_page":   1,
			"last_page":    int(totalRecords) / size,
			"total":        totalRecords,
		},
	})
}

func PostBook(ctx *gin.Context) {
	var req requests.BookPostRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	image := scripts.SaveImage(ctx)

	claims := scripts.GetUserClaims(ctx)

	book := models.Book{
		Name:        req.Name,
		Description: req.Description,
		Image:       image,
		Price:       req.Price,
		AuthorID:    req.AuthorID,
		UserID:      claims.UserID,
	}

	if err := initializers.DB.Create(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, genreID := range req.GenreIDs {
		bookGenre := models.BookGenre{
			BookID:  int(book.ID),
			GenreID: int(genreID),
		}

		if err := initializers.DB.Create(&bookGenre).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func PutBook(ctx *gin.Context) {
	var req requests.BookPutRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := scripts.GetUserClaims(ctx)

	var image string

	if req.ImagePath != "" {
		image = req.ImagePath
	} else {
		image = scripts.SaveImage(ctx)
	}

	if err := initializers.DB.Model(models.Book{}).Where("id = ?", req.ID).Where("user_id = ?", claims.UserID).Updates(models.Book{
		Name:        req.Name,
		Description: req.Description,
		Image:       image,
		Price:       req.Price,
		AuthorID:    req.AuthorID,
	}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := initializers.DB.Delete(models.BookGenre{}, "book_id = ?", req.ID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, genreID := range req.GenreIDs {
		bookGenre := models.BookGenre{
			BookID:  int(req.ID),
			GenreID: int(genreID),
		}

		if err := initializers.DB.Create(&bookGenre).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
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
