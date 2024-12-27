package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func FavoriteFactory(db *gorm.DB) {
	favorites := []models.Favorite{
		{UserID: 1, BookID: 1},
		{UserID: 1, BookID: 2},
		{UserID: 2, BookID: 1},
		{UserID: 2, BookID: 2},
	}

	for _, favorite := range favorites {
		if err := db.FirstOrCreate(&models.Favorite{}, favorite).Error; err != nil {
			log.Printf("Failed to seed favorite")
		}
	}

	log.Println("Favorite seeded successfully")
}
