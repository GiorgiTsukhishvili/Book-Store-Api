package initializers

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
)

func MigrationsInitializer() {
	err := DB.AutoMigrate(&models.User{}, &models.Author{}, &models.Book{}, &models.Genre{}, &models.Review{}, &models.Favorite{}, &models.Notification{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")
}
