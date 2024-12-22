package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
	"gorm.io/gorm"
)

func UserFactory(db *gorm.DB) {
	userPassword, err := utils.HashPassword("password123")

	if err != nil {
		log.Fatal(err)
	}

	adminPassword, err := utils.HashPassword("admin123")

	if err != nil {
		log.Fatal(err)
	}

	businessPassword, err := utils.HashPassword("business123")

	if err != nil {
		log.Fatal(err)
	}

	users := []models.User{
		{Name: "John Doe", Email: "john@example.com", PhoneNumber: "+995511111111", Password: userPassword, Type: models.UserTypeUser},
		{Name: "Admin User", Email: "admin@example.com", PhoneNumber: "+995511111112", Password: adminPassword, Type: models.UserTypeAdmin},
		{Name: "Business User", Email: "business@example.com", PhoneNumber: "+995511111113", Password: businessPassword, Type: models.UserTypeBusiness},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("failed to seed users: %v", err)
		}
	}
}
