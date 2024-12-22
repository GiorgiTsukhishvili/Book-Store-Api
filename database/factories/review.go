package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func ReviewFactory(db *gorm.DB) {
	reviews := []models.Review{
		{Rating: "5", Comment: "Great", BookID: 1, UserID: 1},
		{Rating: "4", Comment: "Great", BookID: 2, UserID: 1},
	}

	for _, review := range reviews {
		if err := db.Create(&review).Error; err != nil {
			log.Fatalf("failed to seed reviews: %v", err)
		}
	}

	log.Println("Review seeded successfully")
}
