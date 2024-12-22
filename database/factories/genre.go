package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func GenreFactory(db *gorm.DB) {
	genres := []models.Genre{
		{Name: "Action"},
		{Name: "Comedy"},
		{Name: "Drama"},
		{Name: "Horror"},
		{Name: "Sci-Fi"},
	}

	for _, genre := range genres {
		if err := db.FirstOrCreate(&models.Genre{}, genre).Error; err != nil {
			log.Printf("Failed to seed genre '%s': %v", genre.Name, err)
		}
	}
}
