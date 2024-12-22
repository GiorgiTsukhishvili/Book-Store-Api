package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func BookGenreFactory(db *gorm.DB) {
	genres := []models.BookGenre{
		{BookID: 1, GenreID: 1},
		{BookID: 1, GenreID: 2},
		{BookID: 1, GenreID: 3},
		{BookID: 2, GenreID: 4},
		{BookID: 2, GenreID: 5},
	}

	for _, bookGenre := range genres {
		if err := db.Create(&bookGenre).Error; err != nil {
			log.Printf("Failed to insert BookGenre: %v", err)
		}
	}
}
