package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func UserFactory(db *gorm.DB) {
	users := []models.User{
		{Name: "John Doe", Email: "john@example.com", Password: "password123", Type: models.UserTypeUser},
		{Name: "Admin User", Email: "admin@example.com", Password: "admin123", Type: models.UserTypeAdmin},
		{Name: "Business User", Email: "business@example.com", Password: "business123", Type: models.UserTypeBusiness},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("failed to seed users: %v", err)
		}
	}
}
