package initializers

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
)

func MigrationsInitializer() {
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")
}
