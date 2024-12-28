package initializers

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/database/factories"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
)

func SeederInitializer() {
	var userCount int64

	DB.Model(&models.User{}).Count(&userCount)

	log.Println("Seeding database")

	if userCount == 0 {
		factories.UserFactory(DB)
		factories.AuthorFactory(DB)
		factories.BookFactory(DB)
		factories.GenreFactory(DB)
		factories.BookGenreFactory(DB)
		factories.ReviewFactory(DB)
		factories.FavoriteFactory(DB)
		factories.NotificationFactory(DB)
	}

	log.Println("Database seeded successfully")
}
