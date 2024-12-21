package factories

import (
	"log"
	"time"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func AuthorFactory(db *gorm.DB) {
	authors := []models.Author{
		{Name: "Jane Austen", BirthDate: time.Date(1775, 12, 16, 0, 0, 0, 0, time.UTC), Description: "English novelist known for her realism and social commentary.", Image: "jane_austen.jpg", Nationality: "British"},
		{Name: "Mark Twain", BirthDate: time.Date(1835, 11, 30, 0, 0, 0, 0, time.UTC), Description: "American writer and humorist best known for his novels.", Image: "mark_twain.jpg", Nationality: "American"},
	}

	for _, author := range authors {
		if err := db.Create(&author).Error; err != nil {
			log.Fatalf("failed to seed authors: %v", err)
		}
	}
}
