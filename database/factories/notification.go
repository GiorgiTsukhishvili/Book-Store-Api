package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func NotificationFactory(db *gorm.DB) {
	notifications := []models.Notification{
		{IsNew: true, Type: "favorite", BookID: 1, UserID: 1, PersonID: 2},
		{IsNew: true, Type: "review", BookID: 1, UserID: 1, PersonID: 2},
		{IsNew: true, Type: "favorite", BookID: 2, UserID: 2, PersonID: 1},
		{IsNew: true, Type: "review", BookID: 2, UserID: 2, PersonID: 1},
	}

	for _, notification := range notifications {
		if err := db.Create(&notification).Error; err != nil {
			log.Fatalf("failed to seed notifications: %v", err)
		}
	}

	log.Println("Notification seeded successfully")
}
